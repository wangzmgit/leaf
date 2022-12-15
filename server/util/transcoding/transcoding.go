package transcoding

import (
	"bytes"
	"encoding/json"
	"errors"
	"os/exec"
	"strconv"
)

type VideoInfoData struct {
	Stream []Streams `json:"streams"`
	Format Format    `json:"format"`
}

type Streams struct {
	CodecName string `json:"codec_name"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
	PixFmt    string `json:"pix_fmt,omitempty"`
	Duration  string `json:"duration"`
}

type Format struct {
	BitRate string `json:"bit_rate"`
}

// 获取视频信息
func GetVideoInfo(input string) (videoData VideoInfoData, err error) {
	cmd := exec.Command("ffprobe", "-i", input, "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams")
	out, err := runCmd(cmd)
	if err != nil {
		return videoData, err
	}

	// 反序列化
	err = json.Unmarshal(out.Bytes(), &videoData)
	if err != nil {
		return videoData, err
	}

	return videoData, nil
}

// 提取音频
func ExtractingAudio(inputFile, outputDir string) (string, error) {
	output := outputDir + "audio.m4a"
	cmd := exec.Command("ffmpeg", "-hide_banner", "-i", inputFile, "-vn", "-c", "copy", output)
	_, err := runCmd(cmd)
	return output, err
}

// 压制视频
func PressingVideo(inputFile, outputDir string, quality int) ([]string, error) {
	outputFileList := make([]string, 0)
	command := []string{"-hide_banner", "-i", inputFile, "-crf", "20"}
	switch quality {
	case 1080:
		output1080 := outputDir + "tmp_1080p_3000k.mp4"
		outputFileList = append(outputFileList, output1080)
		command = append(command, "-c:v", "libx264", "-an", "-s", "1920x1080", "-r", "30000/1001", "-b:v", "3000k", output1080)
		fallthrough
	case 720:
		output720 := outputDir + "tmp_720p_2000k.mp4"
		outputFileList = append(outputFileList, output720)
		command = append(command, "-c:v", "libx264", "-an", "-s", "1080x720", "-r", "30000/1001", "-b:v", "2000k", output720)
		fallthrough
	case 480:
		output480 := outputDir + "tmp_480p_900k.mp4"
		outputFileList = append(outputFileList, output480)
		command = append(command, "-c:v", "libx264", "-an", "-s", "854x480", "-r", "30000/1001", "-b:v", "900k", output480)
		fallthrough
	case 360:
		output360 := outputDir + "tmp_360p_500k.mp4"
		outputFileList = append(outputFileList, output360)
		command = append(command, "-c:v", "libx264", "-an", "-s", "640x360", "-r", "30000/1001", "-b:v", "500k", output360)
	}

	// 翻转 outputFileList ，从低分辨率到高分辨率
	for i, j := 0, len(outputFileList)-1; i < j; i, j = i+1, j-1 {
		outputFileList[i], outputFileList[j] = outputFileList[j], outputFileList[i]
	}

	_, err := runCmd(exec.Command("ffmpeg", command...))
	return outputFileList, err
}

func GenerateDash(videoFiles []string, audioFile, outputDir, outputName string) error {
	mapCommand := make([]string, 0)
	command := make([]string, 0)

	mpdName := outputDir + "index.mpd"
	initStreamName := outputName + "-init-$RepresentationID$.m4s"
	chunkStreamName := outputName + "-$RepresentationID$-$Number%05d$.m4s"

	// 添加视频
	for i := 0; i < len(videoFiles); i++ {
		command = append(command, "-i", videoFiles[i])
		mapCommand = append(mapCommand, "-map", strconv.Itoa(i))
	}

	// 添加音频
	command = append(command, "-i", audioFile)
	mapCommand = append(mapCommand, "-map", strconv.Itoa(len(videoFiles)))

	// 合并命令
	command = append(command, "-c", "copy")
	command = append(command, mapCommand...)
	command = append(command, "-f", "dash", "-init_seg_name", initStreamName, "-media_seg_name", chunkStreamName, mpdName)

	_, err := runCmd(exec.Command("ffmpeg", command...))
	return err
}

func runCmd(cmd *exec.Cmd) (bytes.Buffer, error) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return out, errors.New(stderr.String())
	}

	return out, nil
}

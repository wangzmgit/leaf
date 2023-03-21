package vo

type EmailConfigVO struct {
	// Debug     bool   `json:"debug"`
	User      string `json:"user"`
	Pass      string `json:"pass"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Addresser string `json:"addresser"`
}

type StorageConfigVO struct {
	MaxImgSize         int  `json:"max_img_size"`
	MaxVideoSize       int  `json:"max_video_size"`
	VideoAdaptationIos bool `json:"video_adaptation_ios"`

	Type      string `json:"type"`
	KeyID     string `json:"key_id"`
	KeySecret string `json:"key_secret"`
	Bucket    string `json:"bucket"`
	Endpoint  string `json:"endpoint"`
	AppID     string `json:"app_id"`
	Region    string `json:"region"`
	Domain    string `json:"domain"`
	Private   bool   `json:"private"`
}

type OtherConfigVO struct {
	AllowOrigin string `json:"allow_origin"`
	Prefix      string `json:"prefix"`
}

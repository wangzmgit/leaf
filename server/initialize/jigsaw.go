package initialize

import "github.com/wangzmgit/jigsaw"

func Jigsaw() {
	j := jigsaw.New()
	j.SetBgDir("./static/images/bg/")
	j.SetMaskPath("./static/images/mask.png")
}

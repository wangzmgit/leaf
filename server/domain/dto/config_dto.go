package dto

type EmailConfigDTO struct {
	Debug     bool
	User      string
	Pass      string
	Host      string
	Port      int
	Addresser string
}

type StorageConfigDTO struct {
	MaxImgSize         int
	MaxVideoSize       int
	VideoAdaptationIos bool

	Type      string
	KeyID     string
	KeySecret string
	Bucket    string
	Endpoint  string
	AppID     string
	Region    string
	Domain    string
	Private   bool
}

type OtherConfigDTO struct {
	AllowOrigin string
	Prefix      string
}

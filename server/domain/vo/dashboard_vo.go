package vo

type CardDataVO struct {
	UserCount        int64 `json:"user_count"`
	VideoCount       int64 `json:"video_count"`
	ReviewVideoCount int64 `json:"review_video_count"`
	NewUserCount     int64 `json:"new_user_count"`
	NewVideoCount    int64 `json:"new_video_count"`
}

type OneDayTrendVO struct {
	User  int64  `json:"user"`
	Video int64  `json:"video"`
	Date  string `json:"date"`
}

type PartitionVideoCount struct {
	Content    string `json:"content"`
	VideoCount int    `json:"video_count"`
}

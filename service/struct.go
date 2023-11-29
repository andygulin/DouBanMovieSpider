package service

type SubjectResponse struct {
	SubjectId  string `json:"subject_id" bson:"subject_id"`
	Title      string `json:"title" bson:"title"`
	CoverImage string `json:"cover_image" bson:"cover_image"`
	Year       int    `json:"year" bson:"year"`

	Director    []string `json:"director" bson:"director"`
	Writer      []string `json:"writer" bson:"writer"`
	LeadingRole []string `json:"leading_role" bson:"leading_role"`
	Type        []string `json:"type" bson:"type"`

	Nation        []string `json:"nation" bson:"nation"`
	Language      []string `json:"language" bson:"language"`
	ReleaseDate   []string `json:"release_date" bson:"release_date"`
	Length        string   `json:"length" bson:"length"`
	AlternateName []string `json:"alternate_name" bson:"alternate_name"`
	IMDB          string   `json:"imdb" bson:"imdb"`

	Intro string `json:"intro" bson:"intro"`

	CommentCount     int     `json:"comment_count" bson:"comment_count"`
	RatingNum        float32 `json:"rating_num" bson:"rating_num"`
	RatingNumPercent struct {
		RatingNum5 float32 `json:"rating_num_5" bson:"rating_num_5"`
		RatingNum4 float32 `json:"rating_num_4" bson:"rating_num_4"`
		RatingNum3 float32 `json:"rating_num_3" bson:"rating_num_3"`
		RatingNum2 float32 `json:"rating_num_2" bson:"rating_num_2"`
		RatingNum1 float32 `json:"rating_num_1" bson:"rating_num_1"`
	} `json:"rating_num_percent" bson:"rating_num_percent"`
}

type CommentResponse struct {
	SubjectId   string `json:"subject_id" bson:"subject_id"`
	Avatar      string `json:"avatar" bson:"avatar"`
	UserName    string `json:"user_name" bson:"user_name"`
	Star        int    `json:"star" bson:"star"`
	CommentDate string `json:"comment_date" bson:"comment_date"`
	Content     string `json:"content" bson:"content"`
	HaveSee     bool   `json:"have_see" bson:"have_see"` // true status=P , false status=F
}

type ReviewResponse struct {
	SubjectId   string `json:"subject_id" bson:"subject_id"`
	Avatar      string `json:"avatar" bson:"avatar"`
	UserName    string `json:"user_name" bson:"user_name"`
	Star        int    `json:"star" bson:"star"`
	CommentDate string `json:"comment_date" bson:"comment_date"`
	Content     string `json:"content" bson:"content"`
	HaveSee     bool   `json:"have_see" bson:"have_see"` // true status=P , false status=F

	Title   string `json:"title" bson:"title"`
	Summary string `json:"summary" bson:"summary"`
}

type PhotoResponse struct {
	SubjectId      string `json:"subject_id" bson:"subject_id"`
	ThumbnailImage string `json:"thumbnail_image" bson:"thumbnailImage"`
	LargeImage     string `json:"large_image" bson:"large_image"`
	RawImage       string `json:"raw_image" bson:"raw_image"`
	RawSize        struct {
		Width  int `json:"width" bson:"width"`
		Height int `json:"height" bson:"height"`
	} `json:"raw_size" bson:"raw_size"`
}

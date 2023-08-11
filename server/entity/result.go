package entity

type Result struct {
	Uid    int64  `json:"uid" validate:"required,number"`
	Wid    int    `json:"wid" validate:"required,number"`
	Title  string `json:"title"  validate:"required,alphanumunicode|ascii"`
	URL    string `json:"url" validate:"required,url"`
	Phrase string `json:"phrase" validate:"required,alphanumunicode|ascii"`
}

type SearchingHistory struct {
	Title  string `redis:"title"`
	URL    string `redis:"url"`
	Phrase string `redis:"phrase"`
}

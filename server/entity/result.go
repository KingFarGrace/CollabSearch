package entity

// SearchingJSON is a json entity to receive and pass user searching data.
type SearchingJSON struct {
	Uid    int64  `json:"uid" validate:"required,number"`
	Wid    int    `json:"wid" validate:"required,number"`
	Phrase string `json:"phrase" validate:"required,alphanumunicode|ascii"`
}

// Result is a json entity to receive and pass user searching result.
type Result struct {
	Uid    int64  `json:"uid" validate:"required,number"`
	Wid    int    `json:"wid" validate:"required,number"`
	Title  string `json:"title"  validate:"required,alphanumunicode|ascii"`
	URL    string `json:"url" validate:"required,url"`
	Phrase string `json:"phrase" validate:"required,alphanumunicode|ascii"`
}

// SearchingHistory ORM Model -- hashset 'result:<wid>:<uid>'
type SearchingHistory struct {
	Title  string `redis:"title"`
	URL    string `redis:"url"`
	Phrase string `redis:"phrase"`
}

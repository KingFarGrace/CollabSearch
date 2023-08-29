package entity

// SearchingJSON is a json entity to receive and pass user searching data.
type SearchingJSON struct {
	Wid    int    `json:"wid" validate:"required,number"`
	Phrase string `json:"phrase" validate:"required,alphanumunicode|ascii"`
}

type ResultIndex struct {
	Wid int    `json:"wid" validate:"required,number"`
	URL string `json:"url" validate:"required,url"`
}

// Result is a json entity to receive and pass user searching result.
type Result struct {
	ResultIndex
	Phrase string `json:"phrase" validate:"required,alphanumunicode|ascii"`
}

// SearchingHistory ORM Model -- hashset 'result:<wid>:<uid>'
type SearchingHistory struct {
	URL    string `redis:"url"`
	Phrase string `redis:"phrase"`
}

type Note struct {
	ResultIndex
	Content  string `json:"content" redis:"content" validate:"omitempty,alphanumunicode|ascii,max=255"`
	Feedback int8   `json:"feedback" redis:"feedback" validate:"omitempty,required,number,gte=0,lte=5"`
}

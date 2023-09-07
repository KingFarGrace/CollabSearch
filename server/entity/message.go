package entity

type Message struct {
	Sender   int64  `json:"sender" validate:"omitempty,number"`
	Receiver int64  `json:"receiver" validate:"required,number"`
	Content  string `json:"content" validate:"required,min=1,max=255"`
	Read     bool   `json:"read" validate:"omitempty,boolean"`
}

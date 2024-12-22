package kafka

type eventBody struct {
	Filename string `json:"filename"`
	UserId   int32  `json:"userId"`
}

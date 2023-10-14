package errs

type MessageObject struct {
	Message string `json:"message"`
}

func NewMessageObject(msg string) MessageObject {
	return MessageObject{msg}
}

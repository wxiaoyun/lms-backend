package api

// Code 0 = silent
// Code 1 = Success
// Code 2 = Error
// Code 3 = Warning
// Code 4 = Info
type Message struct {
	Code string `json:"code"`
	Msg  string `json:"message"`
}

func SilentMessage(message string) Message {
	return Message{
		Code: "0",
		Msg:  message,
	}
}

func SuccessMessage(message string) Message {
	return Message{
		Code: "1",
		Msg:  message,
	}
}

func ErrorMessage(message string) Message {
	return Message{
		Code: "2",
		Msg:  message,
	}
}

func WarningMessage(message string) Message {
	return Message{
		Code: "3",
		Msg:  message,
	}
}

func InfoMessage(message string) Message {
	return Message{
		Code: "4",
		Msg:  message,
	}
}

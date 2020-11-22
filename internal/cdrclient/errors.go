package cdrclient

type Error string

func (e Error) Error() string { return string(e) }

const (
	ErrRecordNotExist = Error("record file is not exist")
)

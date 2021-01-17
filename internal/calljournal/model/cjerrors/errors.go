package cjerrors

type Error string

func (e Error) Error() string { return string(e) }

const (
	ErrEmpty = Error("empty helper error")
)

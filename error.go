package ursa

import "strconv"

type Err struct {
	Status int
	Msg    string
}

func (e Err) Error() string {
	return strconv.Itoa(e.Status) + " " + e.Msg
}

// NewError creates a new Ursa error with status code and message
func NewError(status int, msg string) Err {
	return Err{Status: status, Msg: msg}
}

// NewNFError is deprecated, use NewError instead
// Kept for backward compatibility
func NewNFError(status int, msg string) Err {
	return NewError(status, msg)
}

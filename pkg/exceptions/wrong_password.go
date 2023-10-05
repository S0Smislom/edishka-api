package exceptions

type WrongPasswordError struct {
	Msg string
}

func (r *WrongPasswordError) Error() string {
	if r.Msg == "" {
		r.Msg = "Wrong password"
	}
	return r.Msg
}

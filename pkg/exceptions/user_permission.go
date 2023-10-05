package exceptions

type UserPermissionError struct {
	Msg string
}

func (r *UserPermissionError) Error() string {
	if r.Msg == "" {
		r.Msg = "Forbidden"
	}
	return r.Msg
}

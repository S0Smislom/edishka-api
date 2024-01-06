package exceptions

type UnauthorizedError struct {
	Msg string
}

func (r *UnauthorizedError) Error() string {
	if r.Msg == "" {
		r.Msg = "Unauthorized"
	}
	return r.Msg
}

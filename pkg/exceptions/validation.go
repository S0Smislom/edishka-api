package exceptions

type ValidationError struct {
	Msg string
}

func (r *ValidationError) Error() string {
	if r.Msg == "" {
		r.Msg = "Validation error"
	}
	return r.Msg
}

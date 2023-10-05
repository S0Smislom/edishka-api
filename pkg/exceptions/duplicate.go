package exceptions

type DuplicateError struct {
	Msg string
}

func (r *DuplicateError) Error() string {
	if r.Msg == "" {
		r.Msg = "Duplicate"
	}
	return r.Msg
}

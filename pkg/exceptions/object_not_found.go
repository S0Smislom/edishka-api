package exceptions

type ObjectNotFoundError struct {
	Msg string
}

func (r *ObjectNotFoundError) Error() string {
	if r.Msg == "" {
		r.Msg = "Object not found"
	}
	return r.Msg
}

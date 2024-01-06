package exceptions

type LogicError struct {
	Msg string
}

func (r *LogicError) Error() string {
	if r.Msg == "" {
		r.Msg = "Logic error"
	}
	return r.Msg
}

package exceptions

import "encoding/json"

type ValidationError struct {
	Err error
}

func (r *ValidationError) Error() string {
	b, _ := json.Marshal(r.Err)
	return string(b)
}

func (r *ValidationError) ErrorMap() []byte {
	b, _ := json.Marshal(r.Err)
	return b
}

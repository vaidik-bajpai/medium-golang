package validator

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddErrors(field string, errorMessage string) {
	if _, exists := v.Errors[field]; !exists {
		v.Errors[field] = errorMessage
	}
}

func (v *Validator) Check(ok bool, field string, errorMessage string) {
	if !ok {
		v.Errors[field] = errorMessage
	}
}

package validator

type Validator struct {
	Errors map[string]string
}

// New returns a new instance of Validator
func New() *Validator { return &Validator{Errors: make(map[string]string)} }

// Valid returns true if no errors exists
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// AddError adds a new error if the key does not exist
func (v *Validator) AddError(key, msg string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = msg
	}
}

// Check adds a new error if ok is false
func (v *Validator) Check(ok bool, key, msg string) {
	if !ok {
		v.AddError(key, msg)
	}
}

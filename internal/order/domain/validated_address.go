package domain

type ValidatedAddress struct {
	Address
	isValidated bool
}

func (v *ValidatedAddress) IsValid() bool {
	return v.isValidated
}

func NewValidatedAddress(address *Address) (*ValidatedAddress, error) {
	if err := address.validate(); err != nil {
		return nil, err
	}
	return &ValidatedAddress{Address: *address, isValidated: true}, nil
}

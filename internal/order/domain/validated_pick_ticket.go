package domain

type ValidatedPickTicket struct {
	PickTicket
	isValidated bool
}

func (v *ValidatedPickTicket) IsValid() bool {
	return v.isValidated
}

func NewValidatedPickTicket(pickTicket *PickTicket) (*ValidatedPickTicket, error) {
	if err := pickTicket.validate(); err != nil {
		return nil, err
	}
	return &ValidatedPickTicket{PickTicket: *pickTicket, isValidated: true}, nil
}

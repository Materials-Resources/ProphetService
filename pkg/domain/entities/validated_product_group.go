package entities

type ValidatedProductGroup struct {
	ProductGroup
	isValidated bool
}

func (v *ValidatedProductGroup) IsValid() bool {
	return v.isValidated
}

func NewValidatedProductGroup(productGroup *ProductGroup) (*ValidatedProductGroup, error) {
	if err := productGroup.validate(); err != nil {
		return nil, err
	}
	return &ValidatedProductGroup{ProductGroup: *productGroup, isValidated: true}, nil
}

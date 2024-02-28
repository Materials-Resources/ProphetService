package domain

type ValidatedProductSupplier struct {
	ProductSupplier
	isValidated bool
}

func NewValidatedProductSupplier(p *ProductSupplier) (*ValidatedProductSupplier, error) {
	return &ValidatedProductSupplier{
		ProductSupplier: *p,
		isValidated:     true,
	}, nil
}

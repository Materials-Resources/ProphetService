package domain

type ProductSupplier struct {
	ProductId  string
	SupplierId string
	SupplierSn string
	IsPrimary  bool
}

func (p *ProductSupplier) validate() error {
	return nil
}

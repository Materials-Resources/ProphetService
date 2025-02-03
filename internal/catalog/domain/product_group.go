package domain

type ProductGroup struct {
	Create bool

	Sn   string  `avro:"sn" validate:"required,lte=8"`
	Name *string `avro:"name" validate:"required_if=Create true,omitempty,gte=4,lte=80"`
}

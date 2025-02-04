package domain

type Product struct {
	Create           bool
	Id               string  `avro:"id"`
	Sn               string  `avro:"sn"`
	Name             *string `avro:"name" validate:"required_if=Create true,gte=4,lte=40"`
	Description      *string `avro:"description" validate:"omitempty,lte=255"`
	ProductGroupSn   *string `avro:"product_group_sn"`
	ProductGroupName string
}

package avro

import _ "embed"

var (
	//go:embed product.avsc
	ProductAvro string
	//go:embed group.avsc
	GroupAvro string
)

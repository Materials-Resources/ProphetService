package repository

type ReceivingRepository interface {
	ReadReceivingReceipt() error
}

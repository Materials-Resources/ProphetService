package repositories

type CustomerRepository interface {
	FindCustomerById(id string)
}

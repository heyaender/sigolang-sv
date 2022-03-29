package models

import (
	"sigolang-sv/databases"
	"sigolang-sv/schemas"

	"gorm.io/gorm"
)

func CreateCustomer(customer schemas.Customer) *gorm.DB {
	result := databases.DB.Create(&customer)
	return result
}

func GetAllCustomer(limit int, offset int) []schemas.Customer {
	customers := []schemas.Customer{}
	databases.DB.
		Select("id, name, age").
		Limit(limit).
		Offset(offset).
		Find(&customers)
	return customers
}

func GetCustomerByID(CustomerID string) (schemas.Customer, *gorm.DB) {
	var customer schemas.Customer
	result := databases.DB.
		Select("customers.id, customers.name, customers.age, Profile.*").
		Joins("Profile").
		Where("customers.id = ?", CustomerID).
		First(&customer)
	return customer, result
}
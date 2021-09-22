package database

import (
	"myproperty-api/config"
	model "myproperty-api/models"
)

func GetTransactions() []model.Transaction {
	var transactions []model.Transaction
	config.DB.Find(&transactions)
	return transactions
}

func GetTransactionByID(id string) model.Transaction {
	var transaction model.Transaction
	config.DB.Where("id = ?", id).Find(&transaction)
	return transaction
}

func CreateTransaction(transaction model.Transaction) model.Transaction {
	config.DB.Create(&transaction)
	return transaction
}

func DeleteTransactionByID(id string) {
	var transaction model.Transaction
	config.DB.Where("id = ?", id).Delete(&transaction)
}

func UpdateTransactionByID(id string, transaction model.Transaction) {
	config.DB.Where("id = ?", id).Updates(&transaction)
}

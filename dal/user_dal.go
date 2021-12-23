package dal

import (
	"advanced_programming/clients"
	"advanced_programming/models"
)

func CreateUser(app *models.User) {
	var DB = clients.DB
	DB.Create(app)
}

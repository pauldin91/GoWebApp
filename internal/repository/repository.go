package repository

import "github.com/pauldin91/GoWebApp/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	Insert(res models.Reservation) error
}

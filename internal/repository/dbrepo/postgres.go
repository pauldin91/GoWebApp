package dbrepo

import (
	"context"
	"time"

	"github.com/pauldin91/GoWebApp/internal/models"
)

func (m *postgresDbRepo) AllUsers() bool {
	return true
}

func (m *postgresDbRepo) Insert(res models.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	stmt := `insert into reservations (first_name, last_name,phone, start_date,end_date,room_id,created_at,updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8,$9)`

	defer cancel()

	_, err := m.DB.ExecContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartedDate,
		res.EndDate,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}
	return nil
}

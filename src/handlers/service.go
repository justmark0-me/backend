package handlers

import (
	"github.com/jackc/pgx/v5"
	"main/src/internal/app"
)

// Implementation implements
type Implementation struct {
	saveRequestUseCase SaveRequestUseCase
}

// NewClient return new instance of Implementation.
func NewClient(db *pgx.Conn) *Implementation {
	return &Implementation{
		saveRequestUseCase: app.NewSaveRequestUseCase(db),
	}
}

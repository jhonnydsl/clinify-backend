package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jhonnydsl/clinify-backend/src/dtos"
	"github.com/jhonnydsl/clinify-backend/src/utils"
)

type PatientRepository struct{}

func (r *PatientRepository) CreatePatient(ctx context.Context, patient dtos.PatientInput, birthDate time.Time, clientID uuid.UUID) (uuid.UUID, error) {
	query := `INSERT INTO patients (full_name, email, password_hash, phone, birth_date, client_id)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id`

	var id uuid.UUID

	err := DB.QueryRowContext(
		ctx,
		query,
		patient.FullName,
		patient.Email,
		patient.Password,
		patient.Phone,
		birthDate,
		clientID,
	).Scan(&id)
	if err != nil {
		utils.LogError("PatientRepository (erro ao criar user paciente)", err)
		return uuid.UUID{}, utils.InternalServerError("error creating user patient")
	}

	return id, nil
}
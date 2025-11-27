package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jhonnydsl/clinify-backend/src/dtos"
	"github.com/jhonnydsl/clinify-backend/src/utils"
)

type AdminRepository struct{}

func (r *AdminRepository) CreateAdmin(ctx context.Context, admin dtos.AdminInput, birthDate time.Time) (uuid.UUID, error) {
	query := `INSERT INTO clients (full_name, email, password_hash, birth_date, crp, bio, profile_image_url, office_address, phone, public_slug)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING id`

	var id uuid.UUID

	err := DB.QueryRowContext(
		ctx,
		query, 
		admin.FullName, 
		admin.Email, 
		admin.Password, 
		birthDate,
		admin.Crp,
		admin.Bio,
		admin.ProfileImage,
		admin.OfficeAddress,
		admin.Phone,
		admin.PublicSlug,
	).Scan(&id)
	if err != nil {
		utils.LogError("createAdmin (INSERT clients)", err)
		return uuid.UUID{}, utils.InternalServerError("error creating user admin")
	}

	return id, nil
}
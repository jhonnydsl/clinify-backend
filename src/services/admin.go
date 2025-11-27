package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/jhonnydsl/clinify-backend/src/dtos"
	"github.com/jhonnydsl/clinify-backend/src/repository"
	"github.com/jhonnydsl/clinify-backend/src/utils"
)

type AdminService struct {
	Repo *repository.AdminRepository
}

func (services *AdminService) CreateAdmin(ctx context.Context, admin dtos.AdminInput) (uuid.UUID, error) {
	if err := utils.ValidateAdminInput(admin); err != nil {
		utils.LogError("CreatingAdmin service (error validating admin input)", err)
		return uuid.UUID{}, utils.BadRequestError(err.Error())
	}

	hashedPassword, err := utils.HashPassword(admin.Password)
	if err != nil {
		utils.LogError("HashPassword (error to hash password)", err)
		return uuid.UUID{}, utils.InternalServerError("error creating user admin")
	}
	
	admin.Password = hashedPassword

	id, err := services.Repo.CreateAdmin(ctx, admin)
	if err != nil {
		utils.LogError("CreateAdmin service (error to call createAdmin repository)", err)
		return uuid.UUID{}, utils.InternalServerError("error create user admin")
	}

	return id, nil
}
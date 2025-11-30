package services

import (
	"context"

	"github.com/jhonnydsl/clinify-backend/src/dtos"
	"github.com/jhonnydsl/clinify-backend/src/repository"
	"github.com/jhonnydsl/clinify-backend/src/utils"
)

type LoginService struct {
	Repo *repository.LoginRepository
}

func (service *LoginService) LoginUser(ctx context.Context, email, password string) (dtos.LoginOutput, error) {
	admin, err := service.Repo.GetAdminByEmail(ctx, email)
	if err == nil {
		if err := utils.CheckPassword(admin.PasswordHash, password); err != nil {
			return dtos.LoginOutput{}, utils.BadRequestError("email or password incorrect")
		}

		token, err := utils.GenerateJWT(admin.ID.String(), admin.FullName, admin.Email, "admin")
		if err != nil {
			utils.LogError("LoginUser service (error generating token)", err)
			return dtos.LoginOutput{}, utils.InternalServerError("failed authentication")
		}

		return dtos.LoginOutput{
			ID: admin.ID.String(),
			FullName: admin.FullName,
			Email: admin.Email,
			Role: "admin",
			Token: token,
		}, nil
	}

	patient, err := service.Repo.GetPatientByEmail(ctx, email)
	if err == nil {
		if err := utils.CheckPassword(patient.PasswordHash, password); err != nil {
			return dtos.LoginOutput{}, utils.BadRequestError("email or password incorrect")
		}

		token, err := utils.GenerateJWT(patient.ID.String(), patient.FullName, patient.Email, "patient")
		if err != nil {
			utils.LogError("loginUser service (error generating token)", err)
			return dtos.LoginOutput{}, utils.InternalServerError("failed authentication")
		}

		return dtos.LoginOutput{
			ID: patient.ID.String(),
			FullName: patient.FullName,
			Email: patient.Email,
			Role: "patient",
			Token: token,
		}, nil
	}

	return dtos.LoginOutput{}, utils.BadRequestError("email or password incorrect")
}
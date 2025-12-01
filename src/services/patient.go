package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jhonnydsl/clinify-backend/src/dtos"
	"github.com/jhonnydsl/clinify-backend/src/repository"
	"github.com/jhonnydsl/clinify-backend/src/utils"
)

type PatientService struct {
	Repo *repository.PatientRepository
	AdminRepo *repository.AdminRepository
}

func (service *PatientService) CreatePatient(ctx context.Context, patient dtos.PatientInput) (uuid.UUID, error) {
	if err := utils.ValidatePatientInput(patient); err != nil {
		utils.LogError("createPatient service (error validating patient input)", err)
		return uuid.UUID{}, utils.BadRequestError(err.Error())
	}

	clientUUID, err := service.AdminRepo.FindAdminIDBySlug(ctx, patient.PublicSlug)
	if err != nil {
		utils.LogError("createPatient service (error get client_id by public_slug)", err)
		return uuid.UUID{}, utils.BadRequestError("invalid admin url")
	}

	hashedPassword, err := utils.HashPassword(patient.Password)
	if err != nil {
		utils.LogError("hashPassword (error to hash password)", err)
		return uuid.UUID{}, utils.InternalServerError("error creating user patient")
	}

	patient.Password = hashedPassword

	parsedDate, err := time.Parse("2006-01-02", patient.BirthDate)
	if err != nil {
		utils.LogError("createAdmin service (error parsing birth date)", err)
		return uuid.UUID{}, utils.BadRequestError("invalid birth date format, expected YYYY-MM-DD")
	}

	id, err := service.Repo.CreatePatient(ctx, patient, parsedDate, clientUUID)
	if err != nil {
		utils.LogError("createPatient service (error to call createPatient repository)", err)
		return uuid.UUID{}, utils.InternalServerError("error creating user patient")
	}

	return id, nil
}
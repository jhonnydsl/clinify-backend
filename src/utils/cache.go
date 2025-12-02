package utils

import (
	"time"

	"github.com/jhonnydsl/clinify-backend/src/dtos"
	"github.com/patrickmn/go-cache"
)

type PatientsCache struct {
	Data []dtos.PatientOutput
	Total int
}

var Cache = cache.New(30*time.Second, 1*time.Minute)
package address

import (
	"github.com/merembablas/catat-pendatang/models"
)

// Usecase represent contract address
type Usecase interface {
	Provinces() []*models.Province
	Province(ID string) *models.Province
}

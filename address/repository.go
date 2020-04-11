package address

import (
	"github.com/merembablas/catat-pendatang/models"
)

// Repository represent address contract
type Repository interface {
	GetProvinces() (res []*models.Province, err error)
	GetProvince(ID string) (res *models.Province)
}

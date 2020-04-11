package usecase

import (
	"github.com/merembablas/catat-pendatang/address"
	"github.com/merembablas/catat-pendatang/models"
)

type addressUsecase struct {
	addressRepo address.Repository
}

// NewAddressUsecase create address usecase object
func NewAddressUsecase(addressRepo address.Repository) address.Usecase {
	return &addressUsecase{addressRepo}
}

func (addr *addressUsecase) Provinces() []*models.Province {
	provinces, _ := addr.addressRepo.GetProvinces()

	return provinces
}

func (addr *addressUsecase) Province(ID string) *models.Province {
	province := addr.addressRepo.GetProvince(ID)

	return province
}

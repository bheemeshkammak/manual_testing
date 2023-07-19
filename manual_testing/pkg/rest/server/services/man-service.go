package services

import (
	"github.com/bheemeshkammak/manual_testing/manual_testing/pkg/rest/server/daos"
	"github.com/bheemeshkammak/manual_testing/manual_testing/pkg/rest/server/models"
)

type ManService struct {
	manDao *daos.ManDao
}

func NewManService() (*ManService, error) {
	manDao, err := daos.NewManDao()
	if err != nil {
		return nil, err
	}
	return &ManService{
		manDao: manDao,
	}, nil
}

func (manService *ManService) CreateMan(man *models.Man) (*models.Man, error) {
	return manService.manDao.CreateMan(man)
}

func (manService *ManService) UpdateMan(id int64, man *models.Man) (*models.Man, error) {
	return manService.manDao.UpdateMan(id, man)
}

func (manService *ManService) DeleteMan(id int64) error {
	return manService.manDao.DeleteMan(id)
}

func (manService *ManService) ListMen() ([]*models.Man, error) {
	return manService.manDao.ListMen()
}

func (manService *ManService) GetMan(id int64) (*models.Man, error) {
	return manService.manDao.GetMan(id)
}

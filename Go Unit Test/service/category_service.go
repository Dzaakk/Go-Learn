package service

//3
import (
	"errors"
	"go_unit_test/entity"
	"go_unit_test/repository"
)

type CategoryService struct {
	Repository repository.CateogryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}

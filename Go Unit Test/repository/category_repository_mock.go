package repository

//4
import (
	"go_unit_test/entity"

	"github.com/stretchr/testify/mock"
)

type CateogryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CateogryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}

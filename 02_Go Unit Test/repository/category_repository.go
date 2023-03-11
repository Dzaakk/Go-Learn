package repository

//2
import "go_unit_test/entity"

type CateogryRepository interface {
	FindById(id string) *entity.Category
}

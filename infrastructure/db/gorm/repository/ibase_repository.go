package repository

type IBaseRepository[T any] interface {
	FindAll() ([]*T, error)
	FindByID(id uint) (*T, error)
	Create(entity *T) (*T, error)
	Update(entity *T) (*T, error)
	Delete(id uint) error
}

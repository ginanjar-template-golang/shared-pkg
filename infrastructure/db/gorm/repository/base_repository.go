package repository

import (
	"fmt"

	"github.com/ginanjar-template-golang/shared-pkg/errors"
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	DB        *gorm.DB
	TableName string
}

func NewBaseRepository[T any](db *gorm.DB, tableName string) *BaseRepository[T] {
	return &BaseRepository[T]{DB: db, TableName: tableName}
}

func (r *BaseRepository[T]) FindAll() ([]*T, error) {
	var results []*T
	err := r.DB.Find(&results).Error
	if err != nil {
		return nil, errors.GeneralError(fmt.Sprintf("FindAll %s", r.TableName), err)
	}
	return results, nil
}

func (r *BaseRepository[T]) FindByID(id uint) (*T, error) {
	var result T
	err := r.DB.First(&result, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ResourceNotFound(fmt.Sprintf("%s id=%d", r.TableName, id), nil)
		}
		return nil, errors.GeneralError(fmt.Sprintf("FindByID %s", r.TableName), err)
	}
	return &result, nil
}

func (r *BaseRepository[T]) Create(entity *T) (*T, error) {
	err := r.DB.Create(entity).Error
	if err != nil {
		return nil, errors.CreateResourceError(r.TableName, err)
	}
	return entity, nil
}

func (r *BaseRepository[T]) Update(entity *T) (*T, error) {
	err := r.DB.Save(entity).Error
	if err != nil {
		return nil, errors.UpdateResourceError(r.TableName, err)
	}
	return entity, nil
}

func (r *BaseRepository[T]) Delete(id uint) error {
	err := r.DB.Delete(new(T), id).Error
	if err != nil {
		return errors.DeleteResourceError(fmt.Sprintf("%s id=%d", r.TableName, id), err)
	}
	return nil
}

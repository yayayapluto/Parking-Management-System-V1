package repos

import (
	"context"
	"gorm.io/gorm"
	"parking-management-system-v1/internal/dto/requests"
)

type BaseRepository[T any] interface {
	Create(ctx context.Context, entity *T) error
	GetByID(ctx context.Context, id uint) (*T, error)
	GetAll(ctx context.Context) ([]T, error)
	GetWithPagination(ctx context.Context, p requests.PaginationRequest, searchableColumns ...string) ([]T, int64, error)
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id uint) error
	WithTx(tx *gorm.DB) BaseRepository[T]
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{db: db}
}

func (r *baseRepository[T]) WithTx(tx *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{db: tx}
}

func (r *baseRepository[T]) Create(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Create(entity).Error
}

func (r *baseRepository[T]) GetByID(ctx context.Context, id uint) (*T, error) {
	var entity T
	if err := r.db.WithContext(ctx).First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *baseRepository[T]) GetAll(ctx context.Context) ([]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Find(&entities).Error
	return entities, err
}

func (r *baseRepository[T]) GetWithPagination(ctx context.Context, p requests.PaginationRequest, searchableColumns ...string) ([]T, int64, error) {
	var entities []T
	var total int64

	db := r.db.WithContext(ctx).Model(new(T))

	// 1. Handle Global Search (OR condition)
	if p.Search != "" && len(searchableColumns) > 0 {
		searchQuery := ""
		searchValues := []interface{}{}
		for i, col := range searchableColumns {
			searchQuery += col + " LIKE ?"
			searchValues = append(searchValues, "%"+p.Search+"%")
			if i < len(searchableColumns)-1 {
				searchQuery += " OR "
			}
		}
		db = db.Where(searchQuery, searchValues...)
	}

	// 2. Handle Specific Filters (AND condition)
	if len(p.Filters) > 0 {
		for column, value := range p.Filters {
			if value != "" {
				db = db.Where(column+" = ?", value)
			}
		}
	}

	// 3. Count Total after filters
	db.Count(&total)

	// 4. Sort, Limit, Offset
	if p.Sort != "" {
		db = db.Order(p.Sort)
	} else {
		db = db.Order("created_at DESC")
	}

	err := db.Limit(p.GetLimit()).Offset(p.GetOffset()).Find(&entities).Error
	return entities, total, err
}

func (r *baseRepository[T]) Update(ctx context.Context, entity *T) error {
	// Save akan mengupdate semua field berdasarkan Primary Key
	return r.db.WithContext(ctx).Save(entity).Error
}

func (r *baseRepository[T]) Delete(ctx context.Context, id uint) error {
	var entity T
	return r.db.WithContext(ctx).Delete(&entity, id).Error
}

package dbs

import (
	"context"
	"time"

	"gorm.io/gorm"
)

const (
	DatabaseDeadLine = time.Second * 60
)

type IDatabase interface {
	Create(ctx context.Context, doc any) error
	CreateInBatches(ctx context.Context, docs any, batchSize int) error
	Find(ctx context.Context, results any, opts ...fnOption) error
	FindByID(ctx context.Context, id string, doc any) error
	FindOne(ctx context.Context, result any, opts ...fnOption) error
	Save(ctx context.Context, doc any) error
	Delete(ctx context.Context, value any, opts ...fnOption) error
	Migrate(models ...any) error
}

type Database struct {
	DB *gorm.DB
}

type Query struct {
	query string
	args  []any
}

func (d *Database) Create(ctx context.Context, doc any) error {
	_, cancel := context.WithTimeout(ctx, DatabaseDeadLine)
	defer cancel()

	return d.DB.Create(doc).Error
}

func (d *Database) CreateInBatches(ctx context.Context, docs any, batchSize int) error {
	_, cancel := context.WithTimeout(ctx, DatabaseDeadLine)
	defer cancel()

	return d.DB.CreateInBatches(docs, batchSize).Error
}

func (d *Database) FindByID(ctx context.Context, id string, result any) error {
	_, cancel := context.WithTimeout(ctx, DatabaseDeadLine)
	defer cancel()

	if err := d.DB.Where("id = ?", id).First(result).Error; err != nil {
		return err
	}

	return nil
}

func (d *Database) Save(ctx context.Context, doc any) error {
	_, cancel := context.WithTimeout(ctx, DatabaseDeadLine)
	defer cancel()

	return d.DB.Save(doc).Error
}

func (d *Database) Migrate(models ...any) error {
	return d.DB.AutoMigrate(models...)
}

func (d *Database) Find(ctx context.Context, results any, opts ...fnOption) error {
	_, cancel := context.WithTimeout(ctx, DatabaseDeadLine)
	defer cancel()

	query := d.applyOptions(opts...)
	if err := query.Find(results).Error; err != nil {
		return err
	}

	return nil
}

func (d *Database) FindOne(ctx context.Context, result any, opts ...fnOption) error {
	_, cancel := context.WithTimeout(ctx, DatabaseDeadLine)
	defer cancel()

	query := d.applyOptions(opts...)
	if err := query.First(result).Error; err != nil {
		return err
	}

	return nil
}

func (d *Database) Delete(ctx context.Context, value any, opts ...fnOption) error {
	_, cancel := context.WithTimeout(ctx, DatabaseDeadLine)
	defer cancel()

	query := d.applyOptions(opts...)
	if err := query.Delete(value).Error; err != nil {
		return err
	}
	return nil
}

func BuildQuery(query string, args ...any) []Query {
	return []Query{
		{
			query: query,
			args:  args,
		},
	}
}

func (d *Database) applyOptions(opts ...fnOption) *gorm.DB {
	query := d.DB

	opt := NewOption(opts...)
	if opt.query != nil {
		for _, q := range opt.query {
			query = query.Where(q.query, q.args...)
		}
	}

	if opt.Order != nil {
		query = query.Order(opt.Order)
	}

	if opt.limit != 0 {
		query = query.Limit(opt.limit)
	}

	return query
}

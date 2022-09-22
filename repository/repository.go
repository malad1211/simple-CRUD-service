package repository

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"inspiredlab/domain"
	appError "inspiredlab/domain/error"
	"log"
)

type Repository[M domain.Model, E Entity] struct {
	Parser Parser[M, E]
	db     *gorm.DB
}

func New[M domain.Model, E Entity](db *gorm.DB) Repository[M, E] {
	e := new(E)
	err := db.AutoMigrate(e)
	if err != nil {
		panic(err)
	}

	repo := Repository[M, E]{
		db: db,
	}

	repo.Parser = DefaultParser[M, E]{}
	return repo
}

func (r Repository[M, E]) WithParser(parser Parser[M, E]) Repository[M, E] {
	r.Parser = parser
	return r
}

func (r Repository[M, E]) Create(model *M) (*M, error) {
	entity := new(E)
	r.Parser.CopyToEntity(model, entity)
	rs := r.db.Create(entity)
	if e := r.AppError(rs.Error); e != nil {
		return nil, e
	}
	r.Parser.CopyToDomain(entity, model)
	return model, nil
}

func (r Repository[M, E]) FindOne(model *M) (*M, error) {
	entity := new(E)
	r.Parser.CopyToEntity(model, entity)
	tx := r.db.Where(entity)
	rs := tx.Last(entity)
	if e := r.AppError(rs.Error); e != nil {
		return nil, e
	}
	r.Parser.CopyToDomain(entity, model)
	return model, nil
}

func (r Repository[M, E]) Save(model *M) (*M, error) {
	entity := new(E)
	r.Parser.CopyToEntity(model, entity)
	rs := r.db.Model(entity).Save(entity)
	if e := r.AppError(rs.Error); e != nil {
		return nil, e
	}
	r.Parser.CopyToDomain(entity, model)
	return model, nil
}

func (r Repository[M, E]) Delete(model *M) error {
	entity := new(E)
	r.Parser.CopyToEntity(model, entity)
	tx := r.db.Where(entity).Delete(entity)
	r.Parser.CopyToDomain(entity, model)
	return tx.Error
}

func (r Repository[M, E]) Update(where *M, update *M) (*M, error) {
	entityWhere := new(E)
	r.Parser.CopyToEntity(where, entityWhere)
	entityUpdate := new(E)
	r.Parser.CopyToEntity(update, entityUpdate)
	rs := r.db.Where(entityWhere).Updates(entityUpdate)
	if rs.RowsAffected == 0 {
		return update, appError.ErrNoChanged
	}
	r.Parser.CopyToDomain(entityWhere, where)
	r.Parser.CopyToDomain(entityUpdate, update)
	return update, r.AppError(rs.Error)
}

func (r Repository[M, E]) AppError(dbErr error) error {
	switch dbErr {
	case nil:
		return nil
	case gorm.ErrRecordNotFound:
		return appError.ErrDbRecordNotExisted
	default:
		var mysqlErr *mysql.MySQLError
		if errors.As(dbErr, &mysqlErr) && mysqlErr.Number == 1062 {
			return appError.ErrDbDuplicateEntry
		}
		log.Println(dbErr)
		return appError.ErrDb
	}
}

func (r Repository[M, E]) DB() *gorm.DB {
	return r.db.Model(new(E))
}

func (r Repository[M, E]) GetParser() Parser[M, E] {
	return r.Parser
}

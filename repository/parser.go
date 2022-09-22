package repository

import (
	"github.com/jinzhu/copier"
	"inspiredlab/domain"
)

type Parser[M domain.Model, E Entity] interface {
	CopyToDomain(fromValue *E, toValue *M)
	CopyToEntity(fromValue *M, toValue *E)
}

type DefaultParser[M domain.Model, E Entity] struct {
}

func (DefaultParser[M, E]) CopyToDomain(fromValue *E, toValue *M) {
	cpErr := copier.Copy(toValue, fromValue)
	if cpErr != nil {
		panic(cpErr)
	}
}

func (DefaultParser[M, E]) CopyToEntity(fromValue *M, toValue *E) {
	cpErr := copier.Copy(toValue, fromValue)
	if cpErr != nil {
		panic(cpErr)
	}
}

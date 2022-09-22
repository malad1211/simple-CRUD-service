package appError

import "errors"

var (
	ErrDbRecordNotExisted = errors.New("record does not existed")
	ErrNoChanged          = errors.New("no record has been changed")
	ErrDbDuplicateEntry   = errors.New("duplicate entry")
	ErrDb                 = errors.New("database error")
)

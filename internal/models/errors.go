package models

import "errors"

type StorageQueryResult error

var (
	ErrorNotExists StorageQueryResult = errors.New("not exists")
)

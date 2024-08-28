package Errors

import (
	"errors"
	"gorm.io/gorm"
)

func ErrorDbWrapperUtils(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return NoRecordsFound
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return Conflict
	}
	if errors.Is(err, gorm.ErrInvalidDB) {
		return IncorrectDb
	}
	if errors.Is(err, gorm.ErrInvalidField) {
		return InvalidFiled
	}
	if errors.Is(err, gorm.ErrInvalidTransaction) {
		return InvalidTransaction
	}
	return Unknown
}

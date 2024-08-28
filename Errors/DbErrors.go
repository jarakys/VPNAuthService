package Errors

import "errors"

var (
	NoRecordsFound     = errors.New("No records found")
	Conflict           = errors.New("Conflict")
	Unknown            = errors.New("Unknown")
	IncorrectDb        = errors.New("IncorrectDb")
	InvalidFiled       = errors.New("InvalidFiled")
	InvalidTransaction = errors.New("InvalidTransaction")
)

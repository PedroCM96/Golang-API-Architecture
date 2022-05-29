package errors

import (
	"Duna/database/models"
	"fmt"
)

type CreateDummyError struct {
	Err   error
	Dummy models.Dummy
}

func (e CreateDummyError) Error() string {
	return fmt.Sprintf("Error during dummy creation. Error obtained: %d. Dummy: %#v", e.Err, e.Dummy)
}

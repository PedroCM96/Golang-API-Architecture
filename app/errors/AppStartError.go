package errors

type AppStartError struct {
	Err error
}

func (e AppStartError) Error() string {
	return e.Err.Error()
}

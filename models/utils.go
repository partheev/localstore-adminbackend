package models

type DbError struct {
	Message string
}

func (err DbError) Error() string {
	return err.Message
}

func ThrowDbError(message string) DbError {
	err := DbError{
		Message: message,
	}
	return err
}

package exception

type UnauthorizedError struct{
	Error string
}

func NewUnauthorizedError(error string) UnauthorizedError{
	return UnauthorizedError{Error: error}
}


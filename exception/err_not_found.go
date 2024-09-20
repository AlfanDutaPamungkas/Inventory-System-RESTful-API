package exception

type NotFoundErr struct{
	Error string
}

func NewNotFoundErr(error string) NotFoundErr {
	return NotFoundErr{
		Error: error,
	}
}


package exception

type ForbiddenErr struct{
	Error string
}

func NewForbiddenErr(error string) ForbiddenErr {
	return ForbiddenErr{
		Error: error,
	}
}


package exception

type BadReqErr struct{
	Error string
}

func NewBadReqErr(error string) BadReqErr {
	return BadReqErr{
		Error: error,
	}
}


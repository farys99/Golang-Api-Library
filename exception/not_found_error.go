package exception

type NotFoundErr struct {
	Error string
}

func NotFoundError(error string) NotFoundErr {
	return NotFoundErr{Error: error}
}

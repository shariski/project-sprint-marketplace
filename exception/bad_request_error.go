package exception

type BadRequestError struct {
	Message string
}

func (BadRequestError BadRequestError) Error() string {
	return BadRequestError.Message
}

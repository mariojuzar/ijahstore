package exception

type notFoundException struct {

}

func (notFoundException) Error() string {
	return "Data not found"
}

func NewNotFoundException() error {
	return notFoundException{}
}

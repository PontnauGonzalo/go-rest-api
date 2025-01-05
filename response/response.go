package response

type Response interface {
	StatusCode() int
	GetBody() ([]byte, error)
	Error() string
	GetData() interface{}
}
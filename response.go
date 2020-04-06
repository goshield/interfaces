package interfaces

// Response is a HTTP response
type Response interface {
	Status() int
	WithStatus(status int) Response

	Header(key string) string
	WithHeader(key string, value string) Response

	Headers() map[string]string
	WithHeaders(headers map[string]string) Response

	Body() interface{}
	WithBody(body interface{}) Response
}

package x

// ErrorResponse for http response body
type ErrorResponse struct {
	Message string
	Code    string
	Details interface{}
}

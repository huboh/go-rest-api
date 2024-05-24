package main

// Status is the response status.
type Status string

// recognized response Status
const (
	StatusError   = Status("error")
	StatusSuccess = Status("success")
)

// Response represent the api response object.
type Response struct {
	// Status is the response status. either "success" or "error".
	Status Status `json:"status"`

	// Message is the response message. defaults to http.StatusText(Code).
	Message string `json:"message"`

	// StatusCode is the http response status code.
	StatusCode int `json:"statusCode"`

	// Data is the response data.
	Data any `json:"data,omitempty"`

	// Error is the response error. this field is omitted from the response if it is nil.
	Error *ResponseError `json:"error,omitempty"`
}

// Response represent the api error response object.
type ResponseError struct {
	// Name is the name of the error.
	Name string `json:"name,omitempty"`

	// Cause is the error's cause.
	Cause string `json:"cause,omitempty"`

	// Stack is the error stack trace.
	Stack []string `json:"stack,omitempty"`

	// Message is the error message
	Message string `json:"message,omitempty"`
}

func NewJsonResponseError(msg string, name string, cause string, stack []string) *ResponseError {
	return &ResponseError{
		Message: msg,
		Cause:   cause,
		Stack:   stack,
		Name:    name,
	}
}

// Error makes ResponseError meets the error interface
func (re *ResponseError) Error() string {
	return re.Message
}

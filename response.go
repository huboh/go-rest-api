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
	// Code is the http response status code.
	Code int `json:"code"`

	// Status is the response status. either "success" or "error".
	Status Status `json:"status"`

	// Message is the response message. defaults to http.StatusText(Code).
	Message string `json:"message"`

	// Data is the response data.
	Data any `json:"data"`

	// Error is the response error. this field is omitted from the response if it is nil.
	Error *ResponseError `json:"error,omitempty"`
}

// Response represent the api error response object.
type ResponseError struct {
	// Name is the name of the error.
	Name string `json:"name"`

	// Cause is the error's cause.
	Cause string `json:"cause"`

	// Stack is the error stack trace.
	Stack []string `json:"stack"`

	// Message is the error message
	Message string `json:"message"`
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

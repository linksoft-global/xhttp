package xhttp

type Err string

func (e Err) Error() string {
	return string(e)
}

const (
	ErrInvalidStatusCode   = Err("Invalid Status Code")
	ErrFailedToMarshalJSON = Err("Failed to marshal JSON")
	ErrWritingResponse     = Err("Error writing to http.ResponseWriter")
	ErrDataOmitted         = Err("Data payload is nil")
)

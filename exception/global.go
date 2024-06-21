package exception

import "fmt"

func ErrNotFound(message string) *ApiException {
	return NewApiException(404, message)
}

func ErrServerInternal(format string, a ...any) *ApiException {
	return NewApiException(500000, fmt.Sprintf(format, a...))
}

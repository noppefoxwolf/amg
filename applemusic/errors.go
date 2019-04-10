package applemusic

import "fmt"

type APIError struct {
	Errors []ErrorDetail `json:"errors"`
}

type ErrorDetail struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e APIError) Error() string {
	if len(e.Errors) > 0 {
		err := e.Errors[0]
		return fmt.Sprintf("twitter: %d %v", err.Code, err.Message)
	}
	return ""
}

func (e APIError) Empty() bool {
	if len(e.Errors) == 0 {
		return true
	}
	return false
}

func relevantError(httpError error, apiError APIError) error {
	if httpError != nil {
		return httpError
	}
	if apiError.Empty() {
		return nil
	}
	return apiError
}

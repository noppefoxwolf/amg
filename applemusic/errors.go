package applemusic

import (
	"fmt"
)

type APIError struct {
	Errors []ErrorDetail `json:"errors"`
}

type ErrorDetail struct {
	Code   string
	Detail string
	Id     string
	Error  ErrorSource
	Status string
	Title  string
}

func (e APIError) Error() string {
	if len(e.Errors) > 0 {
		err := e.Errors[0]
		return fmt.Sprintf("applemusic: %d %v", err.Code, err.Title)
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

package api

import "github.com/noppefoxwolf/amg/applemusic/models"

func relevantError(httpError error, apiError models.APIError) error {
	if httpError != nil {
		return httpError
	}
	if apiError.Empty() {
		return nil
	}
	return apiError
}

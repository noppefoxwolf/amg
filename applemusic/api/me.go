package api

import (
	"net/http"

	"github.com/noppefoxwolf/amg/applemusic/models"

	"github.com/dghubble/sling"
)

type MeService struct {
	sling *sling.Sling
}

func newMeService(sling *sling.Sling) *MeService {
	return &MeService{
		sling: sling.Path("me/"),
	}
}

type GetMeStorefrontParams struct {
	L string `url:"l,omitempty"`
}

func (s *MeService) GetStorefront(params *GetMeStorefrontParams) (*models.StorefrontResponse, *http.Response, error) {
	storefronts := new(models.StorefrontResponse)
	apiError := new(models.APIError)
	resp, err := s.sling.New().Get("storefront").QueryStruct(params).Receive(storefronts, apiError)
	return storefronts, resp, relevantError(err, *apiError)
}

type GetMeLibraryRecentryAddedParams struct {
	Limit  int    `url:"limit,omitempty"`
	Offset string `url:"offset,omitempty"`
}

func (s *MeService) GetLibraryRecentryAdded(params *GetMeLibraryRecentryAddedParams) (*models.VariableResponse, *http.Response, error) {
	response := new(models.VariableResponse)
	apiError := new(models.APIError)
	resp, err := s.sling.New().Get("library/recently-added").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMeHeavyRotationContentParams struct {
	L      string `url:"l,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Offset int    `url:"offset,omitempty"`
}

func (s *MeService) GetMeHeavyRotationContent(params *GetMeHeavyRotationContentParams) (*models.ResponseRoot, *http.Response, error) {
	albums := new(models.ResponseRoot)
	apiError := new(models.APIError)
	resp, err := s.sling.New().Get("history/heavy-rotation").QueryStruct(params).Receive(albums, apiError)
	return albums, resp, relevantError(err, *apiError)
}

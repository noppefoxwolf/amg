package api

import (
	"net/http"

	"github.com/noppefoxwolf/amg/applemusic/models"

	"github.com/dghubble/sling"
)

type StorefrontsService struct {
	sling *sling.Sling
}

// newTrendsService returns a new TrendsService.
func newStorefrontsService(sling *sling.Sling) *StorefrontsService {
	return &StorefrontsService{
		sling: sling.Path("storefronts/"),
	}
}

type GetAllStorefrontsParams struct {
	L      string `url:"l,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Offset string `url:"offset,omitempty"`
}

func (s *StorefrontsService) GetAll(params *GetAllStorefrontsParams) (*models.StorefrontResponse, *http.Response, error) {
	storefronts := new(models.StorefrontResponse)
	apiError := new(models.APIError)
	resp, err := s.sling.New().Get("").QueryStruct(params).Receive(storefronts, apiError)
	return storefronts, resp, relevantError(err, *apiError)
}

type GetMultipleStorefrontsParams struct {
	L   string   `url:"l,omitempty"`
	Ids []string `url:"ids,comma"`
}

func (s *StorefrontsService) GetMultipleStorefronts(params *GetMultipleStorefrontsParams) (*models.StorefrontResponse, *http.Response, error) {
	storefronts := new(models.StorefrontResponse)
	apiError := new(models.APIError)
	resp, err := s.sling.New().Get("").QueryStruct(params).Receive(storefronts, apiError)
	return storefronts, resp, relevantError(err, *apiError)
}

type GetStorefrontParams struct {
	Id string
}

func (s *StorefrontsService) GetStorefront(params *GetStorefrontParams) (*models.StorefrontResponse, *http.Response, error) {
	storefronts := new(models.StorefrontResponse)
	apiError := new(models.APIError)
	resp, err := s.sling.New().Get(params.Id).Receive(storefronts, apiError)
	return storefronts, resp, relevantError(err, *apiError)
}

package applemusic

import (
	"github.com/dghubble/sling"
	"net/http"
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

type Storefronts struct {
	Data []Storefront
}

type Storefront struct {
	Id string
	Type string
	Href string
	Attributes StorefrontAttribute
}

type StorefrontAttribute struct {
	DefaultLanguageTag string
	Name string
	SupportedLanguageTags []string
	ExplicitContentPolicy string
}

type GetAllStorefrontsParams struct {
	L string `url:"l,omitempty"`
	Limit int `url:"limit,omitempty"`
	Offset string `url:"offset,omitempty"`
}

func (s *StorefrontsService)GetAll(params *GetAllStorefrontsParams) (*Storefronts, *http.Response, error) {
	storefronts := new(Storefronts)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("").QueryStruct(params).Receive(storefronts, apiError)
	return storefronts, resp, relevantError(err, *apiError)
}

type GetMultipleStorefrontsParams struct {
	L string `url:"l,omitempty"`
	Ids []string `url:"ids,comma"`
}

func (s *StorefrontsService)GetMultipleStorefronts(params *GetMultipleStorefrontsParams) (*Storefronts, *http.Response, error) {
	storefronts := new(Storefronts)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("").QueryStruct(params).Receive(storefronts, apiError)
	return storefronts, resp, relevantError(err, *apiError)
}

type GetStorefrontParams struct {
	Id string
}

func (s *StorefrontsService)GetStorefront(params *GetStorefrontParams)  (*Storefronts, *http.Response, error) {
	storefronts := new(Storefronts)
	apiError := new(APIError)
	resp, err := s.sling.New().Get(params.Id).QueryStruct(params).Receive(storefronts, apiError)
	return storefronts, resp, relevantError(err, *apiError)
}


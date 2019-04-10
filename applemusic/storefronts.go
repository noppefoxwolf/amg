package applemusic

import (
	"net/http"

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

type StorefrontResponse struct {
	ResponseRoot
	Data []Storefront
}

type Storefront struct {
	Resource
	Attributes StorefrontAttributes //The attributes for the storefront.
	Type       string               //(Required) This value will always be storefronts. Value: storefronts
}

type StorefrontAttributes struct {
	DefaultLanguageTag    string   //(Required) The default language for the storefront, represented as a language tag.
	Name                  string   //(Required) The localized name of the storefront.
	SupportedLanguageTags []string // (Required) The localizations that the storefront supports, represented as an array of language tags.
	ExplicitContentPolicy string   //Undocumented
}

type GetAllStorefrontsParams struct {
	L      string `url:"l,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Offset string `url:"offset,omitempty"`
}

func (s *StorefrontsService) GetAll(params *GetAllStorefrontsParams) (*StorefrontResponse, *http.Response, error) {
	storefronts := new(StorefrontResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("").QueryStruct(params).Receive(storefronts, apiError)
	return storefronts, resp, relevantError(err, *apiError)
}

type GetMultipleStorefrontsParams struct {
	L   string   `url:"l,omitempty"`
	Ids []string `url:"ids,comma"`
}

func (s *StorefrontsService) GetMultipleStorefronts(params *GetMultipleStorefrontsParams) (*StorefrontResponse, *http.Response, error) {
	storefronts := new(StorefrontResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("").QueryStruct(params).Receive(storefronts, apiError)
	return storefronts, resp, relevantError(err, *apiError)
}

type GetStorefrontParams struct {
	Id string
}

func (s *StorefrontsService) GetStorefront(params *GetStorefrontParams) (*StorefrontResponse, *http.Response, error) {
	storefronts := new(StorefrontResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get(params.Id).QueryStruct(params).Receive(storefronts, apiError)
	return storefronts, resp, relevantError(err, *apiError)
}

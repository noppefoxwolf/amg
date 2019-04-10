package applemusic

import (
	"github.com/dghubble/sling"
	"net/http"
)

type MeService struct {
	sling *sling.Sling
}

// newTrendsService returns a new TrendsService.
func newMeService(sling *sling.Sling) *MeService {
	return &MeService{
		sling: sling.Path("me/"),
	}
}

type GetMeStorefrontParams struct {
	L string `url:"l,omitempty"`
}

func (s *MeService)GetStorefront(params *GetMeStorefrontParams)  (*Storefronts, *http.Response, error) {
	storefronts := new(Storefronts)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("storefront").QueryStruct(params).Receive(storefronts, apiError)
	return storefronts, resp, relevantError(err, *apiError)
}
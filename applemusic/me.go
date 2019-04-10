package applemusic

import (
	"net/http"

	"github.com/dghubble/sling"
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

func (s *MeService) GetStorefront(params *GetMeStorefrontParams) (*StorefrontResponse, *http.Response, error) {
	storefronts := new(StorefrontResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("storefront").QueryStruct(params).Receive(storefronts, apiError)
	return storefronts, resp, relevantError(err, *apiError)
}

type GetMeLibraryRecentryAddedParams struct {
	Limit  int    `url:"limit,omitempty"`
	Offset string `url:"offset,omitempty"`
}

func (s *MeService) GetLibraryRecentryAdded(params *GetMeLibraryRecentryAddedParams) (*LibraryAlbums, *http.Response, error) {
	albums := new(LibraryAlbums)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("library/recently-added").QueryStruct(params).Receive(albums, apiError)
	return albums, resp, relevantError(err, *apiError)
}

type GetMeHeavyRotationContentParams struct {
	L      string `url:"l,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Offset int    `url:"offset,omitempty"`
}

func (s *MeService) GetMeHeavyRotationContent(params *GetMeHeavyRotationContentParams) (*LibraryAlbums, *http.Response, error) {
	albums := new(LibraryAlbums)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("history/heavy-rotation").QueryStruct(params).Receive(albums, apiError)
	return albums, resp, relevantError(err, *apiError)
}

type LibraryAlbums struct {
	Data []Resource
	Href string
	Next string
}

type AlbumAttribute struct {
	ArtistName string
	Artwork    Artwork
	Name       string
	PlayParams PlayParams
	TrackCount int
}

type PlayParams struct {
	Id        string
	IsLibrary bool
	Kind      string
}

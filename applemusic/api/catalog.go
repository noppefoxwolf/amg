package api

import (
	"net/http"

	"github.com/dghubble/sling"
	"github.com/noppefoxwolf/amg/applemusic/models"
)

type CatalogService struct {
	sling *sling.Sling
}

func newCatalogService(sling *sling.Sling) *CatalogService {
	return &CatalogService{
		sling: sling.Path("catalog/"),
	}
}

type GetAlbumParams struct {
	Storefront string   `url:"-"`
	Id         string   `url:"-"`
	L          string   `url:"l"`
	Include    []string `url:"include"`
}

func (s *CatalogService) GetAlbum(params *GetAlbumParams) (*models.VariableResponse, *http.Response, error) {
	response := new(models.VariableResponse)
	apiError := new(models.APIError)
	resp, err := s.sling.New().Get(params.Storefront+"/albums/"+params.Id).QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetAlbumRelationshipParams struct {
	Id           string `url:"-"`
	Relationship string `url:"-"`
	Storefront   string `url:"-"`
	L            string `url:"l"`
	Limit        int    `url:"limit"`
}

func (s *CatalogService) GetAlbumRelationship(params *GetAlbumRelationshipParams) (*models.VariableResponse, *http.Response, error) {
	response := new(models.VariableResponse)
	apiError := new(models.APIError)
	resp, err := s.sling.New().Get(params.Storefront+"/albums/"+params.Id+"/"+params.Relationship).QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultipleAlbumsParams struct {
	Storefront string   `url:"-"`
	Ids        []string `url:"ids,comma"`
	Include    []string `url:"include,omitempty"`
	L          string   `url:"l,omitempty"`
}

func (s *CatalogService) GetMultipleAlbums(params *GetMultipleAlbumsParams) (*models.VariableResponse, *http.Response, error) {
	response := new(models.VariableResponse)
	apiError := new(models.APIError)
	resp, err := s.sling.New().Get(params.Storefront+"/albums").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

//Get a Library Album
////https://api.music.apple.com/v1/me/library/albums/{id}
//id
//string
//l
//string
//include
//[string]

//Get a Library Album's Relationship Directly by Name
//https://api.music.apple.com/v1/me/library/albums/{id}/{relationship}
//id
//string
//relationship
//string
//limit
//number
//offset
//string

//Get Multiple Library Albums
//https://api.music.apple.com/v1/me/library/albums
//ids
//[string]
//l
//string
//include
//[string]

//Get All Library Albums
//https://api.music.apple.com/v1/me/library/albums
//include
//[string]
//l
//string
//limit
//number
//offset
//string

//Get a Catalog Artist
//https://api.music.apple.com/v1/catalog/{storefront}/artists/{id}
//id
//string
//storefront
//string
//l
//string
//include
//[string]

//Get a Catalog Artist's Relationship Directly by Name
//https://api.music.apple.com/v1/catalog/{storefront}/artists/{id}/{relationship}
//id
//string
//relationship
//string
//storefront
//string
//l
//string
//limit
//number

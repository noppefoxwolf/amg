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

func (s *CatalogService) GetAlbum(params *GetAlbumParams) (*models.AlbumResponse, *http.Response, error) {
	response := new(models.AlbumResponse)
	apiError := new(models.APIError)
	resp, err := s.sling.New().Get(params.Storefront+"/albums/"+params.Id).QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

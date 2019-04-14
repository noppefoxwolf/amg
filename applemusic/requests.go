package applemusic

import (
	"net/http"

	"github.com/dghubble/sling"
)

type StorefrontsAndLocalizationService struct {
	sling *sling.Sling
}

func newStorefrontsAndLocalizationService(sling *sling.Sling) *StorefrontsAndLocalizationService {
	return &StorefrontsAndLocalizationService{
		sling: sling.Path("/"),
	}
}

type GetAUsersStorefrontParams struct {
	L string `url:"l,omitempty"` // The localization to use, specified by a language tag. Any supported language tag may be used here. If none is specified, en-US is used.

}

func (s *StorefrontsAndLocalizationService) GetAUsersStorefront(params *GetAUsersStorefrontParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/storefront").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetAStorefrontParams struct {
	Id string `url:"-"`           // The identifier (an ISO 3166 alpha-2 country code) for the storefront you want to fetch.
	L  string `url:"l,omitempty"` // The localization to use, specified by a language tag. Any supported language tag may be used here. If none is specified, en-US is used.

}

func (s *StorefrontsAndLocalizationService) GetAStorefront(params *GetAStorefrontParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/storefronts/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultipleStorefrontsParams struct {
	Ids []string `url:"ids"`         // A list of the identifiers (ISO 3166 alpha-2 country codes) for the storefronts you want to fetch.
	L   string   `url:"l,omitempty"` // The localization to use, specified by a language tag. Any supported language tag may be used here. If none is specified, en-US is used.

}

func (s *StorefrontsAndLocalizationService) GetMultipleStorefronts(params *GetMultipleStorefrontsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/storefronts").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetAllStorefrontsParams struct {
	L      string `url:"l,omitempty"`      // The localization to use, specified by a language tag. Any supported language tag may be used here. If none is specified, en-US is used.
	Limit  int    `url:"limit,omitempty"`  // The limit on the number of objects, or number of objects in the specified relationship, that are returned.
	Offset string `url:"offset,omitempty"` // The next page or group of objects to fetch. See Fetch Resources by Page.

}

func (s *StorefrontsAndLocalizationService) GetAllStorefronts(params *GetAllStorefrontsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/storefronts").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type CommonObjectsService struct {
	sling *sling.Sling
}

func newCommonObjectsService(sling *sling.Sling) *CommonObjectsService {
	return &CommonObjectsService{
		sling: sling.Path("/"),
	}
}

type AlbumsService struct {
	sling *sling.Sling
}

func newAlbumsService(sling *sling.Sling) *AlbumsService {
	return &AlbumsService{
		sling: sling.Path("/"),
	}
}

type GetACatalogAlbumParams struct {
	Id         string   `url:"-"`                 // The unique identifier for the album.
	Storefront string   `url:"-"`                 // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	L          string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include    []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *AlbumsService) GetACatalogAlbum(params *GetACatalogAlbumParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/albums/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetACatalogAlbumsRelationshipDirectlyByNameParams struct {
	Id           string `url:"-"`               // The unique identifier for the album.
	Relationship string `url:"-"`               // The name of the relationship you want to fetch for this resource.
	Storefront   string `url:"-"`               // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	L            string `url:"l,omitempty"`     // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit        int    `url:"limit,omitempty"` // The limit on the number of objects, or number of objects in the specified relationship, that are returned.

}

func (s *AlbumsService) GetACatalogAlbumsRelationshipDirectlyByName(params *GetACatalogAlbumsRelationshipDirectlyByNameParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/albums/"+params.Id+"/"+params.Relationship+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultipleCatalogAlbumsParams struct {
	Storefront string   `url:"-"`                 // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	Ids        []string `url:"ids"`               // The unique identifiers for the albums. The maximum fetch limit is 100.
	Include    []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L          string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *AlbumsService) GetMultipleCatalogAlbums(params *GetMultipleCatalogAlbumsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/albums").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetALibraryAlbumParams struct {
	Id      string   `url:"-"`                 // The unique identifier for the album.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *AlbumsService) GetALibraryAlbum(params *GetALibraryAlbumParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/albums/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetALibraryAlbumsRelationshipDirectlyByNameParams struct {
	Id           string `url:"-"`                // A unique identifier for the library album.
	Relationship string `url:"-"`                // The name of the relationship you want to fetch for this resource.
	Limit        int    `url:"limit,omitempty"`  // The limit on the number of objects, or number of objects in the specified relationship, that are returned.
	Offset       string `url:"offset,omitempty"` // The next page or group of objects to fetch.

}

func (s *AlbumsService) GetALibraryAlbumsRelationshipDirectlyByName(params *GetALibraryAlbumsRelationshipDirectlyByNameParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/albums/"+params.Id+"/"+params.Relationship+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultipleLibraryAlbumsParams struct {
	Ids     []string `url:"ids"`               // The unique identifiers for the library albums. The maximum fetch limit is 100.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *AlbumsService) GetMultipleLibraryAlbums(params *GetMultipleLibraryAlbumsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/albums").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetAllLibraryAlbumsParams struct {
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit   int      `url:"limit,omitempty"`   // The limit on the number of objects, or number of objects in the specified relationship, that are returned. The default value is 25 and the maximum value is 100.
	Offset  string   `url:"offset,omitempty"`  // The next page or group of objects to fetch.

}

func (s *AlbumsService) GetAllLibraryAlbums(params *GetAllLibraryAlbumsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/albums").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

//type AddAResourceToALibraryParams struct {
//	Ids []string `url:"ids"` // The unique catalog identifiers for the resources. To indicate the type of resource to be added, ids must be followed by one of the allowed values. Multiple types can be added in the same request.
//
//}

func (s *AlbumsService) AddAResourceToALibrary(params *AddAResourceToALibraryParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("/v1/me/library").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type ArtistsService struct {
	sling *sling.Sling
}

func newArtistsService(sling *sling.Sling) *ArtistsService {
	return &ArtistsService{
		sling: sling.Path("/"),
	}
}

type GetACatalogArtistParams struct {
	Id         string   `url:"-"`                 // The unique identifier for the artist.
	Storefront string   `url:"-"`                 // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	L          string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include    []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *ArtistsService) GetACatalogArtist(params *GetACatalogArtistParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/artists/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultipleCatalogArtistsParams struct {
	Storefront string   `url:"-"`                 // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	Ids        []string `url:"ids"`               // The unique identifiers for the artists. The maximum fetch limit is 25.
	L          string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include    []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *ArtistsService) GetMultipleCatalogArtists(params *GetMultipleCatalogArtistsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/artists").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetACatalogArtistsRelationshipDirectlyByNameParams struct {
	Id           string `url:"-"`               // A unique identifier for the artist.
	Relationship string `url:"-"`               // The name of the relationship you want to fetch for this resource.
	Storefront   string `url:"-"`               // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	L            string `url:"l,omitempty"`     // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit        int    `url:"limit,omitempty"` // The limit on the number of objects, or number of objects in the specified relationship, that are returned.

}

func (s *ArtistsService) GetACatalogArtistsRelationshipDirectlyByName(params *GetACatalogArtistsRelationshipDirectlyByNameParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/artists/"+params.Id+"/"+params.Relationship+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetALibraryArtistParams struct {
	Id      string   `url:"-"`                 // The unique identifier for the artist.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *ArtistsService) GetALibraryArtist(params *GetALibraryArtistParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/artists/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetAllLibraryArtistsParams struct {
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit   int      `url:"limit,omitempty"`   // The limit on the number of objects, or number of objects in the specified relationship, that are returned. The default value is 25 and the maximum value is 100.
	Offset  string   `url:"offset,omitempty"`  // The next page or group of objects to fetch.

}

func (s *ArtistsService) GetAllLibraryArtists(params *GetAllLibraryArtistsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/artists").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultipleLibraryArtistsParams struct {
	Ids     []string `url:"ids"`               // The unique identifiers for the artists. The maximum fetch limit is 25.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *ArtistsService) GetMultipleLibraryArtists(params *GetMultipleLibraryArtistsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/artists").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetALibraryArtistsRelationshipDirectlyByNameParams struct {
	Id           string `url:"-"`               // A unique identifier for the library artist.
	Relationship string `url:"-"`               // The name of the relationship you want to fetch for this resource.
	Limit        int    `url:"limit,omitempty"` // The limit on the number of objects, or number of objects in the specified relationship, that are returned.

}

func (s *ArtistsService) GetALibraryArtistsRelationshipDirectlyByName(params *GetALibraryArtistsRelationshipDirectlyByNameParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/artists/"+params.Id+"/"+params.Relationship+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type SongsService struct {
	sling *sling.Sling
}

func newSongsService(sling *sling.Sling) *SongsService {
	return &SongsService{
		sling: sling.Path("/"),
	}
}

type GetACatalogSongParams struct {
	Id         string   `url:"-"`                 // The unique identifier for the song.
	Storefront string   `url:"-"`                 // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	L          string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include    []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *SongsService) GetACatalogSong(params *GetACatalogSongParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/songs/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultipleCatalogSongsByIdParams struct {
	Storefront string   `url:"-"`                 // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	Ids        []string `url:"ids"`               // The unique identifiers for the songs. You can substitute filter[isrc] for ids, or use it in conjunction with ids for additional filtering. The maximum fetch limit is 300.
	L          string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include    []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *SongsService) GetMultipleCatalogSongsById(params *GetMultipleCatalogSongsByIdParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/songs").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultipleCatalogSongsByIsrcParams struct {
	Storefront string   `url:"-"`                 // An Apple Music territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	Filter     []string `url:"filter"`            // The ISRC values for the songs. You can substitute filter[isrc] for ids, or use it in conjunction with ids for additional filtering. Note that one ISRC value may return more than one song. The maximum fetch limit is 25.
	L          string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include    []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *SongsService) GetMultipleCatalogSongsByIsrc(params *GetMultipleCatalogSongsByIsrcParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/songs").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetACatalogSongsRelationshipDirectlyByNameParams struct {
	Id           string `url:"-"`               // A unique identifier for the song.
	Relationship string `url:"-"`               // The name of the relationship you want to fetch for this resource.
	Storefront   string `url:"-"`               // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	L            string `url:"l,omitempty"`     // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit        int    `url:"limit,omitempty"` // The limit on the number of objects, or number of objects in the specified relationship, that are returned.

}

func (s *SongsService) GetACatalogSongsRelationshipDirectlyByName(params *GetACatalogSongsRelationshipDirectlyByNameParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/songs/"+params.Id+"/"+params.Relationship+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetALibrarySongParams struct {
	Id      string   `url:"-"`                 // The unique identifier for the library song.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *SongsService) GetALibrarySong(params *GetALibrarySongParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/songs/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetAllLibrarySongsParams struct {
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit   int      `url:"limit,omitempty"`   // The limit on the number of objects, or number of objects in the specified relationship, that are returned. The default value is 25 and the maximum value is 100.
	Offset  string   `url:"offset,omitempty"`  // The next page or group of objects to fetch.

}

func (s *SongsService) GetAllLibrarySongs(params *GetAllLibrarySongsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/songs").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultipleLibrarySongsParams struct {
	Ids     []string `url:"ids"`               // The unique identifiers for the library songs. The maximum fetch limit is 300. For possible values, get all the library songs by sending this endpoint without the ids parameter.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *SongsService) GetMultipleLibrarySongs(params *GetMultipleLibrarySongsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/songs").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetALibrarySongsRelationshipDirectlyByNameParams struct {
	Id           string `url:"-"`               // A unique identifier for the library song.
	Relationship string `url:"-"`               // The name of the relationship you want to fetch for this resource.
	Limit        int    `url:"limit,omitempty"` // The limit on the number of objects, or number of objects in the specified relationship, that are returned.

}

func (s *SongsService) GetALibrarySongsRelationshipDirectlyByName(params *GetALibrarySongsRelationshipDirectlyByNameParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/songs/"+params.Id+"/"+params.Relationship+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type MusicVideosService struct {
	sling *sling.Sling
}

func newMusicVideosService(sling *sling.Sling) *MusicVideosService {
	return &MusicVideosService{
		sling: sling.Path("/"),
	}
}

type GetACatalogMusicVideoParams struct {
	Id         string   `url:"-"`                 // The unique identifier for the music video.
	Storefront string   `url:"-"`                 // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	L          string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include    []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *MusicVideosService) GetACatalogMusicVideo(params *GetACatalogMusicVideoParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/music-videos/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetACatalogMusicVideosRelationshipDirectlyByNameParams struct {
	Id           string `url:"-"`               // The unique identifier for the music video.
	Relationship string `url:"-"`               // The name of the relationship you want to fetch for this resource.
	Storefront   string `url:"-"`               // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	L            string `url:"l,omitempty"`     // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit        int    `url:"limit,omitempty"` // The limit on the number of objects, or number of objects in the specified relationship, that are returned.

}

func (s *MusicVideosService) GetACatalogMusicVideosRelationshipDirectlyByName(params *GetACatalogMusicVideosRelationshipDirectlyByNameParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/music-videos/"+params.Id+"/"+params.Relationship+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultipleCatalogMusicVideosByIdParams struct {
	Storefront string   `url:"-"`                 // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	Ids        []string `url:"ids"`               // The unique identifiers for the music videos. Optionally, you can substitute filter[isrc] for ids, or use it in conjunction with ids for additional filtering. The maximum fetch limit is 100.
	L          string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include    []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *MusicVideosService) GetMultipleCatalogMusicVideosById(params *GetMultipleCatalogMusicVideosByIdParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/music-videos").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultipleCatalogMusicVideosByIsrcParams struct {
	Storefront string   `url:"-"`                 // An Apple Music territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	Filter     []string `url:"filter"`            // The ISRC values for the music videos. You can substitute filter[isrc] for ids, or use it in conjunction with ids for additional filtering. Note that one ISRC value may return more than one music video. The maximum fetch limit is 25.
	L          string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include    []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *MusicVideosService) GetMultipleCatalogMusicVideosByIsrc(params *GetMultipleCatalogMusicVideosByIsrcParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/music-videos").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetALibraryMusicVideoParams struct {
	Id      string   `url:"-"`                 // The unique identifier for the library music video.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *MusicVideosService) GetALibraryMusicVideo(params *GetALibraryMusicVideoParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/music-videos/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetALibraryMusicVideosRelationshipDirectlyByNameParams struct {
	Id           string `url:"-"`               // The unique identifier for the library music video.
	Relationship string `url:"-"`               // The name of the relationship you want to fetch for this resource.
	Limit        int    `url:"limit,omitempty"` // The limit on the number of objects, or number of objects in the specified relationship, that are returned.

}

func (s *MusicVideosService) GetALibraryMusicVideosRelationshipDirectlyByName(params *GetALibraryMusicVideosRelationshipDirectlyByNameParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/music-videos/"+params.Id+"/"+params.Relationship+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultipleLibraryMusicVideosParams struct {
	Ids     []string `url:"ids"`               // The unique identifiers for the library music videos. The maximum fetch limit is 100. For possible values, get all the library music videos by sending this endpoint without the ids parameter.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *MusicVideosService) GetMultipleLibraryMusicVideos(params *GetMultipleLibraryMusicVideosParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/music-videos").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetAllLibraryMusicVideosParams struct {
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit   int      `url:"limit,omitempty"`   // The limit on the number of objects, or number of objects in the specified relationship, that are returned. The default value is 25 and the maximum value is 100.
	Offset  string   `url:"offset,omitempty"`  // The next page or group of objects to fetch.

}

func (s *MusicVideosService) GetAllLibraryMusicVideos(params *GetAllLibraryMusicVideosParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/music-videos").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type PlaylistsService struct {
	sling *sling.Sling
}

func newPlaylistsService(sling *sling.Sling) *PlaylistsService {
	return &PlaylistsService{
		sling: sling.Path("/"),
	}
}

type GetACatalogPlaylistParams struct {
	Id         string   `url:"-"`                 // The unique identifier for the playlist.
	Storefront string   `url:"-"`                 // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	L          string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include    []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *PlaylistsService) GetACatalogPlaylist(params *GetACatalogPlaylistParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/playlists/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetACatalogPlaylistsRelationshipDirectlyByNameParams struct {
	Id           string `url:"-"`               // A unique identifier for the playlist.
	Relationship string `url:"-"`               // The name of the relationship you want to fetch for this resource.
	Storefront   string `url:"-"`               // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	L            string `url:"l,omitempty"`     // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit        int    `url:"limit,omitempty"` // The limit on the number of objects, or number of objects in the specified relationship, that are returned.

}

func (s *PlaylistsService) GetACatalogPlaylistsRelationshipDirectlyByName(params *GetACatalogPlaylistsRelationshipDirectlyByNameParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/playlists/"+params.Id+"/"+params.Relationship+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultipleCatalogPlaylistsParams struct {
	Storefront string   `url:"-"`                 // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	Ids        []string `url:"ids"`               // The unique identifiers for the playlists. The maximum fetch limit is 25.
	L          string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include    []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *PlaylistsService) GetMultipleCatalogPlaylists(params *GetMultipleCatalogPlaylistsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/playlists").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetALibraryPlaylistParams struct {
	Id      string   `url:"-"`                 // A unique identifier for the library playlist.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *PlaylistsService) GetALibraryPlaylist(params *GetALibraryPlaylistParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/playlists/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetALibraryPlaylistsRelationshipDirectlyByNameParams struct {
	Id           string `url:"-"`               // The unique identifier for the library playlist.
	Relationship string `url:"-"`               // The name of the relationship you want to fetch for this resource.
	Limit        int    `url:"limit,omitempty"` // The limit on the number of objects, or number of objects in the specified relationship, that are returned.

}

func (s *PlaylistsService) GetALibraryPlaylistsRelationshipDirectlyByName(params *GetALibraryPlaylistsRelationshipDirectlyByNameParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/playlists/"+params.Id+"/"+params.Relationship+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultipleLibraryPlaylistsParams struct {
	Ids     []string `url:"ids"`               // The unique identifiers for the playlists. The maximum fetch limit is 25.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *PlaylistsService) GetMultipleLibraryPlaylists(params *GetMultipleLibraryPlaylistsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/playlists").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetAllLibraryPlaylistsParams struct {
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit   int      `url:"limit,omitempty"`   // The limit on the number of objects, or number of objects in the specified relationship, that are returned. The default value is 25 and the maximum value is 100.
	Offset  string   `url:"offset,omitempty"`  // The next page or group of objects to fetch.

}

func (s *PlaylistsService) GetAllLibraryPlaylists(params *GetAllLibraryPlaylistsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/playlists").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type CreateANewLibraryPlaylistParams struct {
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *PlaylistsService) CreateANewLibraryPlaylist(params *CreateANewLibraryPlaylistParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("/v1/me/library/playlists").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type AddTracksToALibraryPlaylistParams struct {
	Id string `url:"url"`
}

func (s *PlaylistsService) AddTracksToALibraryPlaylist(params *AddTracksToALibraryPlaylistParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("/v1/me/library/playlists/"+params.Id+"/tracks").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type AddAResourceToALibraryParams struct {
	Ids []string `url:"ids"` // The unique catalog identifiers for the resources. To indicate the type of resource to be added, ids must be followed by one of the allowed values. Multiple types can be added in the same request.

}

func (s *PlaylistsService) AddAResourceToALibrary(params *AddAResourceToALibraryParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("/v1/me/library").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type AppleMusicStationsService struct {
	sling *sling.Sling
}

func newAppleMusicStationsService(sling *sling.Sling) *AppleMusicStationsService {
	return &AppleMusicStationsService{
		sling: sling.Path("/"),
	}
}

type GetACatalogStationParams struct {
	Id         string   `url:"-"`                 // The unique identifier for the station.
	Storefront string   `url:"-"`                 // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	Include    []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L          string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *AppleMusicStationsService) GetACatalogStation(params *GetACatalogStationParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/stations/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultipleCatalogStationsParams struct {
	Storefront string   `url:"-"`                 // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	Ids        []string `url:"ids"`               // The unique identifiers for the stations. The maximum fetch limit is 100.
	L          string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include    []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *AppleMusicStationsService) GetMultipleCatalogStations(params *GetMultipleCatalogStationsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/stations").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type SearchService struct {
	sling *sling.Sling
}

func newSearchService(sling *sling.Sling) *SearchService {
	return &SearchService{
		sling: sling.Path("/"),
	}
}

type SearchForCatalogResourcesParams struct {
	Storefront string   `url:"-"`                // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	Term       string   `url:"term,omitempty"`   // The entered text for the search with ‘+’ characters between each word, to replace spaces (for example term=james+br).
	L          string   `url:"l,omitempty"`      // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit      int      `url:"limit,omitempty"`  // The limit on the number of objects, or number of objects in the specified relationship, that are returned.
	Offset     string   `url:"offset,omitempty"` // The next page or group of objects to fetch.
	Types      []string `url:"types,omitempty"`  // The list of the types of resources to include in the results.

}

func (s *SearchService) SearchForCatalogResources(params *SearchForCatalogResourcesParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/search").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetCatalogSearchHintsParams struct {
	Storefront string   `url:"-"`               // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	Term       string   `url:"term,omitempty"`  // The entered text for the search with ‘+’ characters between each word, to replace spaces (for example term=james+br).
	L          string   `url:"l,omitempty"`     // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit      int      `url:"limit,omitempty"` // The limit on the number of objects, or number of objects in the specified relationship, that are returned.
	Types      []string `url:"types,omitempty"` // The list of the types of resources to include in the results.

}

func (s *SearchService) GetCatalogSearchHints(params *GetCatalogSearchHintsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/search/hints").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type SearchForLibraryResourcesParams struct {
	Term   string   `url:"term,omitempty"`   // The text for the search, with a plus sign (+) between each word to replace spaces (for example, term=james+br).
	Types  []string `url:"types,omitempty"`  // The list of the types of resources to include in the results. The possible values are library-albums, library-songs, library-playlists, library-artists, and library-music-videos.
	Limit  int      `url:"limit,omitempty"`  // The limit on the number of objects, or number of objects in the specified relationship, that are returned. The default value is 5 and the maximum value is 25.
	Offset string   `url:"offset,omitempty"` // The next page or group of objects to fetch.

}

func (s *SearchService) SearchForLibraryResources(params *SearchForLibraryResourcesParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/search").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type RatingsService struct {
	sling *sling.Sling
}

func newRatingsService(sling *sling.Sling) *RatingsService {
	return &RatingsService{
		sling: sling.Path("/"),
	}
}

type GetAPersonalAlbumRatingParams struct {
	Id      string   `url:"-"`                 // The unique identifier for the album.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *RatingsService) GetAPersonalAlbumRating(params *GetAPersonalAlbumRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/ratings/albums/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetAPersonalMusicVideoRatingParams struct {
	Id      string   `url:"-"`                 // The unique identifier for the music video.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *RatingsService) GetAPersonalMusicVideoRating(params *GetAPersonalMusicVideoRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/ratings/music-videos/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetAPersonalPlaylistRatingParams struct {
	Id      string   `url:"-"`                 // The unique identifier for the playlist.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *RatingsService) GetAPersonalPlaylistRating(params *GetAPersonalPlaylistRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/ratings/playlists/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetAPersonalSongRatingParams struct {
	Id      string   `url:"-"`                 // The unique identifier for the song.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *RatingsService) GetAPersonalSongRating(params *GetAPersonalSongRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/ratings/songs/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetAPersonalStationRatingParams struct {
	Id      string   `url:"-"`                 // The unique identifier for the station.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *RatingsService) GetAPersonalStationRating(params *GetAPersonalStationRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/ratings/stations/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultiplePersonalAlbumRatingsParams struct {
	Ids     []string `url:"ids"`               // The unique identifiers for the albums.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *RatingsService) GetMultiplePersonalAlbumRatings(params *GetMultiplePersonalAlbumRatingsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/ratings/albums").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultiplePersonalMusicVideoRatingsParams struct {
	Ids     []string `url:"ids"`               // The unique identifiers for the music videos.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *RatingsService) GetMultiplePersonalMusicVideoRatings(params *GetMultiplePersonalMusicVideoRatingsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/ratings/music-videos").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultiplePersonalPlaylistRatingsParams struct {
	Ids     []string `url:"ids"`               // The unique identifiers for the playlists.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *RatingsService) GetMultiplePersonalPlaylistRatings(params *GetMultiplePersonalPlaylistRatingsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/ratings/playlists").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultiplePersonalSongRatingsParams struct {
	Ids     []string `url:"ids"`               // The unique identifiers for the songs.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *RatingsService) GetMultiplePersonalSongRatings(params *GetMultiplePersonalSongRatingsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/ratings/songs").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultiplePersonalStationRatingsParams struct {
	Ids     []string `url:"ids"`               // The unique identifiers for the stations.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *RatingsService) GetMultiplePersonalStationRatings(params *GetMultiplePersonalStationRatingsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/ratings/stations").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type AddAPersonalAlbumRatingParams struct {
	Id string `url:"-"` // The unique identifier for the album.

}

func (s *RatingsService) AddAPersonalAlbumRating(params *AddAPersonalAlbumRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Put("/v1/me/ratings/albums/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type AddAPersonalMusicVideoRatingParams struct {
	Id string `url:"-"` // The unique identifier for the music video.

}

func (s *RatingsService) AddAPersonalMusicVideoRating(params *AddAPersonalMusicVideoRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Put("/v1/me/ratings/music-videos/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type AddAPersonalPlaylistRatingParams struct {
	Id string `url:"-"` // The unique identifier for the playlist.

}

func (s *RatingsService) AddAPersonalPlaylistRating(params *AddAPersonalPlaylistRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Put("/v1/me/ratings/playlists/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type AddAPersonalSongRatingParams struct {
	Id string `url:"-"` // The unique identifier for the song.

}

func (s *RatingsService) AddAPersonalSongRating(params *AddAPersonalSongRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Put("/v1/me/ratings/songs/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type AddAPersonalStationRatingParams struct {
	Id string `url:"-"` // The unique identifier for the station.

}

func (s *RatingsService) AddAPersonalStationRating(params *AddAPersonalStationRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Put("/v1/me/ratings/stations/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type DeleteAPersonalAlbumRatingParams struct {
	Id string `url:"-"` // The unique identifier for the album.

}

func (s *RatingsService) DeleteAPersonalAlbumRating(params *DeleteAPersonalAlbumRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Delete("/v1/me/ratings/albums/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type DeleteAPersonalMusicVideoRatingParams struct {
	Id string `url:"-"` // The unique identifier for the music video.

}

func (s *RatingsService) DeleteAPersonalMusicVideoRating(params *DeleteAPersonalMusicVideoRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Delete("/v1/me/ratings/music-videos/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type DeleteAPersonalPlaylistRatingParams struct {
	Id string `url:"-"` // The unique identifier for the playlist.

}

func (s *RatingsService) DeleteAPersonalPlaylistRating(params *DeleteAPersonalPlaylistRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Delete("/v1/me/ratings/playlists/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type DeleteAPersonalSongRatingParams struct {
	Id string `url:"-"` // The unique identifier for the song.

}

func (s *RatingsService) DeleteAPersonalSongRating(params *DeleteAPersonalSongRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Delete("/v1/me/ratings/songs/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type DeleteAPersonalStationRatingParams struct {
	Id string `url:"-"` // The unique identifier for the station.

}

func (s *RatingsService) DeleteAPersonalStationRating(params *DeleteAPersonalStationRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Delete("/v1/me/ratings/stations/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetAPersonalLibraryMusicVideoRatingParams struct {
	Id      string   `url:"-"`                 // The unique identifier for the library music video.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *RatingsService) GetAPersonalLibraryMusicVideoRating(params *GetAPersonalLibraryMusicVideoRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/ratings/library-music-videos/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetAPersonalLibraryPlaylistRatingParams struct {
	Id      string   `url:"-"`                 // The unique identifier for the library playlist.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *RatingsService) GetAPersonalLibraryPlaylistRating(params *GetAPersonalLibraryPlaylistRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/ratings/library-playlists/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetAPersonalLibrarySongRatingParams struct {
	Id      string   `url:"-"`                 // The unique identifier for the library song.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *RatingsService) GetAPersonalLibrarySongRating(params *GetAPersonalLibrarySongRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/ratings/library-songs/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultiplePersonalLibraryMusicVideoRatingsParams struct {
	Ids     []string `url:"ids"`               // The unique identifiers for the library music videos.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *RatingsService) GetMultiplePersonalLibraryMusicVideoRatings(params *GetMultiplePersonalLibraryMusicVideoRatingsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/ratings/library-music-videos").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultiplePersonalLibraryPlaylistRatingsParams struct {
	Ids     []string `url:"ids"`               // The unique identifiers for the library playlists.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *RatingsService) GetMultiplePersonalLibraryPlaylistRatings(params *GetMultiplePersonalLibraryPlaylistRatingsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/ratings/library-playlists").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultiplePersonalLibrarySongsRatingsParams struct {
	Ids     []string `url:"ids"`               // The unique identifiers for the library songs.
	Include []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L       string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *RatingsService) GetMultiplePersonalLibrarySongsRatings(params *GetMultiplePersonalLibrarySongsRatingsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/ratings/library-songs").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type AddAPersonalLibraryMusicVideoRatingParams struct {
	Id string `url:"-"` // The unique identifier for the library music video.

}

func (s *RatingsService) AddAPersonalLibraryMusicVideoRating(params *AddAPersonalLibraryMusicVideoRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Put("/v1/me/ratings/library-music-videos/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type AddAPersonalLibraryPlaylistRatingParams struct {
	Id string `url:"-"` // The unique identifier for the library playlist.

}

func (s *RatingsService) AddAPersonalLibraryPlaylistRating(params *AddAPersonalLibraryPlaylistRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Put("/v1/me/ratings/library-playlists/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type AddAPersonalLibrarySongRatingParams struct {
	Id string `url:"-"` // The unique identifier for the library song.

}

func (s *RatingsService) AddAPersonalLibrarySongRating(params *AddAPersonalLibrarySongRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Put("/v1/me/ratings/library-songs/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type DeleteAPersonalLibraryMusicVideoRatingParams struct {
	Id string `url:"-"` // The unique identifier for the library music video.

}

func (s *RatingsService) DeleteAPersonalLibraryMusicVideoRating(params *DeleteAPersonalLibraryMusicVideoRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Delete("/v1/me/ratings/library-music-videos/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type DeleteAPersonalLibraryPlaylistRatingParams struct {
	Id string `url:"-"` // The unique identifier for the library playlist.

}

func (s *RatingsService) DeleteAPersonalLibraryPlaylistRating(params *DeleteAPersonalLibraryPlaylistRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Delete("/v1/me/ratings/library-playlists/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type DeleteAPersonalLibrarySongRatingParams struct {
	Id string `url:"-"` // The unique identifier for the library song.

}

func (s *RatingsService) DeleteAPersonalLibrarySongRating(params *DeleteAPersonalLibrarySongRatingParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Delete("/v1/me/ratings/library-songs/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type ChartsService struct {
	sling *sling.Sling
}

func newChartsService(sling *sling.Sling) *ChartsService {
	return &ChartsService{
		sling: sling.Path("/"),
	}
}

type GetCatalogChartsParams struct {
	Storefront string   `url:"-"`                // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	Types      []string `url:"types,omitempty"`  // A list of the types of charts to include in the results. The possible values are albums, songs, and music-videos.
	L          string   `url:"l,omitempty"`      // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Chart      string   `url:"chart,omitempty"`  // The chart to fetch for the specified types. For possible values, get all the charts by sending this endpoint without the chart parameter. The possible values for this parameter are the chart attributes of the Chart objects in the response.
	Genre      string   `url:"genre,omitempty"`  // The identifier for the genre to use in the chart results.
	Limit      int      `url:"limit,omitempty"`  // The number of resources to include per chart. The default value is 20 and the maximum value is 50.
	Offset     string   `url:"offset,omitempty"` // (Optional; only appears when chart is specified) The next page or group of objects to fetch.

}

func (s *ChartsService) GetCatalogCharts(params *GetCatalogChartsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/charts").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type MusicGenresService struct {
	sling *sling.Sling
}

func newMusicGenresService(sling *sling.Sling) *MusicGenresService {
	return &MusicGenresService{
		sling: sling.Path("/"),
	}
}

type GetACatalogGenreParams struct {
	Id         string `url:"-"`           // The unique identifier for the genre.
	Storefront string `url:"-"`           // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	L          string `url:"l,omitempty"` // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *MusicGenresService) GetACatalogGenre(params *GetACatalogGenreParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/genres/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetACatalogGenresRelationshipDirectlyByNameParams struct {
	Id           string `url:"-"`                // The unique identifier for the genre.
	Relationship string `url:"-"`                // The name of the relationship you want to fetch for this resource.
	Storefront   string `url:"-"`                // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	L            string `url:"l,omitempty"`      // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit        int    `url:"limit,omitempty"`  // The limit on the number of objects, or number of objects in the specified relationship, that are returned.
	Offset       string `url:"offset,omitempty"` // The next page or group of objects to fetch.

}

func (s *MusicGenresService) GetACatalogGenresRelationshipDirectlyByName(params *GetACatalogGenresRelationshipDirectlyByNameParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/genres/"+params.Id+"/"+params.Relationship+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultipleCatalogGenresParams struct {
	Storefront string   `url:"-"`           // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	Ids        []string `url:"ids"`         // The unique identifiers for the catalog genres. For possible values, get all the genres for the current top charts by sending this endpoint without the ids parameter.
	L          string   `url:"l,omitempty"` // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *MusicGenresService) GetMultipleCatalogGenres(params *GetMultipleCatalogGenresParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/genres").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetCatalogTopChartsGenresParams struct {
	Storefront string `url:"-"`                // An Apple Music territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	L          string `url:"l,omitempty"`      // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit      int    `url:"limit,omitempty"`  // The limit on the number of objects, or number of objects in the specified relationship, that are returned.
	Offset     string `url:"offset,omitempty"` // The next page or group of objects to fetch.

}

func (s *MusicGenresService) GetCatalogTopChartsGenres(params *GetCatalogTopChartsGenresParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/genres").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type CuratorsService struct {
	sling *sling.Sling
}

func newCuratorsService(sling *sling.Sling) *CuratorsService {
	return &CuratorsService{
		sling: sling.Path("/"),
	}
}

type GetACatalogCuratorParams struct {
	Id         string   `url:"-"`                 // The unique identifier for the curator.
	Storefront string   `url:"-"`                 // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	L          string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include    []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *CuratorsService) GetACatalogCurator(params *GetACatalogCuratorParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/curators/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetACatalogCuratorsRelationshipDirectlyByNameParams struct {
	Id           string `url:"-"`               // A unique identifier for the curator.
	Relationship string `url:"-"`               // The name of the relationship you want to fetch for this resource.
	Storefront   string `url:"-"`               // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	L            string `url:"l,omitempty"`     // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit        int    `url:"limit,omitempty"` // The limit on the number of objects, or number of objects in the specified relationship, that are returned.

}

func (s *CuratorsService) GetACatalogCuratorsRelationshipDirectlyByName(params *GetACatalogCuratorsRelationshipDirectlyByNameParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/curators/"+params.Id+"/"+params.Relationship+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultipleCatalogCuratorsParams struct {
	Storefront string   `url:"-"`                 // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	Ids        []string `url:"ids"`               // The unique identifiers for the curators. The maximum fetch limit is 100.
	Include    []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L          string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *CuratorsService) GetMultipleCatalogCurators(params *GetMultipleCatalogCuratorsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/curators").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetACatalogAppleCuratorParams struct {
	Id         string   `url:"-"`                 // The unique identifier for the Apple curator.
	Storefront string   `url:"-"`                 // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	L          string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include    []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *CuratorsService) GetACatalogAppleCurator(params *GetACatalogAppleCuratorParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/apple-curators/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetACatalogAppleCuratorsRelationshipDirectlyByNameParams struct {
	Id           string `url:"-"`               // A unique identifier for the Apple curator.
	Relationship string `url:"-"`               // The name of the relationship you want to fetch for this resource.
	Storefront   string `url:"-"`               // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	L            string `url:"l,omitempty"`     // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit        int    `url:"limit,omitempty"` // The limit on the number of objects, or number of objects in the specified relationship, that are returned.

}

func (s *CuratorsService) GetACatalogAppleCuratorsRelationshipDirectlyByName(params *GetACatalogAppleCuratorsRelationshipDirectlyByNameParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/apple-curators/"+params.Id+"/"+params.Relationship+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultipleCatalogAppleCuratorsParams struct {
	Storefront string   `url:"-"`                 // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	Ids        []string `url:"ids"`               // The unique identifiers for the Apple curators. The maximum fetch limit is 100.
	Include    []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L          string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *CuratorsService) GetMultipleCatalogAppleCurators(params *GetMultipleCatalogAppleCuratorsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/apple-curators").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type RecommendationsService struct {
	sling *sling.Sling
}

func newRecommendationsService(sling *sling.Sling) *RecommendationsService {
	return &RecommendationsService{
		sling: sling.Path("/"),
	}
}

type GetARecommendationParams struct {
	Id     string `url:"-"`                // The ID for the recommendation.
	L      string `url:"l,omitempty"`      // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit  int    `url:"limit,omitempty"`  // The limit on the number of objects, or number of objects in the specified relationship, that are returned. The default value and the maximum value are both 10.
	Offset int    `url:"offset,omitempty"` // The next page or group of objects to fetch.

}

func (s *RecommendationsService) GetARecommendation(params *GetARecommendationParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/recommendations/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultipleRecommendationsParams struct {
	Ids    []string `url:"ids"`              // The identifiers for the resource types to fetch. For possible values, get all the recommendations by sending this endpoint without the ids parameter.
	L      string   `url:"l,omitempty"`      // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit  int      `url:"limit,omitempty"`  // The limit on the number of objects, or number of objects in the specified relationship, that are returned. The default value is 10 and the maximum value is 10.
	Offset int      `url:"offset,omitempty"` // The next page or group of objects to fetch.

}

func (s *RecommendationsService) GetMultipleRecommendations(params *GetMultipleRecommendationsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/recommendations").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetDefaultRecommendationsParams struct {
	L      string `url:"l,omitempty"`      // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit  int    `url:"limit,omitempty"`  // The limit on the number of objects, or number of objects in the specified relationship, that are returned. The default value and the maximum value are both 10.
	Offset int    `url:"offset,omitempty"` // The next page or group of objects to fetch.
	Type   string `url:"type,omitempty"`   // The type of recommendations to retrieve. The value can be playlists or albums.

}

func (s *RecommendationsService) GetDefaultRecommendations(params *GetDefaultRecommendationsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/recommendations").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type ActivitiesService struct {
	sling *sling.Sling
}

func newActivitiesService(sling *sling.Sling) *ActivitiesService {
	return &ActivitiesService{
		sling: sling.Path("/"),
	}
}

type GetACatalogActivityParams struct {
	Id         string   `url:"-"`                 // The unique identifier for the activity.
	Storefront string   `url:"-"`                 // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	L          string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Include    []string `url:"include,omitempty"` // Additional relationships to include in the fetch.

}

func (s *ActivitiesService) GetACatalogActivity(params *GetACatalogActivityParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/activities/"+params.Id+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetACatalogActivitysRelationshipDirectlyByNameParams struct {
	Id           string `url:"-"`               // A unique identifier for the activity.
	Relationship string `url:"-"`               // The name of the relationship you want to fetch for this resource.
	Storefront   string `url:"-"`               // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	Limit        int    `url:"limit,omitempty"` // The limit on the number of objects, or number of objects in the specified relationship, that are returned.

}

func (s *ActivitiesService) GetACatalogActivitysRelationshipDirectlyByName(params *GetACatalogActivitysRelationshipDirectlyByNameParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/activities/"+params.Id+"/"+params.Relationship+"").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetMultipleCatalogActivitiesParams struct {
	Storefront string   `url:"-"`                 // An iTunes Store territory, specified by an ISO 3166 alpha-2 country code. The possible values are the id attributes of Storefront objects.
	Ids        []string `url:"ids"`               // The unique identifiers for the activities. The maximum fetch limit is 100.
	Include    []string `url:"include,omitempty"` // Additional relationships to include in the fetch.
	L          string   `url:"l,omitempty"`       // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.

}

func (s *ActivitiesService) GetMultipleCatalogActivities(params *GetMultipleCatalogActivitiesParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/catalog/"+params.Storefront+"/activities").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type HistoryService struct {
	sling *sling.Sling
}

func newHistoryService(sling *sling.Sling) *HistoryService {
	return &HistoryService{
		sling: sling.Path("/"),
	}
}

type GetHeavyRotationContentParams struct {
	L      string `url:"l,omitempty"`      // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit  int    `url:"limit,omitempty"`  // The limit on the number of objects, or number of objects in the specified relationship, that are returned.
	Offset int    `url:"offset,omitempty"` // The next page or group of objects to fetch.

}

func (s *HistoryService) GetHeavyRotationContent(params *GetHeavyRotationContentParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/history/heavy-rotation").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetRecentlyPlayedResourcesParams struct {
	L      string `url:"l,omitempty"`      // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit  int    `url:"limit,omitempty"`  // The limit on the number of objects, or number of objects in the specified relationship, that are returned.
	Offset int    `url:"offset,omitempty"` // The next page or group of objects to fetch.

}

func (s *HistoryService) GetRecentlyPlayedResources(params *GetRecentlyPlayedResourcesParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/recent/played").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetRecentlyPlayedStationsParams struct {
	L      string `url:"l,omitempty"`      // The localization to use, specified by a language tag. The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront. Otherwise, the storefront’s defaultLanguageTag is used.
	Limit  int    `url:"limit,omitempty"`  // The limit on the number of objects, or number of objects in the specified relationship, that are returned.
	Offset int    `url:"offset,omitempty"` // The next page or group of objects to fetch.

}

func (s *HistoryService) GetRecentlyPlayedStations(params *GetRecentlyPlayedStationsParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/recent/radio-stations").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

type GetRecentlyAddedResourcesParams struct {
	Limit  int    `url:"limit,omitempty"`  // The limit on the number of objects, or number of objects in the specified relationship, that are returned.
	Offset string `url:"offset,omitempty"` // The next page or group of objects to fetch.

}

func (s *HistoryService) GetRecentlyAddedResources(params *GetRecentlyAddedResourcesParams) (*VariableResponse, *http.Response, error) {
	response := new(VariableResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("/v1/me/library/recently-added").QueryStruct(params).Receive(response, apiError)
	return response, resp, relevantError(err, *apiError)
}

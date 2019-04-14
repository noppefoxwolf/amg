package applemusic

import (
	"net/http"
)
import "github.com/dghubble/sling"

const applemusicAPI = "https://api.music.apple.com"

type Client struct {
	sling                      *sling.Sling
	StorefrontsAndLocalization *StorefrontsAndLocalizationService
	CommonObjects              *CommonObjectsService
	Albums                     *AlbumsService
	Artists                    *ArtistsService
	Songs                      *SongsService
	MusicVideos                *MusicVideosService
	Playlists                  *PlaylistsService
	AppleMusicStations         *AppleMusicStationsService
	Search                     *SearchService
	Ratings                    *RatingsService
	Charts                     *ChartsService
	MusicGenres                *MusicGenresService
	Curators                   *CuratorsService
	Recommendations            *RecommendationsService
	Activities                 *ActivitiesService
	History                    *HistoryService
}

func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(applemusicAPI)
	return &Client{
		sling:                      base,
		StorefrontsAndLocalization: newStorefrontsAndLocalizationService(base.New()),
		CommonObjects:              newCommonObjectsService(base.New()),
		Albums:                     newAlbumsService(base.New()),
		Artists:                    newArtistsService(base.New()),
		Songs:                      newSongsService(base.New()),
		MusicVideos:                newMusicVideosService(base.New()),
		Playlists:                  newPlaylistsService(base.New()),
		AppleMusicStations:         newAppleMusicStationsService(base.New()),
		Search:                     newSearchService(base.New()),
		Ratings:                    newRatingsService(base.New()),
		Charts:                     newChartsService(base.New()),
		MusicGenres:                newMusicGenresService(base.New()),
		Curators:                   newCuratorsService(base.New()),
		Recommendations:            newRecommendationsService(base.New()),
		Activities:                 newActivitiesService(base.New()),
		History:                    newHistoryService(base.New()),
	}
}

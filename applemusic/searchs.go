package applemusic

type SearchResponse struct {
	ResponseRoot
	Results SearchResults //(Required) The results included in the response for a search request.
}

type SearchResults struct {
	Activities    []ActivityResponse     //The activities returned for the search query.
	Albums        []AlbumResponse        //The albums returned for the search query.
	AppleCurators []AppleCuratorResponse //The Apple curators returned for the search query.
	Artists       []ArtistResponse       //The artists returned for the search query.
	Curators      []CuratorResponse      //The curators returned for the search query.
	MusicVideos   []MusicVideoResponse   //The music videos returned for the search query.
	Playlists     []PlaylistResponse     //The playlists returned for the search query.
	Songs         []SongResponse         //The songs returned for the search query.
	Stations      []StationResponse      //The stations returned for the search query.
}

//=======

type SearchHintsResponse struct {
	ResponseRoot
	Results SearchHints //(Required) The results included in the response for a search hints request.
}

type SearchHints struct {
	Terms []string //(Required) The autocomplete options derived from the search hint.
}

//==========

type LibrarySearchResponse struct {
	ResponseRoot
	Results LibrarySearchResults //(Required) The results included in the response for a library search request.
}

type LibrarySearchResults struct {
	LibraryAlbums      []LibraryAlbum            //The library albums returned for the search query.
	LibraryArtists     []LibraryArtist           //The library artists returned for the search query.
	LibraryMusicVideos []LibraryMusicVideo       //The library music videos returned for the search query.
	LibraryPlaylists   []LibraryPlaylistResponse //The library playlists returned for the search query.
	LibrarySongs       []LibrarySong             //The library songs returned for the search query.
}

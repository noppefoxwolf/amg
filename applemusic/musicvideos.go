package applemusic

type MusicVideoResponse struct {
	ResponseRoot
	Data []MusicVideo //(Required) The data included in the response for a music video object request.
}

type MusicVideo struct {
	Href          string
	Id            string
	Meta          ResourceMeta
	Attributes    MusicVideoAttributes    //The attributes for the music video.
	Relationships MusicVideoRelationships //The relationships for the music video.
}

type MusicVideoAttributes struct {
	AlbumName        string         //The name of the album the music video appears on.
	ArtistName       string         //(Required) The artist’s name.
	Artwork          Artwork        //(Required) The artwork for the music video’s associated album.
	ContentRating    string         //The Recording Industry Association of America (RIAA) rating of the content. The possible values for this rating are clean and explicit. No value means no rating.
	DurationInMillis int            //The duration of the music video in milliseconds.
	EditorialNotes   EditorialNotes //The editorial notes for the music video.
	GenreNames       []string       //(Required) The music video’s associated genres.
	Isrc             string         //(Required) The International Standard Recording Code (ISRC) for the music video.
	Name             string         //(Required) The localized name of the music video.
	PlayParams       PlayParameters //The parameters to use to play back the music video.
	Previews         []Preview      //(Required) The preview assets for the music video.
	ReleaseDate      ReleaseDate    //(Required) The release date of the music video in YYYY-MM-DD format.
	TrackNumber      int            //The number of the music video in the album’s track list.
	Url              string         //(Required) A URL for sharing the music video.
	VideoSubType     string         //The video subtype associated with the content.
	hHasHDR          bool           //(Required) Whether the music video has HDR10-encoded content.
	Has4K            bool           //(Required) Whether the music video has 4K content.
}

type MusicVideoRelationships struct {
	Relationship
	Albums  AlbumRelationship  //The albums associated with the music video. By default, albums includes identifiers only.  Fetch limits: 10 default, 10 maximum
	Artists ArtistRelationship //The artists associated with the music video. By default, artists includes identifiers only.  Fetch limits: 10 default, 10 maximum
	Genres  GenreRelationship  //The genres associated with the music video. By default, genres is not included.  Fetch limits: None
}

type MusicVideoRelationship struct {
	Relationship
	data []MusicVideo //(Required) The data for the music video included in the relationship.
}

// ==========

type LibraryMusicVideoResponse struct {
	ResponseRoot
	Data []LibraryMusicVideo //(Required) The data included in the response for a library music video object request.
}

type LibraryMusicVideo struct {
	attributes    LibraryMusicVideoAttributes    //The attributes for the library music video.
	relationships LibraryMusicVideoRelationships //The relationships for the library music video.
	Type          string                         //(Required) This value will alway be libraryMusicVideos.  Value: libraryMusicVideos
}

type LibraryMusicVideoAttributes struct {
	AlbumName        string         //(Required) The name of the album the music video appears on.
	ArtistName       string         //(Required) The artist’s name.
	Artwork          Artwork        //(Required) The artwork for the music video’s associated album.
	ContentRating    string         //The Recording Industry Association of America (RIAA) rating of the content. The possible values for this rating are clean and explicit. No value means no rating.
	DurationInMillis int            //The duration of the music video in milliseconds.
	Name             string         //(Required) The localized name of the music video.
	PlayParams       PlayParameters //The parameters to use to playback the music video.
	TrackNumber      int            //The number of the music video in the album’s track list.
}

type LibraryMusicVideoRelationships struct {
	Albums  LibraryAlbumRelationship  //The library albums associated with the music video. By default, albums is not included.  Fetch limits: 10 default, 10 maximum
	Artists LibraryArtistRelationship //The library artists associated with the music video. By default, artists is not included.  Fetch limits: 10 default, 10 maximum
}

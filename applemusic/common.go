package applemusic

import "time"

type ResponseRoot struct {
	Data []Resource
	Errors []Error
	Href string
	Meta ResponseRootMeta
	Next string
	Results ResponseRootResults
}

type ResponseRootMeta struct {

}

type Error struct {
	Code string
	Detail string
	Id string
	Error ErrorSource
	Status string
	Title string
}

type ErrorSource struct {
	Parameter string
	Pointer string //JSONPointer
}

type ResponseRootResults struct {
}

type HistoryResponse struct {
	Data Resource
}

type Resource struct {
	Attributes ResourceAttributes
	Href string
	Id string
	Relationships ResourceRelationships
	Type string
	Meta ResourceMeta
}

type ResourceMeta struct {

}

type ResourceAttributes struct {

}

type ResourceRelationships struct {
	Data []Resource
	Href string
	Meta RelationshipMeta
	Next string
}

type RelationshipMeta struct {

}


type Artwork struct {
	BgColor string
	Height int
	Width int
	TextColor1 string
	TextColor2 string
	TextColor3 string
	TextColor4 string
	Url string
}

type EditorialNotes struct {
	Short string
	Standard string
}

type PlayParameters struct {
	Id string
	Kind string
}

type Preview struct {
	Artwork Artwork
	Url string
}

type AlbumResponse struct {
	ResponseRoot
	Data []Album
}

type LibraryAlbumResponse struct {
	ResponseRoot
	Data []LibraryAlbum //(Required) The data included in the response for a library album object request.
}

type ArtistResponse struct {
	ResponseRoot
	Data []Artist //(Required) The data included in the response for an artist object request.
}

type LibraryArtistResponse struct {
	ResponseRoot
	Data []LibraryArtist //(Required) The data included in the response for a library artist object request.
}

type SongResponse struct {
	ResponseRoot
	Data []Song //(Required) The data included in the response for a song object request.
}

type LibrarySongResponse struct {
	ResponseRoot
	Data []LibrarySong //(Required) The data included in the response for a library song object request.
}

type Song struct {
	Resource
	Attributes SongAttributes //The attributes for the song.
	Relationships SongRelationships //The relationships for the song.
	Type string //(Required) This value will always be songs. Value: songs
}

type LibrarySong struct {
	Resource
	Attributes LibrarySongAttributes //The attributes for the library song.
	Relationships LibrarySongRelationships //The relationships for the library song.
	Type string //(Required) This value will alway be librarySongs. Value: librarySongs
}

type LibrarySongAttributes struct {
	AlbumName string //(Required) The name of the album the song appears on.
	ArtistName string //(Required) The artist’s name.
	Artwork Artwork //(Required) The album artwork.
	ContentRating string //The Recording Industry Association of America (RIAA) rating of the content. The possible values for this rating are clean and explicit. No value means no rating.
	DiscNumber int //(Required) The disc number the song appears on.
	DurationInMillis int //The approximate length of the song in milliseconds.
	Mame string //(Required) The localized name of the song.
	PlayParams PlayParameters //The parameters to use to playback the song.
	TrackNumber int //(Required) The number of the song in the album’s track list.
}

type LibrarySongRelationships struct {
	Albums LibraryAlbumRelationship //The library albums associated with the song. By default, albums is not included. Fetch limits: 10 default, 10 maximum
	Artists LibraryArtistRelationship //The library artists associated with the song. By default, artists is not included. Fetch limits: 10 default, 10 maximum
}

type Album struct {
	Resource
	Attributes AlbumAttributes
	Relationships AlbumRelationships
	Type string
}

type LibraryAlbum struct {
	Resource
	Attributes LibraryAlbumAttributes //The attributes for the library album.
	Relationships LibraryAlbumRelationships //The relationships for the library album.
	Type string //(Required) This value will always be libraryAlbums.  Value: libraryAlbums
}

type AlbumAttributes struct {
	AlbumName string //(Required) The name of the album the music video appears on.
	ArtistName string //(Required) The artist’s name.
	Artwork Artwork //The album artwork.
	ContentRating string //The Recording Industry Association of America (RIAA) rating of the content. The possible values for this rating are clean and explicit. No value means no rating.
	Copyright string //The copyright text.
	EditorialNotes EditorialNotes //The notes about the album that appear in the iTunes Store.
	GenreNames []string //(Required) The names of the genres associated with this album.
	IsComplete bool //(Required) Indicates whether the album is complete. If true, the album is complete; otherwise, it is not. An album is complete if it contains all its tracks and songs.
	IsSingle bool //(Required) Indicates whether the album contains a single song.
	Name string //(Required) The localized name of the album.
	PlayParams PlayParameters //The parameters to use to play back the tracks of the album.
	RecordLabel string //(Required) The name of the record label for the album.
	ReleaseDate time.Time //(Required) The release date of the album in YYYY-MM-DD format.
	TrackCount int //(Required) The number of tracks.
	Url string //(Required) The URL for sharing the album in the iTunes Store.
	IsMasteredForItunes bool //(Required) Indicates whether the album has been delivered as Mastered for iTunes.
}

type SongAttributes struct {
	AlbumName string //(Required) The name of the album the song appears on.
	ArtistName string  //(Required) The artist’s name.
	Artwork Artwork  //(Required) The album artwork.
	ComposerName string //The song’s composer.
	ContentRating string //The Recording Industry Association of America (RIAA) rating of the content. The possible values for this rating are clean and explicit. No value means no rating.
	DiscNumber int //(Required) The disc number the song appears on.
	DurationInMillis int //The duration of the song in milliseconds.
	EditorialNotes EditorialNotes //The notes about the song that appear in the iTunes Store.
	GenreNames []string //(Required) The genre names the song is associated with.
	Isrc string //(Required) The International Standard Recording Code (ISRC) for the song.
	MovementCount int //(Classical music only) The movement count of this song.
	MovementName string //(Classical music only) The movement name of this song.
	MovementNumber int //(Classical music only) The movement number of this song.
	Name string //(Required) The localized name of the song.
	PlayParams PlayParameters  //The parameters to use to playback the song.
	Previews []Preview //(Required) The preview assets for the song.
	ReleaseDate time.Time //(Required) The release date of the song in YYYY-MM-DD format.
	TrackNumber int //(Required) The number of the song in the album’s track list.
	Url string //(Required) The URL for sharing a song in the iTunes Store.
	WorkName string //(Classical music only) The name of the associated work.
}

type SongRelationships struct {
	Albums AlbumRelationship //The albums associated with the song. By default, albums includes identifiers only. Fetch limits: 10 default, 10 maximum
	Artists ArtistRelationship //The artists associated with the song. By default, artists includes identifiers only. Fetch limits: 10 default, 10 maximum
	Genres GenreRelationship //The genres associated with the song. By default, genres is not included. Fetch limits: None
	Station StationRelationship //The station associated with the song. By default, station is not included. Fetch limits: None (one station)
}

type AlbumRelationships struct {
	Artists ArtistRelationship //The artists associated with the album. By default, artists includes identifiers only. Fetch limits: 10 default, 10 maximum
	Genres GenreRelationship //The genres for the album. By default, genres is not included. Fetch limits: None
	Tracks TrackRelationship //The songs and music videos on the album. By default, tracks includes objects.  Fetch limits: 300 default, 300 maximum
}

type LibraryAlbumAttributes struct {
	ArtistName string //(Required) The artist’s name.
	Artwork Artwork //(Required) The album artwork.
	ContentRating string //The Recording Industry Association of America (RIAA) rating of the content. The possible values for this rating are clean and explicit. No value means no rating.
	Name string //(Required) The localized name of the album.
	PlayParams PlayParameters //The parameters to use to playback the tracks of the album.
	TrackCount int //(Required) The number of tracks.
}

type LibraryAlbumRelationships struct {
	Artists LibraryArtistRelationship //The library artists associated with the album. By default, artists is not included. Fetch limits: 10 default, 10 maximum
	Tracks LibraryTrackRelationship //The library songs and library music videos on the album. Only available when fetching single library album resource by ID. By default, tracks includes objects.  Fetch limits: 300 default, 300 maximum
}

type Relationship struct {

}

type Artist struct {
	Resource
	Attributes ArtistAttributes //The attributes for the artist.
	Relationships ArtistRelationships // The relationships for the artist.
	Type string //(Required) This value will always be artists.  Value: artists
}

type LibraryArtist struct {
	Resource
	Attributes LibraryArtistAttributes //The attributes for the library artist.
	Relationships LibraryArtistRelationships // The relationships for the library artist.
	Type string //(Required) This value will alway be libraryArtists. Value: libraryArtists
}

type LibraryArtistAttributes struct {
	Name string //(Required) The artist’s name.
}

type LibraryArtistRelationships struct {
	Albums LibraryAlbumRelationship //The library albums associated with the artist. By default, albums is not included. It is available only when fetching a single library artist resource by ID. Fetch limits: 25 default, 100 maximum
}

type ArtistAttributes struct {
	EditorialNotes EditorialNotes //The notes about the artist that appear in the iTunes Store.
	GenreNames []string //(Required) The names of the genres associated with this artist.
	Name string //(Required) The localized name of the artist.
	Url string //(Required) The URL for sharing an artist in the iTunes Store.
}

type ArtistRelationships struct {
	albums AlbumRelationship //The albums associated with the artist. By default, albums includes identifiers only. Fetch limits: 25 default, 100 maximum
	genres GenreRelationship //The genres associated with the artist. By default, genres is not included. Fetch limits: None
	musicVideos MusicVideoRelationship //The music videos associated with the artist. By default, musicVideos is not included.  Fetch limits: 25 default, 100 maximum
	Playlists PlaylistRelationship //The playlists associated with the artist. By default, playlists is not included. Fetch limits: 10 default, 10 maximum
	Station StationRelationship //The station associated with the artist. By default, station is not included. Fetch limits: None (one station)
}

type AlbumRelationship struct {
	Relationship
	Data []Album //(Required) The data for the album included in the relationship.
}

type GenreRelationship struct {
	Relationship
	Data []Genre //(Required) The data for the genre included in the relationship.
}

type MusicVideoRelationship struct {
	Relationship
	data []MusicVideo //(Required) The data for the music video included in the relationship.
}

type PlaylistRelationship struct {
	Relationship
	Data []Playlist //(Required) The data for the playlist included in the relationship.
}

type StationRelationship struct {
	Relationship
	Data Station //(Required) The data for the station included in the relationship.
}

type ArtistRelationship struct {
	Relationship
	Data []Artist //(Required) The data for the artist included in the relationship.
}

type TrackRelationship struct {
	Relationship
	Data []interface{} //(Song | MusicVideo)] (Required) The data for the track included in the relationship.
}

type LibraryArtistRelationship struct {
	Relationship
	Data []LibraryArtist //(Required) The data for the library artist included in the relationship.
}

type LibraryTrackRelationship struct {
	Relationship
	Data interface{} //[(LibrarySong | LibraryMusicVideo)] (Required) The data for the library track included in the relationship.
}

type LibraryAlbumRelationship struct {
	Relationship
	Data []LibraryAlbum //(Required) The data for the library album included in the relationship.
}

type Genre struct {

}

type MusicVideo struct {

}

type Playlist struct {

}

type Station struct {

}
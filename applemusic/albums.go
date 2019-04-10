package applemusic

import "time"

type AlbumResponse struct {
	ResponseRoot
	Data []Album //(Required) The data included in the response for an album object request.
}

type Album struct {
	Resource
	Attributes AlbumAttributes //The attributes for the album.
	Relationships AlbumRelationships //The relationships for the album.
	Type string //(Required) This value will always be albums. Value: albums
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

type AlbumRelationship struct {
	Relationship
	Data []Album //(Required) The data for the album included in the relationship.
}

type AlbumRelationships struct {
	Artists ArtistRelationship //The artists associated with the album. By default, artists includes identifiers only. Fetch limits: 10 default, 10 maximum
	Genres GenreRelationship //The genres for the album. By default, genres is not included. Fetch limits: None
	Tracks TrackRelationship //The songs and music videos on the album. By default, tracks includes objects.  Fetch limits: 300 default, 300 maximum
}

// ==============

type LibraryAlbumResponse struct {
	ResponseRoot
	Data []LibraryAlbum //(Required) The data included in the response for a library album object request.
}

type LibraryAlbum struct {
	Resource
	Attributes LibraryAlbumAttributes //The attributes for the library album.
	Relationships LibraryAlbumRelationships //The relationships for the library album.
	Type string //(Required) This value will always be libraryAlbums.  Value: libraryAlbums
}

type LibraryAlbumAttributes struct {
	ArtistName string //(Required) The artist’s name.
	Artwork Artwork //(Required) The album artwork.
	ContentRating string //The Recording Industry Association of America (RIAA) rating of the content. The possible values for this rating are clean and explicit. No value means no rating.
	Name string //(Required) The localized name of the album.
	PlayParams PlayParameters //The parameters to use to playback the tracks of the album.
	TrackCount int //(Required) The number of tracks.
}


type LibraryAlbumRelationship struct {
	Relationship
	Data []LibraryAlbum //(Required) The data for the library album included in the relationship.
}

type LibraryAlbumRelationships struct {
	Artists LibraryArtistRelationship //The library artists associated with the album. By default, artists is not included. Fetch limits: 10 default, 10 maximum
	Tracks LibraryTrackRelationship //The library songs and library music videos on the album. Only available when fetching single library album resource by ID. By default, tracks includes objects.  Fetch limits: 300 default, 300 maximum
}
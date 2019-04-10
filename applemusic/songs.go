package applemusic

import "time"

type SongResponse struct {
	ResponseRoot
	Data []Song //(Required) The data included in the response for a song object request.
}

type Song struct {
	Resource
	Attributes SongAttributes //The attributes for the song.
	Relationships SongRelationships //The relationships for the song.
	Type string //(Required) This value will always be songs. Value: songs
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


//===================

type LibrarySongResponse struct {
	ResponseRoot
	Data []LibrarySong //(Required) The data included in the response for a library song object request.
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

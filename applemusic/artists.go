package applemusic

type ArtistResponse struct {
	ResponseRoot
	Data []Artist //(Required) The data included in the response for an artist object request.
}

type Artist struct {
	Href          string
	Id            string
	Meta          ResourceMeta
	Attributes    ArtistAttributes    //The attributes for the artist.
	Relationships ArtistRelationships // The relationships for the artist.
}

type ArtistAttributes struct {
	EditorialNotes EditorialNotes //The notes about the artist that appear in the iTunes Store.
	GenreNames     []string       //(Required) The names of the genres associated with this artist.
	Name           string         //(Required) The localized name of the artist.
	Url            string         //(Required) The URL for sharing an artist in the iTunes Store.
}

type ArtistRelationships struct {
	albums      AlbumRelationship      //The albums associated with the artist. By default, albums includes identifiers only. Fetch limits: 25 default, 100 maximum
	genres      GenreRelationship      //The genres associated with the artist. By default, genres is not included. Fetch limits: None
	musicVideos MusicVideoRelationship //The music videos associated with the artist. By default, musicVideos is not included.  Fetch limits: 25 default, 100 maximum
	Playlists   PlaylistRelationship   //The playlists associated with the artist. By default, playlists is not included. Fetch limits: 10 default, 10 maximum
	Station     StationRelationship    //The station associated with the artist. By default, station is not included. Fetch limits: None (one station)
}

type ArtistRelationship struct {
	Relationship
	Data []Artist //(Required) The data for the artist included in the relationship.
}

//========

type LibraryArtistResponse struct {
	ResponseRoot
	Data []LibraryArtist //(Required) The data included in the response for a library artist object request.
}

type LibraryArtist struct {
	Href          string
	Id            string
	Meta          ResourceMeta
	Attributes    LibraryArtistAttributes    //The attributes for the library artist.
	Relationships LibraryArtistRelationships // The relationships for the library artist.
}

type LibraryArtistAttributes struct {
	Name string //(Required) The artist’s name.
}

type LibraryArtistRelationships struct {
	Albums LibraryAlbumRelationship //The library albums associated with the artist. By default, albums is not included. It is available only when fetching a single library artist resource by ID. Fetch limits: 25 default, 100 maximum
}

type LibraryArtistRelationship struct {
	Relationship
	Data []LibraryArtist //(Required) The data for the library artist included in the relationship.
}

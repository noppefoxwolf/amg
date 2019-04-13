package models

import "time"

type PlaylistResponse struct {
	ResponseRoot
	Data []Playlist //(Required) The data included in the response for a playlist object request.
}

type Playlist struct {
	Href          string
	Id            string
	Meta          ResourceMeta
	Attributes    PlaylistAttributes    //The attributes for the playlist.
	Relationships PlaylistRelationships //The relationships for the playlist.
}

type PlaylistAttributes struct {
	artwork          Artwork        //The playlist artwork.
	curatorName      string         //The display name of the curator.
	description      EditorialNotes //A description of the playlist.
	lastModifiedDate time.Time      //(Required) The date the playlist was last modified.
	name             string         //(Required) The localized name of the album.
	playParams       PlayParameters //The parameters to use to play back the tracks in the playlist.
	playlistType     string
	//(Required) The type of playlist. Possible values are:
	//user-shared: A playlist created and shared by an Apple Music user.
	//editorial: A playlist created by an Apple Music curator.
	//external: A playlist created by a non-Apple curator or brand.
	//personal-mix: A personalized playlist for an Apple Music user.
	//Possible values: user-shared, editorial, external, personal-mix
	Url string //(Required) The URL for sharing an album in the iTunes Store.
}

type PlaylistRelationships struct {
	Curator CuratorRelationship //The curator that created the playlist. By default, curator includes identifiers only.  Fetch limits: None
	Tracks  TrackRelationship   //The songs and music videos included in the playlist. By default, tracks includes objects. Fetch limits: 100 default, 300 maximum
}

type PlaylistRelationship struct {
	Relationship
	Data []Playlist //(Required) The data for the playlist included in the relationship.
}

//=======

type LibraryPlaylistResponse struct {
	ResponseRoot
	Data []LibraryPlaylist //(Required) The data included in the response for a library playlist object request.
}

type LibraryPlaylist struct {
	Attributes    LibraryPlaylistAttributes    //The attributes for the library playlist.
	Relationships LibraryPlaylistRelationships //The relationships for the library playlist.
	Type          string                       //(Required) This value will alway be libraryPlaylists. Value: libraryPlaylists
}

type LibraryPlaylistAttributes struct {
	Artwork     Artwork        //The playlist artwork.
	Description string         //A description of the playlist.
	Name        string         //(Required) The localized name of the album.
	PlayParams  PlayParameters //The parameters to use to play back the tracks in the playlist.
	CanEdit     bool           //(Required) Indicates whether the playlist can be edited.
}

type LibraryPlaylistRelationships struct {
	Tracks LibraryTrackRelationship //The library songs and library music videos included in the playlist. By default, tracks is not included. Only available when fetching a single library playlist resource by ID. Fetch limits: 100 default, 100 maximum
}

type LibraryTrackRelationship struct {
	Relationship
	Data interface{} //[(LibrarySong | LibraryMusicVideo)] (Required) The data for the library track included in the relationship.
}

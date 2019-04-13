package models

type CuratorResponse struct {
	ResponseRoot
	Data []Curator //(Required) The data included in the response for a curator object request.
}

type Curator struct {
	Href          string
	Id            string
	Meta          ResourceMeta
	Attributes    CuratorAttributes    //The attributes for the curator.
	Relationships CuratorRelationships //The relationships for the curator.
}

type CuratorAttributes struct {
	Artwork        Artwork        //(Required) The curator artwork.
	EditorialNotes EditorialNotes //The notes about the curator.
	Name           string         //(Required) The localized name of the curator.
	Url            string         //(Required) The URL for sharing a curator in Apple Music.
}

type CuratorRelationships struct {
	Playlists PlaylistRelationship //The playlists associated with the curator. By default, playlists includes identifiers only.  Fetch limits: 10 default, 10 maximum
}

type CuratorRelationship struct {
	Relationship
	Data []Curator //(Required) The data for the curator included in the relationship.
}

//======

type AppleCuratorResponse struct {
	ResponseRoot
	Data []AppleCurator //(Required) The data included in the response for an Apple curator object request.
}

type AppleCurator struct {
	Href          string
	Id            string
	Meta          ResourceMeta
	Attributes    AppleCuratorAttributes    //The attributes for the Apple curator.
	Relationships AppleCuratorRelationships //The relationships for the Apple curator.
}

type AppleCuratorAttributes struct {
	Artwork        Artwork        //(Required) The curator artwork.
	EditorialNotes EditorialNotes //The notes about the curator that appear in the iTunes Store.
	Name           string         //(Required) The localized name of the curator.
	Url            string         //(Required) The URL for sharing the curator in the iTunes Store.
}

type AppleCuratorRelationships struct {
	Playlists PlaylistRelationship //The playlists associated with this curator. By default, playlists includes identifiers only. Fetch limits: 10 default, 10 maximum
}

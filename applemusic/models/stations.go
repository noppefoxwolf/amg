package models

type StationResponse struct {
	ResponseRoot
	Data []Station //(Required) The data included in the response for a station object request.
}

type Station struct {
	Href          string
	Id            string
	Relationships ResourceRelationships
	Meta          ResourceMeta
	Attributes    StationAttributes //The attributes for the station.
}

type StationAttributes struct {
	Artwork          Artwork        //(Required) The radio station artwork.
	DurationInMillis int            //The duration of the stream. This value is not emitted for 'live' or programmed stations.
	EditorialNotes   EditorialNotes //The notes about the station that appear in Apple Music.
	EpisodeNumber    int            //The episode number of the station. This value is emitted when the station represents an episode of a show or other content.
	IsLive           bool           //(Required) Whether the station is a live stream.
	Name             string         //(Required) The localized name of the station.
	Url              string         //(Required) The URL for sharing a station in Apple Music.
}

type StationRelationship struct {
	Relationship
	Data Station //(Required) The data for the station included in the relationship.
}

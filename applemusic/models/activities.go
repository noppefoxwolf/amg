package models

type ActivityResponse struct {
	ResponseRoot
	Data []Activity //(Required) The data included in the response to an activity object request.
}

type Activity struct {
	Attributes    ActivityAttributes    //The attributes for the activity.
	Relationships ActivityRelationships //The relationships for the activity.
	Type          string                // (Required) Always activities. Value: activities
}

type ActivityAttributes struct {
	Artwork        Artwork        //(Required) The activity artwork.
	EditorialNotes EditorialNotes //The notes about the activity that appear in the iTunes Store.
	Name           string         //(Required) The localized name of the activity.
	Url            string         //(Required) The URL for sharing an activity in the iTunes Store.
}

type ActivityRelationships struct {
	Playlists PlaylistRelationship //The playlists associated with this activity. By default, playlists includes identifiers only. Fetch limits: 10 default, 10 maximum
}

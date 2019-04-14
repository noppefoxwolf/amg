package applemusic

type RatingResponse struct {
	ResponseRoot
	Data []Rating //(Required) The data included in the response for a rating object request.
}

type Rating struct {
	Href          string
	Id            string
	Relationships ResourceRelationships
	Meta          ResourceMeta
	Attributes    RatingAttributes //The attributes for the rating.
}

type RatingAttributes struct {
	Value int //(Required) The value for the resource's rating. The possible values for the value key are 1 and -1. Possible values: 1, -1
}

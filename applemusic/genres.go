package applemusic

type GenreResponse struct {
	ResponseRoot
	Data []Genre //(Required) The data included in the response for a genre object request.
}

type Genre struct {
	Attributes GenreAttributes //The attributes for the genre.
	Type string //(Required) This value will always be genres.  Value: genres
}

type GenreAttributes struct {
	Name string //(Required) The localized name of the genre.
}

type GenreRelationship struct {
	Relationship
	Data []Genre //(Required) The data for the genre included in the relationship.
}
package applemusic

type ResponseRoot struct {
	Data    []Resource
	Errors  []Error
	Href    string
	Meta    ResponseRootMeta
	Next    string
	Results ResponseRootResults
}

type ResponseRootMeta struct {
}

type Error struct {
	Code   string
	Detail string
	Id     string
	Error  ErrorSource
	Status string
	Title  string
}

type ErrorSource struct {
	Parameter string
	Pointer   string //JSONPointer
}

type ResponseRootResults struct {
}

type Resource struct {
	Attributes    ResourceAttributes
	Href          string
	Id            string
	Relationships ResourceRelationships
	Type          string
	Meta          ResourceMeta
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
	BgColor    string
	Height     int
	Width      int
	TextColor1 string
	TextColor2 string
	TextColor3 string
	TextColor4 string
	Url        string
}

type EditorialNotes struct {
	Short    string
	Standard string
}

type PlayParameters struct {
	Id        string
	Kind      string
	IsLibrary bool //UnDocumented
}

type Preview struct {
	Artwork Artwork
	Url     string
}

type Relationship struct {
}

type TrackRelationship struct {
	Relationship
	Data []interface{} //(Song | MusicVideo)] (Required) The data for the track included in the relationship.
}

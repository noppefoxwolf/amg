package models

type StorefrontResponse struct {
	ResponseRoot
	Data []Storefront
}

type Storefront struct {
	Href          string
	Id            string
	Relationships ResourceRelationships
	Meta          ResourceMeta
	Attributes    StorefrontAttributes //The attributes for the storefront.
}

type StorefrontAttributes struct {
	DefaultLanguageTag    string   //(Required) The default language for the storefront, represented as a language tag.
	Name                  string   //(Required) The localized name of the storefront.
	SupportedLanguageTags []string // (Required) The localizations that the storefront supports, represented as an array of language tags.
	ExplicitContentPolicy string   //Undocumented
}

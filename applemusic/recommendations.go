package applemusic

import "time"

type RecommendationResponse struct {
	ResponseRoot
	Data []Recommendation //(Required) The data included in the response for a recommendation object request.
}

type Recommendation struct {
	Meta          ResourceMeta
	Attributes    RecommendationAttributes    //The attributes for the recommendation.
	Relationships RecommendationRelationships //For contents, a list of recommended candidates that are a mixture of albums and playlists. For recommendations, a list of recommendations within a group.
	Next          string                      //(Required) The URL for the next page.
	Id            string                      //(Required) The ID of the recommendation.
	Href          string                      //(Required) The URL for the recommendation.
}

type RecommendationAttributes struct {
	IsGroupRecommendation bool      //(Required) Whether the recommendation is of group type.
	NextUpdateDate        time.Time //(Required) The date in UTC format when the recommendation is updated.
	Reason                string    //(Required) The localized reason for the recommendation.
	ResourceTypes         []string  //(Required) The resource types supported by the recommendation.
	Title                 string    //(Required) The localized title for the recommendation.
}

type RecommendationRelationships struct {
	Contents        Relationship //The contents associated with the content recommendation type. By default, contents includes objects.  Fetch limits: 10 default, 10 maximum
	Recommendations Relationship //The recommendations associated with the group recommendation type. By default, recommendations includes objects. Fetch limits: None
}

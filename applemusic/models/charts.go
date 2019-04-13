package models

type ChartResponse struct {
	Albums      Chart //The albums returned when fetching charts.
	MusicVideos Chart //The music videos returned when fetching charts.
	Songs       Chart //The songs returned when fetching charts.
}

type Chart struct {
	Chart string             //(Required) The chart identifier.
	Data  []VariableResource //(Required) An array of the requested objects, ordered by popularity. For example, if songs were specified as the chart type in the request, the array contains Song objects.
	Href  string             //(Required) The URL for the chart.
	Name  string             //(Required) The localized name for the chart.
	Next  string             //The URL for the next page.
}

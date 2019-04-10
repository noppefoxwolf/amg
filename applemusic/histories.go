package applemusic

type HistoryResponse struct {
	ResponseRoot
	Data []Resource //(Required) The data included in the response for a history object request.
}

type RecentlyAddedResponse struct {
	ResponseRoot
	Data []Resource //(Required) The data included in the response for a recently added object request.
}

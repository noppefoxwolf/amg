package models

type HistoryResponse struct {
	ResponseRoot
	Data []VariableResource //(Required) The data included in the response for a history object request.
}

type RecentlyAddedResponse struct {
	ResponseRoot
	Data []VariableResource //(Required) The data included in the response for a recently added object request.
}

package models

// Response structure for api reponse ROOT
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Results interface{} `json:"results"`
}

// NewResponse : create new Response
func NewResponse(status string, message string, results interface{}) Response {
	return Response{Status: status, Message: message, Results: results}
}

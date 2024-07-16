package enums

type GenericResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"statusCode"`
}

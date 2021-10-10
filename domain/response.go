package domain

type Response struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

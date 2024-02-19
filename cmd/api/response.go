package api

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Status  int    `json:"status"`
	Error   bool   `json:"error"`
}

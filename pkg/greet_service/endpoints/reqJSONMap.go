package endpoints

type GreetRequest struct {
	Name string `json:"name"`
}

type GreetResponse struct {
	Message string `json:"message"`
	Err     string `json:"error"`
}

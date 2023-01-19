package endpoints

type GiveNameRequest struct {
	Name string `json:"name"`
}

type GiveNameResponse struct {
	Message string `json:"message"`
	Err     string `json:"error"`
}

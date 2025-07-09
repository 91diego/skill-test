package models

type ValidationErrorResponse struct {
	Error  string             `json:"error"`
	Detail []ValidationDetail `json:"detail"`
}

type ValidationDetail struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}

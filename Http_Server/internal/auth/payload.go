package auth

type LoginRequest struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

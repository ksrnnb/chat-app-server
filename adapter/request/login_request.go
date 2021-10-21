package request

type LoginRequest struct {
	LoginId string  `json:"loginId"`
	Password string `json:"password"`
}
package controller

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/ksrnnb/chat-app-server/adapter/response"
)

type TokenController struct{}

func NewTokenController() TokenController {
	return TokenController{}
}

func (controller TokenController) GetToken(session ISession) *response.Response {
	digit := 32
	randByte := make([]byte, digit)
	_, err := rand.Read(randByte)

	if err != nil {
		return &response.Response{Code: http.StatusInternalServerError, Params: nil}
	}

	// The slice should now contain random bytes instead of only zeroes.
	token := hex.EncodeToString(randByte)
	session.Set("csrf_token", token)
	err = session.Save()

	if err != nil {
		return &response.Response{Code: http.StatusInternalServerError, Params: nil}
	}

	return &response.Response{
		Code: http.StatusOK,
		Params: map[string]string{
			"token": token,
		},
	}
}
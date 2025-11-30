package userhandle

import "github.com/gin-gonic/gin"

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Token string `json:"token"`
}

func Login(ctx *gin.Context, req LoginRequest) LoginResponse {
	var resp LoginResponse
	return resp
}

package v1

import (
	"net/http"

	"github.com/begenov/register-service/pb"
	"github.com/gin-gonic/gin"
)

func (h *Handler) registerLoadRouter(api *gin.RouterGroup) {
	register := api.Group("/registers")
	{
		register.POST("/sign-up", h.signUp)
		register.POST("/sign-in", h.signIn)
		register.POST("/refresh-token", h.refreshToken)
	}
}

type signUpRequest struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

func (h *Handler) signUp(ctx *gin.Context) {
	var inpHttp signUpRequest

	err := ctx.BindJSON(&inpHttp)
	if err != nil {
		h.newResponseError(ctx, http.StatusBadRequest, "invalid input")
		return
	}

	// TODO check inpHttp

	inpRpc := &pb.RequestRegister{
		Email:    inpHttp.Email,
		Phone:    inpHttp.Phone,
		Role:     inpHttp.Role,
		Address:  inpHttp.Address,
		Password: inpHttp.Password,
	}

	res, err := h.registerClient.SignUp(ctx, inpRpc)
	if err != nil {
		h.newResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := ResponseMessage{
		Message: res,
	}

	ctx.JSON(http.StatusOK, response)
}

type signInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (h *Handler) signIn(ctx *gin.Context) {
	var inpHttp signInRequest

	err := ctx.BindJSON(&inpHttp)
	if err != nil {
		h.newResponseError(ctx, http.StatusBadRequest, "invalid input")
		return
	}

	// TODO check inpHttp

	inpRpc := &pb.RequestSignIn{
		Email:    inpHttp.Email,
		Password: inpHttp.Password,
		Role:     inpHttp.Role,
	}

	res, err := h.registerClient.SignIn(ctx, inpRpc)
	if err != nil {
		h.newResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := ResponseMessage{
		Message: res,
	}

	ctx.JSON(http.StatusOK, response)
}

type requestRefreshToken struct {
	RefreshToken string `json:"refresh_token"`
	Role         string `json:"role"`
}

func (h *Handler) refreshToken(ctx *gin.Context) {
	var inpHttp requestRefreshToken

	err := ctx.BindJSON(&inpHttp)
	if err != nil {
		h.newResponseError(ctx, http.StatusBadRequest, "invalid input")
		return
	}

	// TODO check inpHttp

	inpRpc := &pb.RequestToken{
		RefreshToken: inpHttp.RefreshToken,
		Role:         inpHttp.Role,
	}
	res, err := h.registerClient.RefreshToken(ctx, inpRpc)
	if err != nil {
		h.newResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := ResponseMessage{
		Message: res,
	}

	ctx.JSON(http.StatusOK, response)
}

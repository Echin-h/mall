package ctl

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	Error   string      `json:"error"`
	TrackId string      `json:"track_id"`
}

type TrackedErrorResponse struct {
	Response
	TrackId string `json:"track_id"`
}

func RespSuccess(ctx *gin.Context, data interface{}, code ...int) *Response {
	trackId, _ := getTrackIdFromCtx(ctx)
}

func RespError(ctx *gin.Context, error error, msg string, code ...int) *Response {
	trackId, _ := getTrackIdFromCtx(ctx)
}
func getTrackIdFromCtx(ctx *gin.Context) (string, error) {
	trackId, ok := ctx.Get("track_id")
	if !ok {
		return "", errors.New("track_id not found")
	}
}

package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"time"
)

type errorResponse struct {
	Msg string `json:"err_msg"`
}

type okResponse struct {
	UnixTime int `json:"unix_time"`
}

func main() {
	ctx := context.Background()
	e := echo.New()
	e.GET("/unix-time", unixTimeHandler)

	err := e.Start(":8080")
	slog.InfoContext(ctx, "start simple server", slog.String("err", err.Error()))
}

func unixTimeHandler(c echo.Context) error {
	d := c.QueryParam("date")
	t := c.QueryParam("time")

	dateTime, err := time.ParseInLocation(time.DateTime, fmt.Sprintf("%s %s", d, t), time.Local)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, okResponse{
		UnixTime: int(dateTime.Unix()),
	})
}

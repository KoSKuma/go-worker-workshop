package handler

import (
	"net/http"

	"github.com/koskuma/go-worker-workshop/pkg/adapter"
	"github.com/koskuma/go-worker-workshop/pkg/log"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type IProbe interface {
	DBReadyCheck(c echo.Context) error
}

type probe struct {
	mongoDBAdapter adapter.IMongoDBAdapter
	logger         log.ILogger
}

func NewProbe(mongoDBAdapter adapter.IMongoDBAdapter, logger log.ILogger) IProbe {
	return &probe{
		mongoDBAdapter,
		logger,
	}
}

func (h probe) DBReadyCheck(c echo.Context) error {
	err := h.mongoDBAdapter.Ping(c.Request().Context(), readpref.Primary())
	if err != nil {
		return c.NoContent(http.StatusServiceUnavailable)
	}
	return c.NoContent(http.StatusOK)
}

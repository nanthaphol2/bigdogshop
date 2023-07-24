package monitorHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nanthaphol2/bigdogshop/config"
	"github.com/nanthaphol2/bigdogshop/modules/entities"
	"github.com/nanthaphol2/bigdogshop/modules/monitor"
)

type IMontitorHandler interface {
	HealthCheck(c *fiber.Ctx) error
}

type monitorHandler struct {
	cfg config.IConfig
}

func MonitorHandler(cfg config.IConfig) IMontitorHandler {
	return &monitorHandler{
		cfg: cfg,
	}
}

func (h *monitorHandler) HealthCheck(c *fiber.Ctx) error {
	res := &monitor.Monitor{
		Name:    h.cfg.App().Name(),
		Version: h.cfg.App().Version(),
	}
	// return c.Status(fiber.StatusOK).JSON(res)
	return entities.NewResponse(c).Success(fiber.StatusOK, res).Res()
}

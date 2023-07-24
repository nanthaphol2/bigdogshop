package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nanthaphol2/bigdogshop/modules/middelwares/middlewaresHandlers"
	"github.com/nanthaphol2/bigdogshop/modules/middelwares/middlewaresRepositories"
	"github.com/nanthaphol2/bigdogshop/modules/middelwares/middlewaresUsecases"
	"github.com/nanthaphol2/bigdogshop/modules/monitor/monitorHandlers"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	r   fiber.Router
	s   *server
	mid middlewaresHandlers.IMiddlewaresHandler
}

func InitModule(r fiber.Router, s *server, mid middlewaresHandlers.IMiddlewaresHandler) IModuleFactory {
	return &moduleFactory{
		r:   r,
		s:   s,
		mid: mid,
	}
}

func InitMiddlewares(s *server) middlewaresHandlers.IMiddlewaresHandler {
	repository := middlewaresRepositories.MiddlewaresRepository(s.db)
	usecase := middlewaresUsecases.MiddlewaresUsecase(repository)
	return middlewaresHandlers.MiddlewaresHandler(s.cfg, usecase)
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.s.cfg)

	m.r.Get("/", handler.HealthCheck)
}

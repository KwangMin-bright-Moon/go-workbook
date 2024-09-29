package app

import (
	"eCommerce/config"
	"eCommerce/repository"
	"eCommerce/router"
	"eCommerce/service"
)

type App struct {
	config  *config.Config

	router  *router.Router
	repository *repository.Repository
	service *service.Service
}

func NewApp(config *config.Config)*App {
	a := &App{
		config: config,
	}

	var err error

	if a.repository, err = repository.NewRepository(config); err != nil {
		panic(err)
	}

	if a.service, err = service.NewService(config); err != nil {
		panic(err)
	}

	if a.router, err = router.NewRouter(config); err != nil {
		panic(err)
	}

	return a

}

func (a *App) Run() error {
	return a.router.Run()
}
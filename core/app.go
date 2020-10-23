package core

import (
	"github.com/faruoqi/elemental/api"
	"github.com/faruoqi/elemental/config"
)

type App struct {
	/* @todo db interface type */
	api api.APIHandler
	/* @todo logs interface type */
	appConf config.Config
}

func NewApp(env string) *App {

	//init api engine
	var apiEngine api.APIHandler
	apiEngine = api.NewAPIHandler()

	var appConf config.Config
	appConf = config.NewConfig(env)

	return &App{api: apiEngine, appConf: appConf}

}

func (a *App) Start() {
	//initialize functions
	a.init()

	//run http handler
	a.api.Run(a.appConf.GetConfig("api_port"))
}

func (a *App) init() {
	//init router api
	a.api.RouteInit()

	//init config
	a.appConf.Init()
}

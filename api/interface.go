package api

type APIHandler interface {
	Run(port string)
	RouteInit()
	Info() string
}

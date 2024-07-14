package ioc

type Controller struct {
	*MapController
}

func NewController() *Controller {
	var con Controller
	con.MapController = newMapController()
	con.name = "controller"
	return &con
}

var ControllerImpl = NewController()

func NewApiController() *Controller {
	var con Controller
	con.MapController = newMapController()
	con.name = "api"
	return &con
}

var Api = NewApiController()

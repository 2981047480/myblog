package ioc

type Controller struct {
	*MapController
}

func NewController() *Controller {
	var con Controller
	con.MapController = newMapController()
	return &con
}

var ControllerImpl = NewController()

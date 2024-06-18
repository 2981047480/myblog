package ioc

import (
	"fmt"
	"log"
)

func newMapController() *MapController {
	return &MapController{
		storage: make(map[string]Object),
	}
}

type MapController struct {
	storage map[string]Object
}

func (i *MapController) Register(name string, obj Object) {
	i.storage[name] = obj
}

func (i *MapController) Get(name string) any {
	return i.storage[name]
}

func (i *MapController) Init() error {
	for name, obj := range i.storage {
		err := obj.Init()
		if err != nil {
			return fmt.Errorf("%s init ERROR, error:\n %v", name, err)
		}
	}
	log.Println("Congradulactions, all model init done!")
	return nil
}

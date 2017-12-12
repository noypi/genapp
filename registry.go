package genapp

import (
	"github.com/noypi/router"
)

type Registry struct {
	h map[interface{}]WrappedHandler
	o map[interface{}]interface{}
}

var registry = Registry{
	h: map[interface{}]WrappedHandler{},
	o: map[interface{}]interface{}{},
}

func RegisterH(k interface{}, h router.Handler) {
	registry.SetHandler(k, h)
}

func RegisterO(k, o interface{}) {
	registry.SetObject(k, o)
}

func GetH(k interface{}) WrappedHandler {
	return registry.GetHandler(k)
}

func GetO(k interface{}) interface{} {
	return registry.GetObject(k)
}

func (r Registry) GetHandler(k interface{}) WrappedHandler {
	if h, has := r.h[k]; has {
		return h
	}
	return nopHandler
}

func (r *Registry) GetObject(k interface{}) interface{} {
	if o, has := r.o[k]; has {
		return o
	}
	return struct{}{}
}

func (r *Registry) SetHandler(k interface{}, h router.Handler) {
	r.h[k] = WrappedHandler(router.GinWrapC(h))
}

func (r *Registry) SetObject(k, o interface{}) {
	r.o[k] = o
}

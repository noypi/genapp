package genapp

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noypi/router"
)

func init() {
	router.DisableDebugging()
}

type App struct {
	mux *router.EngineStd
}

type Result struct {
	Out *bytes.Buffer
}

func New() *App {
	return &App{
		mux: router.New(),
	}
}

func (app *App) Execute(name string) (*Result, error) {
	req, err := http.NewRequest("PATCH", name, nil)
	if nil != err {
		return nil, err
	}
	writer := newWriter()
	app.mux.ServeHTTP(writer, req)
	return &Result{Out: writer.buf}, nil
}

func (app *App) Handle(name string, seq Sequence) {
	hs := make([]interface{}, len(seq))
	for i, s := range seq {
		hs[i] = gin.HandlerFunc(registry.GetHandler(s))
	}

	app.mux.Any(name, hs...)
}

func EnableDebugging() {
	router.EnableDebugging()
}

package main

import (
	"context"
	"fmt"

	"log"

	"github.com/noypi/genapp"
	"github.com/noypi/router"
)

func main() {

	app := genapp.New()

	genapp.RegisterH("hello", func(c context.Context) {
		log.Println("hello")
	})

	genapp.RegisterH("hi", func(c context.Context) {
		w := router.GetWriter(c)
		w.Write([]byte("some hi"))
		log.Println("hi")
	})

	app.Handle("/", genapp.H("hello", "hi"))

	res, err := app.Execute("/")
	if nil != err {
		log.Fatal(err)
	}

	fmt.Println("result=", res.Out.String())
}

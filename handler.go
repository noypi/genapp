package genapp

import (
	"github.com/gin-gonic/gin"
)

type WrappedHandler gin.HandlerFunc

var nopHandler = func(*gin.Context) {}

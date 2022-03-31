package session

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

//go:generate genny   -pkg $GOPACKAGE     -in=../../../../genny/sarulabsdi/func-types.go -out=gen-func-$GOFILE gen "FuncType=GetSession,GetSessionStore"

type (
	// GetSession ...
	GetSession func(c echo.Context) *sessions.Session
	// GetSessionStore ...
	GetSessionStore func() sessions.Store
)

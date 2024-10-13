package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/pauldin91/GoWebApp/internal/cfg"
)

func TestRoutes(t *testing.T) {
	var app cfg.AppConfig

	mux := routes(&app)
	switch v := mux.(type) {
	case *chi.Mux:
	default:
		t.Error(fmt.Sprintf("type is %T and not type *chi.Mux", v))
	}
}

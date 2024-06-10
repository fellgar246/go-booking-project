package main

import (
	"fmt"
	"testing"

	"github.com/fellgar246/go-booking-project/cmd/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	// t.Error("not implemented")
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing; test passed
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, but is %T", v))
	}
}

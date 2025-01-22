package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/choonhong/hotel-data-merge/adapter"
	"github.com/choonhong/hotel-data-merge/database"
	"github.com/choonhong/hotel-data-merge/internal"
	"github.com/choonhong/hotel-data-merge/provider"
	"github.com/choonhong/hotel-data-merge/restapi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	ctx := context.Background()

	// connect to db
	db, err := database.Connect()
	if err != nil {
		log.Fatal("failed to connect to db:", err)
	}

	// create hotel service
	service := internal.HotelService{
		HotelRepo: &adapter.HotelRepository{Client: db},
		Providers: []internal.Provider{
			&provider.Acme{URL: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/acme"},
			&provider.Paperflies{URL: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/paperflies"},
			&provider.Patagonia{URL: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia"},
		},
	}

	// fetch hotel data on startup
	if _, err := service.FetchAndMergeHotels(ctx); err != nil {
		log.Println("failed to fetch and merge hotels:", err)
	}

	// fetch hotel data every 15 minutes
	ticker := time.NewTicker(15 * time.Minute)
	go func() {
		for range ticker.C {
			if _, err := service.FetchAndMergeHotels(ctx); err != nil {
				log.Println("failed to fetch and merge hotels:", err)
			}
		}
	}()

	// start http server
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/v1", func(r chi.Router) {
		restapi.HandlerFromMux(&service, r)
	})

	// for deploying on heroku
	if http.ListenAndServe(":80", r) != nil {
		http.ListenAndServe(":"+os.Getenv("PORT"), r) //nolint:errcheck
	}
}

package main

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
)

func main() {
	ctx := context.Background()
	inmemRepo := repository.NewInmemRepository()
	svc := service.NewTripService(inmemRepo)

	// Example usage of the tripService
	// You can create a new trip and handle it as needed

	fare := &domain.RideFareModel{
		UserId: "user123",
		// Set other fare attributes as needed
	}

	trip, err := svc.CreateTrip(ctx, fare)
	if err != nil {
		log.Println(err)
	}
	log.Println(trip)

	for {
		time.Sleep(time.Second)
	}
}

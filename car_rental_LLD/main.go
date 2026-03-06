package main

import (
	"car_rental_LLD/models"
	"car_rental_LLD/strategy"
	"fmt"
)

func main() {
	fmt.Println("--- Starting Car Rental LLD Dry Run ---")
	
	// 1. Initialize the Reservation Manager
	rm := models.NewReservationManager()

	// 2. Initialize and add Users
	user1 := models.NewUser(1, "Alice", "alice@example.com")
	user2 := models.NewUser(2, "Bob", "bob@example.com")
	rm.AddUser(user1)
	rm.AddUser(user2)
	fmt.Println("- Added users: Alice, Bob")

	// 3. Initialize and add Vehicles
	car1 := models.NewCar("CAR-1234", true, "X5", "BMW", 5000.0, models.SUV)
	car2 := models.NewCar("CAR-5678", true, "City", "Honda", 2000.0, models.SEDAN)
	bike1 := models.NewBike("BIKE-001", true, "R15", "Yamaha", 1000.0, models.BIKE1)
	
	rm.AddVehicle(car1)
	rm.AddVehicle(car2)
	rm.AddVehicle(bike1)
	fmt.Println("- Added vehicles: BMW X5 (SUV), Honda City (Sedan), Yamaha R15 (Bike)")

	// 4. Book a Reservation for Alice (BMW SUV for 4 days)
	fmt.Println("\n--- Booking a SUV for Alice ---")
	res1 := rm.BookReservation(models.CAR, models.SUV, user1, 1, 5)
	if res1 != nil {
		fmt.Printf("Success! Booking ID: %d\n", res1.GetId())
	} else {
		fmt.Println("Failure! Vehicle unavailable.")
	}

	// 5. Check if another user can book the same SUV type while it's booked
	fmt.Println("\n--- Booking a SUV for Bob (Should fail if no other SUV is available) ---")
	res2 := rm.BookReservation(models.CAR, models.SUV, user2, 2, 6)
	if res2 != nil {
		fmt.Printf("Success! Booking ID: %d\n", res2.GetId())
	} else {
		fmt.Println("Failure! SUV is currently unavailable.")
	}

	// 6. Pay for Alice's reservation
	if res1 != nil {
		fmt.Println("\n--- Paying for Alice's Reservation ---")
		upiProcessor := &strategy.Upi{}
		success := rm.Payprice(res1.GetId(), upiProcessor)
		if success {
			fmt.Println("Payment successful! The vehicle is now free.")
		} else {
			fmt.Println("Payment failed!")
		}
	}

	// 7. Check if Bob can now book the SUV since it has been freed
	fmt.Println("\n--- Booking a SUV for Bob again (Should succeed now) ---")
	res3 := rm.BookReservation(models.CAR, models.SUV, user2, 5, 8)
	if res3 != nil {
		fmt.Printf("Success! Booking ID: %d\n", res3.GetId())
	} else {
		fmt.Println("Failure! Vehicle unavailable.")
	}
	
	fmt.Println("\n--- Dry Run Completed ---")
}

package main

import (
	"fmt"
	"log"

	"github.com/dimka90/powertoken/internal"
)

func main() {
	fmt.Println("=== PowerToken Electricity Billing System ===\n")

	// Create billing system
	bs := internal.NewBillingSystem()

	// Register meters
	fmt.Println("Registering meters...")
	meter1 := internal.NewPrepaidMeter(
		"0101240273561",
		"MADUGU MANGAI JAMES",
		"OPP SENIOR STAFF QUATERS RING ROAD",
		223.88,
	)

	meter2 := internal.NewPrepaidMeter(
		"0101240273562",
		"DIMKA GODWIN",
		"JOS PLATEAU STATE",
		250.00,
	)

	meter3 := internal.NewPostpaidMeter(
		"0101240273563",
		"JOHN DOE",
		"BUKURU TOWN",
		0,
		200.00,
	)

	bs.RegisterMeter(meter1)
	bs.RegisterMeter(meter2)
	bs.RegisterMeter(meter3)
	fmt.Println("✓ 3 meters registered\n")

	// Process payments
	fmt.Println("=== Processing Payments ===\n")

	// Payment 1
	fmt.Printf("Meter Balance Before: %.2f kWh\n", meter1.GetBalance())
	transaction1, err := bs.ProcessPayment("0101240273561", 3000, "OWealth")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(transaction1)
	fmt.Printf("Meter Balance After: %.2f kWh\n\n", meter1.GetBalance())

	// Payment 2
	fmt.Printf("Meter Balance Before: %.2f kWh\n", meter2.GetBalance())
	transaction2, err := bs.ProcessPayment("0101240273562", 5000, "Bank Transfer")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(transaction2)
	fmt.Printf("Meter Balance After: %.2f kWh\n\n", meter2.GetBalance())

	// Try invalid payment (too small)
	fmt.Println("=== Testing Error Handling ===\n")
	_, err = bs.ProcessPayment("0101240273561", 50, "Cash")
	if err != nil {
		fmt.Printf("❌ Error (expected): %v\n\n", err)
	}

	// List all transactions
	fmt.Println("=== All Transactions ===\n")
	err = bs.ListTransactions()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Save to file
	fmt.Println("\n=== Saving to File ===")
	err = bs.SaveTransactionsToFile("transactions.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("✓ Transactions saved to transactions.json")

	fmt.Println("\n=== System Complete ===")
}

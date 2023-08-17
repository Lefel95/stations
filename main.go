package main

import (
	"fmt"
	"log"
	"os"
	"stations/v2/models"
)

func main() {

	printSuccess := func(success bool, successMessage, errorMessage string) {
		if success {
			fmt.Println(successMessage)
		} else {
			fmt.Println(errorMessage)
		}
	}
	p := &models.Parking{}

	fillSlots := func(vtype string, nn int) {
		for i := 0; i < nn; i++ {
			var n = &models.Slot{
				Prior:   nil,
				Type:    vtype,
				Used:    false,
				Vehicle: nil,
				Next:    nil,
			}
			p.AddParkingSlot(n)
		}
	}

	fillSlots(models.BIKE, 10)
	fillSlots(models.CAR, 10)
	fillSlots(models.VAN, 10)

	var n = -1
	fmt.Println("Hello! We have 10 slots for each type!")
	fmt.Println("Type:")
	for n != 0 {

		fmt.Println("1. To Park a Bike")
		fmt.Println("2. To Park a Car")
		fmt.Println("3. To Park a Van")
		fmt.Println("4. To release a bike")
		fmt.Println("5. To release a car")
		fmt.Println("6. To release a van")
		fmt.Println("7. To See if the parking is full")
		fmt.Println("8. To See if the parking is empty")
		fmt.Println("9. To See if all parking slots for bikes are occupied")
		fmt.Println("10. To See if all parking slots for cars are occupied")
		fmt.Println("11. To See if all parking slots for vans are occupied")
		fmt.Println("12. To See how many parking slots vans are occupied")
		fmt.Println("0. To exit")

		_, err := fmt.Scanln(&n)
		if err != nil {
			log.Fatal(err)
		}

		switch n {
		case 1:
			v := &models.Vehicle{
				MyType: models.BIKE,
			}
			success := p.Park(v)

			printSuccess(success, "Vehicle parked successfully", "Error Parking Vehicle, please see if the parking slot is full")
		case 2:
			v := &models.Vehicle{
				MyType: models.CAR,
			}
			success := p.Park(v)
			printSuccess(success, "Vehicle parked successfully", "Error Parking Vehicle, please see if the parking slot is full")
		case 3:
			v := &models.Vehicle{
				MyType: models.VAN,
			}
			success := p.Park(v)
			printSuccess(success, "Vehicle parked successfully", "Error Parking Vehicle, please see if the parking slot is full")
		case 4:
			success := p.Release(models.BIKE)
			printSuccess(success, "Vehicle released successfully", "Error Releasing Vehicle, please see if the any vehicle of that type is in there")
		case 5:
			success := p.Release(models.CAR)
			printSuccess(success, "Vehicle released successfully", "Error Releasing Vehicle, please see if the any vehicle of that type is in there")
		case 6:
			success := p.Release(models.VAN)
			printSuccess(success, "Vehicle released successfully", "Error Releasing Vehicle, please see if the any vehicle of that type is in there")
		case 7:
			isFull := p.Full()
			printSuccess(isFull, "The Parking Slot is full", "The Parking slot is not full")
		case 8:
			isEmpty := p.Empty()
			printSuccess(isEmpty, "The Parking Slot is empty", "The Parking slot is not empty")
		case 9:
			countBike := p.Count(models.BIKE)
			printSuccess(countBike == 10, "All Bike slots are occupied", "Not All bike slots are occupied")
		case 10:
			countCar := p.Count(models.CAR)
			printSuccess(countCar == 10, "All Car slots are occupied", "Not All Car slots are occupied")
		case 11:
			countVan := p.Count(models.VAN)
			printSuccess(countVan == 10, "All Van slots are occupied", "Not All Van slots are occupied")
		case 12:
			countVan := p.CountByVehicleType(models.VAN)
			printSuccess(true, fmt.Sprintf("There are %d slot occupied by vans", countVan), "")

		case 0:
			fmt.Println("Ok Bye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid Code, please give another!")
		}
	}

}

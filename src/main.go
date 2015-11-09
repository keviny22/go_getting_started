package main

import (
	"fmt"
	"strings"
)

func main() {
	plants := []PowerPlant{
		PowerPlant{hydro, 300, active},
		PowerPlant{wind, 30, active},
		PowerPlant{wind, 25, inactive},
		PowerPlant{wind, 35, active},
		PowerPlant{solar, 45, unavailable},
		PowerPlant{solar, 40, inactive},
	}

	grid := PowerGrid{300, plants}

	//plantCapacities := []float64{30, 30, 30, 60, 60, 100}
	//activePlants := []int{0, 1}
	//gridLoad := 75.
	fmt.Println("1) Generate power plant report")
	fmt.Println("2) Generate power grid report")
	fmt.Print("Please choose and option: ")

	var option string
	fmt.Scanln(&option)
	switch option {
	case "1":
		//generatePlantCapacityReport(plantCapacities...)
		grid.generatePlantreport()
	case "2":
		grid.generateGridReport()
		//generatePowerGridReport(activePlants, plantCapacities, gridLoad)
	default:
		fmt.Println("Unknown option")
	}
}

func generatePlantCapacityReport(plantCapacities ...float64) {
	for i, cap := range plantCapacities {
		fmt.Printf("Plant %d capacity: %.0f\n", i, cap)
	}
}

func generatePowerGridReport(activePlants []int, plantCapacities []float64, gridLoad float64) {
	capacity := 0.
	for _, plantId := range activePlants {
		capacity += plantCapacities[plantId]
	}

	fmt.Printf("%-20s%.0f\n", "Capacity: ", capacity)
	fmt.Printf("%-20s%.0f\n", "Load: ", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utilization: ", gridLoad/capacity*100)
}

type PlantType string

const (
	hydro PlantType = "Hydro"
	wind  PlantType = "Wind"
	solar PlantType = "Solar"
)

type PlantStatus string

const (
	active      PlantStatus = "Active"
	inactive    PlantStatus = "Inactive"
	unavailable PlantStatus = "Unavialable"
)

type PowerPlant struct {
	plantType PlantType
	capacity  float64
	status    PlantStatus
}

type PowerGrid struct {
	load   float64
	plants []PowerPlant
}

func (pg *PowerGrid) generatePlantreport() {
	for i, p := range pg.plants {
		label := fmt.Sprintf("%s%d", "Plant #", i)
		fmt.Println(label)
		fmt.Println(strings.Repeat("-", len(label)))
		fmt.Printf("%-20s%s\n", "Type: ", p.plantType)
		fmt.Printf("%-20s%.0f\n", "Capacity: ", p.capacity)
		fmt.Printf("%-20s%s\n", "Load: ", p.status)
		fmt.Printf("")
	}
}

func (pg *PowerGrid) generateGridReport() {
	capacity := 0.
	for _, p := range pg.plants {
		if p.status == active {
			capacity += p.capacity
		}
	}

	label := "Power Grid Report"
	fmt.Println(label)
	fmt.Println(strings.Repeat("-", len(label)))
	fmt.Printf("%-20s%.0f\n", "Capacity: ", capacity)
	fmt.Printf("%-20s%.0f\n", "Load: ", pg.load)
	fmt.Printf("%-20s%.2f%%\n", "Utilization: ", pg.load/capacity*100)

}

//
//func main() {
//	var plantCapacities []float64
//
//	plantCapacities = []float64{30, 30, 30, 60, 60, 100}
//
//	var capacity float64 = plantCapacities[0] + plantCapacities[1] +
//		plantCapacities[2] + plantCapacities[3] + plantCapacities[4] +
//		plantCapacities[5]
//
//	var gridLoad = 75.
//
//	utilization := gridLoad / capacity
//	fmt.Printf("%-20s%.0f\n","Capacity: ", capacity)
//	fmt.Printf("%-20s%.0f\n","Load: ", gridLoad)
//	fmt.Printf("%-20s%.1f%%\n","Utilization: ", utilization * 100)
//
//}

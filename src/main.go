package main

import "fmt"

func main() {

	plantCapacities := []float64{30, 30, 30, 60, 60, 100}
	activePlants := []int{0, 1}
	gridLoad := 75.
	fmt.Println("1) Generate power plant report")
	fmt.Println("2) Generate power grid report")
	fmt.Print("Please choose and option: ")

	var option string
	fmt.Scanln(&option)
	switch option {
	case "1":
		for i, cap := range plantCapacities {
			fmt.Printf("Plant %d capacity: %.0f\n", i, cap)
		}
	case "2":
		capacity := 0.
		for _, plantId := range activePlants {
			capacity += plantCapacities[plantId]
		}


		fmt.Printf("%-20s%.0f\n", "Capacity: ", capacity)
		fmt.Printf("%-20s%.0f\n", "Load: ", gridLoad)
		fmt.Printf("%-20s%.1f%%\n", "Utilization: ", gridLoad / capacity * 100)
	default:
		fmt.Println("Unknown option")
	}

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

package lasagna

import "fmt"

const OvenTime = 40

func RemainingOvenTime(actualMinutesInOven int) int {
    result := OvenTime - actualMinutesInOven
	return result
}

func PreparationTime(numberOfLayers int) int {
	
    result := numberOfLayers * 2
    return result
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	result := (numberOfLayers * 2) + actualMinutesInOven
    return result
}

func main(){
    fmt.Println(RemainingOvenTime(30))
    fmt.Println(PreparationTime(2))
    fmt.Println(ElapsedTime(3,20))
    
}
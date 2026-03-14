package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minutesPerLayer int) int{

    if minutesPerLayer == 0{
        minutesPerLayer = 2
    }
    
    return len(layers) * minutesPerLayer
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64){
    noodlesCount := 0
    layersOfNoodlesNeeded := 50

    sauceCount := 0.0
    layersOfSauceNeeded := 0.2

    for _, value := range layers{
        if value == "noodles"{
        	noodlesCount += 1 * layersOfNoodlesNeeded
		}
        if value == "sauce"{
        	sauceCount += 1 * layersOfSauceNeeded
        }
    }
    return noodlesCount, sauceCount
}
// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(secretIngredients []string, ingredients []string){

    secretIngredient := secretIngredients[len(secretIngredients)-1]

    ingredients[len(ingredients)-1] = secretIngredient
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, numberOfPortions int) []float64{

    portion := float64(numberOfPortions) / 2.0
    
    portions := []float64{}

    for _, value := range amounts{
		desiredPortions := value * portion
        portions = append(portions, desiredPortions)
    }

    return portions
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

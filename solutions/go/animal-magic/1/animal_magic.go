package chance
import "math/rand"

//n := a + rand.Intn(b-a+1) // a ≤ n ≤ b

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	d := 1 + rand.Intn(20-1+1)
    return d
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	f := rand.Float64() * 12.0
    return f
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{"ant","beaver","cat","dog","elephant","fox","giraffe","hedgehog"}
    swap := func(i,j int){
        animals[i],animals[j] = animals[j], animals[i]
    }
    rand.Shuffle(len(animals),swap)
    return animals
}

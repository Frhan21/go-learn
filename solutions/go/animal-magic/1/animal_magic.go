package chance
import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	a := 1 
    b := 20 
    num := a + rand.Intn(b-a+1)
    return num
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	num := rand.Float64() * 12
    return num
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})

	return animals
}

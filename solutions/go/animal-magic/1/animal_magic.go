package chance
import (
    "math/rand"
    "time"
    )
// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
    return 1 + rand.Intn(20)
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
    f := rand.Float64() * 12.0
    return f
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	// Initialize the eight animals
	animals := []string{
		"ant", "beaver", "cat", "dog", 
		"elephant", "fox", "giraffe", "hedgehog",
	}
	
	// Create a new random source with current time as seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	
	for i := len(animals) - 1; i > 0; i-- {
		j := r.Intn(i + 1) // Random index from 0 to i
		animals[i], animals[j] = animals[j], animals[i] // Swap
	}
	
	return animals
}


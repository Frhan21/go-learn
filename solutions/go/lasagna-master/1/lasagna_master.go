package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
    if timePerLayer == 0 {
        timePerLayer = 2 
    }

    estimatedTime := len(layers) * timePerLayer
    return estimatedTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    noodles := 0 
    sauces := 0.0 

    for _, layer := range layers {
        switch layer {
            case "noodles": 
            	noodles += 50
            case "sauce": 
            	sauces += 0.2
        }
    }

    return noodles, sauces
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string){
    myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {

    factor := float64(portions) / 2

    results := make([]float64, len(amounts))

    for i, amount := range amounts {
        results[i] = amount * factor
    }

    
    return results
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

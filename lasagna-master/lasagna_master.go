package lasagnam

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTime int) int {
  if avgTime == 0 {
		avgTime = 2
	}
	return len(layers) * avgTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
  for _, layer := range layers {
    if layer == "noodles" {
      noodles += 50
    } else if layer == "sauce" {
      sauce += 0.2
    }
  }
  return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
  myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amtFor2portions []float64, portions int) []float64 {
  scaledAmounts := make([]float64, len(amtFor2portions))
  for i, amnt2 := range amtFor2portions {
    scaledAmounts[i] = float64(portions) * amnt2 / 2
  }
  return scaledAmounts 
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.

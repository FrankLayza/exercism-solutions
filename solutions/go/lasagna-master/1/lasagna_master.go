package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int{
    if time == 0{
        time = 2
    }
    return len(layers) * time
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string)(int, float64){
    var amountOfNoodles int
    var amountOfSauce float64
    for i := 0; i < len(layers); i ++ {
        if(layers[i] == "sauce"){
            amountOfSauce += 0.2
        }
        if(layers[i] == "noodles"){
            amountOfNoodles += 50
        }
    }
    return amountOfNoodles, amountOfSauce
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
    secretIngredient := friendsList[len(friendsList)-1]
    myList[len(myList)-1] = secretIngredient
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numberOfPortions int) []float64{
	newQuantities := make([]float64, len(quantities))
    factor := float64(numberOfPortions)/2
    for i := 0; i < len(newQuantities); i++ {
        newQuantities[i] = quantities[i]
        newQuantities[i] *= factor
    }
    return newQuantities
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

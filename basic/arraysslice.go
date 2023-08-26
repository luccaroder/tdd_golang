package tddgo

func Sum(numbers [5]int) int {
	// for i := 0; i < 5; i++ {
	//     soma += numeros[i]
	// }
	// return soma
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

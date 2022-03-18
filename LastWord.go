package main

func main() {
	//fmt.Println(division(6, 2))
	//fmt.Println(convert(6))
	//fmt.Println(nextNumber(6))
	//fmt.Println(triangle(3, 3))
	//fmt.Println(powerCalculator(230, 10))
	//fmt.Println(firstValue(230, 10, 34))
	//fmt.Println(ageConversion(65))
	//fmt.Println(sumPolygon(6))
	//fmt.Println(nameString("Mubashir" + "Edabit"))

}

func division(x, y int) int {

	return x / y
}

func convert(a int) int {
	return a * 60
}
func nextNumber(a int) int {
	return a + 1
}
func triangle(a, b float32) float32 {
	return (a * b) / 2
}
func powerCalculator(a, b int) int {
	return a * b
}
func firstValue(a, b, c int) int {
	return a
}
func ageConversion(a int) int {
	return a * 365
}

func sumPolygon(a int) int {
	return (a - 2) * 180
}
func nameString(a string) string {
	return "Mubashir" + "Edabit"
}

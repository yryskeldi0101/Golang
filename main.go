package main
import "fmt"
import "math"
func main() {
	var a float64
	var b float64
	var c float64

fmt.Println("Решите квадратное уровнение")
fmt.Println("Введите a: ")
fmt.Scan(&a)
fmt.Println("Введите b:")
fmt.Scan(&b)
fmt.Println("Введите c:")
fmt.Scan(&c)
D := (b*b) - 4*(a*c)
if D > 0 {
	var x1 float64
	var x2 float64
	x1 = (-b + math.Sqrt(D)) / (2*a)
	x2 = (-b - math.Sqrt(D)) / (2*a)
	fmt.Println("Ваше уровнение имеет 2 корня\nD = " + fmt.Sprint(D))
	fmt.Println("X1: " + fmt.Sprint(x1) + "\nX2: " + fmt.Sprint(x2))
} else if D == 0 {
	var x float64
	x = (-b) / (2*a)
	fmt.Println("Ваше уровнение имеет 1 корнь\nD = 0")
	fmt.Println("X: " + fmt.Sprint(x))
} else if D < 0 {
	fmt.Println("Ваше уровнение не имеет корня\nD < 0\nD = " + fmt.Sprint(D))
}
}
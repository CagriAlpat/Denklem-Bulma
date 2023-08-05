package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf(" A sayisini giriniz = ")
	scanner.Scan()
	a, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Printf("B sayisini girinzi =  ")
	scanner.Scan()
	b, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Printf("C saysiini giriniz ")
	scanner.Scan()
	c, _ := strconv.ParseFloat(scanner.Text(), 64)

	delta := (math.Pow(b, 2)) - 4*a*c // pow üs alma fonksiyonu

	kokbir := (-b - (math.Sqrt(delta))) / 2 * a // sqrt karekök alma fonksiyonu

	kokiki := (-b + (math.Sqrt(delta))) / 2 * a

	fmt.Printf("1.kök= %f  \n   2.kök=  %f  \n   delta = %f", kokbir, kokiki, delta)

}

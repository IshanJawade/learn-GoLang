package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza App")
	fmt.Println("Please rate our Pizaa between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thansk for rating, ", input)

	// converting string into float
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// error handing boiler-plate code
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to the rating: ", numRating+1)
	}
}

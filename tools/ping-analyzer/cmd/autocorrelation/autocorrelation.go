package main

import (
	"fmt"

	"github.com/HT4w5/ping-lab/tools/ping-analyzer/pkg/ping_parser"
)

func main() {
	fmt.Print("k              Autocorrelation_replied    Autocorrelation_lost\n")
	fmt.Print("====================================================================\n")
	p1 := ping_parser.New("../../../../data/www.ntu.edu.sg-2025-08-01_10-06-34-ping.txt")
	for i := -10; i <= 10; i++ {
		fmt.Printf("%v              %v         %v\n", i, p1.AutocorrelationReplied(i), p1.AutocorrelationLost(i))
	}
	fmt.Print("====================================================================\n")
	fmt.Print("Delivery rate\n")
	fmt.Print("====================================================================\n")
	fmt.Printf("%v", p1.DeliveryRate())
}

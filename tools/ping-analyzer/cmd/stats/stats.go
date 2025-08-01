package main

import (
	"fmt"

	"github.com/HT4w5/ping-lab/tools/ping-analyzer/pkg/ping_parser"
)

func main() {
	fmt.Print("Host                      Deliver_rate       Longest_consecutive   Longest_burst_lost   Min_RTT   Max_RTT\n")
	fmt.Print("=============================================================================================================\n")
	p1 := ping_parser.New("../../../../data/41.186.255.86-2025-08-01_09-43-32-ping.txt")
	fmt.Printf("41.186.255.86             %v            %v          %v                    %v            %v\n", p1.DeliveryRate(), p1.LongestConsecutive(), p1.LongestLostBurst(), p1.MinRTT(), p1.MaxRTT())

	p2 := ping_parser.New("../../../../data/www.canterbury.ac.nz-2025-08-01_09-43-07-ping.txt")
	fmt.Printf("www.canterbury.ac.nz      %v            %v          %v                    %v            %v\n", p2.DeliveryRate(), p2.LongestConsecutive(), p2.LongestLostBurst(), p2.MinRTT(), p2.MaxRTT())

	p3 := ping_parser.New("../../../../data/www.ntu.edu.sg-2025-08-01_10-06-34-ping.txt")
	fmt.Printf("www.ntu.edu.sg            %v            %v          %v                    %v            %v\n", p3.DeliveryRate(), p3.LongestConsecutive(), p3.LongestLostBurst(), p3.MinRTT(), p3.MaxRTT())
}

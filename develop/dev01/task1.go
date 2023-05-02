package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"time"
)

func main() {
	ntpTime, err := ntp.Time("time.nist.gov")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("time with ntp: %v\n", ntpTime)
	fmt.Printf("time with time: %v\n", time.Now())
}

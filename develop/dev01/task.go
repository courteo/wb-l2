package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	
	NtpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org") // get time from documentation for example
	tmNow := time.Now()
	if err != nil {
		fmt.Printf("Time ger Error %v \n", err.Error())
		os.Exit(1)
	}
	fmt.Println("current time 	", tmNow)
	fmt.Println("Ntp time 	", NtpTime)
}
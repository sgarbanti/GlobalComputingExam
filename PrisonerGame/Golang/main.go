// Tommaso Sgarbanti - 1185058

package main

import (
	"fmt"
	"os"
	"strconv"
)
// Pcounter behaviour
func  main() {
	var n int = 0

	if len(os.Args) < 2 {
		for n == 0 {
			fmt.Printf( "Inserire il numero di prigionieri: \t")
			_, err := fmt.Scanf("%d", &n)
			if err != nil {
	    	    fmt.Println(err)
	        	os.Exit(2)
	        }
	    }
	} else {
		nTmp, err := strconv.Atoi(os.Args[1])
		if err != nil {
        	fmt.Println(err)
        	os.Exit(2)
    	}
    	n = nTmp
    }

    prisoners := make([]Ptic, 2*(n-1)-1)
	var prisonerCounter Pcounter

	// Bulb input channels
	var turnOFF = make(chan bool)
	var turnON = make(chan bool)
	var continue_ = make(chan bool)
	var finish = make(chan bool)

	// Prisoner input channels
	var enterON = make(chan bool)
	var enterOFF = make(chan bool)
	var exit = make(chan bool)

	pChannels := prisonerChannels {
		enterON,
		enterOFF,
		exit}

	// Prisoners initialization
	for i:=0; i < n-1; i++ {
		prisoners[i] = Ptic {
			i,
        	pChannels}
	}
	prisonerCounter =  Pcounter {
		"counter",
		pChannels}

	// Bulb initialization
	var bulb Bulb = Bulb {
    	[2]string{"LIGHTon", "LIGHToff"},
    	turnOFF,
    	turnON,
    	continue_,
    	finish}

    // Start concurrency
  	for i:=0; i < n-1; i++ {
    	go prisoners[i].Run(turnOFF, turnON, continue_, finish)
  	}
   	go prisonerCounter.Run(turnOFF, turnON, continue_, finish, n)

  	bulb.Run(enterON, enterOFF, exit)

}
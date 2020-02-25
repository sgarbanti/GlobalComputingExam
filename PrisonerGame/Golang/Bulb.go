// Tommaso Sgarbanti - 1185058

package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Bulb struct {
  states [2]string
  turnOFF chan bool
  turnON chan bool
  continue_ chan bool
  finish chan bool
}

// Bulb behaviour
func (self *Bulb) Run (enterON, enterOFF, exit chan bool) {
	// Inizializzazione casuale dello stato della luce
	rand.Seed(time.Now().UTC().UnixNano())
	rdn := rand.Intn(2)
	state := self.states[rdn]

	fmt.Printf("The initial state of the light is: %s\n", state)
	for {
		if state == "LIGHTon" {
		  	// LIGHTon
		  	enterON <- true
		  	select {
		    	case exit <- true:
		    		state = self.states[0]
		    	case  <- self.turnOFF:
		    		exit <- true
		      		select {
		        		case <- self.continue_:
		          			state = self.states[1]
		        		case <- self.finish:
		        			fmt.Printf("Game ended!\n")
		          			return // END
		      		}
		  	}
		} else {
		  	// LIGHToff
		  	enterOFF <- true
		  	select {
		    	case exit <- true:
		        	state = self.states[1]
		    	case  <- self.turnON:
		      		exit <- true
		      		state = self.states[0]
		  	}
		}
	}
}
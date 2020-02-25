// Tommaso Sgarbanti - 1185058

package main

import "fmt"

type prisonerChannels struct {
  enterON chan bool
  enterOFF chan bool
  exit chan bool
}

type Pcounter struct {
  name string
  prisonerChannels
}

type Ptic struct {
  name int
  prisonerChannels
}

// Pcounter behaviour
func (self *Pcounter) Run (turnOFF, turnON, continue_, finish chan bool, n int) {
  counter := 0
  for {
    select {
      case <- self.enterON:
        fmt.Printf( "Prisoner %s enters the room with the light on.\n", self.name)
        fmt.Printf( "Prisoner %s turns off the light for the %d time.\n", self.name, counter+1)
        turnOFF <- true
        fmt.Printf( "Prisoner %s exits the room.\n", self.name)
        <- self.exit
        counter = counter + 1
        if counter == 2*(n-1) {
          fmt.Printf( "Prisoner %s says that all the prisoners have already entered the room at least once.\n", self.name)
          finish <- true
          return
        } else {
          continue_ <- true
        }
      case  <- self.enterOFF:
        fmt.Printf( "Prisoner %s enters the room with the light off.\n", self.name)
        fmt.Printf( "Prisoner %s does nothing and exits the room.\n", self.name)
        <- self.exit
    }
  }
}

// Ptic behaviour
func (self *Ptic) Run (turnOFF, turnON, continue_, finish chan bool) {
  counter := 0
  for {
      select {
        case  <- self.enterON:
          fmt.Printf( "Prisoner %d enters the room with the light on.\n", self.name)
          fmt.Printf( "Prisoner %d does nothing and exits the room.\n", self.name)
          <- self.exit
        case <- self.enterOFF:
          fmt.Printf( "Prisoner %d enters the room with the light off.\n", self.name)
          if counter == 2 { // Ptic2
            fmt.Printf( "Prisoner %d does nothing and exits the room.\n", self.name)
            <- self.exit
          } else{
            fmt.Printf( "Prisoner %d turns on the light for the %d time.\n", self.name, counter+1)
            turnON <- true
            fmt.Printf( "Prisoner %d exits the room.\n", self.name)
            <- self.exit
            counter = counter + 1
          }
      }
  }
}
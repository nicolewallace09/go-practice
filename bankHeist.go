package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  // generating random numbers
  // seeding & using time method 
  rand.Seed(time.Now().UnixNano())

  // variable to keep track if the heist is sucessful or not
  isHeistOn := true 

  // variable containing our random number
  eludedGuards := rand.Intn(100)

  if eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    isHeistOn = false
    fmt.Println("Plan a better disguse next time?")
  }

  openVault := rand.Intn(100)

  if isHeistOn && openVault >= 70 {
    fmt.Println("Grab and GO!")
  } else if isHeistOn {
    isHeistOn = false
    fmt.Println("The vault can't be opened!")
  }

  leftSafely := rand.Intn(5)

  if isHeistOn {
    switch leftSafely {
      case 0: 
        isHeistOn = false
        fmt.Println("You failed, try again next time!")
      case 1: 
        isHeistOn = false
        fmt.Println("Bad luck this time!")
      case 2: 
        isHeistOn = false
        fmt.Println("Looks like you can't get the vault open!")
      case 3: 
        isHeistOn = false
        fmt.Println("Turns out vault doors don't open from the inside")
        default: 
        fmt.Println("Start the getaway car!")
    }
  }

  if isHeistOn {
    amtStolen := 10000 + rand.Intn(1000000)
    fmt.Println(amtStolen, "was taken!")
  }

  fmt.Println("isHeistOn is currently:", isHeistOn)
}


package main

import "fmt"

func main() {


    var conferenceName = "Go Conference"
    const conferenceTickets = 50
    var remainingTickets = 50
    
    
    fmt.Println("Welcome to %v booking application!", conferenceName)
    fmt.Println("We have a total of %v, tickets and %v, are still available \n", conferenceTickets,remainingTickets)
    fmt.Println("Get your tickets here to attend")
}
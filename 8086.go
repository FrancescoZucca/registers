package main


import (
)


// proc8086 is the center of everything.
// This is where we process the requests.
func proc8086() { 
	// TODO: process requests from receiveInstruc()
}

// receiveInstruc receives instructions.
// this is where we receive them,
// process them, then send them to the emulated
// 8086 processor
func receiveInstruc(channel <- chan string){

} 

// sendInstruc sends instructions to our emulated
// 8086 processor
func sendInstruc(channel <- chan string) {
	
}

// The main function
func main() {
}
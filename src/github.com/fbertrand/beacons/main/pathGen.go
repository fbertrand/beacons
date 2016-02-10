// beacons project main.go
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fbertrand/beacons/beacons"
)

func main() {
	fmt.Println("Simulation started")
	bcns := []beacons.Beacon{beacons.Beacon{Id: 1, Position: beacons.MakePosition(15, 10), Power: 1000},
		beacons.Beacon{Id: 2, Position: beacons.MakePosition(5, 15), Power: 1000}}
	indoor := &beacons.Room{bcns}
	user := beacons.MakeUser("user1", 0, 0, beacons.NORTH_EAST, indoor)
	//myLog := log.New(os.Stdout, "", 0)
	file, err := os.OpenFile("path.dat", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	go user.Run( /*os.Stdout*/ file, 100)
	time.Sleep(2000 * time.Millisecond)
	user.Link <- beacons.SOUTH_EAST
	time.Sleep(500 * time.Millisecond)
	user.Link <- beacons.SOUTH_WEST
	time.Sleep(1000 * time.Millisecond)
	user.Link <- beacons.NORTH_WEST
	time.Sleep(500 * time.Millisecond)
	user.Link <- beacons.NORTH
	time.Sleep(500 * time.Millisecond)
	user.Link <- beacons.WEST
	time.Sleep(500 * time.Millisecond)
	close(user.Link)
	fmt.Println("Simulation stopped")
}

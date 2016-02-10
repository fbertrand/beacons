// User.go
package beacons

import (
	"fmt"
	"io"
	"time"
)

type User struct {
	id        string
	position  Position
	direction Position
	Link      chan Command
	playgnd   *Room
}

// Cstr
func MakeUser(name string, x_init, y_init int, dir_init Command, env *Room) User {
	return User{id: name, position: Position{x_init, y_init},
		direction: directions[dir_init], Link: make(chan Command, 1), playgnd: env}
}

func (user *User) Run(log io.Writer, tempo time.Duration) {
	trace := true
	move := true
	for trace {
		distance := make([]float64, len(user.playgnd.Beacons))
		power := make([]float64, len(user.playgnd.Beacons))
		fmt.Fprintf(log, "%v %v ", user.position.x, user.position.y)
		for i, beacon := range user.playgnd.Beacons {
			distance[i] = comp_distance(user.position, beacon.Position)
			power[i] = beacon.signal_power(distance[i])
			fmt.Fprintf(log, "%v ", power[i])
		}
		if move {
			user.position.x += user.direction.x
			user.position.y += user.direction.y
			fmt.Fprintf(log, "%v %v\n", user.direction.x, user.direction.y)
		} else {
			fmt.Fprintf(log, "0 0\n")
		}
		time.Sleep(tempo * time.Millisecond)
		select {
		case msg, ok := <-user.Link:
			switch msg {
			case MOVE, WAIT:
				move = !move
			default:
				user.direction = directions[msg]
			}
			if !ok {
				trace = false // exit
			}
		default: // no msg
		}
	}
}

package robot

import (
	"fmt"
)

const (
	N Dir = iota
	S
	E
	W
)

func (d Dir) String() string {
	switch d {
	case N:
		return "N"
	case S:
		return "S"
	case E:
		return "E"
	case W:
		return "W"
	}
	return ""
}

func Advance() {
	fmt.Printf("%s advance (%d, %d) -> ", Step1Robot.Dir, Step1Robot.X, Step1Robot.Y)
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case E:
		Step1Robot.X++
	case W:
		Step1Robot.X--
	case S:
		Step1Robot.Y--
	}
	fmt.Printf("(%d, %d)\n", Step1Robot.X, Step1Robot.Y)
}

func Right() {
	fmt.Printf("%s turn right (%d, %d) -> ", Step1Robot.Dir, Step1Robot.X, Step1Robot.Y)
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = E
	case E:
		Step1Robot.Dir = S
	case W:
		Step1Robot.Dir = N
	case S:
		Step1Robot.Dir = W
	}
	fmt.Printf("(%d, %d)\n", Step1Robot.X, Step1Robot.Y)
}

func Left() {
	fmt.Printf("%s turn left (%d, %d) -> ", Step1Robot.Dir, Step1Robot.X, Step1Robot.Y)
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = W
	case E:
		Step1Robot.Dir = N
	case W:
		Step1Robot.Dir = S
	case S:
		Step1Robot.Dir = E
	}
	fmt.Printf("(%d, %d)\n", Step1Robot.X, Step1Robot.Y)
}

type Action func()

func StartRobot(cmdChan chan Command, actChan chan Action) {
	fmt.Println("start robot")
	for cmd := range cmdChan {
		//fmt.Println("receive command:", string(cmd))
		switch cmd {
		case Command('R'):
			actChan <- Right
		case Command('A'):
			actChan <- Advance
		case Command('L'):
			actChan <- Left
		}
	}
	//fmt.Println("closing actChan")
	close(actChan)
	//fmt.Println("quit robot")
}

func Room(extent Rect, robot Step2Robot, actChan chan Action, report chan Step2Robot) {
	fmt.Println("start room")
	Step1Robot.X = 0
	Step1Robot.Y = 0
	Step1Robot.Dir = robot.Dir

	for act := range actChan {
		//fmt.Println("\nrobot before", robot)
		//fmt.Printf("robot inside %s (%d, %d)\n", Step1Robot.Dir, Step1Robot.X, Step1Robot.Y)
		act()

		robot.Dir = Step1Robot.Dir
		robot.Easting += RU(Step1Robot.X)
		robot.Northing += RU(Step1Robot.Y)
		//fmt.Println("robot after", robot)

		if robot.Easting > extent.Max.Easting {
			robot.Easting = extent.Max.Easting
		}
		if robot.Northing > extent.Max.Northing {
			robot.Northing = extent.Max.Northing
		}
		if robot.Easting < extent.Min.Easting {
			robot.Easting = extent.Min.Easting
		}
		if robot.Northing < extent.Min.Northing {
			robot.Northing = extent.Min.Northing
		}

		Step1Robot.X = 0
		Step1Robot.Y = 0
		//fmt.Printf("robot inside after %s (%d, %d)\n", Step1Robot.Dir, Step1Robot.X, Step1Robot.Y)
	}

	fmt.Println("report", robot)
	report <- robot
}

type Action3 func()

func StartRobot3(name, script string, actChan chan Action3, log chan string) {
	fmt.Println("start robot")
	for _, cmd := range script {
		//fmt.Println("receive command:", string(cmd))
		switch Command(cmd) {
		case Command('R'):
			actChan <- Right
		case Command('A'):
			actChan <- Advance
		case Command('L'):
			actChan <- Left
		}
	}
	//fmt.Println("closing actChan")
	close(actChan)

	//cmd := make(chan Command)

	//go StartRobot(cmd, actChan)

}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, report chan []Step3Robot, log chan string) {

}

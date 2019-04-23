package robot

const (
	N = iota
	W
	S
	E
)

func (d Dir) String() string {
	return string(d)
}

func Right() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = E
	case E:
		Step1Robot.Dir = S
	case S:
		Step1Robot.Dir = W
	case W:
		Step1Robot.Dir = N
	}
}

func Left() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = W
	case E:
		Step1Robot.Dir = N
	case S:
		Step1Robot.Dir = E
	case W:
		Step1Robot.Dir = S
	}
}

func Advance() {
	switch Step1Robot.Dir {
	case E:
		Step1Robot.X++
	case N:
		Step1Robot.Y++
	case S:
		Step1Robot.Y--
	case W:
		Step1Robot.X--
	}
}

type Action func()

func StartRobot(commands chan Command, act chan Action) {
	command := <-commands

	for command != ' ' {
		switch command {
		case 'A':
			act <- Advance
		case 'R':
			act <- Right
		case 'L':
			act <- Left
		default:
			break
		}
	}

}

func Room(extent Rect, robot Step2Robot, act chan Action, rep chan Step2Robot) {
	for action := range act {
		action()
	}

}

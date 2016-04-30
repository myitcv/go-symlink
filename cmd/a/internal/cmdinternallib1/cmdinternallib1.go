package cmdinternallib1

import "github.com/pborman/uuid"

func DoSomething() uuid.UUID {
	return uuid.NewRandom()
}

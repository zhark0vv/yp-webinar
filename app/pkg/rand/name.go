package rand

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var names = []string{
	"Vlad",
	"Vladimir",
	"Anton",
	"Alex",
	"Alexey",
	"Alexandr",
	"Ilya",
	"Natasha",
	"Olga",
	"Oleg",
	"Igor",
	"Viktor",
	"Vitaliy",
}

func Name() string {
	return names[rand.Intn(len(names))]
}

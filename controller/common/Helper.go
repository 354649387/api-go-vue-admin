package common

import (
	"math/rand"
	"time"
)

func RandWord() string {

	slices := []string{"中文", "中文1", "中文2", "中文3", "中文4"}

	len := len(slices)

	rand.Seed(time.Now().UnixNano())

	r := rand.Intn(len)

	return slices[r]

}

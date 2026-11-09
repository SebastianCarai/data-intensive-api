package utils

import "math/rand"

func GetRandomInt (min int, max int) int{
	return rand.Intn(max - min) + min
}
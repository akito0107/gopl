package ex07

import (
	"math/rand"
	"testing"
)

var randSmall int
var randMid int
var randLarge int

func TestMain(m *testing.M) {
	randSmall = rand.Intn(100)
	randMid = rand.Intn(10000)
	randLarge = rand.Intn(10000000)
	m.Run()
}

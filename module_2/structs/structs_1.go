package main

import "fmt"

type Gangster struct {
	On          bool
	Ammo, Power int
}

func (g *Gangster) Shoot() bool {
	if g.On && g.Ammo > 0 {
		g.Ammo--
		return true
	} else {
		return false
	}
}

func (g *Gangster) RideBike() bool {
	if g.On && g.Power > 0 {
		g.Power--
		return true
	} else {
		return false
	}
}

func main() {
	testStruct := &Gangster{}
	testStruct.On = true
	testStruct.Ammo, testStruct.Power = 10, 20
	fmt.Println(testStruct.Shoot(), testStruct.Ammo)
	fmt.Println(testStruct.RideBike(), testStruct.Power)
	fmt.Println(testStruct.Shoot(), testStruct.Ammo)
	fmt.Println(testStruct.RideBike(), testStruct.Power)
}

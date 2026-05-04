package main

import "fmt"

type Auto interface {
	StepOnGaz()
	StopOnBrake()
}

type BMW struct {
}

type Zhiga struct {
}

func (b BMW) StepOnGaz() {
	fmt.Println("Я BMW! 550 лошадинных сил! Жмем на газ!")
}
func (b BMW) StopOnBrake() {
	fmt.Println("Я BMW! торможу!")
}



func (zhiga Zhiga) StepOnGaz() {
	fmt.Println("Я zhiga! 220 лошадинных сил! Жмем на газ!")
}

func (zhiga Zhiga) StopOnBrake() {
	fmt.Println("Я zhiga! тормоза отказали")
}

func ride(auto Auto) {
	fmt.Println("Я водитель!")
	fmt.Println("Я сажусь в машину")
	fmt.Println("И я нажимаю на газ...")
	auto.StepOnGaz()
	auto.StopOnBrake()
}

func main() {
	bmw := BMW{}
	zhiga := Zhiga{}
	

	ride(bmw)
	ride(zhiga)
}
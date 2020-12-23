package main

// TODO:
//  * check pin mapping in rpio.go


import (
    "fmt"
    "github.com/stianeikeland/go-rpio"
)


// map physical pins
//
// Mapping from pinout command:
//
// (16) GPIO23
// (18) GPIO24
// (22) GPIO25
//
var (
	pinIn1 = rpio.Pin(24)
	pinIn2 = rpio.Pin(23)
	pinEn = rpio.Pin(25)
)


func main() {

	err := rpio.Open()
	if err != nill {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	// shutdown memory access if we are done
	defer rpio.Close()

	// perform initial pin configuration
	pinIn1.Output()
	pinIn2.Output()
	pinEn.Output()

	// power down on start
	pinIn1.Low()
	pinIn2.Low()

	// initialize PWM
	pwm = pinEn.Pwm()
	// We want to have 1000Hz at the output of the PWM with 25% duty cycle
	pwm.Freq(1000 * 4)
	pmw.DutyCycle(1,4)


	// forward
	pinIn1.High()
	pinIn2.Low()

	// backward
	pinIn1.Low()
	pinIn2.High()
}

package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
	g "gobot.io/x/gobot/platforms/dexter/gopigo3"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

// hi this is marie
//robotRunLoop is the main function for the robot, the gobot framework
//will spawn a new thread in the NewRobot factory functin and run this
//function in that new thread. Do all of your work in this function and
//in other functions that this function calls. don't read from sensors or
//use actuators frmo main or you will get a panic.
//add
func robotRunLoop(lightSensor *aio.GroveLightSensorDriver, lightSensor2 *aio.GroveLightSensorDriver, gpg *g.Driver) {
	ledOn := true

	// hey this is a test ~~~

	for {
		sensorVal, err := lightSensor.Read()
		if err != nil {
			fmt.Errorf("Error reading light sensor %+v", err)
		}
		//soundSensorVal, err := soundSensor.Read()
		sensorVal2, err := lightSensor.Read()
		if err != nil {
			fmt.Errorf("Error reading from Sound Sensor %+v", err)
		}
		fmt.Println("Light (AD_2_1 Value is ", sensorVal)
		fmt.Println("Light (AD_1_1 Value is ", sensorVal2)
		//fmt.Println("Sound Value is ", sensorVal2)
		time.Sleep(time.Second)
		gpg.SetLED(1, 200, 200, 200)

		//Professor's code
		if ledOn {
			gpg.SetLED(1, 200, 200, 200) // light on
			gpg.SetLED(2, 200, 200, 200) // light on
			gpg.SetLED(3, 200, 200, 200) // light on
			gpg.SetLED(4, 200, 200, 200) // light on
			ledOn = !ledOn
		} else {
			gpg.SetLED(1, 0, 0, 0) // light off
			gpg.SetLED(3, 0, 0, 0) // light off
			gpg.SetLED(2, 0, 0, 0) // light off
			gpg.SetLED(4, 0, 0, 0) // light off
			ledOn = true
		}

		// assign code below

		if sensorVal <= 1000 {
			gpg.SetLED(1, 200, 200, 200)
			ledOn = true
		} else {
			ledOn = false
		}

		if sensorVal > 1000 {
			gpg.SetMotorDps(g.MOTOR_LEFT, 20)
		}

		//gpg.SetMotorDps(g,  20)
		//gpg.Start()
	}
}

func main() {
	//We create the adaptors to connect the GoPiGo3 board with the Raspberry Pi 3
	//also create any sensor drivers here
	raspiAdaptor := raspi.NewAdaptor()
	gopigo3 := g.NewDriver(raspiAdaptor)
	lightSensor := aio.NewGroveLightSensorDriver(gopigo3, "AD_2_1") //AnalogDigital Port 1 is "AD_1_1" this is port 2
	lightSensor2 := aio.NewGroveLightSensorDriver(gopigo3, "AD_1_1")
	//soundSensor := aio.NewGroveSoundSensorDriver(gopigo3, "AD_1_1")
	//end create hardware drivers

	//here we create an anonymous function assigned to a local variable
	//the robot framework will create a new thread and run this function
	//I'm calling my robot main loop here. Pass any of the variables we created
	//above to that function if you need them
	mainRobotFunc := func() {
		//robotRunLoop(lightSensor, soundSensor, gopigo3)
		robotRunLoop(lightSensor, lightSensor2, gopigo3)
	}

	//this is the crux of the gobot framework. The factory function to create a new robot
	//struct (go uses structs and not objects) It takes four parameters
	robot := gobot.NewRobot("gopigo3sensorChecker", //first a name
		[]gobot.Connection{raspiAdaptor},                   //next a slice of connections to one or more robot controllers
		[]gobot.Device{gopigo3, lightSensor, lightSensor2}, //next a slice of one or more sensors and actuators for the robots
		mainRobotFunc, //the variable holding the function to run in a new thread as the main function
	)

	robot.Start() //actually run the function
}

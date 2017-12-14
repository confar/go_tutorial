package main

import "fmt"

const usixeenbitmax float64  = 65525
const kmh_multiply float64  = 1.60934

type car struct {

	gas_pedal uint64 // min 0 max 65535
	break_pedal uint64
	steering_wheel int16 // -32k -- + 32k
	top_speed_kmh float64

}

func (c car) kmh() float64  {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixeenbitmax)
}

func (c car) mph() float64  {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixeenbitmax/kmh_multiply)
}

func (c *car) new_top_speed(newspeed float64)  {
	c.top_speed_kmh = newspeed
	}

func main() {
	//a_car := car{gas_pedal: 3225,
	//			break_pedal: 0,
	//			steering_wheel:1256,
	//			top_speed_kmh: 225}
	a_car := car{32255,0,1256,225}
	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
	a_car.new_top_speed(500)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
	}

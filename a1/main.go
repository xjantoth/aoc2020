package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal      uint16 // min 0 max 65535
	break_pedal    uint16
	steering_wheel int16
	top_speed_kmh  float64
}

// value receiver because it does not need to modify anything
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax / kmh_multiple)
}

// pointer receiver
func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

func main() {
	a_car := car{
		gas_pedal:      65345,
		break_pedal:    0,
		steering_wheel: 12434,
		top_speed_kmh:  140.0,
	}

	fmt.Println(a_car.gas_pedal)
	fmt.Println("To kilometers: ", a_car.kmh())
	fmt.Println("To mile per hour: ", a_car.mph())

	a_car.new_top_speed(500)

	fmt.Println("(Reinitialized) To kilometers: ", a_car.kmh())
	fmt.Println("(Reinitialized) To mile per hour: ", a_car.mph())
}

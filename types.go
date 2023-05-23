package main

type (
	Meter int //typedef style
	Cm    int
)

func MeterToCm(m Meter) Cm {
	return Cm(m * 100)
}

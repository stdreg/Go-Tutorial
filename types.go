package main

type (
	Meter int //typedef style
	Cm    int
)

func MeterToCm(m Meter) Cm {
	return Cm(m * 100)
}

type adress struct {
	street string
	city   string
}

// composition
type user struct {
	name   string
	adress //anonym field
}

// Additional Tags
type PublicAdress struct {
	Street string `json:"street"`
	City   string `json:"city`
}

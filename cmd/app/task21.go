package main

type Animal interface {
	Sound() string
}

type Dog struct{}

func (d *Dog) Bark() string {
	return "Гав"
}

func (d *Dog) Sound() string {
	return d.Bark()
}

type Piano struct{}

func (*Piano) Play() string {
	return "Играю музыку"
}

type PianoAdapter struct {
	*Piano
}

func (p *PianoAdapter) Sound() string {
	return p.Play()
}

func NewPianoAdapter(p *Piano) Animal {
	return &PianoAdapter{p}
}

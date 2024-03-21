package main

type Human struct {
	Age  int
	Race string
}

func (*Human) do() string {
	return "do"
}

type Action struct {
	Human
}
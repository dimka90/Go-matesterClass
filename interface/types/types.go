package types

type Printer interface {
	String() string
}

type Person struct {
	Name string
	Age  int
}

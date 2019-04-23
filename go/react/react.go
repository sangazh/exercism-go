package react

import "fmt"

func New() Reactor {
	return new(reactiveSystem)
}

type reactiveSystem struct {
}

func (s *reactiveSystem) CreateInput(i int) InputCell {
	ic := new(inputCell)
	ic.SetValue(i)
	return ic
}

func (s *reactiveSystem) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	fmt.Println("before:", c.Value())
	cc := new(computerCell)
	cc.Cell = s.CreateInput(f(c.Value()))
	fmt.Println("after", cc.Cell.Value())
	callback := func(value int) {
		cc.Cell = s.CreateInput(f(value))
	}
	cc.AddCallback(callback)

	return cc
}

func (s *reactiveSystem) CreateCompute2(c1, c2 Cell, f func(int, int) int) ComputeCell {
	newValue := f(c1.Value(), c2.Value())
	cc := new(computerCell)
	cc.Cell = s.CreateInput(newValue)

	return cc
}

type cell struct {
	value int
}

func (c *cell) Value() int {
	return c.value
}

type inputCell struct {
	cell
}

func (i *inputCell) SetValue(input int) {
	i.cell.value = input
}

type computerCell struct {
	Cell
}

func (c *computerCell) AddCallback(f func(int)) Canceler {
	f(c.Value())

	return new(canceller)
}

type canceller func()

func (c canceller) Cancel() {
	c()
}

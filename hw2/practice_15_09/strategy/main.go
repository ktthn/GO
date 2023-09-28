package main

import "fmt"

// паттерны --> стартегия ()

type Strategy interface {
	Process() error
	DoSomething() error
}

type FirstAlgorithm struct {
}

func (f *FirstAlgorithm) Process() error { //имплементация интерфейса процесс для структуры ферсталго
	fmt.Println("first algorithm is in progress")
	//first algo is not strategy yet cuz doesnt implement DoSmth method also
	// когда стракт реализует все методы интерфейса, то тогда он будет стратегией, короче имплементриует интерфейс
	//тут какой-то код будет

	return nil
}
func (f *FirstAlgorithm) DoSomething() error {
	//fmt.Println("first algorithm is in progress")
	// когда стракт реализует все методы интерфейса, то тогда он будет стратегией, короче имплементриует интерфейс
	return nil
}

type SecondAlgorithm struct {
}

func (s *SecondAlgorithm) Process() error {
	fmt.Println("second algorithm is in process")
	return nil
}
func (s *SecondAlgorithm) DoSomething() error {
	fmt.Println("second algorithm is in process")
	return nil
}

type Context struct {
	activeStrategy Strategy //какая та тсруктура окторая будет реализовывать интерфейс стратеджи
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.activeStrategy = strategy
}

func (c *Context) Process() error {
	err := c.activeStrategy.Process()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	context := Context{}
	context.SetStrategy(&FirstAlgorithm{})

	err := context.Process()
	if err != nil {
		return
	}

	context.SetStrategy(&SecondAlgorithm{})
	err = context.Process()
	if err != nil {
		return
	}
}

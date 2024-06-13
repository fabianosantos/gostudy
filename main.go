package main

import (
	"fmt"
	"time"
	"math/rand"
	"gostudy/puppy"
)

func init() {
	fmt.Println("Inicializando a aplicação...")
}

func main() {
	fmt.Println("Hello World!!!")
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())

	if x := 10; x>1 {
		fmt.Println("x > 1")
	}

	fmt.Println("Rand:", rand.Intn(10))

	ch1 := make(chan time.Duration) // criando channel de retorno tipo time.Duration
	ch2 := make(chan time.Duration)

	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))

	//Criando duas goroutines
	//O método go precisa receber uma chamada de função
	//Cada função domirá por um tempo random
	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- d1
	}()
	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- d2
	}()

	select {
	case v1 := <- ch1:
		fmt.Println("Valor channel 1", v1)
	case v2 := <- ch2:
		fmt.Println("Valor channel 2", v2)
	}
	select {
	case v1 := <- ch1:
		fmt.Println("Valor channel 1", v1)
	case v2 := <- ch2:
		fmt.Println("Valor channel 2", v2)
	}
}
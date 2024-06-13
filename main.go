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
	//EXMPLO DE UTILIZAÇÃO DE PACKAGES
	fmt.Println("Hello World!!!")
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())


	//EXEMPLO DE IF STATEMENT
	fmt.Println("\n========= IF X > 1 ========\n")
	if x := 10; x>1 {
		fmt.Println("x > 1")
	}


	//EXEMPLO DE MATH RAND
	fmt.Println("\n========= MATH RAND ========\n")
	fmt.Println("Rand:", rand.Intn(10))


	//FOR RANGE SLICE
	fmt.Println("\n========= FOR RANGE SLICE ========\n")
	s := []int{5,6,7,8,9,10}
	for i, v := range s {
		fmt.Println("Ranging over a slice", i, v)
	}

	//FOR RANGE MAP
	fmt.Println("\n========= FOR RANGE MAP ========\n")
	m := map[string]string{
		"name": "James",
		"surname": "Richard",
		"age": "42",
	}
	for k,v := range m {
		fmt.Printf("%s: %s \n", k, v)
	}



	//EXEMPLO DE GOROUTINES E UTILIZAÇÃO DE CHANNELS
	fmt.Println("\n========= GOROUTINES ========\n")
	ch1 := make(chan time.Duration) // criando channel de retorno tipo time.Duration
	ch2 := make(chan time.Duration)

	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))

	//Criando duas goroutines
	//O método go precisa receber uma chamada de função
	//Cada função dormirá por um tempo random
	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- d1
	}()
	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- d2
	}()

	//O metodo select é muito parecido com um switch statement
	//porem o switch é de uso comum, já o select é para uso especifico de codigo concorrente utilizando channels
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
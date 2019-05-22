package main

import(

	"time"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func f(s string){
	defer wg.Done() 
	time.Sleep(500 * time.Millisecond)
	fmt.Println(s)


}

func main(){

	fmt.Println("--- Normal")
	 for i := 0; i < 4;i++{
		wg.Add(3) 
		f("1")
		f("2")
		f("3")

	} 

	fmt.Println("--- Goroutines")
	 for i := 0; i < 4;i++{
		wg.Add(3) 
		go f("1")
		go f("2")
		go f("3")

	} 

	wg.Wait()
	fmt.Println("DONE")



}
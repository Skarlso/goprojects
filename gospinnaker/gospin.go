package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

//ExecuteProcesses executes a bunch of processes in parallel.
func ExecuteProcesses() {

	fmt.Println("Starting processes in parallel.")
	//Using a waitGroup here instead of a done <- 1 channel.
	//Either is fine as I would have only thrown away the <-done.
	var wg sync.WaitGroup
	wg.Add(50)
	//Fire off 50 processes
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			command := exec.Command("bash", "-c", "for i in {1..10}; do echo Hello; done")
			command.Stderr = os.Stderr
			command.Stdout = os.Stdout
			//I'm starting here because Run would wait for the process to finish.
			//This is only starting it and returning immediatly.
			err := command.Start()
			if err != nil {
				log.Fatal("Problem occured during running a command: ", err)
			}
		}()
	}
	wg.Wait()
}

func main() {
	ExecuteProcesses()
}

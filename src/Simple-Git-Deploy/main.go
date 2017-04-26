package main

import (
	"fmt"
	"os/exec"
	"log"
)

func main()  {
	cmd, err := exec.Command( "/bin/bash", "cmd.sh").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", cmd)
}

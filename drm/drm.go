package main

import (
	"fmt"
	"os/exec"
	"os"
	"strings"
)
func main(){
	// out, err := exec.Command("C:/Program Files/Docker/Docker/resources/bin/docker stop $(docker ps -a -q)", "Stopping docker containers").Output()	
	// out2, err2 := exec.Command("C:/Program Files/Docker/Docker/resources/bin/docker rm $(docker ps -a -q)", "Stopping docker containers").Output()

	cmdName := "docker"
	cmdArgs := []string{"ps", "-a","-q"}
	
	var (
		cmdOut []byte
		err    error
	)

	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running docker command: ", err)
		os.Exit(1)
	}

	containerIds := strings.Split(string(cmdOut), "\n")
	

	for _, el := range containerIds{
		
		if el == "" {
			break
		}

		err := exec.Command("docker", []string{"stop", el}...).Run()
		fmt.Println("Stoped Container:", el)
		if err != nil{
			fmt.Println(err)
		}

		err2 := exec.Command("docker", []string{"rm", el}...).Run()
		fmt.Println("Removed Container:",el)
		if err2 != nil{
			fmt.Println(err)
		} 
	}

}


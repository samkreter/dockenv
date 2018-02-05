package util


/*
TODOs
- Have better logging and user reporting
- Check if a port is in use
- Check if container name in use
- container defaults / predified solutions
- d start "context"
- d add "container" -c "context"  --temp (doesn't save in context)
- d add "container" 
*/


import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"io/ioutil"
	"path/filepath"
	"gopkg.in/yaml.v2"

	homedir "github.com/mitchellh/go-homedir"
)

var (
	containerCmd = "docker"
	configPath string
	contexts []Context
)

type Context struct {
	Name string `yaml:"name"`
	Containers 	[]Container `yaml:"containers"`
}

type Container struct {
	Name 	string 	`yaml:"name"`
	Image 	string 	`yaml:"image"`
	Port	string 	`yaml:"port"`
}

func init(){
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	configPath = filepath.Join(home, ".dockdev")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		os.Mkdir(configPath, os.ModePerm)
	}
}

func RemoveContext(name string) error {
	err := os.Remove(filepath.Join(configPath, name + ".yaml"))
	if err != nil {
		return err
	}

	return nil
}

func Create(name string) error {
	context := Context {
		Name: name,
		Containers: []Container{

		},
	}
	saveContext(context)

	return nil
}

func Start(name string) error {
	return fmt.Errorf("Not Implemented")
}

func Add(containerName string) error {
	return fmt.Errorf("Not Implemented")
}

func List() error {
	filepath.Walk(configPath, contextWalker)
	
	if (len(contexts) == 0){
		fmt.Println("There are no contexts avalible.")
	}
	
	for _, context := range contexts {
		fmt.Println(context.Name)
	}

	return nil
}

func Show(name string) error {
	context := getContext(name)

	fmt.Println(context)

	return nil
}

func Clean() error{
	getContainerIDsCmdArgs := []string{"ps", "-a","-q"}
	
	var (
		cmdOut []byte
		err    error
	)

	if cmdOut, err = exec.Command(containerCmd, getContainerIDsCmdArgs...).Output(); err != nil {
		log.Fatal(err)
	}

	containerIds := strings.Split(string(cmdOut), "\n")

	for _, containerID := range containerIds{
		
		if containerID == "" {
			break
		}

		err := exec.Command(containerCmd, []string{"stop", containerID}...).Run()
		log.Println("Stoped Container:", containerID)
		if err != nil{
			log.Println(err)
		}

		err = exec.Command(containerCmd, []string{"rm", containerID}...).Run()
		log.Println("Removed Container:", containerID)
		if err != nil{
			log.Println(err)
		} 
	}

	return nil
}


func createContainer(container Container) error {
	runCmdArgs := make([]string, 5)
	runCmdArgs = append(runCmdArgs, "run")

	if container.Name != "" {
		runCmdArgs = append(runCmdArgs, []string{"--name ", container.Name}...)
	}

	if container.Port != "" {
		runCmdArgs = append(runCmdArgs, []string{"-p ", container.Port}...)
	}

	runContainerCmdArgs := []string{"run", container.Image}

	cmdOut, err := exec.Command(containerCmd, runContainerCmdArgs...).Output()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(cmdOut)

	return fmt.Errorf("Not Implemented")
}

func getContextFromFilePath(path string) string{
	filename := filepath.Base(path)
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func contextWalker(path string, info os.FileInfo, err error) error {
	if !info.IsDir(){
		contextName := getContextFromFilePath(path)
		context := getContext(contextName)
		contexts = append(contexts, context)
	}
	return nil
}

func saveContext (context Context){
	data, err := yaml.Marshal(context)
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.Join(configPath, context.Name) + ".yaml"

	if err := ioutil.WriteFile(path, data, 0644); err != nil {
		log.Fatal(err)
	}
}

func getContext(name string) Context{
	filePath := filepath.Join(configPath, name + ".yaml")

	rawContextData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Context does not exist.")
	}

	var context Context

	if err = yaml.Unmarshal(rawContextData, &context); err != nil {
		log.Fatal("Could not parse context. Invalid syntax.")
	}

	return context
}

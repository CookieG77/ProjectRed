package main

import (
	"PPR/Scene"
	"os"
)

func main() {
	Scene.CreatePlayerWindow()
	//Smenu.SmenuRender()

}

func getArgs() []string {
	return os.Args[1:]
}

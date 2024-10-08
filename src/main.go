package main

import (
	PPR "PPR/Scene"
	"os"
)

func main() {
	PPR.CreatePlayerWindow()

}

func getArgs() []string {
	return os.Args[1:]
}

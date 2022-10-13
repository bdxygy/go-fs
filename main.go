package main

import "os"

func main() {
	readFile, err := os.ReadFile("learn")
	if err != nil {
		return
	}

	err = os.Mkdir("public", 0750)
	if err != nil {
		return
	}

	err = os.WriteFile("public/learn-thing", readFile, 0666)
	if err != nil {
		return
	}

	err = os.Remove("learn")
	if err != nil {
		return
	}
}

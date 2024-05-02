package main

import (
	"fmt"
	"os"

	croc "github.com/foverzar/crocgodyl"
)

func main() {
	app, _ := croc.NewApp(os.Getenv("CROC_URL"), os.Getenv("CROC_KEY"))

	egg, err := app.GetEgg(1, 1)
	if err != nil {
		handleError(err)
		return
	}

	fmt.Printf("%v\n", egg)
}

func handleError(err error) {
	if errs, ok := err.(*croc.ApiError); ok {
		for _, e := range errs.Errors {
			fmt.Println(e.Error())
		}
	} else {
		fmt.Println(err.Error())
	}
}

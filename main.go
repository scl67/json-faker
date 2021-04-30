package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/google/uuid"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ipsum   = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut quis quam eu dolor vestibulum congue. Vivamus ornare est nec est consequat porttitor. Donec id volutpat ex. Praesent sit amet efficitur magna. Curabitur ultrices ut nibh id placerat. Morbi vulputate justo sit amet molestie aliquet. Etiam sapien dui, finibus non condimentum quis, semper at erat. Maecenas dignissim rutrum fermentum. Cras erat enim, tincidunt id ullamcorper id, congue vitae lacus. Interdum et malesuada fames ac ante ipsum primis in faucibus. Quisque faucibus in nisi sit amet molestie. Cras erat risus, ornare nec metus ut, tincidunt aliquet sem. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Nunc sit amet purus at nunc pretium sodales. Aenean nec eros nulla. Donec nec nisi risus. Mauris est nulla, lacinia id rhoncus eu, faucibus eu turpis. Nullam sed dolor ultricies, mattis sem ve."
)

type fakeJSON struct {
	Id    uuid.UUID
	Name  string
	Age   int
	Blurb string
	Truth bool
}

func main() {
	args := os.Args[1:]

	switch len(args) {
	case 1:
		fmt.Println("Size of json: ", args[0], "kb")
		if i, err := strconv.Atoi(args[0]); err != nil {
			fmt.Println("Dun goofed on the input")
		} else {
			createJSON(i)
		}

	default:
		fmt.Printf("Incorrect number of args")
	}
}

func createJSON(size int) {
	fakes := []fakeJSON{}
	for i := 0; i < size; i++ {
		fake := fakeJSON{}

		id, err := uuid.NewRandom()
		if err != nil {
			fmt.Println("UUID ERROR")
			break
		}

		fake.Id = id
		fake.Age = i
		fake.Name = fmt.Sprintf("%d", i)
		fake.Blurb = ipsum
		fake.Truth = (i%2 == 0)
		fakes = append(fakes, fake)
	}

	if j, err := json.Marshal(fakes); err != nil {
		fmt.Println("JSON ERROR!")
	} else {
		if err := ioutil.WriteFile("fake.json", j, 0644); err != nil {
			fmt.Println("IO ERROR")
		} else {
			fmt.Println("Fake json created")
			fi, err := os.Stat("fake.json")
			if err != nil {
				fmt.Println("FILE SIZE ERROR")
			} else {
				fmt.Println("FILE SIZE: ", fi.Size())
			}
		}
	}
}

package main

import (
	"fmt"

	"github.com/satori/go.uuid"
)

func main() {
	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("UUIDv4: %s\n", u2)
	u3, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}
	fmt.Printf("Successfully parsed: %s\n", u3)
}

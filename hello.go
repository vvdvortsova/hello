package hello

import (
    "fmt"
    "github.com/vvdvortsova/greetings"
)

func Main1() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}

// main.go
package main

import (
    "fmt"
    "github.com/estebanhorn/password/generator"
    "github.com/estebanhorn/password/validate"
)

func main() {
    fmt.Println("Low complexity password:", generator.GeneratePasswordLow())
    fmt.Println("Medium complexity password:", generator.GeneratePasswordMedium())
    fmt.Println("High complexity password:", generator.GeneratePasswordHard())

    password := "YourPassword123!"
    score := validate.ValidatePassword(password)
    fmt.Printf("Password strength score: %d\n", score)
}

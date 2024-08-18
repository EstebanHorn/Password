// validate/validatePassword.go
package validate

import (
    "strings"
)

const (
    mayus   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    minus   = "abcdefghijklmnopqrstuvwxyz"
    numbers = "0123456789"
    symbols = "!@#$%^&*()_+-=[]{}|;:,.<>?/~`"
)

func ValidatePassword(password string) int {
    var score int
    if len(password) >= 8 {
        score += 1
    }
    if len(password) >= 12 {
        score += 1
    }
    if len(password) >= 16 {
        score += 1
    }
    var hasMayus, hasNumber, hasSymbol bool

    for _, char := range password {
        if strings.ContainsRune(mayus, char) && !hasMayus {
            score += 2
            hasMayus = true
        }
        if strings.ContainsRune(numbers, char) && !hasNumber {
            score += 2
            hasNumber = true
        }
        if strings.ContainsRune(symbols, char) && !hasSymbol {
            score += 3
            hasSymbol = true
        }
    }
    return score
}

// generator/passwordGenerator.go
package generator

import (
    "math/rand"
    "time"
    "unicode/utf8"
)

const (
    mayus   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    minus   = "abcdefghijklmnopqrstuvwxyz"
    numbers = "0123456789"
    symbols = "!@#$%^&*()_+-=[]{}|;:,.<>?/~`"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func GeneratePasswordLow() string {
    return GeneratePassword(8, false, false, true)
}

func GeneratePasswordMedium() string {
    return GeneratePassword(12, true, false, true)
}

func GeneratePasswordHard() string {
    return GeneratePassword(16, true, true, true)
}

func GeneratePassword(length int, useMayus bool, useSymbols bool, useNumbers bool) string {
    var charset string
    if useMayus {
        charset += mayus
    }
    if useSymbols {
        charset += symbols
    }
    if useNumbers {
        charset += numbers
    }
    charset += minus // Always include lowercase characters

    if len(charset) == 0 {
        return ""
    }

    var charsetRunes []rune
    for _, char := range charset {
        charsetRunes = append(charsetRunes, char)
    }

    var password []rune
    if useMayus {
        password = append(password, rune(mayus[rand.Intn(len(mayus))]))
    }
    if useSymbols {
        password = append(password, rune(symbols[rand.Intn(len(symbols))]))
    }
    if useNumbers {
        password = append(password, rune(numbers[rand.Intn(len(numbers))]))
    }

    for len(password) < length {
        password = append(password, charsetRunes[rand.Intn(len(charsetRunes))])
    }

    rand.Shuffle(len(password), func(i, j int) {
        password[i], password[j] = password[j], password[i]
    })

    return string(password)
}

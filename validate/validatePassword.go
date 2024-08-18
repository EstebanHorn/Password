package validate

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

const (
    mayus   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    minus   = "abcdefghijklmnopqrstuvwxyz"
    numbers = "0123456789"
    symbols = "!@#$%^&*()_+-=[]{}|;:,.<>?/~`"
)

// LoadCommonPasswords loads a list of common passwords from a file.
func LoadCommonPasswords() ([]string, error) {
    file, err := os.Open("validate/common-passwords.txt")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var commonPasswords []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        commonPasswords = append(commonPasswords, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return commonPasswords, nil
}

// ValidatePassword evaluates the strength of a password and checks if it is in the list of common passwords.
func ValidatePassword(password string) (int) {
    // Evaluate password strength
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

    // Check if password is common
    commonPasswords, err := LoadCommonPasswords()
    if err != nil {
        fmt.Println("Error loading common passwords:", err)
        return score
    }

    for _, common := range commonPasswords {
        if strings.Contains(password, common) {
            return 0
        }
    }


    return score
}

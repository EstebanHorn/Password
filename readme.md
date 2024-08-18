
---
# Password Generator and Validator

This Go package provides functionalities to generate and validate passwords with varying levels of complexity. It includes a password generator that creates passwords of different strengths and a validator to assess the strength of a given password.

## Features

- **Password Generation**:
  - Low complexity passwords (8 characters, only lowercase letters and numbers)
  - Medium complexity passwords (12 characters, includes uppercase letters and numbers)
  - High complexity passwords (16 characters, includes uppercase letters, symbols, and numbers)

- **Password Validation**:
  - Evaluates password strength based on length and character types
  - Scores passwords based on their complexity

## Installation

To use this package in your Go project, follow these steps:

1. **Install the package**:

   ```bash
   go get github.com/EstebanHorn/Password
   ```

2. **Import the package** in your Go code:

   ```go
   import "github.com/EstebanHorn/Password/generator"
   import "github.com/EstebanHorn/Password/validate"
   ```

## Usage

### Password Generation

Use the functions provided to generate passwords of different complexities:

- **Generate a low complexity password**:

  ```go
  passwordLow := generator.GeneratePasswordLow()
  fmt.Println("Low complexity password:", passwordLow)
  ```

- **Generate a medium complexity password**:

  ```go
  passwordMedium := generator.GeneratePasswordMedium()
  fmt.Println("Medium complexity password:", passwordMedium)
  ```

- **Generate a high complexity password**:

  ```go
  passwordHard := generator.GeneratePasswordHard()
  fmt.Println("High complexity password:", passwordHard)
  ```

### Password Validation

Use the `validatePassword` function to assess the strength of a password:

```go
password := "YourPassword123!"
score := validate.ValidatePassword(password)
fmt.Printf("Password strength score: %d\n", score)
```

## Functions

### `generator.GeneratePasswordLow() string`
Generates a low complexity password (8 characters, lowercase letters and numbers).

### `generator.GeneratePasswordMedium() string`
Generates a medium complexity password (12 characters, includes uppercase letters and numbers).

### `generator.GeneratePasswordHard() string`
Generates a high complexity password (16 characters, includes uppercase letters, symbols, and numbers).

### `validate.ValidatePassword(password string) int`
Evaluates the strength of a given password and returns a score based on its length and character types.

## Contributing

Contributions to this project are welcome! If you have suggestions or improvements, please submit a pull request or open an issue.

1. Fork the repository
2. Create a new branch for your feature or fix
3. Commit your changes
4. Push to the branch
5. Create a pull request


---

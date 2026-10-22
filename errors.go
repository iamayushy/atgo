package main

import (
	"errors"
	"fmt"
	"strconv"
)

// sentinel error: a package-level value callers can compare against
var ErrDivideByZero = errors.New("divide by zero")

// custom error type: any type with an Error() string method is an error
type InsufficientFundsError struct {
	balance, amount int
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("need %d, have %d", e.amount, e.balance)
}

// errors are returned as the last value; nil means success
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func withdraw(balance, amount int) (int, error) {
	if amount > balance {
		return balance, &InsufficientFundsError{balance, amount}
	}
	return balance - amount, nil
}

// %w wraps an error so callers can still inspect the original
func parseAge(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parseAge %q: %w", s, err)
	}
	return n, nil
}

// defer runs when the function returns; recover stops a panic
func safeIndex(s []int, i int) (v int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	return s[i], nil
}

func Errors() {
	// check err right after the call
	if q, err := divide(10, 2); err == nil {
		fmt.Println("10 / 2 =", q)
	}

	// errors.Is compares against a sentinel, even through wrapping
	_, err := divide(1, 0)
	fmt.Println("error:", err, "is ErrDivideByZero:", errors.Is(err, ErrDivideByZero))

	// errors.As extracts a specific error type
	_, err = withdraw(50, 80)
	var fundsErr *InsufficientFundsError
	if errors.As(err, &fundsErr) {
		fmt.Println("withdraw failed:", fundsErr, "short by:", fundsErr.amount-fundsErr.balance)
	}

	// wrapped errors keep the chain; Unwrap returns the inner one
	_, err = parseAge("abc")
	fmt.Println("wrapped:", err)
	fmt.Println("unwrapped:", errors.Unwrap(err))

	// panics are for unrecoverable bugs; recover turns one into an error
	_, err = safeIndex([]int{1, 2, 3}, 5)
	fmt.Println("safeIndex:", err)
}

package mux

import "fmt"

// BadRouteError creates error for a bad route
type BadRouteError struct {
	r RouteInterface
	s string
}

func NewBadRouteError(r RouteInterface, s string) *BadRouteError {
	return &BadRouteError{
		r: r,
		s: s,
	}
}

func (bre BadRouteError) Error() string {
	return fmt.Sprintf("Route -> Method: %s Path: %s Error: %s", bre.r.GetMethodName(), bre.r.GetPath(), bre.s)
}

// BadMethodError creates error for bad method
type BadMethodError struct {
	s string
}

func (bme *BadMethodError) Error() string { return fmt.Sprintf("Method not vaild (%s)", bme.s) }

// NewBadMethodError returns an error that formats as the given text.
func NewBadMethodError(text string) error {
	return &BadMethodError{text}
}

// BadPathError creates error for bad path
type BadPathError struct {
	s string
}

func (bme *BadPathError) Error() string { return fmt.Sprintf("Path is invaild (%s)", bme.s) }

// NewBadPathError returns an error that formats as the given text.
func NewBadPathError(text string) error {
	return &BadPathError{s: text}
}

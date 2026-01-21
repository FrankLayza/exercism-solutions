package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	val := float64(nb.Number())
	return fmt.Sprintf("This is a box containing the number %.1f", val)
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	res, ok := fnb.(FancyNumber)
	if !ok{
		return 0
	}
	number, err := strconv.Atoi(res.Value())
	if err != nil {
		return 0
	}
	return number
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	res, ok := fnb.(FancyNumber)
	if !ok {
		return "This is a fancy box containing the number 0.0"
	}
	val, err := strconv.Atoi(res.Value())
	if err != nil {
		return ""
	}
	decimal := float64(val)
	return fmt.Sprintf("This is a fancy box containing the number %.1f", decimal )
	
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	switch v := i.(type) {
	case int:
		return DescribeNumber(float64(v))
	case float64:
		return DescribeNumber(v)
	case NumberBox:
		return DescribeNumberBox(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	default:
		return "Return to sender"
	}
}

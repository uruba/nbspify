package nbspers

import (
	"testing"
)

const nbsperInsideExpectedCode = "inside"
const nbsperInsideExpectedResult = "Lorem ipsum dolor sit amet, consectetur adipiscing&nbsp;elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."

func TestNbsperInsideCode(t *testing.T) {
	testNbsperCode(t, nbsperInside, nbsperInsideExpectedCode)
}

func TestNbsperInsideApply(t *testing.T) {
	testNbsperApply(t, nbsperInside, nbsperInsideExpectedResult)
}

package nbspers

import (
	"testing"
)

const nbsperAfterExpectedCode = "after"
const nbsperAfterExpectedResult = "Lorem&nbsp;ipsum dolor&nbsp;sit amet, consectetur adipiscing elit, sed&nbsp;do eiusmod tempor incididunt ut labore et&nbsp;dolore magna aliqua."
const nbsperAfterExpectedResultCaseSensitive = "Lorem ipsum dolor&nbsp;sit amet, consectetur adipiscing elit, sed&nbsp;do eiusmod tempor incididunt ut labore et&nbsp;dolore magna aliqua."

func TestNbsperAfterCode(t *testing.T) {
	testNbsperCode(t, nbsperAfter, nbsperAfterExpectedCode)
}

func TestNbsperAfterApply(t *testing.T) {
	testNbsperApply(t, nbsperAfter, nbsperAfterExpectedResult)
}

func TestNbsperApplyCaseSensitive(t *testing.T) {
	testNbsperApplyCaseSensitive(t, nbsperAfter, nbsperAfterExpectedResultCaseSensitive)
}

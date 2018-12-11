package nbspers

import (
	"testing"
)

const nbsperAroundExpectedCode = "around"
const nbsperAroundExpectedResult = "Lorem ipsum&nbsp;dolor&nbsp;sit amet, consectetur adipiscing elit,&nbsp;sed&nbsp;do eiusmod tempor incididunt ut labore&nbsp;et&nbsp;dolore magna aliqua."

func TestNbsperAroundCode(t *testing.T) {
	testNbsperCode(t, nbsperAround, nbsperAroundExpectedCode)
}

func TestNbsperAroundApply(t *testing.T) {
	testNbsperApply(t, nbsperAround, nbsperAroundExpectedResult)
}

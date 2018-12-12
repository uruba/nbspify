package nbspers

import (
	"testing"
)

const nbsperAroundExpectedCode = "around"
const nbsperAroundExpectedResult = "Lorem&nbsp;ipsum&nbsp;dolor&nbsp;sit&nbsp;amet, consectetur&nbsp;adipiscing elit,&nbsp;sed&nbsp;do eiusmod tempor incididunt ut labore&nbsp;et&nbsp;dolore magna&nbsp;aliqua."
const nbsperAroundExpectedResultCaseSensitive = "Lorem ipsum&nbsp;dolor&nbsp;sit&nbsp;amet, consectetur&nbsp;adipiscing elit,&nbsp;sed&nbsp;do eiusmod tempor incididunt ut labore&nbsp;et&nbsp;dolore magna&nbsp;aliqua."

func TestNbsperAroundCode(t *testing.T) {
	testNbsperCode(t, nbsperAround, nbsperAroundExpectedCode)
}

func TestNbsperAroundApply(t *testing.T) {
	testNbsperApply(t, nbsperAround, nbsperAroundExpectedResult)
}

func TestNbsperAroundApplyCaseSensitive(t *testing.T) {
	testNbsperApplyCaseSensitive(t, nbsperAround, nbsperAroundExpectedResultCaseSensitive)
}

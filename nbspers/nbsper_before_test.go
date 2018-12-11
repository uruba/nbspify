package nbspers

import (
	"testing"
)

const nbsperBeforeExpectedCode = "before"
const nbsperBeforeExpectedResult = "Lorem ipsum&nbsp;dolor sit&nbsp;amet, consectetur&nbsp;adipiscing elit,&nbsp;sed do eiusmod tempor incididunt ut labore&nbsp;et dolore magna&nbsp;aliqua."

func TestNbsperBeforeCode(t *testing.T) {
	testNbsperCode(t, nbsperBefore, nbsperBeforeExpectedCode)
}

func TestNbsperBeforeApply(t *testing.T) {
	testNbsperApply(t, nbsperBefore, nbsperBeforeExpectedResult)
}

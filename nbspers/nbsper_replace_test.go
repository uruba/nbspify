package nbspers

import (
	"testing"
)

const nbsperReplaceExpectedCode = "replace"
const nbsperReplaceExpectedResult = "&nbsp; ipsum &nbsp; sit &nbsp;, consect&nbsp;ur &nbsp;, &nbsp; do eiusmod tempor incididunt ut labore &nbsp; &nbsp;e magna &nbsp;."

func TestNbsperReplaceCode(t *testing.T) {
	testNbsperCode(t, nbsperReplace, nbsperReplaceExpectedCode)
}

func TestNbsperReplaceApply(t *testing.T) {
	testNbsperApply(t, nbsperReplace, nbsperReplaceExpectedResult)
}

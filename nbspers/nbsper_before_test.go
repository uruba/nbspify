package nbspers

import (
	"os"
	"testing"
)

var nbsperBefore *NbsperBefore

const codeExpected = "before"

func TestMain(m *testing.M) {
	nbsperBefore = &NbsperBefore{}

	os.Exit(m.Run())
}

func TestNbsperBeforeCode(t *testing.T) {
	code := nbsperBefore.GetCode()

	if code != codeExpected {
		t.Errorf("The nbsperBefore's code should be '%s', got '%s'", codeExpected, code)
	}
}

func TestNbsperBeforeApply(t *testing.T) {

}

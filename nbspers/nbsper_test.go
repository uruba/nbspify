package nbspers

import (
	"nbspify/config"
	"testing"
)

var nbsperTestInput = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
var nbsperTestMatchSegments = []string{"lorem", "dolor", "amet", "adipiscing elit", "sed", "et", "aliqua"}

var nbsperBefore = &NbsperBefore{}
var nbsperAfter = &NbsperAfter{}
var nbsperAround = &NbsperAround{}
var nbsperInside = &NbsperInside{}
var nbsperReplace = &NbsperReplace{}

//func TestMain(m *testing.M) {
//	os.Exit(m.Run())
//}

func testNbsperCode(t *testing.T, nbsper Nbsper, expectedCode string) {
	code := nbsper.GetCode()

	if code != expectedCode {
		t.Errorf("The nbsper's code should be '%s', got '%s'", expectedCode, code)
	}
}

func testNbsperApply(t *testing.T, nbsper Nbsper, expectedResult string) {
	optionsNbsper := &config.OptionsNbsper{
		CaseSensitive: false,
	}

	result := nbsper.Apply(nbsperTestInput, nbsperTestMatchSegments, optionsNbsper)

	if result != expectedResult {
		t.Errorf("The nbsper's result should be '%s', got '%s'", expectedResult, result)
	}
}

func testNbsperApplyCaseSensitive(t *testing.T, nbsper Nbsper, expectedResult string) {
	optionsNbsper := &config.OptionsNbsper{
		CaseSensitive: true,
	}

	result := nbsper.Apply(nbsperTestInput, nbsperTestMatchSegments, optionsNbsper)

	if result != expectedResult {
		t.Errorf("The nbsper's result should be '%s', got '%s'", expectedResult, result)
	}
}

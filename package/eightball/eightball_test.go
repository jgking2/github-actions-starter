package eightball

import (
	"testing"
)

func TestDetermineMyFortune(t *testing.T) {
	accurateAnswer := DetermineMyFortune("Will lunch be delicious today?")
	if accurateAnswer == "" {
		t.Error("By some act of horror....we can't answer the question")
	}
}

package captcha_test

import (
	"image"
	"testing"
	"vez/captcha"
)

func TestChallengeUserSuccess(t *testing.T) {

	challenger := stubChallenger("42")
	prompt := stubPrompt("42")

	got := captcha.ChallengeUser(challenger, prompt)
	if !got {
		t.Fatal("Expected captcha.ChallengeUser return true")
	}
}

func TestChallengeUserFail(t *testing.T) {

	challenger := stubChallenger("lorem ipsum")
	prompt := stubPrompt("42")

	got := captcha.ChallengeUser(challenger, prompt)
	if got {
		t.Fatal("Expected captcha.ChallengeUser return false")
	}
}

//region Private

type stubChallenger string

func (stub stubChallenger) Challenge() (img image.Image, answer string) {

	return image.NewNRGBA(image.Rect(0, 0, 100, 100)), string(stub)
}

type stubPrompt string

func (prompt stubPrompt) Prompt(_ image.Image) string {

	return string(prompt)
}

//endregion

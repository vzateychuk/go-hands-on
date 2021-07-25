package captcha

import (
	"crypto/subtle"
	"image"
)

// Challenger is implemented by objects that can generate CAPTCHA image challenges.
type Challenger interface {
	Challenge() (img image.Image, imgText string)
}

// Prompter is implemented by objects that display a CAPTCHA image to the user,
// ask them to type their contents and return back their response
type Prompter interface {
	Prompt(img image.Image) string
}

// ChallengeUser requests a challenge from 'challenger' and prompts the user for an answer
// using 'prompter'. If the user's answer matches the challenge then ChallengeUser
// returns true. All comparisons are performed using constant-time checks to prevent information leaks.
func ChallengeUser(challenger Challenger, prompter Prompter) bool {

	img, expectedAnswer := challenger.Challenge()
	userAnswer := prompter.Prompt(img)

	isEqualLength := subtle.ConstantTimeEq(int32(len(expectedAnswer)), int32(len(userAnswer))) == 1
	if !isEqualLength {
		return false
	}

	isEqualContents := subtle.ConstantTimeCompare([]byte(userAnswer), []byte(expectedAnswer)) == 1
	return isEqualContents
}

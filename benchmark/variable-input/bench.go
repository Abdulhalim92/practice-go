package variable_input

import "github.com/Pallinder/go-randomdata"

func generateRandomString(length int) string {
	runes := randomdata.RandStringRunes(length)

	return runes
}

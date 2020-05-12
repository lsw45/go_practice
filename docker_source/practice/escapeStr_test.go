package practice

import (
	"fmt"
	"testing"
)

func TestEscapeStr(t *testing.T) {
	charsToEscape := `();\:`
	upstreamUAEscaped := escapeStr(":upstreamUA(", charsToEscape)
	fmt.Printf("%s UpstreamClient(%s)", "dockerUA", upstreamUAEscaped)
	t.Log(upstreamUAEscaped) //\:upstreamUA\(
}

// returns s with every rune in charsToEscape escaped by a backslash
func escapeStr(s string, charsToEscape string) string {
	var ret string
	for _, currRune := range s {
		appended := false
		for _, escapeableRune := range charsToEscape {
			if currRune == escapeableRune {
				ret += `\` + string(currRune)
				appended = true
				break
			}
		}
		if !appended {
			ret += string(currRune)
		}
	}
	return ret
}

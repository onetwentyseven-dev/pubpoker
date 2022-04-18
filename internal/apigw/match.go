package apigw

import (
	"net/url"
	"strings"
)

func matchOrigin(pattern, origin string) bool {

	if pattern == "" || strings.EqualFold(pattern, origin) {
		return true
	}

	if origin == "" {
		return false
	}

	patternURL, _ := url.Parse(strings.ToLower(pattern))
	originURL, _ := url.Parse(strings.ToLower(origin))

	if patternURL == nil || originURL == nil {
		return false
	}

	return matchSchema(patternURL, originURL) &&
		matchPort(patternURL, originURL) &&
		matchHostname(patternURL, originURL)

}

func matchSchema(pattern, origin *url.URL) bool {
	return pattern.Scheme == origin.Scheme
}

func matchPort(pattern, origin *url.URL) bool {
	return pattern.Port() == origin.Port()
}

func matchHostname(pattern, origin *url.URL) bool {
	patternSplit := strings.Split(pattern.Hostname(), ".")
	originSplit := strings.Split(origin.Hostname(), ".")

	if len(patternSplit) != len(originSplit) {
		return false
	}

	for i := len(patternSplit) - 1; i >= 0; i-- {
		if patternSplit[i] != "*" && patternSplit[i] != originSplit[i] {
			return false
		}
	}

	return false
}

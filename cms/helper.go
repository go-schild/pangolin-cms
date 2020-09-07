package cms

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// stringOr returns the first non-empty string
func stringOr(s ...string) string {
	if len(s) > 0 {
		for _, s0 := range s {
			if s0 != "" {
				return s0
			}
		}
		return s[len(s)-1]
	}
	return ""
}

func unmarshal(url string, s interface{}) error {
	data, err := ioutil.ReadFile(url)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, s)
}

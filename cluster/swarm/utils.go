package swarm

import (
	"regexp"
	"strings"
)

// convertKVStringsToMap converts ["key=value"] to {"key":"value"}
func convertKVStringsToMap(values []string) map[string]*string {
	result := make(map[string]*string, len(values))
	for _, value := range values {
		kv := strings.SplitN(value, "=", 2)
		if len(kv) == 1 {
			result[kv[0]] = nil
		} else {
			result[kv[0]] = &kv[1]
		}
	}

	return result
}

// convertMapToKVStrings converts {"key": "value"} to ["key=value"]
func convertMapToKVStrings(values map[string]*string) []string {
	result := make([]string, len(values))
	i := 0
	for key, value := range values {
		valueString := ""
		if value != nil {
			valueString = *value
		}
		result[i] = key + "=" + valueString
		i++
	}
	return result
}

var imageEngineOSErrorPattern = regexp.MustCompile(`cannot load (.+) image on (.+)`)

func isErrorLoadImageOsMismatch(err string) (match bool, imageOs, engineOs string) {
	matches := imageEngineOSErrorPattern.FindStringSubmatch(err)
	if len(matches) != 3 {
		return false, "", ""
	}
	return true, matches[1], matches[2]
}

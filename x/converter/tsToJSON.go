package converter

import "regexp"

func MudConfigToJSON(mudConfig []byte) []byte {
	temp := string(mudConfig)

	// Remove all imports
	importRegex := regexp.MustCompile(`import [ \"\@\;\/\{\}\,A-Za-z0-9]*\n`)
	temp = importRegex.ReplaceAllString(temp, "")

	// convert mudconfi function to json
	functionRegex := regexp.MustCompile(`export default mudConfig\(`)
	temp = functionRegex.ReplaceAllString(temp, "{config:")

	endFunction := regexp.MustCompile(`\)\;`)
	temp = endFunction.ReplaceAllString(temp, "}")

	// add quotes to keys
	quotesRegex := regexp.MustCompile("([a-zA-Z0-9-]+):")
	temp = quotesRegex.ReplaceAllString(temp, `"$1":`)

	return []byte(temp)
}

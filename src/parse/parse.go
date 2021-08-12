package parse

import (
	"errors"
	"strings"
	"net/url"
)

/*-------------------*/
/* EXPOSED FUNCTIONS */
/*-------------------*/

func CliArgumentsToQueryParams (arguments []string) (url.Values, error){
	queryParams := url.Values{}

	if len(arguments)%2 != 0 {
		return queryParams, errors.New("A search parameter is missing a value")
	}

	/* We iterate over each key,value pair at a time.*/
	for iter:=0; iter<len(arguments); iter+=2 {
		key := arguments[iter]
		value := arguments[iter+1]

		if !strings.HasPrefix(key, "--") {
			return queryParams, errors.New("A search parameter is not prefixed with --.")
		}

		// We take off the --.
		queryParams.Add(key[2:], value)
	}

	return queryParams, nil
}

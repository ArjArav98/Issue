package parse

import (
	"fmt"
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

	/* We iterate over all the values. */
	for iter:=0; iter<len(arguments); iter++ {

		if iter%2 != 0 { //Odd arguments are keys.
			if strings.HasPrefix(arguments[iter], "--") {
				return queryParams, errors.New("A search parameter is not prefixed with --.")
			}
		}

	}
}

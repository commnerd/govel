package response

import (
	"errors"

	"github.com/commnerd/govel/gerror"
)

func View(i interface{}, vars map[string]interface{}) Response {
	var response Response

	switch input := i.(type) {
	case string:
		response = Response{Response: jsonResponseFromString(input)}
	default:
		gerror.Check(errors.New("No handler for passed Json element."))
	}

	return response
}

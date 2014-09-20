package response

import "encoding/json"
import "net/http"

// JSON response.
func JSON(res http.ResponseWriter, val interface{}) {
	var b []byte
	var err error

	if Pretty {
		b, err = json.MarshalIndent(val, "", "  ")
	} else {
		b, err = json.Marshal(val)
	}

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(b)
}

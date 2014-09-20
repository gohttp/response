package response

import "encoding/xml"
import "net/http"

// XML response.
func XML(res http.ResponseWriter, val interface{}) {
	var b []byte
	var err error

	if Pretty {
		b, err = xml.MarshalIndent(val, "", "  ")
	} else {
		b, err = xml.Marshal(val)
	}

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/xml")
	res.Write(b)
}

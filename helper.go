package gopaypal

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func anyToReader(value interface{}) io.Reader {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(value)
	return buffer
}

func readerToAny(reader io.Reader, ptr interface{}) error {
	return json.NewDecoder(reader).Decode(ptr)
}

func catchError(resp *http.Response) error {
	if resp.StatusCode > 399 {
		paypalError := &PayPalError{}
		if err := json.NewDecoder(resp.Body).Decode(paypalError); err != nil {
			return err
		}
		return paypalError
	}

	return nil
}

func anyToJSON(ptr interface{}) string {
	data, _ := json.Marshal(ptr)
	return string(data)
}

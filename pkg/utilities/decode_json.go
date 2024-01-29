package utilities

import "encoding/json"

/*func Decode_json(data []byte) *json.Decoder {
	decoder := json.NewDecoder(data)
	if err := decoder.Decode(&transactions); err != nil {
		log.Fatal(err)
	}

	return decoder
}
*/

func DecodeJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

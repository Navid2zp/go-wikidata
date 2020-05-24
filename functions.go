package gowikidata

import (
	"encoding/json"
	"strconv"
)

func createParam(param string, values []string) string {
	newString := "&" + param + "="
	valuesLen := len(values)
	for i, value := range values {
		newString += value
		if i+1 < valuesLen {
			newString += "|"
		}
	}
	return newString
}

func (r *WikiDataGetEntitiesRequest) setParam(param string, values *[]string) {
	r.URL += createParam(param, *values)
}

func (r *WikiDataGetClaimsRequest) setParam(param string, values *[]string) {
	r.URL += createParam(param, *values)
}

func (r *WikiDataSearchEntitiesRequest) setParam(param string, values *[]string) {
	r.URL += createParam(param, *values)
}

// Unmarshal json to DynamicDataValue
func (d *DynamicDataValue) UnmarshalJSON(b []byte) (err error) {
	s := string(b)

	// If value starts with " and also ends with "
	// Then its string
	if string(s[0]) == "\"" && string(s[len(s)-1]) == "\"" {
		// Remove extra " from both sides of the string.
		cleaned := s[1 : len(s)-1]
		d.Data = cleaned
		d.S = cleaned
		d.Type = "String"
	} else {
		// If its int
		i, err := strconv.Atoi(s)
		if err != nil {
			// If its not int or string
			// Use DataValueFields
			values := DataValueFields{}
			err := json.Unmarshal(b, &values)
			if err != nil {
				return err
			}
			d.Type = "DataValueFields"
			d.ValueFields = values
			d.Data = values
		} else {
			// set value
			d.Type = "Int"
			d.I = i
			d.Data = i
		}
	}
	return
}

// Returns entity description in the given language code
func (e *Entity) GetDescription(languageCode string) string {
	return e.Descriptions[languageCode].Value
}

// Returns entity label in the given language code
func (e *Entity) GetLabel(languageCode string) string {
	return e.Labels[languageCode].Value
}

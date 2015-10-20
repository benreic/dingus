package utils

import (
	"bytes"
	"html/template"
)

func Render(fileName string, data map[string]interface{}) string {
	t, err := template.ParseFiles(fileName)
	buff := bytes.NewBufferString("")
	if err == nil {
		err := t.Execute(buff, data)
		if err == nil {
			return buff.String()
		}
	}

	panic(err)
}

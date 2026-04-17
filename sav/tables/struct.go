package tables

import (
	"fmt"
	"reflect"

	"github.com/jedib0t/go-pretty/v6/table"
)

func EnumerateFields(s any) (fields []any) {
	val := reflect.ValueOf(s)
	for i := 0; i < val.NumField(); i++ {
		fields = append(fields, val.Type().Field(i).Name)
	}
	return
}

func EnumerateFieldValues(s any) (values []any) {
	val := reflect.ValueOf(s)
	for i := 0; i < val.NumField(); i++ {
		if val.Type().Field(i).Name == "Name_" {
			values = append(values, val.MethodByName("Name").Call([]reflect.Value{})[0].Interface())
		} else {
			values = append(values, val.Field(i).Interface())
		}
	}
	return
}

func Tableize(slice []interface{}, labelfunc ...func(uint8) string) {
	if len(slice) == 0 {
		return
	}
	t := table.NewWriter()
	t.SetAutoIndex(true)
	t.SetStyle(table.StyleBold)
	header := EnumerateFields(slice[0])
	if labelfunc != nil {
		header = append([]any{"Label"}, header...)
	}
	t.AppendHeader(header)
	for i, s := range slice {
		row := EnumerateFieldValues(s)
		if labelfunc != nil {
			row = append([]any{labelfunc[0](uint8(i + 1))}, row...)
		}
		t.AppendRow(row)
	}
	fmt.Println(t.Render())
}

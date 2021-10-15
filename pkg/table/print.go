package table

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/olekukonko/tablewriter"
)

func Print(header []string, obj interface{}) error {
	table := tablewriter.NewWriter(os.Stdout)

	switch reflect.TypeOf(obj).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(obj)

		if s.Len() == 0 {
			fmt.Println("No results")
			return nil
		}

		table.SetBorder(false)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetHeader(header)

		for i := 0; i < s.Len(); i++ {
			entity := s.Index(i)
			var row []string
			for j := range header {
				row = append(row, entity.FieldByName(header[j]).String())
			}
			table.Append(row)
		}
		table.Render()
	default:
		return errors.New("attempting to print a non-slice as table")
	}
	return nil
}
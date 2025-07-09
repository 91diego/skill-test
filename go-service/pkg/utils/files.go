package utils

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/sevengit-wq/skill-test/pkg/models"
)

func GeneratePDF(student models.Student, documentTitle string) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, strings.ToTitle(documentTitle))
	pdf.Ln(12)

	val := reflect.ValueOf(student)
	typ := reflect.TypeOf(student)

	pdf.SetFont("Arial", "", 12)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		label := fieldType.Tag.Get("label")
		if label == "" {
			label = fieldType.Name
		}

		// Omitir valores nil
		if field.Kind() == reflect.Ptr {
			if field.IsNil() {
				continue
			}
			field = field.Elem() // Obtener el valor apuntado
		}

		// Omitir strings vacíos
		if field.Kind() == reflect.String && field.String() == "" {
			continue
		}

		// Agregar línea al PDF
		pdf.CellFormat(50, 10, fmt.Sprintf("%s:", label), "0", 0, "L", false, 0, "")
		pdf.CellFormat(0, 10, fmt.Sprintf("%v", field.Interface()), "0", 1, "L", false, 0, "")
	}

	// Exportar a []byte
	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

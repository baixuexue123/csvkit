package csvkit

import (
	"encoding/csv"
	"io"
)

type DictWriter struct {
	writer     *csv.Writer
	fieldNames []string
}

func NewDictWriter(w io.Writer, fields []string) *DictWriter {
	return &DictWriter{
		writer:     csv.NewWriter(w),
		fieldNames: fields,
	}
}

func (w *DictWriter) WriteHeader() error {
	return w.writer.Write(w.fieldNames)
}

func (w *DictWriter) WriteRow(row Record) error {
	values := make([]string, 0, len(row))
	for _, k := range w.fieldNames {
		values = append(values, row[k])
	}
	return w.writer.Write(values)
}

func (w *DictWriter) WriteRows(rows []Record) (err error) {
	for _, r := range rows {
		err = w.WriteRow(r)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *DictWriter) Flush() {
	w.writer.Flush()
}

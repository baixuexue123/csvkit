package csvkit

import (
	"encoding/csv"
	"io"
)

type DictReader struct {
	reader     *csv.Reader
	fieldNames []string
	restVal    string
}

func NewDictReader(r io.Reader) *DictReader {
	return &DictReader{
		reader: csv.NewReader(r),
	}
}

// FieldNames
// If fieldNames is omitted, the values in the first row of reader will be used as the fieldNames
func (w *DictReader) FieldNames() ([]string, error) {
	if len(w.fieldNames) == 0 {
		fieldNames, err := w.reader.Read()
		if err != nil {
			return nil, err
		}
		w.fieldNames = fieldNames
	}
	return w.fieldNames, nil
}

func (w *DictReader) Read() (dict Record, err error) {
	_, err = w.FieldNames()
	if err != nil {
		return nil, err
	}

	record, err := w.reader.Read()
	if err != nil {
		return nil, err
	}

	dict = make(Record, len(w.fieldNames))
	for i, name := range w.fieldNames {
		if len(record) > i {
			dict[name] = record[i]
		} else {
			dict[name] = w.restVal
		}
	}
	return dict, nil
}

func (w *DictReader) ReadAll() (data []Record, err error) {
	_, err = w.FieldNames()
	if err != nil {
		return nil, err
	}

	records, err := w.reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, r := range records {
		d := make(Record, len(w.fieldNames))
		for j, name := range w.fieldNames {
			if len(r) > j {
				d[name] = r[j]
			} else {
				d[name] = w.restVal
			}
		}
		data = append(data, d)
	}
	return data, nil
}

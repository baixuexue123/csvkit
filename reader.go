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

// 如果没有设置字段名, 读取文件第一行作为字段名
func (r *DictReader) FieldNames() ([]string, error) {
	if len(r.fieldNames) == 0 {
		fieldNames, err := r.reader.Read()
		if err != nil {
			return nil, err
		}
		r.fieldNames = fieldNames
	}
	return r.fieldNames, nil
}

func (r *DictReader) Read() (dict Record, err error) {
	_, err = r.FieldNames()
	if err != nil {
		return nil, err
	}

	record, err := r.reader.Read()
	if err != nil {
		return nil, err
	}

	dict = make(Record, len(r.fieldNames))
	for i, name := range r.fieldNames {
		if len(record) > i {
			dict[name] = record[i]
		} else {
			dict[name] = r.restVal
		}
	}
	return dict, nil
}

func (r *DictReader) ReadAll() (data []Record, err error) {
	_, err = r.FieldNames()
	if err != nil {
		return nil, err
	}

	records, err := r.reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, rc := range records {
		d := make(Record, len(r.fieldNames))
		for j, name := range r.fieldNames {
			if len(rc) > j {
				d[name] = rc[j]
			} else {
				d[name] = r.restVal
			}
		}
		data = append(data, d)
	}
	return data, nil
}

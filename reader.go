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
	if r.fieldNames == nil {
		fieldNames, err := r.reader.Read()
		if err != nil {
			return nil, err
		}
		r.fieldNames = fieldNames
	}
	return r.fieldNames, nil
}

func (r *DictReader) ReadLine() (dict Record, err error) {
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

func (r *DictReader) ReadLines(n int) (data []Record, err error) {
	for i := 0; i < n; i++ {
		line, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		data = append(data, line)
	}
	return data, nil
}

func (r *DictReader) ReadAll() (data []Record, err error) {
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

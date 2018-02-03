package csvkit

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := strings.NewReader(in)
	reader := NewDictReader(r)

	t.Log(reader.FieldNames())

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error(err)
		}
		t.Log(record)
	}
}

func TestReadAll(t *testing.T) {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := strings.NewReader(in)
	reader := NewDictReader(r)

	t.Log(reader.FieldNames())

	data, err := reader.ReadAll()
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}

func TestDictWriter_WriteRow(t *testing.T) {
	records := []Record{
		{"first_name": "Rob", "last_name": "Pike", "username": "rob"},
		{"first_name": "Ken", "last_name": "Thompson", "username": "ken"},
		{"first_name": "Robert", "last_name": "Griesemer", "username": "gri"},
	}
	header := []string{"first_name", "last_name", "username"}
	writer := NewDictWriter(os.Stdout, header)
	var err error
	err = writer.WriteHeader()
	for _, r := range records {
		err = writer.WriteRow(r)
		if err != nil {
			t.Error(err)
		}
	}
	writer.Flush()
}

func TestDictWriter_WriteRows(t *testing.T) {
	records := []Record{
		{"first_name": "Rob", "last_name": "Pike", "username": "rob"},
		{"first_name": "Ken", "last_name": "Thompson", "username": "ken"},
		{"first_name": "Robert", "last_name": "Griesemer", "username": "gri"},
	}
	header := []string{"first_name", "last_name", "username"}
	writer := NewDictWriter(os.Stdout, header)
	var err error
	err = writer.WriteHeader()
	if err != nil {
		t.Error(err)
	}
	err = writer.WriteRows(records)
	if err != nil {
		t.Error(err)
	}
	writer.Flush()
}

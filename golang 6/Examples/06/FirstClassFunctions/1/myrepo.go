package myrepo

import (
	"io"
	"bytes"
	"encoding/json"
)

type FullRepositoryLoader interface {
	LoadAll(input io.Reader) (result bytes.Buffer, err error) 
}	

type Filler interface {
	ReadFrom(buffer io.Reader) (err error)
}

type (
	Date struct {
		Year  uint16
		Month byte
		Day   byte
	}
	Team struct {
		Title string
		Goals byte
	}
	Match struct {
		Date     Date
		Host     Team
		Guest    Team
		Overtime bool
	}
)

type MyRepository struct {}

func (repo MyRepository) LoadAll(input io.Reader) (result bytes.Buffer, err error) {
	_, err = result.ReadFrom(input)
	return
}

func (m *Match) ReadFrom(dataRow io.Reader) (err error) {
	err = json.NewDecoder(dataRow).Decode(m)
	return 
}
 

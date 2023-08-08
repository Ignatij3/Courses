package rep

import (
	"encoding/binary"
	"errors"
	"io"
)

type Database struct {
	Data io.ReadWriteSeeker
}

func (db Database) Append(row []byte) error {
	if _, err := db.Data.Seek(0, io.SeekEnd); err != nil {
		return err
	}
	if err := binary.Write(db.Data, binary.LittleEndian, int32(len(row))); err != nil {
		return err
	}
	if _, err := db.Data.Write(row); err != nil {
		return err
	}
	return nil
}

func (db Database) GetNext(row []byte) (int, error) {
	var n int32
	if err := binary.Read(db.Data, binary.LittleEndian, &n); err != nil {
		return 0, nil
	}
	return io.ReadFull(db.Data, row[:n])
}

func (db Database) GetFirst(row []byte) (n int, err error) {
	if _, err := db.Data.Seek(0, io.SeekStart); err != nil {
		return 0, err
	}
	return db.GetNext(row)
}

type EncoderDecoder interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte, interface{}) error
}

type Row map[string]string

type Repository struct {
	DB    Database
	Codec EncoderDecoder
}

type RowEncoder interface{ EncodeInto(Row) }
type RowDecoder interface{ DecodeFrom(Row) }

func (r Repository) Append(e RowEncoder) error {
	row := make(map[string]string)
	e.EncodeInto(row)
	dbRow, err := r.Codec.Encode(row)
	if err != nil {
		return err
	}
	return r.DB.Append(dbRow)
}

func (r Repository) GetFirst(d RowDecoder) error { return fromDB(r.DB.GetFirst, r.Codec, d) }

func (r Repository) GetNext(d RowDecoder) error { return fromDB(r.DB.GetNext, r.Codec, d) }

func (r Repository) NextThat(d RowDecoder, cond func (r RowDecoder) bool) error { 
	if err:= fromDB(r.DB.GetNext, r.Codec, d); err != nil {
		return err
	}         
	for !cond(d) {
		if err:= fromDB(r.DB.GetNext, r.Codec, d); err != nil { 
			return err
		}
	}
	return nil
}

func fromDB(getRow func([]byte) (int, error), dbCodec EncoderDecoder, d RowDecoder) error {
	var dbRow [250]byte // TODO: should be dynamic, but requires db changes
	n, err := getRow(dbRow[:])
	if n==0 {
		return errors.New("Empty row")
	}	
	if err != nil {
		return err
	}
	var row map[string]string
	if err := dbCodec.Decode(dbRow[:n], &row); err != nil {
		return err
	}
	d.DecodeFrom(row)
	return nil
}

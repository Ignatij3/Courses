package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"rep"
)

type SeekableBuffer struct {
	b   []byte
	pos int
}

func (buf *SeekableBuffer) Write(b []byte) (int, error) {
	buf.b = append(buf.b, b...)
	return len(b), nil
}

func (buf *SeekableBuffer) Read(b []byte) (n int, err error) {
	n = copy(b, buf.b[buf.pos:])
	if n < len(b) {
		return 0, io.EOF
	}
	buf.pos += n
	return n, err
}

func (buf *SeekableBuffer) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		buf.pos = int(offset)
	case io.SeekEnd:
		buf.pos = len(buf.b) - int(offset)
	case io.SeekCurrent:
		buf.pos += int(offset)
	}
	return int64(buf.pos), nil
}

type jsonCodec struct{}

func (jsonCodec) Encode(v interface{}) ([]byte, error)    { return json.Marshal(v) }
func (jsonCodec) Decode(data []byte, v interface{}) error { return json.Unmarshal(data, v) }

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

type MatchEncoder struct{ m Match }

func (me MatchEncoder) EncodeInto(r rep.Row) {
	r["year"] = strconv.Itoa(int(me.m.Date.Year))
	r["month"] = strconv.Itoa(int(me.m.Date.Month))
	r["day"] = strconv.Itoa(int(me.m.Date.Day))
	r["hostTeam"] = me.m.Host.Title
	r["hostGoals"] = strconv.Itoa(int(me.m.Host.Goals))
	r["guestTeam"] = me.m.Guest.Title
	r["guestGoals"] = strconv.Itoa(int(me.m.Guest.Goals))
	if me.m.Overtime {
		r["overtime"] = "yes"
	} else {
		r["overtime"] = "no"
	}
}

type MatchDecoder struct{ m *Match }

func (md MatchDecoder) DecodeFrom(r rep.Row) {
	w, _ := strconv.Atoi(r["year"])
	md.m.Date.Year = uint16(w)
	w, _ = strconv.Atoi(r["month"])
	md.m.Date.Month = byte(w)
	w, _ = strconv.Atoi(r["day"])
	md.m.Date.Day = byte(w)
	md.m.Host.Title = r["hostTeam"]
	w, _ = strconv.Atoi(r["hostGoals"])
	md.m.Host.Goals = byte(w)
	md.m.Guest.Title = r["guestTeam"]
	w, _ = strconv.Atoi(r["guestGoals"])
	md.m.Guest.Goals = byte(w)
	md.m.Overtime = r["overtime"] == "yes"
}

func main() {
	f := new(SeekableBuffer)
	var JSON jsonCodec

	db := rep.Database{f}
	repo := rep.Repository{db, JSON}

	fin, err := os.Open("hockey.json")
	if err != nil {
		return
	}

	md := MatchDecoder{&Match{}}
	me := MatchEncoder{Match{}}
	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		s := scanner.Text()
		if err = json.Unmarshal([]byte(s), md.m); err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		me.m = *md.m
		repo.Append(me)
	}
	fin.Close()

	for err := repo.GetFirst(md); err == nil; err = repo.GetNext(md) {
		if err != nil {
			log.Fatal("Get failed:", err)
		}
		fmt.Println("===", *(md.m))
	}

	cond := func(md rep.RowDecoder) bool {
		return (*(md.(MatchDecoder).m)).Date.Day == 3
	}
	for err := repo.GetFirst(md); err == nil; err = repo.NextThat(md, cond) {
		fmt.Println("+++", *(md.m))
	}
}

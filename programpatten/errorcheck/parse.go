package errorcheck

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type Point struct {
	Longitude     float64
	Latitude      float64
	Distance      float64
	ElevationGain float64
	ElevationLoss float64
}

// 例子1
func parseBefore(r io.Reader) (*Point, error) {
	var p Point

	if err := binary.Read(r, binary.BigEndian, &p.Latitude); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.BigEndian, &p.Longitude); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.BigEndian, &p.Distance); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.BigEndian, &p.ElevationGain); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.BigEndian, &p.ElevationLoss); err != nil {
		return nil, err
	}

	return &p, nil
}

// 例子2
func parseAfter(r io.Reader) (*Point, error) {
	var p Point
	var err error
	read := func(data interface{}) {
		if err != nil {
			return
		}
		err = binary.Read(r, binary.BigEndian, data)
	}

	read(&p.Latitude)
	read(&p.Longitude)
	read(&p.Distance)
	read(&p.ElevationGain)
	read(&p.ElevationLoss)

	if err != nil {
		return &p, err
	}
	return &p, nil
}

type Reader struct {
	r   io.Reader
	err error
}

func (r *Reader) read(data interface{}) {
	if r.err == nil {
		r.err = binary.Read(r.r, binary.BigEndian, data)
	}
}

// 例子3
func parseStruct(input io.Reader) (*Point, error) {
	var p Point
	r := Reader{r: input}

	r.read(&p.Latitude)
	r.read(&p.Longitude)
	r.read(&p.Distance)
	r.read(&p.ElevationGain)
	r.read(&p.ElevationLoss)

	if r.err != nil {
		return nil, r.err
	}

	return &p, nil
}

// 例子4

var b = []byte{0x48, 0x61, 0x6f, 0x20, 0x43, 0x68, 0x65, 0x6e, 0x00, 0x00, 0x2c}
var r = bytes.NewReader(b)

type Person struct {
	Name   [10]byte
	Age    uint8
	Weight uint8
	err    error
}

func (p *Person) read(data interface{}) {
	if p.err == nil {
		p.err = binary.Read(r, binary.BigEndian, data)
	}
}

func (p *Person) ReadName() *Person {
	p.read(&p.Age)
	return p
}

func (p *Person) ReadAge() *Person {
	p.read(&p.Age)
	return p
}

func (p *Person) ReadWeight() *Person {
	p.read(&p.Weight)
	return p
}

func (p *Person) Print() *Person {
	if p.err == nil {
		fmt.Printf("Name%s, Age=%d, Weight=%d\n", p.Name, p.Age, p.Weight)
	}
	return p
}

func PrintPerson() {
	p := Person{}
	p.ReadName().ReadAge().ReadWeight().Print()
	fmt.Println(p.err)
}

package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unsafe"

	"github.com/gogo/protobuf/proto"
)

const UINT64 = uint64(unsafe.Sizeof(uint64(0)))

var UINT64B = make([]byte, UINT64)

//go:generate bash -c ./proto.sh
func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing subcommand: list or add")
		os.Exit(1)
	}
	var err error
	switch cmd := flag.Arg(0); cmd {
	case "list":
		err = list()
	case "add":
		err = add(strings.Join(flag.Args()[1:], " "))
	default:
		err = fmt.Errorf("unknown subcommand %s", cmd)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func add(text string) error {
	task := &Task{
		Text: text,
		Done: false,
	}

	b, err := proto.Marshal(task)
	if err != nil {
		return fmt.Errorf("could not encode task: %v", err)
	}

	f, err := os.OpenFile("data.pbf", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("could not open %s", err)
	}

	if err := binary.Write(f, binary.LittleEndian, uint64(len(b))); err != nil {
		return fmt.Errorf("could not encode length of message: %s", err)
	}

	_, err = f.Write(b)
	if err != nil {
		return fmt.Errorf("could not write task to file: %s", err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("could not close file %s", err)
	}
	return nil
}

func list() error {
	file, err := os.Open("data.pbf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return chunck(file, 0)
}

func chunck(file *os.File, off uint64) error {
	var l uint64
	n, err := file.ReadAt(UINT64B, int64(off))
	if n > 0 {
		if err := binary.Read(bytes.NewReader(UINT64B), binary.LittleEndian, &l); err != nil {
			return fmt.Errorf("could not decode message length: %v", err)
		}
		fmt.Println("-----", off, "->", off+UINT64+l, "-----")
		b := make([]byte, l)
		if n, err := file.ReadAt(b, int64(off+UINT64)); err != nil {
			return fmt.Errorf("could not decode message body: %v, %d", err, n)
		}
		if err := task(b); err != nil {
			return fmt.Errorf("could not decode message body: %v", err)
		}
	}
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	return chunck(file, uint64(off)+UINT64+l)
}

func task(b []byte) error {
	var task Task
	if err := proto.Unmarshal(b, &task); err != nil {
		return fmt.Errorf("could not read task: %v", err)
	}
	if task.Done {
		fmt.Printf("👍")
	} else {
		fmt.Printf("😱")
	}
	fmt.Printf(" %s\n", task.Text)
	return nil
}

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type demoRead struct{}

func (d *demoRead) Read(b []byte) (n int, err error) {
	fmt.Println("in <")
	return os.Stdin.Read(b)
}

type demoWrite struct {
}

func (d *demoWrite) Write(b []byte) (n int, err error) {
	fmt.Println("out >")
	return os.Stdout.Write(b)
}

func main() {

	var (
		reader demoRead
		writer demoWrite
	)

	// input := make([]byte, 4096)

	// op, err := reader.Read(input)

	// // fmt.Println(op, err)

	// if err != nil {
	// 	log.Fatalln("unable to read data")
	// } else {
	// 	fmt.Printf("Read %v bytes\n", op)
	// }

	// opp, errr := writer.Write(input)

	// if errr != nil {
	// 	log.Fatalln("unable to wrote data")
	// } else {
	// 	fmt.Printf("write %v bytes\n", opp)
	// }

	// use io.Copy() we do not need to write many conditions using this

	_, errc := io.Copy(&writer, &reader)
	if errc != nil {
		log.Fatalln("unable to read/write")
	}
}

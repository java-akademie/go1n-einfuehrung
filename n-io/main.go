package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/java-akademie/myutils"
)

var fileName = "test.txt"

func main() {
	myutils.H1("Input/Output Files")
	createFile()
	myutils.Wait()
	readFile1()
	myutils.Wait()
	readFile2()
	myutils.Wait()
	readFile3()
	myutils.Wait()
}

func createFile() {
	myutils.H2("Create File")

	file, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File Created: %v \n", fileName)

	defer file.Close()

	for i := 1; i <= 5; i++ {
		_, err := file.WriteString(fmt.Sprintf("zeile nr %v\n", i))

		if err != nil {
			log.Fatal(err)
		}
	}
}

func readFile1() {
	myutils.H2("Read File 1 - small chunks")

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	data := make([]byte, 22)

	for {
		n, err := file.Read(data)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("DataLength: %d; ReadLength: %d\n", len(data), n)
		fmt.Println(string(data[0:n]))
	}
}

func readFile2() {
	myutils.H2("Read File 2 - the entire file")

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data and Read Length: %d bytes\n", len(data))
	fmt.Printf("Data: \n%s", data)
}

func readFile3() {
	myutils.H2("Read File 3 - line by line")

	buf, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = buf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	line := bufio.NewScanner(buf)

	for line.Scan() {
		t := line.Text()
		fmt.Printf("ReadLength: %d\n", len(t))
		fmt.Println(t)
	}

	err = line.Err()
	if err != nil {
		log.Fatal(err)
	}
}

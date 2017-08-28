package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"fmt"
	"io"
)

func main() {
	f, _ := os.Open("Accounts.csv")

    // Create a new reader.
    r := csv.NewReader(bufio.NewReader(f))

    file, err := os.Create("New.csv")
    checkError("Cannot create file", err)
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for {
        record, err := r.Read()
        // Stop at EOF.
        if err == io.EOF {
            break
        }
        // Display record.
        // ... Display record length.
        // ... Display all individual elements of the slice.
        fmt.Println(record)
        fmt.Println(len(record))
        errors := writer.Write(record)
        checkError("Cannot write to file", errors)
        for value := range record {
            fmt.Printf("  %v\n", record[value])
        }
    }

}

func checkError(message string, err error) {
    if err != nil {
        panic(err)
    }
}
package main

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

func main() {
	csvString := "eko,kurniawan,khannedy\n" +
		"budi,rahayu,nugraha\n" +
		"joko,moro,diah\n"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		println(record[0], record[1], record[2])
	}

	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"Eko", "Kurniawan", "Khannedy"})
	_ = writer.Write([]string{"Budi", "Rahayu", "Nugraha"})
	_ = writer.Write([]string{"Joko", "Moro", "Diah"})

	writer.Flush()
}

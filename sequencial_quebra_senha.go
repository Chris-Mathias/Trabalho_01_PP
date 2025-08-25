package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {

	password := "12345678"
	max := 100_000_000
	start := time.Now()
	var result string

	for n := 0; n < max; n++ {
		try := fmt.Sprintf("%08d", n)
		if try == password {
			result = try
			break
		}
	}

	if result == "" {
		fmt.Println("Senha não encontrada!")
		return
	}

	elapsed := time.Since(start)

	fileName := "resultado_sequencial.csv"
	_, err := os.Stat(fileName)
	createHeader := os.IsNotExist(err)

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	writer := csv.NewWriter(f)

	if createHeader {
		writer.Write([]string{"Execucao", "nGoroutines", "Tempo"})
	}

	data, _ := os.ReadFile(fileName)
	lines := 0
	for _, b := range data {
		if b == '\n' {
			lines++
		}
	}
	id := lines
	if id == 0 {
		id = 1
	}

	record := []string{
		fmt.Sprintf("%d", id),
		"1",
		elapsed.String(),
	}
	writer.Write(record)
	writer.Flush()

	fmt.Println("Resultado salvo em", fileName, "| Goroutines: 1 | Tempo:", record[2])
}

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {

	password := "12345678"
	goroutines := 200
	max := 100_000_000

	start := time.Now()
	var wg sync.WaitGroup
	found := make(chan string, 1)

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(startNum int) {
			defer wg.Done()
			for n := startNum; n < max; n += goroutines {
				if fmt.Sprintf("%08d", n) == password {
					found <- password
					return
				}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(found)
	}()

	<-found
	elapsed := time.Since(start)

	fileName := "resultado_paralelo.csv"
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
		fmt.Sprintf("%d", goroutines),
		elapsed.String(),
	}
	writer.Write(record)
	writer.Flush()

	fmt.Println("Resultado salvo em", fileName, "| Goroutines: ", record[1], "| Tempo:", record[2])
}

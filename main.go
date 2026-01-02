package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	correctas := 0

	csvFilename := flag.String("csv", "problems.csv", "Archivo CSV con preguntas")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el CSV:", err)
		return
	}

	timer := time.NewTimer(5 * time.Second)

	for _, record := range records {

		answerCh := make(chan string)
		fmt.Printf("¿%s? ", record[0])

		go func() {
			var respuesta string
			fmt.Scanln(&respuesta)
			answerCh <- respuesta
		}()

		select {
		case respuesta := <-answerCh:
			if respuesta == record[1] {
				fmt.Println("Correctooou")
				correctas++
			} else {
				fmt.Println("Maaal")
			}
		case <-timer.C:
			fmt.Println("\n⏰ ¡Tiempo terminado!")
			fmt.Printf("Respuestas correctas: %d de %d\n", correctas, len(records))
			return
		}

	}

}

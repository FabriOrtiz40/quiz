package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	correctas := 0

	csvFilename := flag.String("csv", "problems.csv", "Archivo CSV con preguntas")
	timeLimit := flag.Int("limit", 30, "Tiempo límite del quiz en segundos")
	shuffle := flag.Bool("shuffle", false, "Mezclar el orden de las preguntas")

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

	if *shuffle {
		rand.Seed(time.Now().UnixNano()) // importante para que sea aleatorio real
		rand.Shuffle(len(records), func(i, j int) {
			records[i], records[j] = records[j], records[i]
		})
	}

	fmt.Println("¡Presione Enter para comenzar el quiz!")
	fmt.Scanln()
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

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
			user := strings.TrimSpace(strings.ToLower(respuesta))
			correcta := strings.TrimSpace(strings.ToLower(record[1]))
			if user == correcta {
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

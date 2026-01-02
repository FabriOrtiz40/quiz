package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	correctas := 0

	file, err := os.Open("problems.csv")
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

	for _, record := range records {
		var respuesta string

		fmt.Printf("¿%s? ", record[0])
		fmt.Scanln(&respuesta)

		if respuesta == record[1] {
			fmt.Printf("Correctooou \n ")
			correctas++
		} else {
			fmt.Printf("Maaal \n ")
		}

	}
	total := len(records)
	fmt.Printf("Respuestas correctas: %d de %d\n", correctas, total)

}

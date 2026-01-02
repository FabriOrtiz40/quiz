package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
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
        fmt.Println("Pregunta:", record[0], "| Respuesta:", record[1])
    }
}
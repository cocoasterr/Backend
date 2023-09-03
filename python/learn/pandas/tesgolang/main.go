package main

import (
	"fmt"
	"time"

	"github.com/tealeg/xlsx"
)

func processSheet(sheet *xlsx.Sheet) [][]string {
    var data [][]string

    for _, row := range sheet.Rows {
        var rowData []string
        for _, cell := range row.Cells {
            text := cell.String()
            rowData = append(rowData, text)
        }
        data = append(data, rowData)
    }

    return data
}

func main() {
    // Nama file Excel yang akan dibaca
    filename := "../Orders-With Nulls.xlsx"

    // Memulai mengukur waktu eksekusi
    startTime := time.Now()

    // Buka file Excel
    xlFile, err := xlsx.OpenFile(filename)
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }

    var allData [][]string

    // Loop melalui setiap sheet dalam file Excel
    for _, sheet := range xlFile.Sheets {
        data := processSheet(sheet)
        allData = append(allData, data...)
    }

    // Menghentikan mengukur waktu eksekusi
    endTime := time.Now()
    elapsedTime := endTime.Sub(startTime)
    fmt.Printf("Waktu Eksekusi: %v\n", elapsedTime)

    // Sekarang Anda memiliki semua data dalam 'allData'
}

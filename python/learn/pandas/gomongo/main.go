package main

import (
	"context"
	"fmt"
	"time"

	"github.com/tealeg/xlsx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func readExcel(filename string) ([][]string, time.Duration, error) {
    // Buka file Excel
    xlFile, err := xlsx.OpenFile(filename)
    if err != nil {
        return nil, 0, err
    }

    var allData [][]string

    // Loop melalui setiap sheet dalam file Excel
    for _, sheet := range xlFile.Sheets {
        data := processSheet(sheet)
        allData = append(allData, data...)
    }

    return allData, 0, nil
}

func insertToMongoDB(data [][]string) (time.Duration, error) {
    // Koneksi ke MongoDB
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        return 0, err
    }

    // Memilih database dan koleksi MongoDB
    database := client.Database("mydb")
    collection := database.Collection("mycollection")

    // Memulai mengukur waktu eksekusi penyisipan ke MongoDB
    startTime := time.Now()

    // Iterasi dan sisipkan data ke MongoDB
    for _, row := range data {
        document := make(map[string]interface{})
        for i, value := range row {
            columnName := fmt.Sprintf("column%d", i+1)
            document[columnName] = value
        }

        _, err := collection.InsertOne(context.TODO(), document)
        if err != nil {
            return 0, err
        }
    }

    // Menghentikan mengukur waktu eksekusi penyisipan
    endTime := time.Now()
    elapsedTime := endTime.Sub(startTime)

    // Menutup koneksi ke MongoDB
    client.Disconnect(context.TODO())

    return elapsedTime, nil
}

func main() {
    // Nama file Excel yang akan dibaca
    filename := "../Orders-With Nulls.xlsx"

    // Memulai mengukur waktu eksekusi
    totalStartTime := time.Now()

    // Baca Excel
    excelData, _, err := readExcel(filename)
    if err != nil {
        fmt.Printf("Error reading Excel: %v\n", err)
        return
    }

    // Menghentikan mengukur waktu eksekusi membaca Excel
    readEndTime := time.Now()
    readElapsedTime := readEndTime.Sub(totalStartTime)
    fmt.Printf("Read Executed in: %v\n", readElapsedTime)

    // Insert ke MongoDB
    insertDuration, err := insertToMongoDB(excelData)
    if err != nil {
        fmt.Printf("Error inserting to MongoDB: %v\n", err)
        return
    }

    // Menghentikan mengukur waktu eksekusi penyisipan
    totalEndTime := time.Now()
    totalElapsedTime := totalEndTime.Sub(totalStartTime)

    fmt.Printf("Migration to Mongodb in: %v\n", insertDuration)
    fmt.Printf("Total Execute time: %v\n", totalElapsedTime)
}

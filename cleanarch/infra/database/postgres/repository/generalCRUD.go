package repository

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"

	"github.com/cocoasterr/cleanarch/domain"
)

func PGGeneralCreate(r RepositoryContract, Listkey reflect.Type, data domain.DataModel) error {
	var key, value string
	var values []interface{}

	var placeIndex int
	num := Listkey.NumField()
	for i := 0; i < num; i++ {
		var v interface{}
		field := Listkey.Field(i)
		if field.Name != "Id" {
			if field.Name == "CreatedAt" || field.Name == "UpdatedAt" {
				v = time.Now()
			} else {
				v = reflect.ValueOf(data).Elem().Field(i).Interface()
			}
			placeIndex++
			key += fmt.Sprintf("%s ", field.Name)
			value += fmt.Sprintf("$%d", placeIndex)
			values = append(values, v)
			if i < num-1 {
				key += ", "
				value += ", "
			}
		}
	}
	tableName := data.TableName()
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, key, value)
	_, err := r.DB().Exec(query, values...)
	return err
}


func PGGeneralGetAll(r RepositoryContract,  page, limit int, data domain.DataModel) (map[string]interface{}, error){
	offset := (page - 1) * limit
	tb_name := data.TableName()
	query := fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", tb_name, limit, offset)
	rows ,_ := r.DB().Query(query)

	dataRes, err := getRes(rows)
	if err != nil{
		return nil, err
	}

	var total int
	queryTotal := fmt.Sprintf("SELECT COUNT(id) FROM %s ", tb_name)
	err = r.DB().QueryRow(queryTotal).Scan(&total)
	if err != nil{
		return nil, err
	}


	result:= map[string]interface{}{
			"message": "success",
			"data":    dataRes,
			"page":    page,
			"total":   total,
	}
	return result, nil
}


func PGGeneralGetById(r RepositoryContract, data domain.DataModel, id int)(map[string]interface{}, error){
	tb_name := data.TableName()
    query := fmt.Sprintf("SELECT * FROM %s WHERE id = %d", tb_name, id)
	rows ,_ := r.DB().Query(query)
	dataRes, err := getRes(rows)
	if err != nil{
		return nil, err
	}

    result := map[string]interface{}{
        "message": "success",
        "data":    dataRes,
        "page":    1,
        "total":   1,
    }
    return result, nil
}

func getRes(rows *sql.Rows) ([]interface{} , error){
	listColumn,_ := rows.Columns()
	var dataRes []interface{}
	for rows.Next() {
		dest := make([]interface{}, len(listColumn))
		for i := range listColumn {
			dest[i] = new(interface{})
		}
		err := rows.Scan(dest...)
		if err != nil {
			return nil, err
		}

		itemMap := make(map[string]interface{})
		for i, colName := range listColumn {
			itemMap[colName] = *dest[i].(*interface{})
		}

		dataRes = append(dataRes, itemMap)
	}
	return dataRes, nil
}
package PGRepository

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
)


func GeneralCreate(ctx context.Context, listkey reflect.Type, tbName string, domain interface{})(string, []interface{}){
	var key, value []string
	var values []interface{}
	
	num := listkey.NumField()
	for i := 0; i < num; i++ {
		field := listkey.Field(i).Name
		key = append(key, field)
		if strings.ToLower(field) == "id"{
			values = append(values, uuid.New().String())
		}else if strings.ToLower(field) == "createdat"||strings.ToLower(field) == "updatedat"{
			values = append(values, time.Now())
		}else if strings.ToLower(field) == "createdby"||strings.ToLower(field) == "updatedby"{
			values = append(values, "admin")
		}else{
			values = append(values, reflect.ValueOf(domain).Elem().Field(i).Interface())
		}
		value = append(value, fmt.Sprintf("$%d", i+1))
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tbName, strings.Join(key, ", "), strings.Join(value, ", "))
	return query, values
}

func GeneralIndex(r *sql.DB,  ctx context.Context, page, limit int, tbName string) (map[string]interface{}, error){
	offset := (page-1)*limit
	query := fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", tbName, limit, offset)
	rows,_ := r.QueryContext(ctx, query)
	data, _ := getResIndex(rows)
	
	var total int
	query_total :=fmt.Sprintf("SELECT COUNT(id) FROM %s", tbName)
	if err := r.QueryRowContext(ctx, query_total).Scan(&total); err != nil{
		return nil, err
	}
	resp:= map[string]interface{}{
		"message": "success!",
		"data": data,
		"total":total,
	}
	return resp, nil
}

func getResIndex(rows *sql.Rows)([]interface{}, error){
	lisColumn ,_ := rows.Columns()
	var res []interface{}
	for rows.Next(){
		dest := make([]interface{}, len(lisColumn))
		for i:= range lisColumn{
			dest[i] = new(interface{})
		}
		err := rows.Scan(dest...)
		if err != nil{
			return nil, err
		}
		itermap:= make(map[string]interface{})
		for i, colName:= range lisColumn{
			itermap[colName] = *dest[i].(*interface{})
		}
		res = append(res, itermap)
	} 
	return res, nil
}
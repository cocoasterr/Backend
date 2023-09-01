package PGRepositories

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GeneralCreate(ctx context.Context, listkey reflect.Type, tbName string, domain interface{})(string, []interface{}){
	var key []string
	var values []interface{}
	var value []string

	num := listkey.NumField()

	for i := 0; i < num; i++ {
		field := listkey.Field(i).Name
		key = append(key, field)
		if strings.ToLower(field) == "id"{
			 values = append(values, uuid.New().String())
		}else if strings.ToLower(field) == "createdat"|| strings.ToLower(field) == "updatedat"{
			values = append(values, time.Now())
		}else if strings.ToLower(field) == "createdby"||strings.ToLower(field) == "updatedby"{
			//bypass user
			values = append(values, "admin")
		}else{
			values = append(values, reflect.ValueOf(domain).Elem().Field(i).Interface())
		}
		value = append(value, fmt.Sprintf("$%d", i+1))
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s)", tbName, strings.Join(key, ", "), strings.Join(value, ", "))
	return query, values
}

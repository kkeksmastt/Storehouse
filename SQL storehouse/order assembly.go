package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type my_data struct {
	order_num        int
	prod_id          int
	amount           int
	prod_name        string
	shelf_name       string
	additional_shelf []string
}

func isAvailable(str string) bool {
	for i := range os.Args {
		if os.Args[i] == str {
			return true
		}
	}
	return false
}

func main() {

	connStr := "user=postgres password=my_pass dbname=123 sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select ord.order_number, ord.product as product_id, ord.amount, goods.product_name, sh.shelf_name, shelfs.additional_shelf " +
		"from orders ord, goods, goods_shelf shelfs, storehouse sh " +
		"where ord.product = goods.product_id and ord.product = shelfs.product and shelfs.main_shelf = sh.shelf_id " +
		"order by shelf_name, product_id")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	bks := make([]my_data, 0)
	for rows.Next() {
		var bk my_data
		err := rows.Scan(&bk.order_num, &bk.prod_id, &bk.amount, &bk.prod_name, &bk.shelf_name, pq.Array(&bk.additional_shelf))
		if err != nil {
			panic(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	if len(bks) > 0 {

		rows_shelf, err := db.Query("select * from storehouse")
		if err != nil {
			panic(err)
		}
		defer rows_shelf.Close()

		var shelfs = map[string]string{}

		for rows_shelf.Next() {
			var arr [2]string
			err := rows_shelf.Scan(&arr[0], &arr[1])
			if err != nil {
				panic(err)
			}
			shelfs[arr[0]] = arr[1]
		}

		var shelf string
		for i := range bks {

			if !isAvailable(strconv.Itoa(bks[i].order_num)) {
				continue
			}

			if shelf != bks[i].shelf_name {
				shelf = bks[i].shelf_name
				println("\n===Стеллаж ", shelf)
			} else {
				fmt.Printf("\n")
			}

			fmt.Printf("%s (id = %d)\nЗаказ %d, %d шт\n", bks[i].prod_name, bks[i].prod_id, bks[i].order_num, bks[i].amount)

			if len(bks[i].additional_shelf) > 0 {
				fmt.Printf("доп стеллаж: ")
				for ii := 0; ii < len(bks[i].additional_shelf); ii++ {
					if name, ok := shelfs[bks[i].additional_shelf[ii]]; ok {
						fmt.Printf("%s ", name)
					}
				}
				fmt.Printf("\n")
			}
		}
	}
}

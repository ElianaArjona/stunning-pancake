package caller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"pancake/services"
)

func ReportBG(c *services.BankConfig) (*services.BgEntries, error) {

	var entries = &services.BgEntries{}

	if c.FileType == "excel" {

		rawData, err := c.ParseExcelFile()
		if err != nil {
			return nil, err
		}

		for _, raw := range rawData {
			entry, err := raw.ProcessRawToEntry(c)
			if err != nil {
				return nil, err
			}

			entries.Entries = append(entries.Entries, entry)
		}
	}

	jsonData, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

	return entries, nil

}

func ConnectionDB() {
	dsn := "postgresql://eliarjona:rr0waTcWxC6ZAire2sTq3w@waning-ocelot-3452.g8z.cockroachlabs.cloud:26257/pancake?sslmode=verify-full"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	// id UUID PRIMARY KEY,
	// bank    COLLATE,
	// month   INT,
	// year    INT,
	// income  FLOAT,
	// expense FLOAT,
	// balance FLOAT,
	// serviceType COLLATE

	if _, err := db.Exec(
		`INSERT INTO report (id, bank, month, year, income, expense,balance, serviceType) 
		VALUES ('1', 'banco general', 1, 2020 , 10.0, 15.0, 5.0, 'service');`); err != nil {
		log.Fatal(err)
	}

	// Select Statement.
	rows, err := db.Query("select id, bank, income FROM report;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {

		var bank string
		var income int64

		if err := rows.Scan(&bank, &income); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Employee Id : %d \t Employee Name : %s\n", bank, income)
	}

	// ctx := context.Background()
	// conn, err := pgx.Connect(ctx, dsn)

	// defer conn.Close(context.Background())
	// if err != nil {
	// 	log.Fatal("failed to connect database", err)
	// }

	// var now time.Time
	// err = conn.QueryRow(ctx, "SELECT NOW()").Scan(&now)
	// if err != nil {
	// 	log.Fatal("failed to execute query", err)
	// }

	// fmt.Println(now)
}

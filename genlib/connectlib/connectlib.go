package connectlib

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/sijms/go-ora"
)

func Read_ConnectSTR(v_file string) (v_host string, v_port int, v_user string, v_password string, v_dbname string) {
	// Open File
	file, err := os.OpenFile(v_file, os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	// Blok sonunda dosyayı kapat
	defer file.Close()

	// Satır satır oku
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		satir := scanner.Text()

		index := strings.Index(satir, "=")

		if satir[:index] == "host" {
			v_host = satir[index+1:]
		}

		if satir[:index] == "port" {
			v_port_x, err := strconv.Atoi(satir[index+1:])

			if err != nil {
				fmt.Println("Hata:", err)
				return
			}
			v_port = v_port_x

		}

		if satir[:index] == "user" {
			v_user = satir[index+1:]
		}

		if satir[:index] == "password" {
			v_password = satir[index+1:]
		}

		if satir[:index] == "dbname" {
			v_dbname = satir[index+1:]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return v_host, v_port, v_user, v_password, v_dbname
}

func Connect_Oracle(password string) (*sql.DB, error) {
	// Replace these values with your database credentials
	username := "system"
	serviceName := "XE"

	// The SSH tunnel is forwarding the remote Oracle port to the local port 2525
	host := "DESKTOP-5N72M6L.local"
	port := 1521

	// Set the number of prefetch rows
	//prefetchRows := 1000

	// Create the connection string
	connStr := fmt.Sprintf(
		"oracle://%s:%s@%s:%d/%s",
		username, password, host, port, serviceName)

	fmt.Println("Opening connection to db using go-ora:")
	fmt.Println(connStr)

	// Connect to the database
	db, err := sql.Open("oracle", connStr)
	if err != nil {
		fmt.Printf("HATA (1: %v", err)
		return nil, err
	}

	fmt.Println("HATA ALMADI....")
	fmt.Println(db)

	rows, err := db.Query("SELECT * FROM dvsys.CODE$")
	if err != nil {
		fmt.Println("HATA ALDIK....")
		fmt.Println(db)
		return nil, err
	}
	defer rows.Close()

	fmt.Println("HATA ALMADI....")
	fmt.Println(rows)
	/*
		// Test the connection
		err = db.Ping()
		if err != nil {
			fmt.Println("HATA ALDIK....")
			fmt.Println(db)
			fmt.Println(err)
			db.Close()
			return nil, err
		}
	*/
	return db, nil
}

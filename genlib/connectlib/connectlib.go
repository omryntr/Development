package connectlib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

package main

import (
	// Not using yet
	"fmt"
	"os"

	"news.com/event/genlib"
	"news.com/event/genlib/connectlib"
)

// Connection Variables
var Host_file = "ConnConfig.cfg"

func main() {
	host, port, user, password, dbname := connectlib.Read_ConnectSTR("ConnConfig.cfg")
	fmt.Fprintln(os.Stdout, "Host:", host, "Port:", port, "user:", user, "password:", password, "dbname:", dbname)
	fmt.Println(genlib.CelToFah(5))
}

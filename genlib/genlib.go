package genlib

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
	"news.com/event/genlib/convertlib"
)

func Check_File_Info(file string) {

	file_info, err := os.Lstat(file)

	if err != nil {
		fmt.Println(err)
		return
	}

	if !strings.Contains(file_info.Name(), "PSG_") {
		fmt.Println("Bu Dosya Aktarım Dosyası değil...")
	}

	if file_info.IsDir() {
		fmt.Println("It is a directory")
	}

}

func OpenandChange_File(file string) {

	read, err := os.ReadFile(file)

	if err != nil {
		panic(err)
	}

	fmt.Println(read)

	newContents := Convert_Data_Type(string(read))

	fmt.Println(read)

	new_file_name := file + "_Changed"

	err = os.WriteFile(new_file_name, []byte(newContents), 0)

	if err != nil {
		panic(err)
	}
}

func OpenFileFromScreen() string {
	result, err := cfdutil.ShowOpenFileDialog(cfd.DialogConfig{
		Title: "Open A File For Migration",
		Role:  "OpenFileExample",
		FileFilters: []cfd.FileFilter{
			{
				DisplayName: "Text Files (*.txt)",
				Pattern:     "*.txt",
			},
			{
				DisplayName: "Image Files (*.jpg, *.png)",
				Pattern:     "*.jpg;*.png",
			},
			{
				DisplayName: "All Files (*.*)",
				Pattern:     "*.*",
			},
		},
		SelectedFileFilterIndex: 2,
		FileName:                "file.txt",
		DefaultExtension:        "txt",
	})
	if err == cfd.ErrorCancelled {
		log.Fatal("Dialog was cancelled by the user.")
	} else if err != nil {
		log.Fatal(err)
	}
	log.Printf("Chosen file: %s\n", result)
	return result
}

// Covert Oracle Data Type to PostgresSQL Data Type

func Convert_Data_Type(file_scope string) string {
	newContents := convertlib.Converter(file_scope, "BFILE", "VARCHAR(255)")
	newContents1 := convertlib.Converter(newContents, "BINARY_FLOAT", "REAL")
	newContents2 := convertlib.Converter(newContents1, "BINARY_DOUBLE", "DOUBLE PRECISION")
	newContents3 := convertlib.Converter(newContents2, "BLOB", "BYTEA")
	newContents4 := convertlib.Converter(newContents3, "CLOB", "TEXT")
	newContents5 := convertlib.Converter(newContents4, "DATE", "TIMESTAMP")
	newContents6 := convertlib.Converter(newContents5, "LONG", "TEXT")
	newContents7 := convertlib.Converter(newContents6, "LONG RAW", "BYTEA")
	newContents8 := convertlib.Converter(newContents7, "NCHAR VARYING", "VARCHAR")
	newContents9 := convertlib.Converter(newContents8, "NCLOB", "TEXT")
	newContents10 := convertlib.Converter(newContents9, "NVARCHAR2", "VARCHAR")
	newContents11 := convertlib.Converter(newContents10, "RAW", "BYTEA")
	newContents12 := convertlib.Converter(newContents11, "ROWID", "CHAR(10)")
	newContents13 := convertlib.Converter(newContents12, "VARCHAR2", "VARCHAR")
	newContents14 := convertlib.Converter(newContents13, "XMLTYPE", "XML")

	fmt.Println(string(file_scope[1]))
	return newContents14
}

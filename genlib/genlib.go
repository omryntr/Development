package genlib

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
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

	newContents := strings.Replace(string(read), "VARCHAR2", "NUMBERrrrrrrrrr", -1)

	fmt.Println(newContents)
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

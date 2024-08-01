package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getfilenames(number int) []string {
	list := []string{}

	for i := 0; i < number; i++ {
		filename := fmt.Sprintf("data/abc%v.txt", i)
		list = append(list, filename)
	}

	return list
}

func createdir(dirname string) error {
	abspath, err := filepath.Abs(dirname)
	if err != nil {
		fmt.Println("An error occured getting abs")
		fmt.Println(err)
		return err
	}

	_, err = os.Stat(abspath)
	if err != nil {

		if !os.IsNotExist(err) {
			fmt.Println("An error occured getting stat")
			fmt.Println(err)
			return err
		}
	}

	if os.IsNotExist(err) {
		filemode := os.ModeDir
		err := os.MkdirAll(abspath, filemode)
		if err != nil {
			fmt.Println("An error occured creating directory")
			fmt.Println(err)
			return err
		}

		return nil
	} else {
		fmt.Println("Directory already existing")
		return nil
	}
}

func createfile(filename string) error {
	dirname := filepath.Dir(filename)
	fmt.Println("Using directory", dirname)
	err := createdir(dirname)
	if err != nil {
		fmt.Println("An error occured creating directory")
		fmt.Println(err)
		return err
	}

	osfile, err := os.Create(filename)
	if err != nil {
		fmt.Println("An error occured")
		fmt.Println(err)
		return err
	}

	_, err = osfile.WriteString("bcd")
	if err != nil {
		fmt.Println("An error occured while writing string")
		fmt.Println(err)
		return err
	}
	return nil
}

func createfiles(lista []string) error {
	fmt.Println("Start creating files")
	for _, filename := range lista {
		fmt.Println("Start creating file", filename)
		createfile(filename)
	}
	fmt.Println("Creating files completed")
	return nil
}

func main() {
	lista := getfilenames(20)
	createfiles(lista)
}

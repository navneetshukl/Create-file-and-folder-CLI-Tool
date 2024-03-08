package helpers

import (
	"fmt"
	"log"
	"os"
)

// CreateFolder function create a folder of specific name
func CreateFolder(name string) error {
	if err := os.Mkdir(name, os.ModePerm); err != nil {
		log.Println(fmt.Sprintf("Error in creating the folder named %s. Error is % v :", name, err.Error()))
		return err
	} else {
		return nil
	}

}

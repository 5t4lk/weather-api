package settings

import (
	"fmt"
	"os"
)

func Eject() error {
	data, err := os.ReadFile("settings.txt")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

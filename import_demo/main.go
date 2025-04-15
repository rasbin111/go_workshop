package main

import (
	"fmt"
	"import_demo/office"
)

func main() {
	office.PrintOfficeName()
	fmt.Println("Imported value:", office.OfficeName)
}

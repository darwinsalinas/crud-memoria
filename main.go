package main

import (
	"crud-memoria/controllers"
)

func main() {
	controllers.PrintHeader()
	controllers.FillInitialData(50)
	controllers.PrintMenu()
}

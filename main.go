package main

import "github.com/nebiros/go-rest-test/models"

func main() {
	fw := models.NewFacturaWeb()
	clientId := "43577753"
	fw.BillByClientId(clientId)
	// fmt.Println(config.ApplicationEnv)
}

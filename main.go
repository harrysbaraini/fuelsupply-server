package main

import (
	"harrysbaraini/fuelsupply/modules/supplies"
	"harrysbaraini/fuelsupply/system"
)

func main() {
	// Register router
	router := system.RegisterRouter()

	// Create version 1 routing
	v1 := router.Group("/v1")
	{
		supplies.Service.Register(v1)
	}

	router.Run(":8888")
}

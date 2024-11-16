package routes

import (
	"fmt"
	"os"
)

func Init() {
	r := Router()
	serverAddr := fmt.Sprintf("0.0.0.0%s", os.Getenv("SERVER_PORT"))
	fmt.Println("Server running on http://" + serverAddr)
	err := r.Run(serverAddr)
	if err != nil {
		panic(err)
	}
}

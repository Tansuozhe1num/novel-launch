package main

import (
	"log"
	"novel-launch/novel/Biu"
	"novel-launch/novel/server/httpServer"
)

func main() {
	r := httpServer.SetupRouter()

	err := Biu.MustInit()
	if err != nil {
		log.Fatal(err)
	}

	addr := ":8080"
	log.Println("listening at http://localhost" + addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}

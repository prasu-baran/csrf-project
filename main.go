package main

import (
"log"
"github.com/prasu-baran/csrf-project/db"
"github.com/prasu-baran/csrf-project/server"
"github.com/prasu-baran/csrf-project/server/middleware/myjwt"

)
var host = "localhost"
var port = "9000"

func main(){
  db.initDB()

  jwtErr := myJwt.InitJWT()
  if jwtErr !=nil{
	log.Println("Error initializing the Jwt !")
	log.Fatal(jwtErr)
  }

  serverErr := server.StartServer(host,port)
  if serverErr !=nil{
	log.Println("Error starting the server !")
	log.Fatal(serverErr)
  }
}
package main

import (
	"database/sql"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting server")
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			listener.Close()
			continue
		}
		go handleConnection(conn)
		go databaseCheck()
	}
}
func handleConnection(conn net.Conn)string {
	defer conn.Close()
	for {
		input := make([]byte, 1024*4)
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			break
		}
		//читаем что пришло
		source := string(input[0:n])
		//проверка на наличие в словаре
		fmt.Println(source)
		conn.Write([]byte(source + "\n"))
		quary := source
		// test
		//conn.Write([]byte("you're send " + source))
		//возвращаем для поиска в бд
		return quary
	}
	return "ok"
}


func databaseCheck(){
	db, err := sql.Open("mysql", "clicklock:clicklock@localhost")
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to db -- ok!")
	var version string
	//db.Query("SELECT hash FROM clicklock.locks")
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to: ", version)
	defer db.Close()
}


//todo
/* добавить gracefull shutdown
добавить поддержку sql
добавить домофон
предохранитель
 */
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/CSProjectsAvatar/distri-systems/utils"
)

func getStudents() []string {
	file, err := os.Open("class.txt")
	utils.CheckErr(err)
	defer file.Close()

	// read file line by line
	scanner := bufio.NewScanner(file)
	var students []string
	for scanner.Scan() {
		// append line to students
		students = append(students, scanner.Text())
	}
	return students
}

func setStudents(students []string) {
	file, err := os.Create("class.txt")
	utils.CheckErr(err)
	defer file.Close()

	// write non-duplicated students to file
	for _, student := range utils.Unique(students) {
		_, err = file.WriteString(student + "\n")
		utils.CheckErr(err)
	}
}

func main() {
	log.Println("Server started.")

	stds := getStudents()

	log.Println("Listening...")
	ln, err := net.Listen("tcp", ":8000")
	utils.CheckErr(err)
	defer ln.Close()

	clientCh := getClientCh(ln)

	for {
		go welcomeStudent(<-clientCh, &stds)
	}
}

// Welcomes a student to the class and saves its name if necessary.
func welcomeStudent(conn net.Conn, students *[]string) {
	defer conn.Close()

	msg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println("couldn't read from connection:", err.Error())
		return
	}
	*students = append(*students, string(msg))
	setStudents(*students)

	_, err = conn.Write([]byte(fmt.Sprintln("Welcome to class,", msg)))
	if err != nil {
		log.Println("couldn't write:", err.Error())
	}
}

// Returns a channel which emites connections from clients when accepted by server.
func getClientCh(ln net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	go func() {
		for {
			conn, err := ln.Accept()

			if err != nil {
				log.Println("couldn't accept connection:", err.Error())
				continue
			}
			log.Println("Accepted connection from", conn.RemoteAddr())
			ch <- conn
		}
	}()
	return ch
}

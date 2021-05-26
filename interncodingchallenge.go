package main

import "fmt"

func getURL(url string) {
	fmt.Println("https://dludtke.zendesk.com/api/v2/tickets.json?page[size]=25")
}

func getUser(username string) string {
	fmt.Scanln("Please enter your username: ")
	return username
}

func getPassword(password string) string {
	fmt.Scanln("Please enter your password: ")
	return password
}

func welcomeMessage(selection string) string {
	fmt.Println("Welcome to the ticket viewer!")
	fmt.Println("To view all tickets, please press enter.")
	fmt.Println("To view a single ticket, please specify the ticket ID")
	fmt.Println("To exit, please press x.")
	fmt.Scanln()

	return selection
}

package main

import "fmt"

type User struct {
	Name, Tempat string
	Age		 int	
}

func showuser(u User)  {
	fmt.Println("Nama:", u.Name)
	fmt.Println("Tempat:", u.Tempat)
	fmt.Println("Umur:", u.Age)	
}

func main() {
	Pengguna1 := User{
		Name:   "Afif",
		Tempat: "Bandung",
		Age:    21,	
	}

	showuser(Pengguna1)
}

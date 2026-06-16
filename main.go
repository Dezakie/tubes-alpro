package main

import (
	"fmt"
	"tubes-alpro/securepass"
)

func main() {
	var input int

	for {
		fmt.Println("\n === Password Manager === \n")
		fmt.Println("Opsi pilihan: ")
		fmt.Println("[1] List akun yang tersimpan")
		fmt.Println("[2] Ganti password aplikasi SecurePass")
		fmt.Println("[3] Keluar dari Aplikasi\n")
		fmt.Print("Masukan opsi pilihan: ")

		fmt.Scanln(&input)
		switch input {
		case 1:
			securepass.ListAkun()
		case 2:
			securepass.GantiPasswordApp()
		case 3:
			fmt.Println("Terima kasih telah menggunakan Password Manager")
			return
		default:
			fmt.Println("Opsi tidak valid, silakan coba lagi")
		}
	}
}

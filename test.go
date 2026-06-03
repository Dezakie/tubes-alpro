package main
import "fmt"

type data struct {
	namaLayanan string
	username string
	password string
	tanggalEdit string
}

var dataAkun []data
var passwordApp string = "admin123"

func main() {
	var input int
	fmt.Println("\n === Password Manager === \n")
	fmt.Println("Opsi pilihan: ")
	fmt.Println("[1] List akun yang tersimpan")
	fmt.Println("[2] Ganti password akun Password Manager")
	fmt.Println("[3] Keluar\n")
	fmt.Print("Masukan opsi pilihan: ")

	fmt.Scanln(&input)
	switch input {
	case 1:
		listAkun()
	case 2:
		gantiPasswordApp()
	case 3:
		fmt.Println("Terima kasih telah menggunakan Password Manager")
		return
	default:
		fmt.Println("Opsi tidak valid, silakan coba lagi.")
	}
}
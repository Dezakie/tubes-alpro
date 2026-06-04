package main
import "fmt"

type data struct {
	namaLayanan string
	username string
	password string
	// tanggalEdit string
}

var dataAkun []data 
var passwordApp string = "admin123"
dataAkun = []data{
	{namaLayanan: "Steam", username: "Dezakie", password "Balikpapan2006" },
	{namaLayanan: "Google", username: "dzakyangasli@gmail.com", password "Balikpapan2006" },
	{namaLayanan: "Instagram", username: "dzaky_ono", password "Balikpapan2006" }
}

func main() {
	var input int

	for {

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
		fmt.Println("Opsi tidak valid, silakan coba lagi")
	}
	}
}

func listAkun() {
	var p string
	
	for p != passwordApp {
		fmt.Print("\nMasukan password PManager")
		fmt.Scanln(&p)

		if p == passwordApp {
			for i := 0, i < len(dataAkun), i++ {
			fmt.Println(i, dataAkun[i].namaLayanan, dataAkun[i].username)
			}

			fmt.Println("\nOpsi pilihan: ")
			fmt.Println("[1] Lihat password tersimpan")
			fmt.Println("[2] Ubah password tersimpan")
			fmt.Println("[3] Tambah akun")
			fmt.Println("[4] Hapus akun tersimpan\n")
			fmt.Print("Masukan opsi pilihan: ")

			for {
			var input string
			fmt.Scanln(&input)

			switch input {
			case 1:
				lihatPassword()
				break
			case 2:
				gantiPassword()
				break
			case 3:
				tambahAkun()
				break
			case 4:
				hapusAkun()
				break
			default:
				fmt.Println("Opsi tidak valid, silakan coba lagi")
			}

			}

		} else {
			fmt.Println("Password salah, coba lagi!")
		}
	}
}

func lihatPassword() {
	var l string
	for {
	fmt.Print("\nMasukan nama layanan: ")
	fmt.Scanln(&l)
	for i := 0; i < len(dataAkun); i++ {
		if l == dataAkun[i].namaLayanan {
			fmt.Println("\n === Hasil ===")
			fmt.Printf("Detail username: %s \n", dataAkun[i].username)
			fmt.Printf("Detail password: %s \n", dataAkun[i].password)
		}
	}

	}
}

func gantiPassword() {

}

func tambahAkun() {

}

func hapusAkun() {

}


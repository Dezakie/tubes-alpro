package securepass

import "fmt"

type Account struct {
	NamaLayanan string
	Username    string
	Password    string
}

var DataAkun = []Account{
	{NamaLayanan: "Steam", Username: "Dezakie", Password: "Balikpapan2006"},
	{NamaLayanan: "Google", Username: "dzakyangasli@gmail.com", Password: "Balikpapan2006"},
	{NamaLayanan: "Instagram", Username: "dzaky_ono", Password: "Balikpapan2006"},
}

var PasswordApp string = "admin123"

func GantiPasswordApp() {
	var p string
	fmt.Print("Masukan password pw_manager saat ini: ")
	fmt.Scanln(&p)
	if p == PasswordApp {
		fmt.Print("Masukan password pw_manager baru: ")
		fmt.Scanln(&PasswordApp)
		fmt.Println("Password pw_manager berhasil diganti")
	} else {
		fmt.Println("Password salah woe, ulangi")
	}
}

func ListAkun() {
	var p string
	for p != PasswordApp {
		fmt.Print("\nMasukkan password PManager: ")
		fmt.Scanln(&p)

		if p == PasswordApp {
			for {
				TampilkanStatistik()

fmt.Println("\n=== Pilihan Menu Utama ===")
				fmt.Println("[1] Lihat password berdasarkan layanan (Sequential + Insertion)")
				fmt.Println("[2] Lihat password berdasarkan username (Binary + Selection)")
				fmt.Println("[3] Sorting berdasarkan alfabet (Selection)")
				fmt.Println("[4] Sorting berdasarkan waktu input (Insertion)")
				fmt.Println("[5] Tambah akun")
				fmt.Println("[6] Ubah password")
				fmt.Println("[7] Hapus akun")
				fmt.Println("[8] Keluar")
				fmt.Print("Masukkan opsi pilihan: ")

				var input int
				fmt.Scanln(&input)

				switch input {
				case 1:
					CariBerdasarkanLayanan()
				case 2:
					CariBerdasarkanUsername()
				case 3:
					SortBerdasarkanAlfabet()
				case 4:
					SortBerdasarkanWaktu()
				case 5:
					TambahAkun()
				case 6:
					GantiPassword()
				case 7:
					HapusAkun()
				case 8:
					// Kembali / Keluar dari perulangan menu
					return
				default:
					fmt.Println("Opsi tidak valid, silakan coba lagi")
				}
			}
		} else {
			fmt.Println("Password salah, coba lagi!")
		}
	}
}

func GantiPassword() {
}

func TambahAkun() {
}

func HapusAkun() {
}

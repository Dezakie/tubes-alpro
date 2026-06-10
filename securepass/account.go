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
				fmt.Println("\n=== Daftar Akun Tersimpan ===")
				if len(DataAkun) == 0 {
					fmt.Println("Belum ada akun yang tersimpan.")
				} else {
					for i := 0; i < len(DataAkun); i++ {
						fmt.Printf("%d. Layanan: %s | Username: %s\n", i+1, DataAkun[i].NamaLayanan, DataAkun[i].Username)
					}
				}

				fmt.Println("\n--- Menu Kelola Akun ---")
				fmt.Println("[1] Cari Data Akun (Search)")
				fmt.Println("[2] Urutkan Data Akun (Sort)")
				fmt.Println("[3] Lihat Statistik & Kekuatan Sandi")
				fmt.Println("[4] Tambah Akun baru")
				fmt.Println("[5] Ubah Password Akun")
				fmt.Println("[6] Hapus Akun Tersimpan")
				fmt.Println("[7] Kembali ke Menu Utama\n")
				fmt.Print("Masukkan opsi pilihan: ")

				var input int
				fmt.Scanln(&input)

				switch input {
				case 1:
					MenuCari() 
				case 2:
					MenuUrut()
				case 3: 
					TampilkanStatistik()
				case 4:
					TambahAkun()
				case 5:
					GantiPassword()
				case 6:
					HapusAkun()
				case 7:
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

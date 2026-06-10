package securepass

import "fmt"

type Account struct {
	NamaLayanan string
	Username    string
	Password    string
}

const MaxAkun = 100
var DataAkun [MaxAkun]Account
var JumlahAkun int
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
				if JumlahAkun == 0 {
					fmt.Println("Belum ada akun yang tersimpan.")
				} else {
					for i := 0; i < JumlahAkun; i++ {
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
				case 4:
					TambahAkun()
				case 3: 
				  TampilkanStatistik()
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
	var u, p string
	fmt.Print("Masukan username atau email: ")
	fmt.Scanln(&u)

	for i := 0; i < JumlahAkun; i++ {
		if DataAkun[i].Username == u {
			fmt.Print("Masukan password pw_manager: ")
			for p != PasswordApp {
				fmt.Scanln(&p)
				if p == DataAkun[i].Password {
					fmt.Print("Masukan password baru: ")
					fmt.Scanln(&DataAkun[i].Password)
					fmt.Println("Password akun berhasil diubah!")
				} else {
					fmt.Println("Password salah woe, ulangi")
				}
			}
		}
	}
}

func TambahAkun() {
	var u, p string
	fmt.Print("Masukan username atau email: ")
	fmt.Scanln(&u)
	fmt.Print("Masukan password pw_manager: ")
	
	for p != PasswordApp {
		fmt.Scanln(&p)
		if p == PasswordApp {
			fmt.Print("Masukan password akun: ")
			fmt.Scanln(&DataAkun[JumlahAkun].Password)
			DataAkun[JumlahAkun].Username = u
			
			fmt.Print("Masukan nama layanan: ")
			fmt.Scanln(&DataAkun[JumlahAkun].NamaLayanan)

			JumlahAkun++ 
			fmt.Println("Akun baru berhasil ditambahkan!")
		} else {
			fmt.Println("Password salah woe, ulangi")
		}
	}
}

func HapusAkun() {
	var u, p string
	fmt.Print("Masukan username atau email: ")
	fmt.Scanln(&u)

	for i := 0; i < JumlahAkun; i++ {
		if DataAkun[i].Username == u {
			fmt.Print("Masukan password pw_manager: ")
			fmt.Scanln(&p)
			if p == PasswordApp {
				DataAkun[i] = Account{} 
				fmt.Println("Akun berhasil dihapus")
			} else {
				fmt.Println("Password salah woe, ulangi")
			}
		}
	}
}

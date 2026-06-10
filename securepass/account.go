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

func init() {
	DataAkun[0] = Account{NamaLayanan: "Steam", Username: "Dezakie", Password: "Balikpapan2006"}
	DataAkun[1] = Account{NamaLayanan: "Google", Username: "dzakyangasli@gmail.com", Password: "Balikpapan2006"}
	DataAkun[2] = Account{NamaLayanan: "Instagram", Username: "dzaky_ono", Password: "Balikpapan2006"}
	JumlahAkun = 3 
}

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
			for i := 0; i < JumlahAkun; i++ {
				fmt.Println(i+1, DataAkun[i].NamaLayanan, DataAkun[i].Username)
			}

			for {
				fmt.Println("\n--- Menu Kelola Akun ---")
				fmt.Println("[1] Lihat password tersimpan")
				fmt.Println("[2] Ubah password tersimpan")
				fmt.Println("[3] Tambah akun")
				fmt.Println("[4] Hapus akun tersimpan")
				fmt.Println("[5] Kembali ke Menu Utama\n")
				fmt.Print("Masukkan opsi pilihan: ")

				var input int
				fmt.Scanln(&input)

				switch input {
				case 1:
					LihatPassword() 
				case 2:
					GantiPassword()
				case 3:
					TambahAkun()
				case 4:
					HapusAkun()
				case 5:
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

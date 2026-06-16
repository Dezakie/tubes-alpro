package securepass

import (
	"fmt"
	"time"
)

type Account struct {
	NamaLayanan string
	Username    string
	Password    string
	TimeEdit    time.Time

	LayananNorm  string
	UsernameNorm string
}

var DataAkun = []Account{
	{NamaLayanan: "Google", Username: "dzakyangasli@gmail.com", Password: "Balikpapan2006"},
	{NamaLayanan: "Google", Username: "dzaky@student.telkomuniversity.ac.id", Password: "EmailKampus999#"},
	{NamaLayanan: "Steam", Username: "Dezakie", Password: "Balikpapan2006"},
	{NamaLayanan: "Steam", Username: "dzaky_smurf", Password: "hanyahuruf"},
	{NamaLayanan: "Instagram", Username: "dzaky_second", Password: "12345678"},
	{NamaLayanan: "Github", Username: "coder_dzaky", Password: "G1thub!Secret99"},
	{NamaLayanan: "TikTok", Username: "vt_dzaky", Password: "Tiktok1#"},
	{NamaLayanan: "Spotify", Username: "musik_dzaky", Password: "LaguEnak2026"},
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
				normSemua()
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
func spasi() {
	fmt.Println("\n============================================================================================ ")
}

func GantiPassword() {
	var target, target1 string

	spasi()
	fmt.Print("\nMasukan nama layanan: ")
	fmt.Scanln(&target)
	target = normalisasi(target)

	fmt.Println(" ")
	fmt.Println("Layanan dan username yang tersedia:")
	cariAkunLayanan(target)
	fmt.Println(" ")

	fmt.Print("Masukan username: ")
	fmt.Scanln(&target1)
	target1 = normalisasi(target1)

	for i := 0; i < len(DataAkun); i++ {
		if DataAkun[i].LayananNorm == target && DataAkun[i].UsernameNorm == target1 {
			fmt.Print("Masukan password baru: ")
			fmt.Scanln(&DataAkun[i].Password)
			DataAkun[i].TimeEdit = time.Now()
			fmt.Println("Password berhasil diubah..")
			return
		}
	}
}

func TambahAkun() {
	var layanan, username, password string

	fmt.Print("\nMasukan jumlah akun yang ingin ditambahkan: ")
	var jumlah int
	fmt.Scanln(&jumlah)

	for i := 0; i < jumlah; i++ {
		spasi()
		fmt.Print("\nMasukan nama layanan: ")
		fmt.Scanln(&layanan)

		fmt.Print("Masukan username: ")
		fmt.Scanln(&username)

		fmt.Print("Masukan password: ")
		fmt.Scanln(&password)

		DataAkun = append(DataAkun, Account{
			NamaLayanan: layanan,
			Username:    username,
			Password:    password,
			TimeEdit:    time.Now(),
			LayananNorm:  normalisasi(layanan),
			UsernameNorm: normalisasi(username),
		})

		fmt.Println("\nAkun berhasil ditambahkan..")
	}
}
func HapusAkun() {
	var target, target1 string

	spasi()
	fmt.Print("\nMasukan nama layanan: ")
	fmt.Scanln(&target)
	target = normalisasi(target)

	fmt.Println(" ")
	fmt.Println("Layanan dan username yang tersedia:")
	cariAkunLayanan(target)
	fmt.Println(" ")

	fmt.Print("Masukan username: ")
	fmt.Scanln(&target1)
	target1 = normalisasi(target1)

	for i := 0; i < len(DataAkun); i++ {
		if DataAkun[i].LayananNorm == target && DataAkun[i].UsernameNorm == target1 {
			DataAkun = append(DataAkun[:i], DataAkun[i+1:]...)
			fmt.Println("Akun berhasil dihapus..")
			return
		}
	}
	fmt.Println("Akun tidak ditemukan..")
}

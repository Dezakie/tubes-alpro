package main
import (
	"fmt"
	"time"
)

type data struct {
	layanan string
	username string
	password string
	tEdit time.Time

	layananNorm string
	usernameNorm string
	klasifikasi string
}

var dataAkun = []data{
	{layanan: "Steam", username: "Dezakie", password: "Bpp2006"}, 
	{layanan: "Steam", username: "gugu", password: "Bpp2006"},
	{layanan: "Google", username: "dezakie@gmail.com", password: "Bpp2006"},
	{layanan: "Google", username: "gugu@gmail.com", password: "Bpp2006"},
	{layanan: "Instagram", username: "deazakie", password: "Bpp2006"},
	{layanan: "Instagram", username: "gugu", password: "Bpp2006"},
}

var appPassword string = "admin123"
var lemah, sedang, kuat int

func nyoba() {
	var input int
	for i := 0; i < len(dataAkun); i++ {
		dataAkun[i].tEdit = time.Now()
	}

	for {
		spasi()
		fmt.Println("\n=== Aplikasi Pengelola Kata Sandi Pribadi (SecurePass) ===\n")
		fmt.Println("Opsi pilihan:")  
		fmt.Println("[1] Lihat akun tersimpan")
		fmt.Println("[2] Ganti password aplikasi SecurePass")
		fmt.Println("[3] Keluar aplikasi\n")
		fmt.Print("Masukan opsi pilihan: ")

		fmt.Scanln(&input)

		switch input {
		case 1:
			lihatAkun()
		case 2:
			gantiAppPassword()
		case 3:
			fmt.Println("\nTerima kasih telah menggunakan SecurePass\n")
			return
		default:
			fmt.Println("\nOpsi tidak valid, silahkan coba lagi!")
		}
	}
}

func gantiAppPassword() {
	var input string

	for (input != appPassword) {
		spasi()
		fmt.Print("\nMasukan password: ")
		fmt.Scanln(&input)

		switch {
		case input == appPassword:
			fmt.Print("Masukan password baru: ")
			fmt.Scanln(&appPassword)
			input = appPassword
			fmt.Println("Password telah diperbarui..")
		default:
			fmt.Println("Password salah, silahkan coba lagi!")
		}
	}
}

func lihatAkun() {
	var input string

	for (input != appPassword) {
		spasi()
		fmt.Print("\nMasukan password: ")
		fmt.Scanln(&input)

		switch {
		case input == appPassword:
			for {
				var input1 int
				normSemua()
				
				spasi()
				fmt.Println("\n=== Panel Statistik ===")
				fmt.Printf("Total akun tersimpan: %d \n", len(dataAkun))
				fmt.Println("\nKlasifikasi kekuatan kata sandi: ")
				fmt.Printf("Lemah  : %d Akun\n", lemah)
				fmt.Printf("Sedang : %d Akun\n", sedang)
				fmt.Printf("Kuat   : %d Akun\n\n", kuat)
				fmt.Println("=== Pilihan Menu Utama ===")
				fmt.Println("[1] Lihat password berdasarkan layanan (Sequential + Selection)")
				fmt.Println("[2] Lihat password berdasarkan username (Binary + Insertion)\n")
				fmt.Println("[3] Sorting berdasarkan alfabet (Selection)")
				fmt.Println("[4] Sorting berdasarkan waktu input (Insertion)\n") //belum
				fmt.Println("[5] Tambah akun") //baikin
				fmt.Println("[6] Ubah password\n")
				fmt.Println("[7] Hapus akun ")
				fmt.Println("[8] Keluar\n")
				fmt.Print("Masukan opsi pilihan: ")

				fmt.Scanln(&input1)

				switch input1 {
				case 1:
					fmt.Println("lihatPassLayanan()")
				case 2:
					fmt.Println("lihatPassUsername()")
				case 3:
					sortingAlfabet()
				case 4:
					fmt.Println("sortingWInput()")
				case 5:
					tambahAkun()
				case 6:
					ubahPassword()
				case 7:
					hapusAkun()
				case 8:
					fmt.Println("Kembali ke halaman utama..")
					return
				default:
					fmt.Println("\nOpsi tidak valid, silahkan coba lagi!")
				}
			}
		default:
			fmt.Println("Password salah, silahkan coba lagi!")
		}
	}
}

func spasi() {
	fmt.Println("\n============================================================================================ ")
}

func hapusAkun() {
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

	for i := 0; i < len(dataAkun); i++ {
		if dataAkun[i].layananNorm == target && dataAkun[i].usernameNorm == target1{
			dataAkun = append(dataAkun[:i], dataAkun[i+1:] ...)
			fmt.Println("Akun berhasil dihapus..")
			return
		}
	}
}

func cariAkunLayanan(n string) {
	for i := 0; i < len(dataAkun); i++ {
		if dataAkun[i].layananNorm == n {
			fmt.Println(dataAkun[i].layanan, "     ", dataAkun[i].username)
		}
	}
}

func normalisasi(s string) string {
	runes := []rune(s)

	for i := 0; i < len (runes); i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 32
		}
		
	}

	return string(runes)
}

func normSemua() {
	for i := 0; i < len(dataAkun); i++ {
		dataAkun[i].layananNorm = normalisasi(dataAkun[i].layanan)
		dataAkun[i].usernameNorm = normalisasi(dataAkun[i].username)
	}
}

func ubahPassword() {
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

	for i := 0; i < len(dataAkun); i++ {
		if dataAkun[i].layananNorm == target && dataAkun[i].usernameNorm == target1{
			fmt.Print("Masukan password baru: ")
			fmt.Scanln(&dataAkun[i].password)
			dataAkun[i].tEdit = time.Now()
			fmt.Println("Password berhasil diubah..")
			return
		}
	}
}

func tambahAkun() {
	var layanan, username, password string

	spasi()
	fmt.Print("\nMasukan nama layanan: ")
	fmt.Scanln(&layanan)

	fmt.Print("Masukan username: ")
	fmt.Scanln(&username)

	fmt.Print("Masukan password: ")
	fmt.Scanln(&password)

	dataAkun = append(dataAkun, data{
		layanan: layanan,
		username: username,
		password: password,
		tEdit: time.Now(),
	})

	fmt.Println("\nAkun berhasil ditambahkan..")
}

func sortingAlfabet() {
	n := len(dataAkun)

	for i := 0; i < n; i++ {
		minIndex := i
		
		for j := i + 1; j < n; j++ {
			if dataAkun[j].layananNorm < dataAkun[minIndex].layananNorm {
				minIndex = j
			}
		}
		
		dataAkun[i], dataAkun[minIndex] = dataAkun[minIndex], dataAkun[i]
	}

	fmt.Println("\nIni hasil sorting berdasarkan alfabet nama layanan:")
	fmt.Printf("%-15s %-20s %-15s\n", "Layanan", "Username", "Terakhir diedit")
	fmt.Println("========================================================")
	for i := 0; i < n; i++ {
		fmt.Printf("%-15s %-20s %-15s\n", 
		dataAkun[i].layanan, 
		dataAkun[i].username, 
		dataAkun[i].tEdit.Format("15:04:05 02-01-2006"))
	}
}
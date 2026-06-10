package securepass

import "fmt"

func TampilkanStatistik() {
	fmt.Println("\n=== Statistik Password Manager ===")
	fmt.Printf("Total akun yang tersimpan: %d akun\n", JumlahAkun)
	
	fmt.Println("\n[Fitur Klasifikasi Kekuatan Sandi belum dibuat]")
}

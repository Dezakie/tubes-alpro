package securepass

import "fmt"

// Sequential Search + Insertion Sort (A-Z)
func CariBerdasarkanLayanan() {
	SortBerdasarkanLayanan() 
	var l string
	fmt.Print("\nMasukan nama layanan: ")
	fmt.Scanln(&l)
	
	ditemukan := false
	fmt.Println("\n === Hasil (Sequential + Insertion) ===")
	for i := 0; i < len(DataAkun); i++ {
		if l == DataAkun[i].NamaLayanan {
			fmt.Printf("Detail layanan: %s \n", DataAkun[i].NamaLayanan)
			fmt.Printf("Detail password: %s \n", DataAkun[i].Password)
			fmt.Print("\n")
			ditemukan = true
		}
	}
	
	if !ditemukan {
		fmt.Println("Akun tidak ditemukan.")
	}
}

// Binary Search + Selection Sort (A-Z)
func CariBerdasarkanUsername() {
	SortBerdasarkanUsername()
	var target string
	fmt.Print("\nMasukan nama username: ")
	fmt.Scanln(&target)
	kiri := 0
	kanan := len(DataAkun) - 1
	ditemukan := false

	fmt.Println("\n === Hasil (Binary + Selection) ===")
	for kiri <= kanan && !ditemukan {
		tengah := (kiri + kanan) / 2

		if DataAkun[tengah].Username == target {
			fmt.Printf("Detail username: %s \n", DataAkun[tengah].Username)
			fmt.Printf("Detail password: %s \n", DataAkun[tengah].Password)
			fmt.Print("\n")
			ditemukan = true
		} else if DataAkun[tengah].Username < target {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	
	if !ditemukan {
		fmt.Println("Akun tidak ditemukan, atau data belum terurut secara alfabetis")
	}
}

// Cari berdasarkan layanan dan username
func cariAkunLayanan(n string) {
	for i := 0; i < len(DataAkun); i++ {
		if DataAkun[i].LayananNorm == n {
			fmt.Println(DataAkun[i].NamaLayanan, "     ", DataAkun[i].Username)
		}
	}
}
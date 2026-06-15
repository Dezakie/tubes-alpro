package securepass

import "fmt"

// Sequential Search
func CariBerdasarkanLayanan() {
	var l string
	fmt.Print("\nMasukan nama layanan: ")
	fmt.Scanln(&l)
	
	ditemukan := false
	
	for i := 0; i < len(DataAkun); i++ {
		if l == DataAkun[i].NamaLayanan {
			fmt.Println("\n === Hasil (Sequential + Insertion) ===")
			fmt.Printf("Detail username: %s \n", DataAkun[i].Username)
			fmt.Printf("Detail password: %s \n", DataAkun[i].Password)
			ditemukan = true
		}
	}
	
	if !ditemukan {
		fmt.Println("Akun tidak ditemukan.")
	}
}

// Binary Search
func CariBerdasarkanUsername() {
	var target string
	fmt.Print("\nMasukan nama username: ")
	fmt.Scanln(&target)
	kiri := 0
	kanan := len(DataAkun) - 1
	ditemukan := false

	for kiri <= kanan && !ditemukan {
		tengah := (kiri + kanan) / 2

		if DataAkun[tengah].Username == target {
			fmt.Println("\n === Hasil (Binary + Selection) ===")
			fmt.Printf("Detail layanan: %s \n", DataAkun[tengah].NamaLayanan)
			fmt.Printf("Detail password: %s \n", DataAkun[tengah].Password)
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


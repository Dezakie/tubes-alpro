package securepass

import (
  "fmt"
  "strings"
  )

func MenuCari() {
	var opsi int
	fmt.Println("\n--- Metode Pencarian ---")
	fmt.Println("[1] Sequential Search")
	fmt.Println("[2] Binary Search")
	fmt.Println("[3] Kembali")
	fmt.Print("Pilih metode: ")
	fmt.Scanln(&opsi)

	switch opsi {
	case 1:
		SequentialSearch()
	case 2:
		BinarySearch()
	case 3:
	  return
	default:
		fmt.Println("Opsi tidak valid.")
	}
}

func SequentialSearch() {
	var l string
	fmt.Print("\nMasukan nama layanan: ")
	fmt.Scanln(&l)
	
	ditemukan := false
	
	for i := 0; i < JumlahAkun; i++ {
		if strings.ToLower(l) == strings.ToLower(DataAkun[i].NamaLayanan) {
			fmt.Println("\n === Hasil (Sequential) ===")
			fmt.Printf("Detail username: %s \n", DataAkun[i].Username)
			fmt.Printf("Detail password: %s \n", DataAkun[i].Password)
			ditemukan = true
		}
	}
	
	if !ditemukan {
		fmt.Println("Akun tidak ditemukan.")
	}
}

func BinarySearch() {
  var target string
  fmt.Print("\nMasukan nama layanan: ")
  fmt.Scanln(&target)
  
  kiri := 0
  kanan := JumlahAkun - 1
  ditemukan := false
  
  targetLayanan := strings.ToLower(target)
  
  for kiri <= kanan && !ditemukan {
    tengah := (kiri+kanan) / 2
    
    tengahLayanan := strings.ToLower (DataAkun[tengah].NamaLayanan)
    
    if tengahLayanan == targetLayanan {
      fmt.Println("\n === Hasil (Binary) ===")
      fmt.Printf("Detail username: %s \n", DataAkun[tengah].Username)
			fmt.Printf("Detail password: %s \n", DataAkun[tengah].Password)
			ditemukan = true
    } else if DataAkun[tengah].NamaLayanan < target {
      kiri = tengah + 1
    } else {
      kanan = tengah - 1
    }
  }
  if !ditemukan {
    fmt.Println("Akun tidak ditemukan, atau data belum terurut secara alfabetis")
  }
}
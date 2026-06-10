package securepass

import "fmt"

func LihatPassword() {
	var l string
	for {
		fmt.Print("\nMasukan nama layanan: ")
		fmt.Scanln(&l)
		for i := 0; i < JumlahAkun; i++ {
			if l == DataAkun[i].NamaLayanan {
				fmt.Println("\n === Hasil ===")
				fmt.Printf("Detail username: %s \n", DataAkun[i].Username)
				fmt.Printf("Detail password: %s \n", DataAkun[i].Password)
			}
		}
	}
}

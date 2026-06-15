package securepass

import "fmt"

func TampilkanStatistik() {
	var lemah, sedang, kuat int
	for _, akun := range DataAkun {
		switch KlasifikasiSandi(akun.Password) {
		case "lemah":
			lemah++
		case "sedang":
			sedang++
		case "kuat":
			kuat++
		}
	}
	fmt.Println("\n=== Statistik Password Manager ===")
	fmt.Printf("Total akun yang tersimpan: %d akun\n", len(DataAkun))
	
	fmt.Println("\nKlasifikasi Kekuatan Sandi: ")
	fmt.Printf("- Sandi Lemah: %d akun\n", lemah)
	fmt.Printf("- Sandi Sedang: %d akun\n", sedang)
	fmt.Printf("- Sandi Kuat: %d akun\n", kuat)
}

func KlasifikasiSandi(password string) string {
	var upper, lower, digit, special int
	if len(password) < 8 {
		return "lemah"
	}

	for _, char := range password {
		switch {
		case char >= 'A' && char <= 'Z':
			upper++
		case char >= 'a' && char <= 'z':
			lower++
		case char >= '0' && char <= '9':
			digit++
		default:
			special++
		}
	}

	if upper > 0 && lower > 0 && digit > 0 && special > 0 {
		return "kuat"
	}

	if (upper > 0 && lower > 0) || (upper > 0 && digit > 0) || (lower > 0 && digit > 0) {
		return "sedang"
	}

	return "lemah"
}
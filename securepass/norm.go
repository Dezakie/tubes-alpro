package securepass

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
	for i := 0; i < len(DataAkun); i++ {
		DataAkun[i].LayananNorm = normalisasi(DataAkun[i].NamaLayanan)
		DataAkun[i].UsernameNorm = normalisasi(DataAkun[i].Username)
	}
}

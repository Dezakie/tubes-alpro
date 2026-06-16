package securepass

import "fmt"

// Selection Sort
func SortBerdasarkanAlfabet() {
	fmt.Println("\n[Fitur Pengurutan berdasarkan alfabet belum dibuat]")
}

// Insertion Sort
func SortBerdasarkanWaktu() {
	fmt.Println("\n[Fitur Pengurutan berdasarkan waktu belum dibuat]")
}

// Insertion Sort + Sequential Search (A-Z)
func SortBerdasarkanLayanan() {
	for i := 1; i < len(DataAkun); i++ {
		key := DataAkun[i]
		j := i - 1
		for j >= 0 && DataAkun[j].NamaLayanan > key.NamaLayanan {
			DataAkun[j+1] = DataAkun[j]
			j--
		}
		DataAkun[j+1] = key
	}
}

// Selection Sort + Binary Search (A-Z)
func SortBerdasarkanUsername() {
	fmt.Println("\n[Fitur Pengurutan berdasarkan username belum dibuat]")
}
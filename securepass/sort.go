package securepass

import "fmt"

// Selection Sort
func SortBerdasarkanAlfabet() {
	n := len(DataAkun)

	for i := 0; i < n; i++ {
		minIndex := i
		
		for j := i + 1; j < n; j++ {
			if DataAkun[j].LayananNorm < DataAkun[minIndex].LayananNorm {
				minIndex = j
			}
		}
		
		DataAkun[i], DataAkun[minIndex] = DataAkun[minIndex], DataAkun[i]
	}

	fmt.Println("\nIni hasil sorting berdasarkan alfabet nama layanan:")
	fmt.Printf("%-15s %-20s %-15s\n", "Layanan", "Username", "Terakhir diedit")
	fmt.Println("========================================================")
	for i := 0; i < n; i++ {
		fmt.Printf("%-15s %-20s %-15s\n", 
		DataAkun[i].NamaLayanan, 
		DataAkun[i].Username, 
		DataAkun[i].TimeEdit.Format("15:04:05 02-01-2006"))
	}
}

// Insertion Sort
func SortBerdasarkanWaktu() {
    n := len(DataAkun)
    for i := 1; i < n; i++ {
        key := DataAkun[i]
        j := i - 1

        for j >= 0 && key.TimeEdit.After(DataAkun[j].TimeEdit	) {
            DataAkun[j+1] = DataAkun[j]
            j--
        }

        DataAkun[j+1] = key
    }

    fmt.Println("\nHasil sorting berdasarkan waktu edit:\n")
    fmt.Printf("%-15s %-20s %-20s\n", "Layanan", "Username", "Terakhir diedit")
    fmt.Println("========================================================")
    for i := 0; i < n; i++ {
        fmt.Printf("%-15s %-20s %-20s\n",
            DataAkun[i].NamaLayanan,
            DataAkun[i].Username,
            DataAkun[i].TimeEdit.Format("15:04:05 02-01-2006"))
    }
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
	for i := 0; i < len(DataAkun); i++ {
		minIndex := i
		for j := i + 1; j < len(DataAkun); j++ {
			if DataAkun[j].Username < DataAkun[minIndex].Username {
				minIndex = j
			}
		}
		DataAkun[i], DataAkun[minIndex] = DataAkun[minIndex], DataAkun[i]
	}
}
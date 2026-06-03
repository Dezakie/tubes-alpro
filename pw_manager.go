package main
import "fmt"

type data struct {
	username string
	password string
}

var dataAkun [999]data
var pwapp string = "admin123"

func main() {
	var input int
	for {
		fmt.Print(
	"Selamat datang di password manager\n\n" +

	"pilih opsi: \n" +
	"1. melihat password tersimpan \n" +
	"2. ganti password tersimpan \n" +
	"3. manambah akun \n" +
	"4. menghapus akun \n" +
	"5. ganti password pw_manager \n" +
	"6. keluar \n\n" +

	"ketikan opsi yang ingin dipilih: ")
	fmt.Scanln(&input)

	switch input {
	case 1:
		lihatPassword()
	case 2:
		gantiPassword()
	case 3:
		tambahAkun()
	case 4:
		hapusAkun()
	case 5:
		gantiPasswordPwManager()
	case 6:
		fmt.Println("Terima kasih sudah menggunakan password manager, sampai jumpa!")
		return

	default:
		fmt.Println("Salah opsi bang, pilih yang bener")
	}

	}
}

func lihatPassword() {
	var p string
	fmt.Print("Masukan password pw_manager saat ini: ")

	for p != pwapp {
	fmt.Scanln(&p)

	if p == pwapp {
		fmt.Println("List akun yang tersimpan: ")
		for i := 0; i < len(dataAkun); i++ {
			fmt.Printf("%d. %s\n", i + 1, dataAkun[i].username)
		}
		
		var u string
		fmt.Print("Masukan username atau email: ")
		fmt.Scanln(&u)

		for i := 0; i < len(dataAkun); i++ {
		if dataAkun[i].username == u {
			fmt.Printf("\nUsername: %s\n" +
			"Password: %s\n", dataAkun[i].username, dataAkun[i].password)
		}
	}
	} else {
		fmt.Println("Password salah woe, ulangi")
	}
	}
}

func gantiPassword() {
	var u, p string
	fmt.Print("Masukan username atau email: ")
	fmt.Scanln(&u)

	for i := 0; i < len(dataAkun); i++ {
		if dataAkun[i].username == u {
			fmt.Print("Masukan password pw_manager: ")
		
			for p != pwapp{
				fmt.Scanln(&p)

				if p == dataAkun[i].password {
					fmt.Print("Masukan password baru: ")
					fmt.Scanln(&dataAkun[i].password)
				} else {
					fmt.Println("Password salah woe, ulangi")
				}
			}
		}
	}
}

func tambahAkun() {
	var u, p string
	fmt.Print("Masukan username atau email: ")
	fmt.Scanln(&u)
	fmt.Print("Masukan password pw_manager: ")
	for p != pwapp {
		fmt.Scanln(&p)
		if p == pwapp {
			fmt.Print("Masukan password akun: ")
			fmt.Scanln(&dataAkun[len(dataAkun) - 1].password)
			dataAkun[len(dataAkun) - 1].username = u
		} else {
			fmt.Println("Password salah woe, ulangi")
		}
	}
}

func hapusAkun() {
	var u, p string
	fmt.Print("Masukan username atau email: ")
	fmt.Scanln(&u)

	for i := 0; i < len(dataAkun); i++ {
		if dataAkun[i].username == u {
			fmt.Print("Masukan password pw_manager: ")
			fmt.Scanln(&p)
			if p == pwapp {
				dataAkun[i] = data{} // Mengosongkan data akun
				fmt.Println("Akun berhasil dihapus")
			} else {
				fmt.Println("Password salah woe, ulangi")
			}
		}
	}
}

func gantiPasswordPwManager() {
	var p string
	fmt.Print("Masukan password pw_manager saat ini: ")
	fmt.Scanln(&p)
	if p == pwapp {
		fmt.Print("Masukan password pw_manager baru: ")
		fmt.Scanln(&pwapp)
		fmt.Println("Password pw_manager berhasil diganti")
	} else {
		fmt.Println("Password salah woe, ulangi")
	}
}

// package main
// import "fmt"

// type data struct {
//     username string
//     password string
// }

// var dataAkun []data
// var pwapp string = "admin123"

// func main() {
//     var input int
//     for {
//         fmt.Print(
//             "\nSelamat datang di password manager\n\n" +
//                 "pilih opsi: \n" +
//                 "1. melihat password tersimpan \n" +
//                 "2. ganti password tersimpan \n" +
//                 "3. menambah akun \n" +
//                 "4. menghapus akun \n" +
//                 "5. ganti password pw_manager \n" +
//                 "6. keluar \n\n" +
//                 "ketikan opsi yang ingin dipilih: ")
//         fmt.Scanln(&input)

//         switch input {
//         case 1:
//             lihatPassword()
//         case 2:
//             gantiPassword()
//         case 3:
//             tambahAkun()
//         case 4:
//             hapusAkun()
//         case 5:
//             gantiPasswordPwManager()
//         case 6:
//             fmt.Println("Terima kasih sudah menggunakan password manager, sampai jumpa!")
//             return
//         default:
//             fmt.Println("Salah opsi bang, pilih yang bener")
//         }
//     }
// }

// func lihatPassword() {
//     var p string
//     fmt.Print("Masukan password pw_manager saat ini: ")

//     for p != pwapp {
//         fmt.Scanln(&p)
//         if p == pwapp {
//             fmt.Println("List akun yang tersimpan: ")
//             for i, akun := range dataAkun {
//                 if akun.username != "" {
//                     fmt.Printf("%d. %s\n", i+1, akun.username)
//                 }
//             }

//             var u string
//             fmt.Print("Masukan username atau email: ")
//             fmt.Scanln(&u)

//             for _, akun := range dataAkun {
//                 if akun.username == u {
//                     fmt.Printf("\nUsername: %s\nPassword: %s\n", akun.username, akun.password)
//                 }
//             }
//         } else {
//             fmt.Println("Password salah woe, ulangi")
//         }
//     }
// }

// func gantiPassword() {
//     var u, p string
//     fmt.Print("Masukan username atau email: ")
//     fmt.Scanln(&u)

//     for i := range dataAkun {
//         if dataAkun[i].username == u {
//             fmt.Print("Masukan password pw_manager: ")
//             for p != pwapp {
//                 fmt.Scanln(&p)
//                 if p == pwapp {
//                     fmt.Print("Masukan password baru: ")
//                     fmt.Scanln(&dataAkun[i].password)
//                     fmt.Println("Password berhasil diganti")
//                 } else {
//                     fmt.Println("Password salah woe, ulangi")
//                 }
//             }
//         }
//     }
// }

// func tambahAkun() {
//     var u, p string
//     fmt.Print("Masukan username atau email: ")
//     fmt.Scanln(&u)
//     fmt.Print("Masukan password pw_manager: ")
//     for p != pwapp {
//         fmt.Scanln(&p)
//         if p == pwapp {
//             var pass string
//             fmt.Print("Masukan password akun: ")
//             fmt.Scanln(&pass)
//             dataAkun = append(dataAkun, data{username: u, password: pass})
//             fmt.Println("Akun berhasil ditambahkan")
//         } else {
//             fmt.Println("Password salah woe, ulangi")
//         }
//     }
// }

// func hapusAkun() {
//     var u, p string
//     fmt.Print("Masukan username atau email: ")
//     fmt.Scanln(&u)

//     for i := range dataAkun {
//         if dataAkun[i].username == u {
//             fmt.Print("Masukan password pw_manager: ")
//             fmt.Scanln(&p)
//             if p == pwapp {
//                 dataAkun[i] = data{} // kosongkan data
//                 fmt.Println("Akun berhasil dihapus")
//             } else {
//                 fmt.Println("Password salah woe, ulangi")
//             }
//         }
//     }
// }

// func gantiPasswordPwManager() {
//     var p string
//     fmt.Print("Masukan password pw_manager saat ini: ")
//     fmt.Scanln(&p)
//     if p == pwapp {
//         fmt.Print("Masukan password pw_manager baru: ")
//         fmt.Scanln(&pwapp)
//         fmt.Println("Password pw_manager berhasil diganti")
//     } else {
//         fmt.Println("Password salah woe, ulangi")
//     }
// }

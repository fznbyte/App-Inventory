package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	var idlist []int
	var tanggallist []string
	var namalist []string
	var jumlahlist []int

	var pilihan int
	var jalan bool = true

	for jalan {

		clearScreen()

		fmt.Println("==========================================")
		fmt.Println("               INVENTORY APP              ")
		fmt.Println("==========================================")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Tampilkan Barang")
		fmt.Println("3. Update Barang")
		fmt.Println("4. Hapus Barang")
		fmt.Println("5. Cari Barang")
		fmt.Println("6. Keluar")
		fmt.Println("==========================================")
		fmt.Print("Masukkan Pilihan: ")
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			clearScreen()
			fmt.Println("==========================================")
			fmt.Println("              TAMBAH BARANG")
			fmt.Println("==========================================")

			for {
				var id, jumlah int
				var nama string

				fmt.Print("Masukkan ID Barang: ")
				fmt.Scanln(&id)
				fmt.Print("Masukkan Nama Barang: ")
				fmt.Scanln(&nama)
				fmt.Print("Masukkan Jumlah Barang: ")
				fmt.Scanln(&jumlah)

				if jumlah <= 0 {
					fmt.Println("\nGAGAL! Jumlah barang harus lebih dari 0.")
					fmt.Println("\nSilahkan masukkan jumlah yang valid.")
					continue
				}

				tanggal := time.Now().Format("02-01-2006 15:04:05")

				index := -1
				for i := 0; i < len(idlist); i++ {
					if idlist[i] == id {
						index = i
						break
					}
				}

				if index == -1 {
					idlist = append(idlist, id)
					tanggallist = append(tanggallist, tanggal)
					namalist = append(namalist, nama)
					jumlahlist = append(jumlahlist, jumlah)
					fmt.Println("\nBERHASIL! Barang baru ditambahkan.")
				} else {
					jumlahlist[index] += jumlah
					tanggallist[index] = tanggal
					fmt.Println("\nBERHASIL! Barang sudah ada, jumlah ditambahkan.")
				}

				var ulang string
				fmt.Print("\nTambah Barang Lagi? (y/n): ")
				fmt.Scanln(&ulang)
				if ulang == "n" {
					break
				}
				clearScreen()
			}

			fmt.Println("\nTekan ENTER untuk kembali ke menu...")
			fmt.Scanln()
		}

		if pilihan == 2 {
			clearScreen()
			fmt.Println("==========================================")
			fmt.Println("             DAFTAR SELURUH BARANG")
			fmt.Println("==========================================")

			if len(idlist) == 0 {
				fmt.Println("Tidak ada barang tersimpan.")
			} else {
				for i := 0; i < len(idlist); i++ {
					fmt.Println("------------------------------------------")
					fmt.Println("ID Barang      :", idlist[i])
					fmt.Println("Tanggal Masuk  :", tanggallist[i])
					fmt.Println("Nama Barang    :", namalist[i])
					fmt.Println("Jumlah Barang  :", jumlahlist[i])
				}
				fmt.Println("------------------------------------------")
			}

			fmt.Println("\nTekan ENTER untuk kembali...")
			fmt.Scanln()
		}

		if pilihan == 3 {
			clearScreen()
			fmt.Println("==========================================")
			fmt.Println("               UPDATE BARANG")
			fmt.Println("==========================================")

			if len(idlist) == 0 {
				fmt.Println("Tidak ada barang untuk diupdate.")
				fmt.Println("\nTekan ENTER untuk kembali...")
				fmt.Scanln()
			} else {
				var id int
				fmt.Print("Masukkan ID Barang: ")
				fmt.Scanln(&id)

				index := -1
				for i := 0; i < len(idlist); i++ {
					if idlist[i] == id {
						index = i
						break
					}
				}

				if index == -1 {
					fmt.Println("ID tidak ditemukan.")
				} else {
					for {
						fmt.Println("\n======= PILIHAN UPDATE =======")
						fmt.Println("1. Update Nama Barang")
						fmt.Println("2. Update Jumlah Barang")
						fmt.Println("3. Update Nama & Jumlah")
						fmt.Println("4. Kembali")
						fmt.Print("Pilih: ")

						var up int
						fmt.Scanln(&up)

						if up == 1 {
							var baru string
							fmt.Print("Nama Baru: ")
							fmt.Scanln(&baru)
							namalist[index] = baru
							fmt.Println("Nama berhasil diperbarui.")

						} else if up == 2 {
							var baruJumlah int
							fmt.Print("Jumlah Baru: ")
							fmt.Scanln(&baruJumlah)

							if baruJumlah <= 0 {
								fmt.Println("Jumlah tidak valid!")
							} else {
								jumlahlist[index] = baruJumlah
								fmt.Println("Jumlah berhasil diperbarui.")
							}

						} else if up == 3 {
							var baruNama string
							var baruJumlah int

							fmt.Print("Nama Baru: ")
							fmt.Scanln(&baruNama)
							fmt.Print("Jumlah Baru: ")
							fmt.Scanln(&baruJumlah)

							if baruJumlah <= 0 {
								fmt.Println("Jumlah tidak valid!")
							} else {
								namalist[index] = baruNama
								jumlahlist[index] = baruJumlah
								fmt.Println("Nama & jumlah berhasil diperbarui!")
							}

						} else if up == 4 {
							break
						} else {
							fmt.Println("Pilihan tidak valid.")
						}
					}
				}
				fmt.Println("\nTekan ENTER untuk kembali...")
				fmt.Scanln()
			}
		}

		if pilihan == 4 {
			clearScreen()
			fmt.Println("==========================================")
			fmt.Println("               HAPUS BARANG")
			fmt.Println("==========================================")

			if len(idlist) == 0 {
				fmt.Println("Tidak ada barang untuk dihapus.")
			} else {
				var id int
				fmt.Print("Masukkan ID Barang: ")
				fmt.Scanln(&id)

				index := -1
				for i := 0; i < len(idlist); i++ {
					if idlist[i] == id {
						index = i
						break
					}
				}

				if index == -1 {
					fmt.Println("ID tidak ditemukan.")
				} else {
					idlist = append(idlist[:index], idlist[index+1:]...)
					namalist = append(namalist[:index], namalist[index+1:]...)
					jumlahlist = append(jumlahlist[:index], jumlahlist[index+1:]...)
					tanggallist = append(tanggallist[:index], tanggallist[index+1:]...)
					fmt.Println("Barang berhasil dihapus!")
				}
			}

			fmt.Println("\nTekan ENTER untuk kembali...")
			fmt.Scanln()
		}

		if pilihan == 5 {
			clearScreen()
			fmt.Println("==========================================")
			fmt.Println("                CARI BARANG")
			fmt.Println("==========================================")

			var id int
			fmt.Print("Masukkan ID Barang: ")
			fmt.Scanln(&id)

			found := false
			for i := 0; i < len(idlist); i++ {
				if idlist[i] == id {
					fmt.Println("\n------------- DATA DITEMUKAN -------------")
					fmt.Println("ID     :", idlist[i])
					fmt.Println("Nama   :", namalist[i])
					fmt.Println("Jumlah :", jumlahlist[i])
					fmt.Println("Tanggal:", tanggallist[i])
					fmt.Println("------------------------------------------")
					found = true
					break
				}
			}

			if !found {
				fmt.Println("\nBarang tidak ditemukan.")
			}

			fmt.Println("\nTekan ENTER untuk kembali...")
			fmt.Scanln()
		}

		if pilihan == 6 {
			clearScreen()
			fmt.Println("==========================================")
			fmt.Println("         TERIMA KASIH TELAH MENGGUNAKAN")
			fmt.Println("              INVENTORY APP")
			fmt.Println("==========================================")
			jalan = false
		}
	}
}

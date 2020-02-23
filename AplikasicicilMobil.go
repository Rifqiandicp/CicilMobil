package main

import "fmt"

const Batas = 100

type Buyer struct {
	nama, alamat                                          string
	tahunBeli, poin, jumBayar, urut, umur                 int
	hargaMobil, lamaKredit, pilihanAngsuran, sisaAngsuran float64
	Waktu                                                 [36]waktuBeli
}

type waktuBeli struct {
	tgl, bln int
}

var (
	List                                               [Batas]Buyer
	nama, pilihan                                      string
	pilihan2, pilihan3, N, x, urutWaktu, jumPembayaran int
	trigg, triggTampilan                               bool
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func main() {

	tampilan()
	penutup()

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func backMenu() {

	trigg = false
	for !trigg {
		fmt.Println("\n\n-------------------------------------------------\n\n")
		fmt.Println("Apakah anda ingin kembali ke menu?\n\n 1. ya. \n 2. tidak. \n\njawaban anda: \n")
		pilihan2 = 3
		fmt.Scanln(&pilihan2)
		if pilihan2 == 1 {
			trigg = true
		} else if pilihan2 == 2 {
			fmt.Println("\n\n--------------------------------------------------\n\n")
			trigg = false
			for !trigg {
				fmt.Println("Apakah anda ingin keluar dari Aplikasi?\n\n 1. ya. \n 2. tidak. \n\njawaban anda: \n")
				pilihan3 = 3
				fmt.Scanln(&pilihan3)
				fmt.Println("\n--------------------------------------------------\n\n")
				if pilihan3 == 1 {
					fmt.Println("\n")
					trigg = true
					triggTampilan = true
				} else if pilihan3 == 2 {
					backMenu()
				} else {
					fmt.Println("Pilih antara angka 1 atau 2.")
				}
				fmt.Println("\n\n--------------------------------------------------\n\n")
			}
		} else {
			fmt.Println("Pilih antara angka 1 atau 2.")
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func penutup() {
	fmt.Println("\n\n//////////////////////////////////////////////////\n")
	fmt.Println("Terima Kasih Telah Berkunjung")
	fmt.Println("\n//////////////////////////////////////////////////\n")
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func tampilan() {
	triggTampilan = false
	for !triggTampilan {
		fmt.Println("\n\n////////////////////////////////////////////////////////////")
		fmt.Println("Selamat datang di aplikasi Kelola Cicilan Mobil!")
		fmt.Println("///////////////////////////////////////////////////////////\n\n")

		fmt.Println("1. Input data.")
		fmt.Println("2. searching.")
		fmt.Println("3. Buy.")
		fmt.Println("4. Print Data.")
		fmt.Println("5. History.")
		fmt.Println("6. Quit.")

		fmt.Println("\n\nSilakan pilih menu yang anda inginkan: \n")
		fmt.Scanln(&pilihan)

		fmt.Println("\n")
		if pilihan == "1" {
			inputData(&N)
		} else if pilihan == "2" {

			fmt.Println("//////////////////////////////////////////////////\n\n")
			fmt.Println("PENCARIAN DATA")
			fmt.Print("Nama yang ingin di cari: ")
			fmt.Scanln(&nama)
			searching(nama)
			x = searching(nama)

			if x == -1 {
				fmt.Println("\n\n-------------------------------------------------\n\n")
				fmt.Println("data tidak ada.")
				fmt.Println("\n\n-------------------------------------------------\n\n")
			} else {
				fmt.Println("\n\n--------------------------------------------------\n\n")
				fmt.Println("Nama\t\t\t: ", List[x].nama)
				fmt.Println("Alamat\t\t\t: ", List[x].alamat)
				fmt.Println("harga mobil\t\t: ", int(List[x].hargaMobil))
				fmt.Println("Umur\t\t\t: ", List[x].umur)
				fmt.Println("Pilihan angsuran\t: ", int(List[x].pilihanAngsuran))
				fmt.Println("Lama kredit\t\t: ", int(List[x].lamaKredit))
				fmt.Println("Tahun pembelian mobil\t: ", List[x].tahunBeli)
				fmt.Println("jumlah Poin saat ini\t: ", List[x].poin)
				fmt.Println("Sisa angsuran\t\t: ", int(List[x].sisaAngsuran))
				fmt.Println("\n\n-------------------------------------------------\n\n")
			}
			backMenu()
		} else if pilihan == "3" {
			bayar()
		} else if pilihan == "4" {

			printData()

		} else if pilihan == "5" {
			fmt.Println("atas nama: ")
			fmt.Scanln(&nama)
			searching(nama)
			x = searching(nama)

			if x == -1 {
				fmt.Println("\n\n//////////////////////////////////////////////////\n\n")
				fmt.Println("data tidak ada.")
				fmt.Println("\n\n//////////////////////////////////////////////////\n\n")
				backMenu()
			} else {
				sortingBulan(&List, x)
				history(&List, x)
				backMenu()
			}
		} else if pilihan == "6" {
			triggTampilan = true
		} else {
			fmt.Println("\n\n-------------------------------------------------\n\n")
			fmt.Println("Pilih menu sesuai dengan nomernya. (antara 1-6) ")
			fmt.Println("\n\n-------------------------------------------------\n\n")
		}
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func inputData(N *int) {

	var (
		triggData, triggData2, selesai bool
		ulang                          string
		a, b                           string
		g, h, d, temp                  int
		c, e, f, i                     float64
	)

	if *N < Batas {

		triggData = false
		selesai = false
		for !triggData {
			fmt.Println("\n\n-------------------------------------------------\n\n")

			fmt.Print("Masukkan nama \t\t\t: ")
			a = "-"
			temp = 1
			fmt.Scanln(&a)

			if a != "-"{
				temp = searching(a)
			}

			for temp != -1 && a != "-"{
				fmt.Println("\n\n-------------------------------------------------")
				fmt.Println("Nama sudah digunakan.")
				fmt.Println("-------------------------------------------------\n\n")
				fmt.Print("Masukkan nama \t\t\t: ")
				a = "-"
				temp = 1
				fmt.Scanln(&a)
				if a != "-"{
					temp = searching(a)
				}
			}	
			for a == "-" && temp !=-1{
				fmt.Println("\n\n-------------------------------------------------")
				fmt.Println("Mohon masukkan nama anda.")
				fmt.Println("-------------------------------------------------\n\n")
				fmt.Print("Masukkan nama \t\t\t: ")
				a = "-"
				temp = 1
				fmt.Scanln(&a)
				if a != "-"{
					temp = searching(a)
				}
				for temp != -1 && a != "-"{
					fmt.Println("\n\n-------------------------------------------------")
					fmt.Println("Nama sudah digunakan.")
					fmt.Println("-------------------------------------------------\n\n")
					fmt.Print("Masukkan nama \t\t\t: ")
					a = "-"
					temp = 1
					fmt.Scanln(&a)
					if a != "-"{
						temp = searching(a)
					}
				}	
			}
			b="-"
			fmt.Print("Masukkan alamat\t\t\t: ")
			fmt.Scanln(&b)
			fmt.Println("Pilihan harga mobil: \n\n 1. Tipe X : Rp. 300000000 \n\n 2. Tipe Y : Rp. 600000000 \n\n 3. Tipe Z : Rp. 900000000\n")
			fmt.Print("Masukkan harga Mobil\t\t: ")
			fmt.Scanln(&c)
			for c != 300000000 && (c != 600000000 && c != 900000000) && c != 1 && (c != 2 && c != 3) {
				fmt.Println("\n\n-------------------------------------------------")
				fmt.Println("Mohon isi harga mobil sesuai ketentuan.")
				fmt.Println("-------------------------------------------------\n\n")
				fmt.Println("Pilihan harga mobil: \n\n 1. Tipe X : Rp. 300000000 \n\n 2.Tipe Y : Rp. 600000000 \n\n 3. Tipe Z : Rp. 900000000\n\n")
				fmt.Print("Masukkan harga Mobil\t\t: ")
				fmt.Scanln(&c)
			}
			if c == 1 {
				c = 300000000
			} else if c == 2 {
				c = 600000000
			} else if c == 3 {
				c = 900000000
			}

			fmt.Print("Masukkan umur\t\t\t: ")
			fmt.Scanln(&d)
			for d < 1 {

				fmt.Println("\n\n-------------------------------------------------")
				fmt.Println("Mohon isi umur anda dengan benar.")
				fmt.Println("-------------------------------------------------\n\n")
				fmt.Print("Masukkan umur\t\t\t: ")
				fmt.Scanln(&d)
			}
			fmt.Println("Pilihan Angsuran: \n\n 1. 25000000 \n 2. 50000000\n")
			fmt.Print("Masukkan pilihan Angsuran \t: ")
			fmt.Scanln(&e)
			for e != 25000000 && (e != 50000000 && e != 1) && e != 2 {
				fmt.Println("\n\n-------------------------------------------------")
				fmt.Println("Mohon isi pilihan angsuran anda dengan benar.")
				fmt.Println("-------------------------------------------------\n\n")
				fmt.Print("Masukkan pilihan Angsuran : ")
				fmt.Scanln(&e)
			}
			if e == 1 {
				e = 25000000
			}
			if e == 2 {
				e = 50000000
			}

			f = c / e

			fmt.Println("Lama kredit anda\t\t:", int(f))
			fmt.Print("Tahun pembelian mobil anda\t: ")
			fmt.Scanln(&g)
			for g < 1 || g > 2020 {
				fmt.Println("\n\n-------------------------------------------------")
				fmt.Println("Mohon isi tahun pembelian mobil anda dengan benar.")
				fmt.Println("-------------------------------------------------\n\n")
				fmt.Print("Tahun pembelian mobil anda\t: ")
				fmt.Scanln(&g)
			}
			fmt.Println("Jumlah poin anda saat ini\t: 0")
			h = 0
			fmt.Println("Sisa Angsuran anda saat ini\t:", int(c))
			i = c

			fmt.Println("\n-------------------------------------------------\n\n")

			triggData2 = false
			triggData = true
			for !triggData2 {
				fmt.Println("Apakah ada data yang ingin anda ubah? (iya/tidak)")
				fmt.Scanln(&ulang)
				if ulang == "tidak" {
					triggData2 = true
					triggData = true
					selesai = true
				} else if ulang != "tidak" && ulang != "iya" {
					fmt.Println("\n\n-------------------------------------------------\n\n")
					fmt.Println("Mohon isi sesuai dengan Iya atau tidak.")
					fmt.Println("\n\n-------------------------------------------------\n\n")
				} else {
					triggData2 = true
					triggData = false
				}
			}
			if selesai {
				List[*N].nama = a
				List[*N].alamat = b
				List[*N].hargaMobil = c
				List[*N].umur = d
				List[*N].pilihanAngsuran = e
				List[*N].lamaKredit = f
				List[*N].tahunBeli = g
				List[*N].poin = h
				List[*N].sisaAngsuran = i
			}
		}
		*N++

		//fmt.Println("\n\n-------------------------------------------------\n\n")

	}
	backMenu()
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func searching(name string) int {
	var i int

	i = 0
	for i < N && name != List[i].nama {
		i++
	}
	if name == List[i].nama {
		return i
	} else {
		return -1
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func printData() {

	var i = 0
	var triggDiatas20 bool
	var pilihanJenis int

	fmt.Println("Pilihan print jenis data pembeli: \n\n 1. Semua. \n 2. Pemilik poin diatas 20.\n")
	fmt.Println("Pilihan anda: ")
	fmt.Scanln(&pilihanJenis)
	for pilihanJenis != 1 && pilihanJenis != 2 {
		fmt.Println("\n\n-------------------------------------------------\n\n")
		fmt.Println("Mohon pilih antara angka 1 dan 2.")
		fmt.Println("\n\n-------------------------------------------------\n\n")
		fmt.Println("Pilihan print jenis data pembeli: \n\n 1. Semua. \n 2. Pemilik poin diatas 20.\n")
		fmt.Println("Pilihan anda: ")
		fmt.Scanln(&pilihanJenis)
	}

	if pilihanJenis == 1 {
		fmt.Println("\n//////////////////////////////////////////////////\n")
		fmt.Println("\nDATA PEMBELI: ")

		for i < N {
			fmt.Println("\n\n//////////////////////////////////////////////////\n\n")
			fmt.Println("DATA ", i+1, " :")

			fmt.Println("Nama\t\t\t: ", List[i].nama)
			fmt.Println("Alamat\t\t\t: ", List[i].alamat)
			fmt.Println("Harga mobil\t\t: ", int(List[i].hargaMobil))
			fmt.Println("Umur\t\t\t: ", List[i].umur)
			fmt.Println("Pilihan angsuran\t: ", int(List[i].pilihanAngsuran))
			fmt.Println("Lama kredit\t\t: ", int(List[i].lamaKredit))
			fmt.Println("Tahun pembelian mobil\t: ", List[i].tahunBeli)
			fmt.Println("jumlah Poin saat ini\t: ", List[i].poin)
			fmt.Println("Sisa angsuran\t\t: ", int(List[i].sisaAngsuran))
			fmt.Println("\n\n")

			i++
		}
		if N == 0 {
			fmt.Println("\n\n - \n\n\n//////////////////////////////////////////////////\n\n")
		}

	}

	if pilihanJenis == 2 {
		i = 0
		fmt.Println("\n\n//////////////////////////////////////////////////\n\n")
		fmt.Println("DATA PEMBELI YANG MEMILIKI POIN DIATAS 20: \n\n")
		for i < N {
			if List[i].poin > 20 {
				fmt.Println("/////////////////////////////////////////////////\n\n")
				fmt.Println("DATA ", i+1, " :")

				fmt.Println("Nama\t\t\t: ", List[i].nama)
				fmt.Println("Alamat\t\t\t: ", List[i].alamat)
				fmt.Println("Harga mobil\t\t: ", int(List[i].hargaMobil))
				fmt.Println("Umur\t\t\t: ", List[i].umur)
				fmt.Println("Pilihan angsuran\t: ", int(List[i].pilihanAngsuran))
				fmt.Println("Lama kredit\t\t: ", int(List[i].lamaKredit))
				fmt.Println("Tahun pembelian mobil\t: ", List[i].tahunBeli)
				fmt.Println("jumlah Poin saat ini\t: ", List[i].poin)
				fmt.Println("Sisa angsuran\t\t: ", int(List[i].sisaAngsuran))
				fmt.Println("\n\n")
				triggDiatas20 = true
			}
			i++
		}
		if triggDiatas20 == false {
			fmt.Println(" - \n\n\n//////////////////////////////////////////////////\n\n")
		}
	}

	backMenu()
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func bayar() {
	var (
		tanggal, bulan int
		triggBulan     bool
	)

	fmt.Println("atas nama: ")
	fmt.Scanln(&nama)
	searching(nama)
	x = searching(nama)

	if x == -1 {
		fmt.Println("\ndata tidak ada.\n")
	} else {
		triggBulan = false
		if List[x].sisaAngsuran < 1 {
			fmt.Println("\n-------------------------------------------------\n")
			fmt.Println("Cicilan sudah lunas.")	
			fmt.Println("\n-------------------------------------------------\n")
			triggBulan = true
		}
		for !triggBulan {
			fmt.Println("Pembayaran dilakukan untuk bulan: ")
			fmt.Scanln(&bulan)
			if bulan > 0 && bulan < 13 {
				List[x].Waktu[List[x].urut].bln = bulan
				triggBulan = true
				fmt.Println("Pembayaran dilakukan untuk tanggal: ")
				fmt.Scanln(&tanggal)
				for tanggal < 1 || tanggal > 30{
					fmt.Println("\n-------------------------------------------------\n")
					fmt.Println("mohon untuk mengisi tanggal dengan angka antara 1 sampai-dengan 30.")
					fmt.Println("\n-------------------------------------------------\n")
					fmt.Println("Pembayaran dilakukan untuk tanggal: ")
					fmt.Scanln(&tanggal)
				}
				List[x].Waktu[List[x].urut].tgl = tanggal
				
				if (tanggal >= 1 && tanggal <= 10) && (bulan > 5 && bulan < 9) {
					List[x].poin = List[x].poin + 10
				} else if tanggal == 15 && bulan == 6 {
					List[x].poin = List[x].poin + 5
				}

				List[x].sisaAngsuran = List[x].sisaAngsuran - List[x].pilihanAngsuran

				List[x].urut++

				jumPembayaran++
				List[x].jumBayar = jumPembayaran
				fmt.Println("\n\n-------------------------------------------------\n\n")
				fmt.Println("\n=================================================\n")
				fmt.Println("sisa angsuran anda saat ini\t: ", int(List[x].sisaAngsuran))
				fmt.Println("Poin anda saat ini\t\t: ", List[x].poin)
				fmt.Println("\n=================================================\n")
			} else {
				fmt.Println("\n-------------------------------------------------\n")
				fmt.Println("mohon untuk mengisi bulan dengan angka antara 1 sampai-dengan 12.")
				fmt.Println("\n-------------------------------------------------\n")

			}
		}
	}
	backMenu()

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func history(List *[Batas]Buyer, x int) {
	var i int

	i = 0
	if i == List[x].urut {
		fmt.Println("Data tidak ada.")
	}
	for i < List[x].urut {
		fmt.Println("Tanggal \t: ", List[x].Waktu[i].tgl)
		fmt.Println("Bulan \t\t: ", List[x].Waktu[i].bln)
		i++
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func sortingBulan(List *[Batas]Buyer, x int) {
	var (
		temp         waktuBeli
		pass, max, i int
	)
	pass = 0
	for pass <= List[x].urut-2 {
		max = pass
		i = pass + 1
		for i <= List[x].urut-1 {
			if List[x].Waktu[max].bln < List[x].Waktu[i].bln {
				max = i
			}
			i++
		}
		temp = List[x].Waktu[max]
		List[x].Waktu[max] = List[x].Waktu[pass]
		List[x].Waktu[pass] = temp
		pass++
	}

}

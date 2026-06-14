package main

import "fmt"

const MAX = 35

type riwayatTidur struct {
	tanggal   int
	jamtidur  int
	jambangun int
	durasi    float64
	saran     string
}

type TabTidur [MAX]riwayatTidur

var dataTidur TabTidur
var nTidur int = 0

func hitungDurasi(jamtidur int, jambangun int) float64 {
	var h1, m1, h2, m2 int

	h1 = jamtidur / 100
	m1 = jamtidur % 100

	h2 = jambangun / 100
	m2 = jambangun % 100

	totalMenitTidur := (h1 * 60) + m1
	totalMenitBangun := (h2 * 60) + m2

	if totalMenitBangun < totalMenitTidur {
		totalMenitBangun += 24 * 60
	}

	selisihMenit := totalMenitBangun - totalMenitTidur
	return float64(selisihMenit) / 60.0
}

func getSaran(durasi float64) string {
	if durasi < 7.0 {
		return "Kurang tidur, perbanyak istirahat."
	} else if durasi <= 9.0 {
		return "Pola tidur sehat."
	}
	return "Terlalu banyak tidur."
}

func sequentialSearch(tgl int) int {
	idx := -1
	i := 0
	for i < nTidur && idx == -1 {
		if dataTidur[i].tanggal == tgl {
			idx = i
		}
		i++
	}
	return idx
}

func binarySearch(tgl int) int {
	kiri := 0
	kanan := nTidur - 1
	idx := -1

	for kiri <= kanan && idx == -1 {
		tengah := (kiri + kanan) / 2
		if dataTidur[tengah].tanggal == tgl {
			idx = tengah
		} else if dataTidur[tengah].tanggal < tgl {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return idx
}

func selectionSortTanggal(asc bool) {
	i := 0
	for i < nTidur-1 {
		idx := i
		j := i + 1
		for j < nTidur {
			kondisi := false
			if asc {
				kondisi = dataTidur[j].tanggal < dataTidur[idx].tanggal
			} else {
				kondisi = dataTidur[j].tanggal > dataTidur[idx].tanggal
			}

			if kondisi {
				idx = j
			}
			j++
		}
		temp := dataTidur[i]
		dataTidur[i] = dataTidur[idx]
		dataTidur[idx] = temp
		i++
	}
}

func insertionSortDurasi(asc bool) {
	i := 1
	for i < nTidur {
		temp := dataTidur[i]
		j := i - 1
		stop := false

		for j >= 0 && !stop {
			kondisi := false
			if asc {
				kondisi = dataTidur[j].durasi > temp.durasi
			} else {
				kondisi = dataTidur[j].durasi < temp.durasi
			}

			if kondisi {
				dataTidur[j+1] = dataTidur[j]
				j--
			} else {
				stop = true
			}
		}
		dataTidur[j+1] = temp
		i++
	}
}

func tambahData() {
	if nTidur < MAX {
		var tgl, bed, wake int
		fmt.Print("Masukkan Tanggal (YYYYMMDD): ")
		fmt.Scan(&tgl)

		validBed := false
		for !validBed {
			fmt.Print("Masukkan Jam Tidur (HHMM): ")
			fmt.Scan(&bed)

			h := bed / 100
			m := bed % 100
			if h >= 0 && h <= 23 && m >= 0 && m <= 59 {
				validBed = true
			} else {
				fmt.Println(">> Waktu tidak valid! Jam harus 00-23 dan Menit 00-59.")
			}
		}

		validWake := false
		for !validWake {
			fmt.Print("Masukkan Jam Bangun (HHMM)(untuk di jam 0 - 9, masukkan 3 digit saja seperi 500 untuk jam 5 pagi): ")
			fmt.Scan(&wake)

			h := wake / 100
			m := wake % 100
			if h >= 0 && h <= 23 && m >= 0 && m <= 59 {
				validWake = true
			} else {
				fmt.Println(">> Waktu tidak valid! Jam harus 00-23 dan Menit 00-59.")
			}
		}

		durasi := hitungDurasi(bed, wake)
		saran := getSaran(durasi)

		dataTidur[nTidur] = riwayatTidur{
			tanggal:   tgl,
			jamtidur:  bed,
			jambangun: wake,
			durasi:    durasi,
			saran:     saran,
		}
		nTidur++
		fmt.Println(">> Data berhasil ditambahkan!")
	} else {
		fmt.Println(">> Kapasitas memori penuh!")
	}
}

func ubahData() {
	var tgl int
	fmt.Print("Masukkan Tanggal data yang ingin diubah (YYYYMMDD): ")
	fmt.Scan(&tgl)

	idx := sequentialSearch(tgl)
	if idx != -1 {
		var bed, wake int

		validBed := false
		for !validBed {
			fmt.Print("Masukkan Jam Tidur Baru (HHMM): ")
			fmt.Scan(&bed)

			h := bed / 100
			m := bed % 100
			if h >= 0 && h <= 23 && m >= 0 && m <= 59 {
				validBed = true
			} else {
				fmt.Println(">> Waktu tidak valid! Jam harus 00-23 dan Menit 00-59.")
			}
		}

		validWake := false
		for !validWake {
			fmt.Print("Masukkan Jam Bangun Baru (HHMM): ")
			fmt.Scan(&wake)

			h := wake / 100
			m := wake % 100
			if h >= 0 && h <= 23 && m >= 0 && m <= 59 {
				validWake = true
			} else {
				fmt.Println(">> Waktu tidak valid! Jam harus 00-23 dan Menit 00-59.")
			}
		}

		durasi := hitungDurasi(bed, wake)
		dataTidur[idx].jamtidur = bed
		dataTidur[idx].jambangun = wake
		dataTidur[idx].durasi = durasi
		dataTidur[idx].saran = getSaran(durasi)
		fmt.Println(">> Data berhasil diubah!")
	} else {
		fmt.Println(">> Data tidak ditemukan.")
	}
}

func hapusData() {
	var tgl int
	fmt.Print("Masukkan Tanggal data yang ingin dihapus (YYYYMMDD): ")
	fmt.Scan(&tgl)

	selectionSortTanggal(true)
	idx := binarySearch(tgl)

	if idx != -1 {
		i := idx
		for i < nTidur-1 {
			dataTidur[i] = dataTidur[i+1]
			i++
		}
		nTidur--
		fmt.Println(">> Data berhasil dihapus!")
	} else {
		fmt.Println(">> Data tidak ditemukan.")
	}
}

func tampilkanLaporan() {
	if nTidur == 0 {
		fmt.Println(">> Belum ada riwayat tidur.")
		return
	}

	fmt.Println("\n--- Data Riwayat Tidur Saat Ini ---")
	i := 0
	for i < nTidur {
		fmt.Printf("%d. Tgl: %d | Tidur: %04d | Bangun: %04d | Durasi: %.2f Jam | Saran: %s\n",
			i+1, dataTidur[i].tanggal, dataTidur[i].jamtidur, dataTidur[i].jambangun, dataTidur[i].durasi, dataTidur[i].saran)
		i++
	}
}

func menuUrutkanData() {
	var pilihan int
	fmt.Println("\nPilih metode pengurutan:")
	fmt.Println("1. Berdasarkan Tanggal (Ascending) - Selection Sort")
	fmt.Println("2. Berdasarkan Tanggal (Descending) - Selection Sort")
	fmt.Println("3. Berdasarkan Durasi (Ascending) - Insertion Sort")
	fmt.Println("4. Berdasarkan Durasi (Descending) - Insertion Sort")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		selectionSortTanggal(true)
		fmt.Println(">> Data diurutkan berdasarkan Tanggal (Naik).")
	} else if pilihan == 2 {
		selectionSortTanggal(false)
		fmt.Println(">> Data diurutkan berdasarkan Tanggal (Turun).")
	} else if pilihan == 3 {
		insertionSortDurasi(true)
		fmt.Println(">> Data diurutkan berdasarkan Durasi (Naik).")
	} else if pilihan == 4 {
		insertionSortDurasi(false)
		fmt.Println(">> Data diurutkan berdasarkan Durasi (Turun).")
	} else {
		fmt.Println(">> Pilihan tidak valid.")
	}
}

func main() {
	selesai := false

	for !selesai {
		fmt.Println("\n=== APLIKASI PEMANTAUAN POLA TIDUR ===")
		fmt.Println("1. Tambah Riwayat Tidur")
		fmt.Println("2. Ubah Riwayat Tidur (Sequential Search)")
		fmt.Println("3. Hapus Riwayat Tidur (Binary Search)")
		fmt.Println("4. Urutkan Data (Selection & Insertion Sort)")
		fmt.Println("5. Tampilkan Laporan")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var menu int
		fmt.Scan(&menu)

		if menu == 1 {
			tambahData()
		} else if menu == 2 {
			ubahData()
		} else if menu == 3 {
			hapusData()
		} else if menu == 4 {
			menuUrutkanData()
		} else if menu == 5 {
			tampilkanLaporan()
		} else if menu == 0 {
			selesai = true
			fmt.Println("Keluar dari program. Tetap jaga kesehatan tidur Anda!")
		} else {
			fmt.Println("Menu tidak valid.")
		}
	}
}

package main

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

package main

import "fmt"

type pengguna struct {
	hariKe  int
	durasi  int
	kalori  int
	tanggal string
	jenis   string
}

const Ndays int = 365

type workout [Ndays]pengguna

var P workout
var N int
var stop bool
var hari int

func main() {
	N = 1
	hari = N
	for !stop { // bila user tidak memilih exit maka menu akan terus ditampilkan
		menu()
	}
}

// prosedur untuk menampilkan menu
func menu() {
	var Nomor int

	fmt.Println("-------------")
	fmt.Println("    Menu")
	fmt.Println("-------------")
	fmt.Println("1. Menambahkan data")
	fmt.Println("2. Edit data")
	fmt.Println("3. Hapus data")
	fmt.Println("4. Rekomendasi workout")
	fmt.Println("5. Cari workout berdasarkan jenis olahraga")
	fmt.Println("6. Cari workout berdasarkan hari")
	fmt.Println("7. Urutkan berdasarkan kalori")
	fmt.Println("8. Urutkan berdasarkan durasi")
	fmt.Println("9. Laporan")
	fmt.Println("10. Exit")
	fmt.Print("Pilih menu 1-10: ")
	fmt.Scan(&Nomor)
	switch Nomor {
	case 1:
		addData(&P, &N)
	case 2:
		editData(&P, N)
	case 3:
		hapusData(&P, &N)
	case 4:
		rekomendasi(P, N)
	case 5:
		cariOlahraga(P, N)
	case 6:
		cariHari(P, N)
	case 7:
		sortKalori(P, N)
	case 8:
		sortDurasi(P, N)
	case 9:
		showData(P, N)
	default:
		stop = true // membuat nilai stop menjadi true sehingga menu akan berhenti ditampilkan
	}

}

// prosedur untuk menambahkan data kedalam array
func addData(w *workout, n *int) {
	var pilih int = 1

	for pilih == 1 {
		fmt.Println("=======================")
		fmt.Println("   Menambahkan Data    ")
		fmt.Println("=======================")
		fmt.Println("Masukkan data hari ke-", hari)
		w[*n].hariKe = hari //memasukkan data hari ke-
		fmt.Print("Masukkan jadwal latihan (DDMMYYYY): ")
		fmt.Scan(&w[*n].tanggal)
		fmt.Print("Masukkan jenis latihan: ")
		fmt.Scan(&w[*n].jenis)
		fmt.Print("Masukkan durasi latihan (menit): ")
		fmt.Scan(&w[*n].durasi)
		fmt.Print("Masukkan total kalori: ")
		fmt.Scan(&w[*n].kalori)
		fmt.Println()
		fmt.Print("Ketik 1 untuk menambah data lagi, 0 untuk keluar: ")
		fmt.Scan(&pilih)
		fmt.Println()

		*n++
		hari++
	}
}

// prosedur untuk mencari olahraga menggunakan sequential search
func cariOlahraga(w workout, n int) {
	var x string
	var idx int
	var found workout
	idx = 0
	fmt.Print("Jenis olahraga yang ingin dicari: ")
	fmt.Scan(&x)
	for i := 0; i < n; i++ {
		if w[i].jenis == x { // bila olahraga yang dicari tersedia maka data akan disimpan ke array found
			found[idx] = w[i]
			idx++
		}
	}
	if idx != 0 { // bila array found memiliki isi, maka data akan dioutputkan
		fmt.Printf("| %-10s | %-9s | %-20s | %-7s | %-7s |\n", "Hari ke-", "Tanggal", "Latihan", "Durasi", "Kalori")
		for i := 0; i < idx; i++ {
			fmt.Printf("| %-10d | %-9s | %-20s | %-7d | %-7d |\n", found[i].hariKe, found[i].tanggal, found[i].jenis, found[i].durasi, found[i].kalori)
		}
	} else { // bila array found tidak memiliki isi maka akan diber10i output "Data tidak ditemukan!"
		fmt.Println("Data tidak ditemukan!")
	}
}

// prosedur untuk mencari workout hari ke- menggunakan binary search
func cariHari(w workout, n int) {
	var left, right, mid, idx int
	var hari int
	left = 0
	right = n - 1
	idx = -1
	fmt.Print("Cari workout hari ke-: ")
	fmt.Scan(&hari)
	for (left <= right) && (idx == -1) {
		mid = (left + right) / 2
		if hari < w[mid].hariKe {
			right = mid - 1
		} else if hari > w[mid].hariKe {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	if idx == -1 {
		fmt.Println("Data tidak ditemukan!")
	} else { // menampilkan data bila data yang dicari tersedia
		fmt.Printf("| %-10s | %-9s | %-20s | %-7s | %-7s |\n", "Hari ke-", "Tanggal", "Latihan", "Durasi", "Kalori")
		fmt.Printf("| %-10d | %-9s | %-20s | %-7d | %-7d |\n", w[idx].hariKe, w[idx].tanggal, w[idx].jenis, w[idx].durasi, w[idx].kalori)
	}
}

// prosedur untuk menampilkan data
func showData(w workout, n int) {
	var total int
	fmt.Printf("| %-10s | %-9s | %-20s | %-7s | %-7s |\n", "Hari ke-", "Tanggal", "Latihan", "Durasi", "Kalori")
	for i := 1; i < n; i++ {
		fmt.Printf("| %-10d | %-9s | %-20s | %-7d | %-7d |\n", w[i].hariKe, w[i].tanggal, w[i].jenis, w[i].durasi, w[i].kalori)
	}
	//total kalori
	for j := 1; j < n; j++ {
		total += w[j].kalori
	}
	fmt.Printf("Total kalori keseluruhan: %d\n", total)
}

// prosedur untuk mengubah data sesuai keinginan user
func editData(w *workout, n int) {
	var idx, pilihan, edit int
	showData(*w, n)
	fmt.Print("Pilih data hari ke- berapa yang ingin diedit: ")
	fmt.Scan(&idx)
	fmt.Printf("| %-10s | %-9s | %-20s | %-7s | %-7s |\n", "Hari ke-", "Tanggal", "Latihan", "Durasi", "Kalori")
	fmt.Printf("| %-10d | %-9s | %-20s | %-7d | %-7d |\n", w[idx].hariKe, w[idx].tanggal, w[idx].jenis, w[idx].durasi, w[idx].kalori)
	fmt.Print("pilih 1 untuk edit keseluruhan, 0 untuk edit sebagian: ")
	fmt.Scan(&edit)
	if edit == 1 {
		fmt.Print("Tanggal baru: ")
		fmt.Scan(&w[idx].tanggal)
		fmt.Print("Jenis olahraga baru: ")
		fmt.Scan(&w[idx].jenis)
		fmt.Print("Durasi baru: ")
		fmt.Scan(&w[idx].durasi)
		fmt.Print("Kalori baru: ")
		fmt.Scan(&w[idx].kalori)
	} else if edit == 0 {
		fmt.Println("Pilih bagian yang ingin diedit: ")
		fmt.Println("1. Ubah tanggal")
		fmt.Println("2. Ubah jenis olahraga")
		fmt.Println("3. Ubah durasi")
		fmt.Println("4. Ubah total kalori")
		fmt.Println("5. Batalkan edit")
		fmt.Print("Pilih menu 1/2/3/4/5: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			fmt.Print("Masukkan tanggal yang baru: ")
			fmt.Scan(&w[idx].tanggal)
		case 2:
			fmt.Print("Masukkan jenis olahraga yang baru: ")
			fmt.Scan(&w[idx].jenis)
		case 3:
			fmt.Print("Masukkan durasi yang baru: ")
			fmt.Scan(&w[idx].durasi)
		case 4:
			fmt.Print("Masukkan jumlah kalori yang baru: ")
			fmt.Scan(&w[idx].kalori)
		}
	}
}

// prosedur untuk menghapus data yang diinginkan user
func hapusData(w *workout, n *int) {
	var idx int
	showData(*w, *n)
	fmt.Print("Pilih data hari ke- berapa yang ingin dihapus: ")
	fmt.Scan(&idx)
	for i := idx; i < *n; i++ {
		w[i] = w[i+1] // menimpa data yang dihapus dengan data selanjutnya
	}
	*n -= 1 // mengurangi jumlah data yang ada di dalam array

}

// memberi rekomendasi workout kepada user berdasarkan 3 workout terakhir
func rekomendasi(w workout, n int) {
	var count int
	fmt.Println("Rekomendasi workout berdasarkan 3 workout terakhir: ")
	count = 1
	if n > 3 { // mengecek apakah data yang tersedia lebih dari 3
		for i := n - 3; i < n; i++ {
			fmt.Printf("%d. %s\n", count, w[i].jenis)
			count++
		}
	} else {
		for i := 0; i < n; i++ {
			fmt.Printf("%d. %s\n", count, w[i].jenis)
			count++
		}
	}
}

// prosedur untuk mengurutkan kalori menggunakan selection sort secara descending
func sortKalori(A workout, n int) {
	var i, idx, pass int
	var temp workout

	pass = 2
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].kalori > A[idx].kalori {
				idx = i
			}
			i += 1
		}
		temp[0] = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp[0]
		pass += 1
	}
	fmt.Println("Data terurut dari kalori terbesar sampai terendah: ")
	showData(A, N) // memberikan output data yang sudah terurut
}

// prosedur untuk mengurutkan data dengan insertion sort secara ascending berdasarkan durasi
func sortDurasi(A workout, n int) {
	var pass, i int
	var temp pengguna
	pass = 2
	for pass <= n-1 {
		i = pass
		temp = A[pass]
		for i > 1 && temp.durasi < A[i-1].durasi {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
	fmt.Println("Data terurut dari durasi workout paling cepat hingga paling lama: ")
	showData(A, n) // menampilkan data yang sudah terurut
}

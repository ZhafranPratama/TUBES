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
	for !stop {
		menu()
	}
}

func menu() {
	var Nomor int

	fmt.Println("-------------")
	fmt.Println("    Menu")
	fmt.Println("-------------")
	fmt.Println("1. Menambahkan data")
	fmt.Println("2. Laporan")
	fmt.Println("3. Edit data")
	fmt.Println("4. Hapus data")
	fmt.Println("5. Rekomendasi workout")
	fmt.Println("6. Cari workout berdasarkan jenis")
	fmt.Println("7. Cari workout berdasarkan tanggal")
	fmt.Println("8. Urutkan berdasarkan jumlah kalori")
	fmt.Println("9. Exit")
	fmt.Print("Pilih menu 1/2/3/4/5...: ")
	fmt.Scan(&Nomor)
	switch Nomor {
	case 1:
		addData(&P, &N)
	case 2:
		showData(P, N)
	case 3:
		editData(&P, N)
	case 4:
		hapusData(&P, &N)
	case 5:
		rekomendasi(P, N)
	case 6:
		cariOlahraga(P, N)
	case 7:
		cariTanggal(P, N)
	case 8:
		sortKalori(P, N)
	default:
		stop = true
	}

}

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

func cariOlahraga(w workout, n int) {
	var x string
	fmt.Print("Jenis olahraga yang ingin dicari: ")
	fmt.Scan(&x)
	fmt.Printf("| %-7s | %-10s | %-8s | %-7s | %-7s |\n", "Hari ke-", "Tanggal", "Latihan", "Durasi", "Kalori")
	for i := 0; i < n; i++ {
		if w[i].jenis == x {
			fmt.Printf("| %-7d | %-10s | %-8s | %-7d | %-7d |\n", w[i].hariKe, w[i].tanggal, w[i].jenis, w[i].durasi, w[i].kalori)
		}
	}
}

func cariTanggal(w workout, n int) {
	var left, right, mid, idx int
	var tanggal string
	left = 0
	right = n - 1
	idx = -1
	fmt.Print("Cari tanggal workout (DDMMYYYY): ")
	fmt.Scan(&tanggal)
	for (left <= right) && (idx == -1){
		mid = (left + right) / 2
		if tanggal < w[mid].tanggal{
			right = mid - 1
		}else if tanggal > w[mid].tanggal{
			left = mid + 1
		}else{
			idx = mid
		}
	}
	if idx == -1 {
		fmt.Println("Data tidak ditemukan!")
	}else{
		fmt.Printf("| %-7s | %-10s | %-8s | %-7s | %-7s |\n", "Hari ke-", "Tanggal", "Latihan", "Durasi", "Kalori")
		fmt.Printf("| %-7d | %-10s | %-8s | %-7d | %-7d |\n", w[idx].hariKe, w[idx].tanggal, w[idx].jenis, w[idx].durasi, w[idx].kalori)
	}
}



func showData(w workout, n int) {
	var total int
	fmt.Printf("| %-7s | %-10s | %-8s | %-7s | %-7s |\n", "Hari ke-", "Tanggal", "Latihan", "Durasi", "Kalori")
	for i := 1; i < n; i++ {
		fmt.Printf("| %-7d | %-10s | %-8s | %-7d | %-7d |\n", w[i].hariKe, w[i].tanggal, w[i].jenis, w[i].durasi, w[i].kalori)
	}
	//total kalori
	for j:=1; j<n; j++ {
		total += w[j].kalori
	}
	fmt.Printf("Total kalori keseluruhan: %d\n", total)
}

func editData(w *workout, n int) {
	var idx, pilihan, edit int
	showData(*w, n)
	fmt.Print("Pilih data hari ke- berapa yang ingin diedit: ")
	fmt.Scan(&idx)
	fmt.Printf("| %-7s | %-10s | %-8s | %-7s | %-7s |\n", "Hari ke-", "Tanggal", "Latihan", "Durasi", "Kalori")
	fmt.Printf("| %-7d | %-10s | %-8s | %-7d | %-7d |\n", w[idx].hariKe, w[idx].tanggal, w[idx].jenis, w[idx].durasi, w[idx].kalori)
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

func hapusData(w *workout, n *int) {
	var idx int
	showData(*w, *n)
	fmt.Print("Pilih data hari ke- berapa yang ingin dihapus: ")
	fmt.Scan(&idx)
	for i := idx; i < *n; i++ {
		w[i] = w[i + 1]
	}
	*n -= 1

}

func rekomendasi(w workout, n int){
	var count int
	fmt.Println("Rekomendasi workout berdasarkan 3 workout terakhir: ")
	count = 1
	if n > 3{
		for i := n - 3; i < n; i++{
			fmt.Printf("%d. %s\n",count, w[i].jenis)
			count++
		}
	}else{
		for i := 0; i < n; i++{
			fmt.Printf("%d. %s\n",count, w[i].jenis)
			count++
		}
	}
}

func sortKalori(A workout, n int){
	var i, idx, pass int
	var temp workout
	
	pass = 2
	for pass < n{
		idx = pass - 1
		i = pass
		for i < n{
			if A[i].kalori > A[idx].kalori{
				idx = i
			}
			i += 1
		}
		temp[0] = A[pass - 1]
		A[pass - 1] = A[idx]
		A[idx] = temp[0]
		pass += 1
	}
	fmt.Println("Data terurut dari kalori terbesar sampai terendah: ")
	showData(A, N)
}

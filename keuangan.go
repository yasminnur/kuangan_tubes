package main

import "fmt"

type layanan_berlangganan struct {
	no                int
	nama_layanan      string
	biaya             int
	metode_pembayaran string
	tgl_pembayaran    string
	status            string
}

const NMAX int = 7
type tabInt [NMAX]layanan_berlangganan

func tambahLayanan(A *tabInt, no *int) {
	A[*no].no = *no + 1
	fmt.Print("Nama Layanan: ")
	fmt.Scan(&A[*no].nama_layanan)
	fmt.Print("biaya : ")
	fmt.Scan(&A[*no].biaya)
	fmt.Print("Metode Pembayaran : ")
	fmt.Scan(&A[*no].metode_pembayaran)
	fmt.Print("Tanggal Pembayaran : ")
	fmt.Scan(&A[*no].tgl_pembayaran)
	fmt.Print("Status : ")
	fmt.Scan(&A[*no].status)
	fmt.Println("======================")
	fmt.Println(A)
	cekIsiArray(A)
	*no++
}
func tampilkanLayanan(A *tabInt, no int) {
	for i := 0; i < no; i++ {
		fmt.Println("ini i = ",i)
		fmt.Print("Nama Layanan: ")
		fmt.Print(A[0].nama_layanan, "\n")
		// fmt.Println()
		fmt.Print("biaya : ")
		fmt.Print(A[0].biaya)
		// fmt.Println()
		fmt.Print("Metode Pembayaran : ")
		fmt.Print(A[0].metode_pembayaran, "\n")
		// fmt.Println()
		fmt.Print("Tanggal Pembayaran : ")
		fmt.Print(A[i].tgl_pembayaran, "\n")
		// fmt.Println()
		fmt.Print("Status : ")
		fmt.Print(A[i].status, "\n")
		// fmt.Println()
	}
}
func editLayanan(A *tabInt, idx int) {
	fmt.Println("======== EDIT ========")
	idx = idx-1
	fmt.Print("Nama Layanan: ")
	fmt.Scan(&A[idx].nama_layanan)
	// fmt.Println()
	fmt.Print("biaya : ")
	fmt.Scan(&A[idx].biaya)
	fmt.Print("Metode Pembayaran : ")
	fmt.Scan(&A[idx].metode_pembayaran)
	// fmt.Println()
	fmt.Print("Tanggal Pembayaran : ")
	fmt.Scan(&A[idx].tgl_pembayaran)
	// fmt.Println()
	fmt.Print("Status : ")
	fmt.Scan(&A[idx].status)
	fmt.Println()
}
func cekIsiArray(A *tabInt) int {
	var jmlh int
	for i := 0; i < NMAX && A[i].no != 0; i++ {
		jmlh++
		// fmt.Println(A[i].no)
	}
	return jmlh
}

// func search(A tabInt, x int)int{
// 	var left, rigth, mid int
// 	var idx int
// 	left = 0
// 	rigth = NMAX-1
// 	idx = -1
// 	for left <= rigth && idx == -1 {
// 		mid = (left+rigth)/2
// 		if x < A[mid] {
// 			rigth = mid-1
// 		} else if x > A[mid] {
// 			left = mid + 1
// 		} else {
// 			idx = mid
// 		}
// 	}
// 	return idx
// }

func menu(A tabInt){
	var pil, isiArr int
	for {
	fmt.Println("===================== MENU ========================")
	fmt.Println("1. Tampilkan Daftar Layanan")
	fmt.Println("2. Tambahkan Daftar Layanan")
	fmt.Println("3. Edit Daftar Layanan")
	fmt.Scan(&pil)
	loadData(&A)
	isiArr = cekIsiArray(&A)
	switch pil{
	case 1 :
		tampilkanLayanan(&A, isiArr);
		fmt.Println("data skrg = ", isiArr)
	case 2 : 
		cekIsiArray(&A)
		tambahLayanan(&A, &isiArr)
		fmt.Println("data skrg = ", isiArr)
	case 3 :
		var index int
		fmt.Print("Baris nomor berapa yang ingin diubah? ")
		fmt.Scan(&index)
		fmt.Println("Sebelum edit ")
		fmt.Println(A)
		editLayanan(&A, index)
		fmt.Println("Sesudah edit ")
		fmt.Println(A)
	}
}
}
func loadData(data *tabInt){
	*data = tabInt{
		{1, "aaa", 1, "ww", "20-302-0", "lunas"},
		{2, "bbb", 2, "ww", "20-302-0", "lunas"},
		{3, "ccc", 3, "ww", "20-302-0", "belum"},
		{4, "ddd", 4, "ww", "20-302-0", "lunas"},
		{5, "eee", 5, "ww", "20-302-0", "lunas"},
	}
}
func main() {
	var data tabInt
	// var isiArray int
	// loadData(&data)
	// fmt.Println(data)
	// cekIsiArray(&data)
	// fmt.Println(data)
	menu(data)
}

// ===== DATA AWAL ========
// --------- data 1
// 1
// nama_Layanan : a
// biaya : 1
// metode_pembayaran : a
// tgl_pembayaran : 1
// status : lunas
// ----------- data 2
// nama_Layanan : a
// biaya : 1
// metode_pembayaran : a
// tgl_pembayaran : 1
// status : lunas

//  ========= EDIT =========
// nama_Layanan :
// biaya :
// metode_pembayaran :
// tgl_pembayaran :
// status :

// NOTE
// angka nggak blh minus, bikin pengecekan ketat

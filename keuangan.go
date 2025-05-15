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

const NMAX int = 100

type tabInt [NMAX]layanan_berlangganan

func tambahLayanan(A *tabInt) {
	for i := 0; i < NMAX; i++ {
		// fmt.Println("masukkan : ")
		fmt.Print("Nama Layanan: ")
		fmt.Scan(&A[i].nama_layanan)
		fmt.Println()
		fmt.Print("biaya : ")
		fmt.Scan(&A[i].biaya)
		fmt.Println()
		fmt.Print("Metode Pembayaran : ")
		fmt.Scan(&A[i].metode_pembayaran)
		fmt.Println()
		fmt.Print("Tanggal Pembayaran : ")
		fmt.Scan(&A[i].tgl_pembayaran)
		fmt.Println()
		fmt.Print("Status : ")
		fmt.Scan(&A[i].status)
		fmt.Println()
		// A[i].no= i+1
		// fmt.Scan(&A[i].nama_layanan)
		// fmt.Scan(&A[i].biaya)
		// fmt.Scan(&A[i].metode_pembayaran)
		// fmt.Scan(&A[i].tgl_pembayaran)
		// fmt.Scan(&A[i].status)
	}
}
func tampilkanLayanan(A *tabInt) {
	for i := 0; i < NMAX; i++ {
		// fmt.Println("print : ")
		// fmt.Println(A[i].nama_layanan)
		// fmt.Println(A[i].biaya)
		// fmt.Println(A[i].metode_pembayaran)
		// fmt.Println(A[i].tgl_pembayaran)
		// fmt.Println(A[i].status)
		fmt.Print("Nama Layanan: ")
		fmt.Print(&A[i].nama_layanan)
		fmt.Println()
		fmt.Print("biaya : ")
		fmt.Print(&A[i].biaya)
		fmt.Println()
		fmt.Print("Metode Pembayaran : ")
		fmt.Print(&A[i].metode_pembayaran)
		fmt.Println()
		fmt.Print("Tanggal Pembayaran : ")
		fmt.Print(&A[i].tgl_pembayaran)
		fmt.Println()
		fmt.Print("Status : ")
		fmt.Print(&A[i].status)
		fmt.Println()
	}
}
func editLayanan(A *tabInt, idx int) {
	fmt.Println("======== EDIT ========")
	fmt.Print("Nama Layanan: ")
	fmt.Scan(&A[idx].nama_layanan)
	fmt.Println()
	fmt.Print("biaya : ")
	fmt.Scan(&A[idx].biaya)
	fmt.Println()
	fmt.Print("Metode Pembayaran : ")
	fmt.Scan(&A[idx].metode_pembayaran)
	fmt.Println()
	fmt.Print("Tanggal Pembayaran : ")
	fmt.Scan(&A[idx].tgl_pembayaran)
	fmt.Println()
	fmt.Print("Status : ")
	fmt.Scan(&A[idx].status)
	fmt.Println()
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
// func menu(){
// 	switch
// }
func main() {
	var data tabInt
	var index int
	// menu()
	tambahLayanan(&data)
	// search(data, x1)
	editLayanan(&data, index)
	fmt.Println(data)
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

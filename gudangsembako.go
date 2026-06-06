package main

import "fmt"

const NMAX = 1000

type barang struct {
	kode string
	nama string
	stok int
	harga int
}
type tabBarang [NMAX] barang

func tambahBarang(data *tabBarang, n *int) {
	var i int
	var found bool
	var x string

	if *n >= NMAX {
		fmt.Println("Data Penuh!")
	}

	found = false

	fmt.Print("Masukkan Kode Barang: ")
	fmt.Scan(&x)

	for i = 0; i < *n; i++ {
		if data[i].kode == x {
			found = true
		}
	}

	if found {
		fmt.Println("Kode Barang Sudah Digunakan!")
	} else {
		data[*n].kode = x

		fmt.Print("Masukkan Nama Barang: ")
		fmt.Scan(&data[*n].nama)

		fmt.Print("Masukkan Stok Barang: ")
		fmt.Scan(&data[*n].stok)

		fmt.Print("Masukan Harga Barang: ")
		fmt.Scan(&data[*n].harga)

		for data[*n].stok < 0 {
			fmt.Println("Stok Tidak Boleh Negatif!")
			fmt.Print("Masukkan Stok Barang: ")
			fmt.Scan(&data[*n].stok)
		}
		for data[*n].harga < 0 {
			fmt.Println("Harga Tidak Boleh Negatif!")
			fmt.Print("Masukkan Harga Barang: ")
			fmt.Scan(&data[*n].harga)
		}
		*n = *n + 1

		fmt.Println("Data Berhasil Ditambahkan!")
	}
}

func ubahBarang(data *tabBarang, n int) {
	var i int
	var x string
	var found bool

	fmt.Print("Masukkan Kode Barang Yang Ingin Diubah: ")
	fmt.Scan(&x)

	found = false
	i = 0

	for i < n && !found {
		if data[i].kode == x {
			found = true
		} else {
			i = i + 1
		}
	}

	if found {
		fmt.Println("Data Ditemukan!")

		fmt.Print("Nama Barang Baru: ")
		fmt.Scan(&data[i].nama)

		fmt.Print("Jumlah Stok Baru: ")
		fmt.Scan(&data[i].stok)

		fmt.Print("Harga Barang Baru: ")
		fmt.Scan(&data[i].harga)

		fmt.Println("Data Berhasil Diubah!")
	} else {
		fmt.Println("Data Tidak Ditemukan!")
	}
}

func hapusBarang(data *tabBarang, n *int) {
	var i, j int
	var x string
	var found bool

	fmt.Print("Masukkan Kode Barang Yang Ingin Dihapus: ")
	fmt.Scan(&x)

	found = false 
	i = 0

	for i < *n && !found {
		if data[i].kode == x {
			found = true
		} else {
			i = i + 1
		}
	}

	if found {
		for j = i; j < *n-1; j++ {
			data[j] = data[j+1]
		}
		*n = *n - 1
		fmt.Println("Data Berhasil Dihapus!")
	} else {
		fmt.Println("Data Tidak Ditemukan!")
	}
}

func transaksiMasuk(data *tabBarang, n int) {
	var i, stok1 int
	var x string
	var found bool

	fmt.Print("Masukkan Kode Barang: ")
	fmt.Scan(&x)

	found = false 
	i = 0

	for i < n && !found {
		if data[i].kode == x {
			found = true
		} else {
			i = i + 1
		}
	}

	if found {
		fmt.Print("Masukkan Jumlah Barang Masuk: ")
		fmt.Scan(&stok1)

		data[i].stok = data[i].stok + stok1

		fmt.Println("Transaksi Barang Masuk Berhasil")
	} else {
		fmt.Println("Barang Tidak Ditemukan!")
	}
}

func transaksiKeluar(data *tabBarang, n int) {
	var i, stok1 int
	var x string
	var found bool

	fmt.Print("Masukkan Kode Barang: ")
	fmt.Scan(&x)

	found = false 
	i = 0

	for i < n && !found {
		if data[i].kode == x {
			found = true
		} else {
			i = i + 1
		}
	}

	if found {
		fmt.Print("Masukkan Jumlah Barang Keluar: ")
		fmt.Scan(&stok1)

		if stok1 <= data[i].stok {
			data[i].stok = data[i].stok - stok1
			fmt.Print("Transaksi barang keluar telah berhasil!")
		} else {
			fmt.Print("Stok barang tidak ditemukan")
		}
	}else {
		fmt.Print("Barang tidak ditemukan!")
	}
}

func printData(data tabBarang, n int) {
	var i int

	if n == 0{
		fmt.Print("Barang kosong")
	}else{
		fmt.Println("==== DATA BARANG ====")
		for i=0; i<n; i++{
			fmt.Println("Data ke: ", i+1)
			fmt.Println("Kode: ", data[i].kode)
			fmt.Println("Nama: ", data[i].nama)
			fmt.Println("Stok: ", data[i].stok)
			fmt.Println("Harga: ", data[i].harga)
		}
		fmt.Println()
	}
}

func sequentialSearch(data tabBarang, n int, x string) int{
	var i int

	for i=0; i<n; i++{
		if data[i].nama == x{
			return i
		}
	}
	return -1
}

func binarySearch(data tabBarang, n int, x string) int{
	var left, right, mid int
	left = 0
	right = n - 1

	for left <= right {
		mid = (left + right) / 2
		if data[mid].kode == x {
			return mid
		} else if x < data[mid].kode {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func selectionSort(data *tabBarang, n int) {
	var i, pass, idxMin int
	var temp barang

	for i = 0; i < n-1; i++ {

		idxMin = i

		for pass = i + 1; pass < n; pass++ {

			if data[pass].kode < data[idxMin].kode {
				idxMin = pass
			}
		}

		temp = data[i]
		data[i] = data[idxMin]
		data[idxMin] = temp
	}
}

func insertionSort(data *tabBarang, n int) {
	var pass, i int
	var temp barang

	pass = 1

	for pass <= n-1 {
		i = pass
		temp = data[pass]

		for i > 0 && temp.stok > data[i-1].stok {
			data[i] = data[i-1]
			i = i - 1
		}

		data[i] = temp
		pass = pass + 1
	}
}

func statistikGudang(data tabBarang, n int) {
	var i int
	var totalAset int
	var idxMin, idxMax int

	if n > 0 {

		idxMin = 0
		idxMax = 0

		for i = 0; i < n; i++ {

			totalAset += data[i].stok * data[i].harga

			if data[i].stok < data[idxMin].stok {
				idxMin = i
			}

			if data[i].stok > data[idxMax].stok {
				idxMax = i
			}
		}

		fmt.Println("===== STATISTIK GUDANG =====")

		fmt.Println("Total aset gudang :", totalAset)

		fmt.Println("Barang stok minimum :")
		fmt.Println(data[idxMin].nama, "-", data[idxMin].stok)

		fmt.Println("Barang stok maksimum :")
		fmt.Println(data[idxMax].nama, "-", data[idxMax].stok)

	} else {
		fmt.Println("Data barang kosong")
	}
}

func main() {
	var data tabBarang
	var n, pilihan, idx int
	var x string

	n = 0

	for pilihan != 12 {
		fmt.Println("===== GUDANGIN =====")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Ubah Barang")
		fmt.Println("3. Hapus Barang")
		fmt.Println("4. Transaksi Barang Masuk")
		fmt.Println("5. Transaksi Barang Keluar")
		fmt.Println("6. Tampilkan Data Barang")
		fmt.Println("7. Cek Menggunakan Nama Barang")
		fmt.Println("8. Cek Menggunakan Kode Barang")
		fmt.Println("9. Mengurutkan Stok Dari Terkecil")
		fmt.Println("10. Mengurutkan Stok Dari Terbesar")
		fmt.Println("11. Statistik Gudang")
		fmt.Println("12. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		fmt.Println()

		switch pilihan {
		case 1:
			tambahBarang(&data, &n)
		case 2:
			ubahBarang(&data, n)
		case 3:
			hapusBarang(&data, &n)
		case 4:
			transaksiMasuk(&data, n)
		case 5:
			transaksiKeluar(&data, n)
		case 6:
			printData(data, n)
		case 7:
			fmt.Print("Masukkan Nama barang yang dicari: ")
			fmt.Scan(&x)
			idx = sequentialSearch(data, n, x)
			if idx != -1 {
				fmt.Println("Barang ditemukan pada urutan ke ", idx)
				fmt.Println("Nama Barang :", data[idx].nama)
			} else {
				fmt.Println("Barang tidak ditemukan")
			}
		case 8:
			fmt.Print("Masukkan kode barang yang dicari: ")
			fmt.Scan(&x)
			idx = binarySearch(data, n, x)
			if idx != -1 {
				fmt.Println("Barang ditemukan pada urutan ke ", idx)
				fmt.Println("Nama Barang :", data[idx].nama)
			} else {
				fmt.Println("Barang tidak ditemukan")
			}
		case 9:
			selectionSort(&data, n)
			fmt.Println("Data berhasil diurutkan dari yang terkecil")
		case 10:
			insertionSort(&data, n)
			fmt.Println("Data berhasil diurutkan dari yang terbesar")
		case 11:
			statistikGudang(data, n)
		case 12:
			fmt.Println("Terima Kasih!")
		default:
			fmt.Println("Menu tidak tersedia")
		}
		fmt.Println()
	}
}
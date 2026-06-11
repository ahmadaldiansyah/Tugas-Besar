package main

import "fmt"

// Kapasitas maksimum data barang yang bisa disimpan
const NMAX = 1000

// Struct barang menyimpan informasi satu barang
type barang struct {
	kode  string // Kode unik barang
	nama  string // Nama barang
	stok  int    // Jumlah stok tersedia
	harga int    // Harga satuan barang
}

// tabBarang adalah tipe array statis berisi data barang
type tabBarang [NMAX]barang

// tambahBarang menambahkan data barang baru ke dalam tabel.
// Kode barang harus unik; stok dan harga tidak boleh negatif.
func tambahBarang(data *tabBarang, n *int) {
	var i int
	var found bool
	var x string

	// Cek apakah data sudah penuh
	if *n >= NMAX {
		fmt.Println("Data Penuh!")
	}

	found = false

	// Masukan Kode Barang baru minimal 4 karakter atau digit
	fmt.Print("Masukkan Kode Barang: ")
	fmt.Scan(&x)

	// Cek apakah kode barang sudah digunakan
	for i = 0; i < *n; i++ {
		if data[i].kode == x {
			found = true
		}
	}

	if found {
		fmt.Println("Kode Barang Sudah Digunakan!")
	} else {
		// Simpan kode dan input data barang baru
		data[*n].kode = x

		fmt.Print("Masukkan Nama Barang: ")
		fmt.Scan(&data[*n].nama)

		fmt.Print("Masukkan Stok Barang: ")
		fmt.Scan(&data[*n].stok)

		fmt.Print("Masukan Harga Barang: ")
		fmt.Scan(&data[*n].harga)

		// Validasi: stok tidak boleh negatif
		for data[*n].stok < 0 {
			fmt.Println("Stok Tidak Boleh Negatif!")
			fmt.Print("Masukkan Stok Barang: ")
			fmt.Scan(&data[*n].stok)
		}

		// Validasi: harga tidak boleh negatif
		for data[*n].harga < 0 {
			fmt.Println("Harga Tidak Boleh Negatif!")
			fmt.Print("Masukkan Harga Barang: ")
			fmt.Scan(&data[*n].harga)
		}

		// Tambah jumlah data
		*n = *n + 1

		fmt.Println("Data Berhasil Ditambahkan!")
	}
}

// ubahBarang mencari barang berdasarkan kode lalu mengubah nama, stok, dan harganya.
func ubahBarang(data *tabBarang, n int) {
	var i int
	var x string
	var found bool

	fmt.Print("Masukkan Kode Barang Yang Ingin Diubah: ")
	fmt.Scan(&x)

	found = false
	i = 0

	// Cari barang dengan kode yang sesuai (berhenti saat ditemukan)
	for i < n && !found {
		if data[i].kode == x {
			found = true
		} else {
			i = i + 1
		}
	}

	if found {
		fmt.Println("Data Ditemukan!")

		// Input data baru untuk menggantikan data lama
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

// hapusBarang menghapus barang berdasarkan kode dengan menggeser elemen-elemen sesudahnya.
func hapusBarang(data *tabBarang, n *int) {
	var i, j int
	var x string
	var found bool

	fmt.Print("Masukkan Kode Barang Yang Ingin Dihapus: ")
	fmt.Scan(&x)

	found = false
	i = 0

	// Cari posisi barang yang akan dihapus
	for i < *n && !found {
		if data[i].kode == x {
			found = true
		} else {
			i = i + 1
		}
	}

	if found {
		// Geser semua elemen setelah indeks i ke kiri satu posisi
		for j = i; j < *n-1; j++ {
			data[j] = data[j+1]
		}
		*n = *n - 1 // Kurangi jumlah data
		fmt.Println("Data Berhasil Dihapus!")
	} else {
		fmt.Println("Data Tidak Ditemukan!")
	}
}

// transaksiMasuk menambah stok barang berdasarkan kode barang yang diinput.
func transaksiMasuk(data *tabBarang, n int) {
	var i, stok1 int
	var x string
	var found bool

	fmt.Print("Masukkan Kode Barang: ")
	fmt.Scan(&x)

	found = false
	i = 0

	// Cari barang berdasarkan kode
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

		// Tambahkan jumlah barang masuk ke stok yang ada
		data[i].stok = data[i].stok + stok1

		fmt.Println("Transaksi Barang Masuk Berhasil")
	} else {
		fmt.Println("Barang Tidak Ditemukan!")
	}
}

// transaksiKeluar mengurangi stok barang; hanya diproses jika stok mencukupi.
func transaksiKeluar(data *tabBarang, n int) {
	var i, stok1 int
	var x string
	var found bool

	fmt.Print("Masukkan Kode Barang: ")
	fmt.Scan(&x)

	found = false
	i = 0

	// Cari barang berdasarkan kode
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

		// Cek apakah stok cukup sebelum mengurangi
		if stok1 <= data[i].stok {
			data[i].stok = data[i].stok - stok1
			fmt.Print("Transaksi barang keluar telah berhasil!")
		} else {
			fmt.Print("Stok barang tidak ditemukan")
		}
	} else {
		fmt.Print("Barang tidak ditemukan!")
	}
}

// printData menampilkan seluruh data barang yang tersimpan.
func printData(data tabBarang, n int) {
	var i int

	if n == 0 {
		fmt.Print("Barang kosong")
	} else {
		fmt.Println("==== DATA BARANG ====")
		for i = 0; i < n; i++ {
			fmt.Println("Data ke: ", i+1)
			fmt.Println("Kode: ", data[i].kode)
			fmt.Println("Nama: ", data[i].nama)
			fmt.Println("Stok: ", data[i].stok)
			fmt.Println("Harga: ", data[i].harga)
		}
		fmt.Println()
	}
}

// sequentialSearch mencari barang berdasarkan nama secara linear.
// Mengembalikan indeks barang jika ditemukan, atau -1 jika tidak.
func sequentialSearch(data tabBarang, n int, x string) int {
	var i int

	for i = 0; i < n; i++ {
		if data[i].nama == x {
			return i
		}
	}
	return -1
}

// binarySearch mencari barang berdasarkan kode secara biner.
// Syarat: data harus sudah diurutkan berdasarkan kode.
// Mengembalikan indeks barang jika ditemukan, atau -1 jika tidak.
func binarySearch(data tabBarang, n int, x string) int {
	var left, right, mid int
	left = 0
	right = n - 1

	for left <= right {
		mid = (left + right) / 2
		if data[mid].kode == x {
			return mid // Kode ditemukan
		} else if x < data[mid].kode {
			right = mid - 1 // Cari di bagian kiri
		} else {
			left = mid + 1 // Cari di bagian kanan
		}
	}

	return -1
}

// selectionSort mengurutkan data barang berdasarkan kode secara ascending
// menggunakan algoritma Selection Sort.
func selectionSort(data *tabBarang, n int) {
	var i, pass, idxMin int
	var temp barang

	for i = 0; i < n-1; i++ {
		idxMin = i // Anggap elemen ke-i adalah yang terkecil

		// Cari elemen terkecil dari posisi i+1 sampai akhir
		for pass = i + 1; pass < n; pass++ {
			if data[pass].kode < data[idxMin].kode {
				idxMin = pass
			}
		}

		// Tukar elemen ke-i dengan elemen terkecil yang ditemukan
		temp = data[i]
		data[i] = data[idxMin]
		data[idxMin] = temp
	}
}

// insertionSort mengurutkan data barang berdasarkan stok secara descending (terbesar ke terkecil)
// menggunakan algoritma Insertion Sort.
func insertionSort(data *tabBarang, n int) {
	var pass, i int
	var temp barang

	pass = 1

	for pass <= n-1 {
		i = pass
		temp = data[pass] // Simpan elemen yang akan disisipkan

		// Geser elemen yang lebih kecil ke kanan untuk memberi tempat
		for i > 0 && temp.stok > data[i-1].stok {
			data[i] = data[i-1]
			i = i - 1
		}

		// Sisipkan elemen pada posisi yang tepat
		data[i] = temp
		pass = pass + 1
	}
}

// statistikGudang menampilkan total aset, barang dengan stok terendah, dan tertinggi.
func statistikGudang(data tabBarang, n int) {
	var i int
	var totalAset int
	var idxMin, idxMax int

	if n > 0 {
		idxMin = 0
		idxMax = 0

		for i = 0; i < n; i++ {
			// Akumulasi total aset: stok * harga tiap barang
			totalAset += data[i].stok * data[i].harga

			// Update indeks barang dengan stok minimum
			if data[i].stok < data[idxMin].stok {
				idxMin = i
			}

			// Update indeks barang dengan stok maksimum
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

// main adalah fungsi utama yang menjalankan program manajemen gudang (GUDANGIN).
func main() {
	var data tabBarang
	var n, pilihan, idx int
	var x string

	n = 0 // Inisialisasi jumlah barang = 0

	// Loop utama: terus tampilkan menu sampai user memilih keluar (12)
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
			// Pencarian nama barang menggunakan sequential search
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
			// Pencarian kode barang menggunakan binary search (data harus terurut by kode)
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
			// Urutkan berdasarkan kode (ascending) menggunakan selection sort
			selectionSort(&data, n)
			fmt.Println("Data berhasil diurutkan dari yang terkecil")
		case 10:
			// Urutkan berdasarkan stok (descending) menggunakan insertion sort
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
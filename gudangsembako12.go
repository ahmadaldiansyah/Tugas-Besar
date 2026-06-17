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

// ===============================================================
// DATA DUMMY UNTUK TESTING AWAL PROGRAM
// ===============================================================

func isiDataDummy(data *tabBarang, n *int) {
	data[*n].kode = "B001"
	data[*n].nama = "Beras"
	data[*n].stok = 50
	data[*n].harga = 75000
	*n = *n + 1

	data[*n].kode = "B002"
	data[*n].nama = "Gula"
	data[*n].stok = 30
	data[*n].harga = 18000
	*n = *n + 1

	data[*n].kode = "B003"
	data[*n].nama = "Minyak"
	data[*n].stok = 25
	data[*n].harga = 22000
	*n = *n + 1

	data[*n].kode = "B004"
	data[*n].nama = "Tepung"
	data[*n].stok = 40
	data[*n].harga = 15000
	*n = *n + 1

	data[*n].kode = "B005"
	data[*n].nama = "Garam"
	data[*n].stok = 60
	data[*n].harga = 5000
	*n = *n + 1
}

// Function untuk mengetahui apakah kode barang sudah ada di dalam data atau belum
func kodeSudahAda(data tabBarang, n int, x string) bool {
    var i int
    for i = 0; i < n; i++ {
        if data[i].kode == x {
            return true
        }
    }
    return false
}

// ===============================================================
// DATA MENAMBAH BARANG
// ===============================================================

// tambahBarang menambahkan data barang baru ke dalam tabel.
// Kode barang harus unik; stok dan harga tidak boleh negatif.
func tambahBarang(data *tabBarang, n *int) {
	var x string

	// Cek apakah data sudah penuh
	if *n >= NMAX {
		fmt.Println("Data Penuh!")
	} else {
		// Input kode barang baru
		fmt.Print("Masukkan Kode Barang: ")
		fmt.Scan(&x)
		
		// Cek Jumlah data saat ini
		fmt.Println("Jumlah data =", *n)

		// Cek apakah kode barang sudah ada
		if kodeSudahAda(*data, *n, x) {
			fmt.Println("Kode Barang Sudah Digunakan!")
		} else {
			fmt.Println("Kode belum ada, data akan ditambahkan")
			
			// Masukkan data baru ke dalam tabel
			data[*n].kode = x

			// Input nama
			fmt.Print("Masukkan Nama Barang: ")
			fmt.Scan(&data[*n].nama)

			// Input stok
			fmt.Print("Masukkan Stok Barang: ")
			fmt.Scan(&data[*n].stok)

			// Input harga
			fmt.Print("Masukkan Harga Barang: ")
			fmt.Scan(&data[*n].harga)

			// Validasi stok agar tidak negatif
			for data[*n].stok < 0 {
				fmt.Println("Stok Tidak Boleh Negatif!")
				fmt.Print("Masukkan Stok Barang: ")
				fmt.Scan(&data[*n].stok)
			}

			// Validasi harga agar tidak negatif
			for data[*n].harga < 0 {
				fmt.Println("Harga Tidak Boleh Negatif!")
				fmt.Print("Masukkan Harga Barang: ")
				fmt.Scan(&data[*n].harga)
			}
			// Tambahkan jumlah data
			*n = *n + 1

			// Tampilkan pesan sukses
			fmt.Println("Data Berhasil Ditambahkan!")
		}
	}
}

// ===============================================================
// DATA MENGUBAH BARANG
// ===============================================================
// ubahBarang mencari barang berdasarkan kode lalu mengubah nama, stok, dan harganya.
func ubahBarang(data *tabBarang, n int) {
	var i int
	var x string
	var found bool

	// Input kode barang yang ingin diubah
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

	// Jika barang ditemukan, input data baru untuk menggantikan data lama
	if found {
		fmt.Println("Data Ditemukan!")

		// Input data baru untuk menggantikan data lama
		fmt.Print("Nama Barang Baru: ")
		fmt.Scan(&data[i].nama)

		// Input stok baru
		fmt.Print("Jumlah Stok Baru: ")
		fmt.Scan(&data[i].stok)

		// Input harga baru
		fmt.Print("Harga Barang Baru: ")
		fmt.Scan(&data[i].harga)

		// Menampilkan pesan sukses setelah data berhasil diubah
		fmt.Println("Data Berhasil Diubah!")
	} else {
		fmt.Println("Data Tidak Ditemukan!") // Menampilkan pesan jika barang dengan kode yang dimasukkan tidak ditemukan
	}
}

// ===============================================================
// DATA MENGHAPUS BARANG
// ===============================================================
// hapusBarang menghapus barang berdasarkan kode dengan menggeser elemen-elemen sesudahnya.
func hapusBarang(data *tabBarang, n *int) {
	var i, j int
	var x string
	var found bool

	// Input kode barang yang ingin dihapus
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

	// Jika barang ditemukan, hapus dengan menggeser elemen-elemen setelahnya ke kiri
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

// ===============================================================
// DATA TRANSAKSI MASUK DAN KELUAR
// ===============================================================

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

// ===============================================================
// DATA UNTUK MENAMPILKAN
// ===============================================================

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

// ===============================================================
// DATA UNTUK MENAMPILKAN DATA DUMMY
// ===============================================================
func printDataDummy(data tabBarang, n int) {
	var i int

	if n == 0 {
		fmt.Print("Barang kosong")
	} else {
		fmt.Println("==== DATA BARANG DUMMY ====")
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

// ===============================================================
// DATA UNTUK MENCARI BARANG BERDASARKAN NAMA DAN KODE
// ===============================================================
// sequentialSearch mencari barang berdasarkan nama secara linear.
// Mengembalikan indeks barang jika ditemukan, atau -1 jika tidak.
func sequentialSearchNama(data tabBarang, n int, x string) int {
	var i int

	for i = 0; i < n; i++ {
		if data[i].nama == x {
			return i
		}
	}
	return -1
}

// ===============================================================
// DATA UNTUK MENCARI BARANG BERDASARKAN HARGA
// ===============================================================
// sequentialSearch mencari barang berdasarkan harga secara linear.
// Mengembalikan indeks barang jika ditemukan, atau -1 jika tidak.
func sequentialSearchHarga(data tabBarang, n int, x int) int {
	var i int

	for i = 0; i < n; i++ {
		if data[i].harga == x {
			return i
		}
	}
	return -1
}

// ===============================================================
// DATA UNTUK MENCARI BARANG BERDASARKAN KODE MENGGUNAKAN BINARY SEARCH
// ===============================================================
// binarySearch mencari barang berdasarkan kode secara biner.
// Syarat: data harus sudah diurutkan berdasarkan kode.
// Mengembalikan indeks barang jika ditemukan, atau -1 jika tidak.
func binarySearchKode(data tabBarang, n int, x string) int {
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

// ==================================================================================
// DATA UNTUK MENCARI BARANG BERDASARKAN KODE MENGGUNAKAN SELECTION SORT (ASCENDING)
// ==================================================================================
// selectionSort mengurutkan data barang berdasarkan kode secara ascending
// menggunakan algoritma Selection Sort.
func selectionSortKode(data *tabBarang, n int) {
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

// =================================================================================
// DATA UNTUK MENCARI BARANG BERDASARKAN STOK MENGGUNAKAN SELECTION SORT (ASCENDING)
// =================================================================================
func selectionSortStok(data *tabBarang, n int) {
	var i, pass, idxMin int
	var temp barang

	for i = 0; i < n-1; i++ {
		idxMin = i // Anggap elemen ke-i adalah yang terkecil

		// Cari elemen terkecil dari posisi i+1 sampai akhir
		for pass = i + 1; pass < n; pass++ {
			if data[pass].stok < data[idxMin].stok {
				idxMin = pass
			}
		}

		// Tukar elemen ke-i dengan elemen terkecil yang ditemukan
		temp = data[i]
		data[i] = data[idxMin]
		data[idxMin] = temp
	}
}

// ==================================================================================
// DATA UNTUK MENCARI BARANG BERDASARKAN HARGA MENGGUNAKAN SELECTION SORT (ASCENDING)
// ==================================================================================
func selectionSortHarga(data *tabBarang, n int) {
	var i, pass, idxMin int
	var temp barang

	for i = 0; i < n-1; i++ {
		idxMin = i // Anggap elemen ke-i adalah yang terkecil

		// Cari elemen terkecil dari posisi i+1 sampai akhir
		for pass = i + 1; pass < n; pass++ {
			if data[pass].harga < data[idxMin].harga {
				idxMin = pass
			}
		}

		// Tukar elemen ke-i dengan elemen terkecil yang ditemukan
		temp = data[i]
		data[i] = data[idxMin]
		data[idxMin] = temp
	}
}
// ===================================================================================
// DATA UNTUK MENCARI BARANG BERDASARKAN KODE MENGGUNAKAN INSERSTION SORT (DESCENDING)
// ===================================================================================
// insertionSort mengurutkan data barang berdasarkan stok secara descending (terbesar ke terkecil)
// menggunakan algoritma Insertion Sort.
func insertionSortKode(data *tabBarang, n int) {
	var pass, i int
	var temp barang

	pass = 1

	for pass <= n-1 {
		i = pass
		temp = data[pass] // Simpan elemen yang akan disisipkan

		// Geser elemen yang lebih kecil ke kanan untuk memberi tempat
		for i > 0 && temp.kode > data[i-1].kode {
			data[i] = data[i-1]
			i = i - 1
		}

		// Sisipkan elemen pada posisi yang tepat
		data[i] = temp
		pass = pass + 1
	}
}

// ===================================================================================
// DATA UNTUK MENCARI BARANG BERDASARKAN STOK MENGGUNAKAN INSERSTION SORT (DESCENDING)
// ===================================================================================
func insertionSortStok(data *tabBarang, n int) {
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

// ====================================================================================
// DATA UNTUK MENCARI BARANG BERDASARKAN HARGA MENGGUNAKAN INSERSTION SORT (DESCENDING)
// ====================================================================================
func insertionSortHarga(data *tabBarang, n int) {
	var pass, i int
	var temp barang

	pass = 1

	for pass <= n-1 {
		i = pass
		temp = data[pass] // Simpan elemen yang akan disisipkan

		// Geser elemen yang lebih kecil ke kanan untuk memberi tempat
		for i > 0 && temp.harga > data[i-1].harga {
			data[i] = data[i-1]
			i = i - 1
		}

		// Sisipkan elemen pada posisi yang tepat
		data[i] = temp
		pass = pass + 1
	}
}

// ===============================================================
// DATA STATISTIK GUDANG
// ===============================================================
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
			totalAset = totalAset + data[i].stok * data[i].harga

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

func pilihanMenu(n int) {
	var i int

	n = 1

	for i = 0; i < n ; i++ {
		fmt.Println("===== GUDANGIN =====")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Ubah Barang")
		fmt.Println("3. Hapus Barang")
		fmt.Println("4. Transaksi Barang Masuk")
		fmt.Println("5. Transaksi Barang Keluar")
		fmt.Println("6. Tampilkan Data Barang")
		fmt.Println("7. Cek Menggunakan Nama Barang")
		fmt.Println("8. Cek Menggunakan Kode Barang")
		fmt.Println("9. Cek Menggunakan Harga Barang")
		fmt.Println("10. Mengurutkan Kode Dari Yang Terkecil (ASCENDING)")
		fmt.Println("11. Mengurutkan Stok Dari Yang Terkecil (ASCENDING)")
		fmt.Println("12. Mengurutkan Harga Dari Yang Terkecil (ASCENDING)")
		fmt.Println("13. Mengurutkan Stok Dari Yang Terbesar (DESCENDING)")
		fmt.Println("14. Mengurutkan Kode Dari Yang Terbesar (DESCENDING)")
		fmt.Println("15. Mengurutkan Harga Dari Yang Terbesar (DESCENDING)")
		fmt.Println("16. Statistik Gudang")
		fmt.Println("17. Keluar")
		fmt.Print("Pilih menu: ")
	}
}

// main adalah fungsi utama yang menjalankan program manajemen gudang (GUDANGIN).
func main() {
	var data tabBarang
	var n, pilihan, idx int
	var x string
	var y int

	n = 0 // Inisialisasi jumlah barang = 0

	// Procedure untuk mengisi data dummy dengan data dummy yang sudah disiapkan sebanyak 5 data barang
	isiDataDummy(&data, &n)

	// Tampilkan data dummya yang sudah diisi untuk memastikan data sudah masuk dengan benar
	printDataDummy(data, n)

	// Loop utama: terus tampilkan menu sampai user memilih keluar (17)
	for pilihan != 17 {
		pilihanMenu(n) // Procedure untuk tampilkan menu dan input pilihan
		fmt.Scan(&pilihan)
		fmt.Println()

	// Proses pilihan menu menggunakan switch-case
		switch pilihan {
		case 1:
			fmt.Println("=========================================================")
			fmt.Printf("%40s\n", "MENAMBAH DATA BARANG BARU")
			fmt.Println("=========================================================")
			tambahBarang(&data, &n)
		case 2:
			fmt.Println("=========================================================")
			fmt.Printf("%40s\n", "MENGUBAH DATA BARANG")
			fmt.Println("=========================================================")
			ubahBarang(&data, n)
		case 3:
			fmt.Println("=========================================================")
			fmt.Printf("%40s\n", "MENGAHAPUS DATA BARANG")
			fmt.Println("=========================================================")
			hapusBarang(&data, &n)
		case 4:
			fmt.Println("=========================================================")
			fmt.Printf("%40s\n", "TRANSAKSI MASUK DATA BARANG")
			fmt.Println("=========================================================")
			transaksiMasuk(&data, n)
		case 5:
			fmt.Println("=========================================================")
			fmt.Printf("%40s\n", "TRANSAKSI KELUAR DATA BARANG")
			fmt.Println("=========================================================")
			transaksiKeluar(&data, n)
		case 6:
			fmt.Println("=========================================================")
			fmt.Printf("%40s\n", "MENAMPILKAN DATA BARANG")
			fmt.Println("=========================================================")
			printData(data, n)
		case 7:
			fmt.Println("=========================================================")
			fmt.Printf("%40s\n", "MENCARI DATA BARANG BERDASARKAN NAMA")
			fmt.Println("=========================================================")
			// Pencarian nama barang menggunakan sequential search
			fmt.Print("Masukkan Nama barang yang dicari: ")
			fmt.Scan(&x)
			idx = sequentialSearchNama(data, n, x)
			if idx != -1 {
				fmt.Println("Barang ditemukan pada urutan ke ", idx)
				fmt.Println("Nama Barang :", data[idx].nama)
			} else {
				fmt.Println("Barang tidak ditemukan")
			}
		case 8:
			fmt.Println("=========================================================")
			fmt.Printf("%40s\n", "MENCARI DATA BARANG BERDASARKAN KODE")
			fmt.Println("=========================================================")
			// Pencarian kode barang menggunakan binary search (data harus terurut by kode)
			fmt.Print("Masukkan kode barang yang dicari: ")
			fmt.Scan(&x)
			idx = binarySearchKode(data, n, x)
			if idx != -1 {
				fmt.Println("Barang ditemukan pada urutan ke ", idx)
				fmt.Println("Nama Barang :", data[idx].nama)
			} else {
				fmt.Println("Barang tidak ditemukan")
			}
		case 9:
			fmt.Println("=========================================================")
			fmt.Printf("%40s\n", "MENCARI DATA BARANG BERDASARKAN HARGA")
			fmt.Println("=========================================================")
			// Pencarian harga barang menggunakan sequential search
			fmt.Print("Masukkan Harga barang yang dicari: ")
			fmt.Scan(&y)
			idx = sequentialSearchHarga(data, n, y)
			if idx != -1 {
				fmt.Println("Barang ditemukan pada urutan ke ", idx)
				fmt.Println("Nama Barang :", data[idx].nama)
			} else {
				fmt.Println("Barang tidak ditemukan")
			}
		case 10:
			fmt.Println("=========================================================")
			fmt.Printf("%40s\n", "MENCARI DATA BARANG BERDASARKAN KODE (ASCENDING)")
			fmt.Println("=========================================================")
			// Urutkan berdasarkan kode (ascending) menggunakan selection sort
			selectionSortKode(&data, n)
			printData(data, n)
			fmt.Println("Data berhasil diurutkan dari Kode yang terkecil (ASCENDING)")
		case 11:
			fmt.Println("=========================================================")
			fmt.Printf("%40s\n", "MENCARI DATA BARANG BERDASARKAN STOK (ASCENDING)")
			fmt.Println("=========================================================")
			// Urutkan berdasarkan kode (ascending) menggunakan selection sort
			selectionSortStok(&data, n)
			printData(data, n)
			fmt.Println("Data berhasil diurutkan dari Stok yang terkecil (ASCENDING)")
		case 12:
			fmt.Println("=========================================================")
			fmt.Printf("%40s\n", "MENCARI DATA BARANG BERDASARKAN HARGA (ASCENDING)")
			fmt.Println("=========================================================")
			// Urutkan berdasarkan kode (ascending) menggunakan selection sort
			selectionSortHarga(&data, n)
			printData(data, n)
			fmt.Println("Data berhasil diurutkan dari Harga yang terkecil (ASCENDING)")
		case 13:
			fmt.Println("=========================================================")
			fmt.Printf("%40s\n", "MENCARI DATA BARANG BERDASARKAN STOK (DESCENDING)")
			fmt.Println("=========================================================")
			// Urutkan berdasarkan stok (descending) menggunakan insertion sort
			insertionSortStok(&data, n)
			printData(data, n)
			fmt.Println("Data berhasil diurutkan dari Stok yang terbesar (DESCENDING)")
		case 14:
			fmt.Println("=========================================================")
			fmt.Printf("%40s\n", "MENCARI DATA BARANG BERDASARKAN KODE (DESCENDING)")
			fmt.Println("=========================================================")
			// Urutkan berdasarkan stok (descending) menggunakan insertion sort
			insertionSortKode(&data, n)
			printData(data, n)
			fmt.Println("Data berhasil diurutkan dari Kode yang terbesar (DESCENDING)")
		case 15:
			fmt.Println("=========================================================")
			fmt.Printf("%40s\n", "MENCARI DATA BARANG BERDASARKAN HARGA (DESCENDING)")
			fmt.Println("=========================================================")
			// Urutkan berdasarkan stok (descending) menggunakan insertion sort
			insertionSortHarga(&data, n)
			printData(data, n)
			fmt.Println("Data berhasil diurutkan dari harga yang terbesar (DESCENDING)")
		case 16:
			fmt.Println("=========================================================")
			fmt.Printf("%40s\n", "STATISTIK GUDANG")
			fmt.Println("=========================================================")
			statistikGudang(data, n)
		case 17:
			fmt.Println("Terima Kasih!")
		default:
			fmt.Println("Menu tidak tersedia")
		}
		fmt.Println()
	}
}
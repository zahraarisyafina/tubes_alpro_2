package main

import "fmt"

const NMAX int = 1500

// Jumlah maksimum data sembako yang dapat disimpan

type daftarSembako struct {
	nama         string
	kode         string
	harga        float64
	stok         int
	kategori     string
	barangMasuk  int
	barangKeluar int
}

// Tipe bentukan untuk menyimpan informasi barang sembako
// meliputi nama, kode, harga, stok, kategori,
// jumlah barang masuk, dan jumlah barang keluar.

type tabSembako [NMAX]daftarSembako

// Array yang menyimpan kumpulan data sembako

var listKategori = [9]string{
	"Beras",
	"Gula Pasir",
	"Minyak Goreng",
	"Daging",
	"Telur Ayam",
	"Susu",
	"Bawang",
	"Gas Elpiji",
	"Garam",
}

// Array kategori sembako yang digunakan untuk pilihan kategori produk yang akan diInputkan

func menuAwal() {
	/*IS :
	- Program telah dijalankan.
	- Tampilan awal aplikasi belum ditampilkan.

	FS :
	- Tampilan awal aplikasi Manajemen Toko Sembako ditampilkan pada layar.*/

	fmt.Println()
	fmt.Println("+==========================================================+")
	fmt.Println("|                                                          |")
	fmt.Println("|              SELAMAT DATANG DI APLIKASI                  |")
	fmt.Println("|                 MANAJEMEN TOKO SEMBAKO                   |")
	fmt.Println("|                                                          |")
	fmt.Println("+==========================================================+")
	fmt.Println()
}

func menuUtama() {
	/* IS :
	- Program sedang berjalan dan siap menampilkan menu utama.
	- Pengguna belum memilih menu.

	FS :
	- Daftar menu aplikasi ditampilkan pada layar.
	- Program meminta pengguna memasukkan pilihan menu.
	*/

	fmt.Println("+==========================================================+")
	fmt.Println("|                    DAFTAR MENU                           |")
	fmt.Println("+==========================================================+")
	fmt.Println("| 1. Tambah Produk Sembako                                 |")
	fmt.Println("| 2. Tampilkan Data Produk                                 |")
	fmt.Println("| 3. Barang Masuk                                          |")
	fmt.Println("| 4. Barang Keluar                                         |")
	fmt.Println("| 5. Cari Produk                                           |")
	fmt.Println("| 6. Edit Produk                                           |")
	fmt.Println("| 7. Hapus Produk                                          |")
	fmt.Println("| 8. Urutkan Produk                                        |")
	fmt.Println("| 9. Laporan Barang Masuk dan Keluar                       |")
	fmt.Println("| 10. Cari Stok Terbanyak dan Tersedikit                   |")
	fmt.Println("| 0. Keluar                                                |")
	fmt.Println("+==========================================================+")
	fmt.Print("Pilih menu: ")
}

func tampilKategori() {
	/* IS :
	- Data kategori sembako telah tersimpan pada array listKategori.
	FS :
	- Seluruh kategori pada listKategori ditampilkan pada layar.
	===========
	membaca isi array listKategori satu per satu menggunakan perulangan for,
	lalu menampilkannya dengan nomor urut mulai dari 1*/

	var categoryIndex int
	fmt.Println()
	fmt.Println("Daftar Kategori:")
	for categoryIndex = 0; categoryIndex < len(listKategori); categoryIndex++ {
		fmt.Printf("%d. %s\n", categoryIndex+1, listKategori[categoryIndex])
	}
}

func dataAwal(dataSembako *tabSembako, jumlahData *int) {
	/* IS : dataSemabko dan jumlahData terdefinisi

	FS : - dataSembako berisi 9 data awal produk sembako.
		- jumlahData = 9.
	*/
	dataSembako[0] = daftarSembako{"Beras_Ramos_5kg", "BR001", 76000, 30, "Beras", 30, 0}
	dataSembako[1] = daftarSembako{"Gula_Pasir_Gulaku_1kg", "GP001", 18000, 45, "Gula Pasir", 45, 0}
	dataSembako[2] = daftarSembako{"Minyak_Goreng_Bimoli_2L", "MG001", 39000, 25, "Minyak Goreng", 25, 0}
	dataSembako[3] = daftarSembako{"Telur_Ayam_Negeri_1kg", "TA001", 31000, 40, "Telur Ayam", 40, 0}
	dataSembako[4] = daftarSembako{"Gas_Elpiji_3kg", "GE001", 22000, 15, "Gas Elpiji", 15, 0}
	dataSembako[5] = daftarSembako{"Daging_Ayam_2kg", "DA001", 35000, 22, "Daging", 22, 0}
	dataSembako[6] = daftarSembako{"Susu_Sapi_1L", "SS001", 55000, 19, "Susu", 19, 0}
	dataSembako[7] = daftarSembako{"Bawang_Merah_2Kg", "BM001", 23000, 27, "Bawang", 27, 0}
	dataSembako[8] = daftarSembako{"Garam_Himalaya_1Kg", "GM001", 50000, 26, "Garam", 26, 0}
	*jumlahData = 9
}

func tambahProduk(dataSembako *tabSembako, jumlahData *int) {
	/*IS : - dataSembako dan JumlahData terdefinisi
	           - jumlahData menyatakan banyaknya produk yang tersimpan.

		FS :
		- Jika data yang dimasukkan valid dan kapasitas array belum penuh, maka produk baru ditambahkan ke dataSembako dan jumlahData
		  bertambah 1.
		- Jika data tidak valid atau array penuh, maka dataSembako dan jumlahData tidak berubah. */

	var pilihanKategori int

	if *jumlahData >= NMAX {
		fmt.Println("Data sudah penuh.")
	} else {
		fmt.Println()
		fmt.Println("+================ TAMBAH PRODUK SEMBAKO =================+")

		fmt.Print("Masukkan kode produk (MAX 5): ")
		fmt.Scan(&dataSembako[*jumlahData].kode)

		if len(dataSembako[*jumlahData].kode) > 5 {
			fmt.Println("KODE TIDAK BOLEH LEBIH DARI 5")
		} else {

			fmt.Print("Masukkan nama produk: ")
			fmt.Scan(&dataSembako[*jumlahData].nama)

			fmt.Print("Masukkan harga produk: ")
			fmt.Scan(&dataSembako[*jumlahData].harga)

			fmt.Print("Masukkan stok awal: ")
			fmt.Scan(&dataSembako[*jumlahData].stok)

			if dataSembako[*jumlahData].stok < 0 {
				fmt.Println("Stok tidak boleh negatif.")
			} else {

				tampilKategori()
				fmt.Print("Pilih kategori: ")
				fmt.Scan(&pilihanKategori)

				if pilihanKategori < 1 || pilihanKategori > len(listKategori) {
					fmt.Println("Kategori tidak valid. Produk gagal ditambahkan.")
				} else {
					dataSembako[*jumlahData].kategori = listKategori[pilihanKategori-1]
					dataSembako[*jumlahData].barangMasuk = dataSembako[*jumlahData].stok
					dataSembako[*jumlahData].barangKeluar = 0

					(*jumlahData)++

					fmt.Println("Produk berhasil ditambahkan.")
					tampilProduk(*dataSembako, *jumlahData)
				}
			}
		}
	}
}

func tampilProduk(dataSembako tabSembako, jumlahData int) {
	/* IS : data sembako dan jumlah data terdefinisi, dengan jumlah data > 0
	FS : Menampilkan data produk meliputi nomor urut, kode, nama produk, kategori, harga, stok, jumlah keluar dan masuknya produk. Jika jumlah
	data yang dimasukkan = 0 maka akan keluar "Belum ada data produk."*/

	var productIndex int
	if jumlahData == 0 {
		fmt.Println("Belum ada data produk.")
	} else {
		fmt.Println()
		fmt.Println("+==============================================================================================================+")
		fmt.Println("| No | Kode | Nama Produk                     | Kategori                  | Harga      | Stok | Masuk | Keluar |")
		fmt.Println("+==============================================================================================================+")

		for productIndex = 0; productIndex < jumlahData; productIndex++ {
			fmt.Printf("| %-2d | %-5s | %-30s | %-25s | %-10.0f | %-4d | %-5d | %-6d |\n",
				productIndex+1,
				dataSembako[productIndex].kode,
				dataSembako[productIndex].nama,
				dataSembako[productIndex].kategori,
				dataSembako[productIndex].harga,
				dataSembako[productIndex].stok,
				dataSembako[productIndex].barangMasuk,
				dataSembako[productIndex].barangKeluar,
			)
		}

		fmt.Println("+==============================================================================================================+")
	}
}

func cariProdukSequential(dataSembako tabSembako, jumlahData int, kodeProduk string) int {
	/* Mengembalikan productIndex berupa bilangan bulat untuk mencari produk secara sequential serta mengembalikan -1 jika produk tidak ditemukan*/
	var productIndex int
	for productIndex = 0; productIndex < jumlahData; productIndex++ {
		if dataSembako[productIndex].kode == kodeProduk {
			return productIndex
		}
	}
	return -1
}

func cariPNamabinary(dataSembako tabSembako, jumlahData int, namaProduk string) int {
	/*dataSembako telah terurut menaik berdasarkan nama produk, jumlahData dan namaProduk terdefinisi. Mengembalikan idx produk yang
	di cari berupa bilangan bulat dan jika produk tidak ditemukan maka akan mengembalikan idx dengan nilai -1*/
	var left, right, mid int
	var idx int

	left = 0
	right = jumlahData - 1
	idx = -1

	for left <= right && idx == -1 {
		mid = (left + right) / 2

		if dataSembako[mid].nama == namaProduk {
			idx = mid
		} else if dataSembako[mid].nama > namaProduk {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return idx
}

func barangMasuk(dataSembako *tabSembako, jumlahData int) {
	/*IS : data sembako dan jumlah data terdefinisi
			FS : jika data berupa kode produk yang dimasukkan valid serta jumlah barang masuk > 0 maka jumlah stok produk tersebut akan
			bertambah sesuai dengan banyaknya barang masuk. Jika kode produk tidak ditemukan atau jumlahBarangMasuk <= 0,
	        maka data tidak berubah.
	*/
	var kodeProduk string
	var jumlahBarangMasuk, productIndex int

	fmt.Println()
	tampilProduk(*dataSembako, jumlahData)
	fmt.Println("+================ BARANG MASUK =================+")
	fmt.Print("Masukkan kode produk: ")
	fmt.Scan(&kodeProduk)

	productIndex = cariProdukSequential(*dataSembako, jumlahData, kodeProduk)

	if productIndex == -1 {
		fmt.Println("Produk tidak ditemukan.")
	} else {
		fmt.Print("Masukkan jumlah barang masuk: ")
		fmt.Scan(&jumlahBarangMasuk)

		if jumlahBarangMasuk <= 0 {
			fmt.Println("Jumlah barang masuk harus lebih dari 0.")
		} else {
			dataSembako[productIndex].stok += jumlahBarangMasuk
			dataSembako[productIndex].barangMasuk += jumlahBarangMasuk

			fmt.Println("Barang masuk berhasil dicatat.")
			fmt.Println("Stok terbaru:", dataSembako[productIndex].stok)
		}
	}
}

func barangKeluar(dataSembako *tabSembako, jumlahData int) {
	/*IS : data sembako dan jumlah data terdefinisi
	FS : jika data berupa kode produk yang dimasukkan valid serta jumlah barang keluar > 0 maka jumlah stok produk tersebut akan
	berkurang di kolom banyak stok, serta mengupdate data di kolom keluar sesuai dengan banyaknya barang keluar.
	Jika kode produk tidak ditemukan atau jumlahBarangKeluar <= 0, maka data tidak berubah.
	*/
	var kodeProduk string
	var jumlahBarangKeluar, productIndex int

	fmt.Println()
	tampilProduk(*dataSembako, jumlahData)

	fmt.Println("+================ BARANG KELUAR =================+")
	fmt.Print("Masukkan kode produk: ")
	fmt.Scan(&kodeProduk)

	productIndex = cariProdukSequential(*dataSembako, jumlahData, kodeProduk)

	if productIndex == -1 {
		fmt.Println("Produk tidak ditemukan.")
	} else {
		fmt.Print("Masukkan jumlah barang keluar: ")
		fmt.Scan(&jumlahBarangKeluar)

		if jumlahBarangKeluar <= 0 {
			fmt.Println("Jumlah barang keluar harus lebih dari 0.")

		} else if jumlahBarangKeluar > dataSembako[productIndex].stok {
			fmt.Println("Stok tidak mencukupi.")
		} else {
			dataSembako[productIndex].stok -= jumlahBarangKeluar
			dataSembako[productIndex].barangKeluar += jumlahBarangKeluar

			fmt.Println("Barang keluar berhasil dicatat.")
			fmt.Println("Stok terbaru:", dataSembako[productIndex].stok)
		}
	}
}

func cariProdukSeq(dataSembako tabSembako, jumlahData int) {
	/* IS :
	   - dataSembako dan jumlahData terdefinisi.

	   FS :
	   - Menampilkan data produk yang memiliki kode sesuai input pengguna.
	   - Jika produk tidak ditemukan, menampilkan pesan bahwa produk tidak ditemukan.
	*/
	var kodeProduk string
	var productIndex int

	fmt.Println()
	fmt.Println("+===========++====== CARI PRODUK =================+")
	fmt.Print("Masukkan kode produk: ")
	fmt.Scan(&kodeProduk)

	productIndex = cariProdukSequential(dataSembako, jumlahData, kodeProduk)

	if productIndex == -1 {
		fmt.Println("Produk tidak ditemukan.")
	} else {
		fmt.Println("Produk ditemukan:")
		fmt.Println("kode        :", dataSembako[productIndex].kode)
		fmt.Println("Nama         :", dataSembako[productIndex].nama)
		fmt.Println("Kategori     :", dataSembako[productIndex].kategori)
		fmt.Println("Harga        :", dataSembako[productIndex].harga)
		fmt.Println("Stok         :", dataSembako[productIndex].stok)
		fmt.Println("Barang Masuk :", dataSembako[productIndex].barangMasuk)
		fmt.Println("Barang Keluar:", dataSembako[productIndex].barangKeluar)
	}
}

func insertionSortKode(dataSembako tabSembako, jumlahData int) tabSembako {
	/*dataSembako dan jumlahData terdefinisi dan jumlahData menyatakan banyaknya data yang akan diurutkan.
	  mengembalikan dataSembako yang telah terurut menaik(ascending) berdasarkan kode produk.*/
	var currentIndex, previousIndex int
	var tempProduct daftarSembako

	for currentIndex = 1; currentIndex < jumlahData; currentIndex++ {
		tempProduct = dataSembako[currentIndex]
		previousIndex = currentIndex - 1

		for previousIndex >= 0 && dataSembako[previousIndex].kode > tempProduct.kode {
			dataSembako[previousIndex+1] = dataSembako[previousIndex]
			previousIndex--
		}

		dataSembako[previousIndex+1] = tempProduct
	}

	return dataSembako
}

func binarySearchPreparationSort(dataSembako tabSembako, jumlahData int) tabSembako {
	/*dataSembako dan jumlahData terdefinisi dan jumlahData menyatakan banyaknya data yang akan diurutkan.
	  mengemvalikan dataSembako yang telah terurut menaik(ascending) berdasarkan nama produk.*/
	var currentIndex, previousIndex int
	var tempProduct daftarSembako

	for currentIndex = 1; currentIndex < jumlahData; currentIndex++ {
		tempProduct = dataSembako[currentIndex]
		previousIndex = currentIndex - 1

		for previousIndex >= 0 && dataSembako[previousIndex].nama > tempProduct.nama {
			dataSembako[previousIndex+1] = dataSembako[previousIndex]
			previousIndex--
		}

		dataSembako[previousIndex+1] = tempProduct
	}

	return dataSembako
}

func cariProduk(dataSembako tabSembako, jumlahData int) {
	/* IS : dataSembako dan jumlahData terdefinisi.
		FS : - Menampilkan data produk yang dicari berdasarkan pilihan pengguna.
	   - Jika memilih pencarian berdasarkan kode, pencarian dilakukan
	     dengan Sequential Search.
	   - Jika memilih pencarian berdasarkan nama, pencarian dilakukan
	     dengan Binary Search.
	   - Jika produk tidak ditemukan, ditampilkan pesan bahwa produk
	     tidak ditemukan.
	*/
	var pilihanSearch, productIndex int
	var namaProduk, kodeProduk string

	fmt.Println()
	fmt.Println("+================ CARI PRODUK =================+")
	fmt.Println("1. Cari produk berdasarkan kode produk dengan Sequential Search")
	fmt.Println("2. Cari produk berdasarkan nama dengan Binary Search")
	fmt.Println("+==============================================+")
	fmt.Print("Pilih jenis pencarian: ")
	fmt.Scan(&pilihanSearch)

	if pilihanSearch == 1 {
		fmt.Print("Masukkan kode produk: ")
		fmt.Scan(&kodeProduk)

		productIndex = cariProdukSequential(dataSembako, jumlahData, kodeProduk)

		if productIndex == -1 {
			fmt.Println("Produk tidak ditemukan.")
		} else {
			fmt.Println("Kode Produk ditemukan:")
			fmt.Println("kode        :", dataSembako[productIndex].kode)
			fmt.Println("Nama         :", dataSembako[productIndex].nama)
			fmt.Println("Kategori     :", dataSembako[productIndex].kategori)
			fmt.Println("Harga        :", dataSembako[productIndex].harga)
			fmt.Println("Stok         :", dataSembako[productIndex].stok)
			fmt.Println("Barang Masuk :", dataSembako[productIndex].barangMasuk)
			fmt.Println("Barang Keluar:", dataSembako[productIndex].barangKeluar)

			fmt.Println("Produk dicari dengan Sequential.")
		}

	} else if pilihanSearch == 2 {

		fmt.Print("Masukkan nama produk: ")
		fmt.Scan(&namaProduk)

		dataSembako = binarySearchPreparationSort(dataSembako, jumlahData)
		productIndex = cariPNamabinary(dataSembako, jumlahData, namaProduk)

		if productIndex == -1 {
			fmt.Println("Produk tidak ditemukan.")
		} else {
			fmt.Println("Nama Produk ditemukan:")
			fmt.Println("kode        :", dataSembako[productIndex].kode)
			fmt.Println("Nama         :", dataSembako[productIndex].nama)
			fmt.Println("Kategori     :", dataSembako[productIndex].kategori)
			fmt.Println("Harga        :", dataSembako[productIndex].harga)
			fmt.Println("Stok         :", dataSembako[productIndex].stok)
			fmt.Println("Barang Masuk :", dataSembako[productIndex].barangMasuk)
			fmt.Println("Barang Keluar:", dataSembako[productIndex].barangKeluar)

			fmt.Println("Produk dicari berdasarkan nama dengan Binary.")
		}
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func editProduk(dataSembako *tabSembako, jumlahData int) {
	/* IS : dataSembako dan jumlahData terdefinisi dan jumlahData menyatakan banyaknya data produk yang tersimpan.

	   FS : Jika kode produk ditemukan, pengguna mengonfirmasi edit, dan kategori yang dipilih valid, maka data produk diperbarui
	     sesuai data baru yang diinputkan.
	   - Jika kode produk tidak ditemukan, ditampilkan pesan "Produk tidak ditemukan" dan data tidak berubah.
	   - Jika pengguna membatalkan edit, data tidak berubah.
	   - Jika kategori tidak valid, kategori produk tidak diubah.
	*/
	var kodeProduk string
	var jawab string
	var pilihanKategori, productIndex, stokbaru int

	fmt.Println()
	fmt.Println("+================ EDIT PRODUK =================+")
	fmt.Print("Masukkan kode produk yang ingin diedit: ")
	fmt.Scan(&kodeProduk)

	productIndex = cariProdukSequential(*dataSembako, jumlahData, kodeProduk)

	if productIndex == -1 {
		fmt.Println("Produk tidak ditemukan.")
	} else {
		fmt.Println("Apakah anda yakin ingin mengedit produk dengan kode :", kodeProduk)
		fmt.Print("Yes/No: ")
		fmt.Scan(&jawab)

		if jawab == "No" {
			fmt.Println("Edit produk dibatalkan.")
		} else if jawab == "Yes" {
			fmt.Print("Masukkan kode produk baru: ")
			fmt.Scan(&dataSembako[productIndex].kode)

			fmt.Print("Masukkan nama produk baru: ")
			fmt.Scan(&dataSembako[productIndex].nama)

			fmt.Print("Masukkan harga produk baru: ")
			fmt.Scan(&dataSembako[productIndex].harga)

			fmt.Print("Masukkan stok produk baru: ")
			fmt.Scan(&stokbaru)

			(*dataSembako)[productIndex].stok = stokbaru
			(*dataSembako)[productIndex].barangMasuk = stokbaru

			tampilKategori()
			fmt.Print("Pilih kategori baru: ")
			fmt.Scan(&pilihanKategori)

			if pilihanKategori >= 1 && pilihanKategori <= len(listKategori) {
				dataSembako[productIndex].kategori = listKategori[pilihanKategori-1]

				fmt.Println("Data produk berhasil diedit.")
				fmt.Println("Silahkan ketik 2 untuk melihat data terbaru.")
			} else {
				fmt.Println("Kategori tidak valid.")
			}
		} else {
			fmt.Println("Pilihan hanya boleh Yes atau No.")
		}
	}
}

func hapusProduk(dataSembako *tabSembako, jumlahData *int) {
	/* IS : dataSembako dan jumlahData terdefinisi dan jumlahData menyatakan banyaknya data produk yang tersimpan.
	   FS : -Jika produk ditemukan dan pengguna mengonfirmasi penghapusan, maka produk dihapus dari dataSembako dan jumlahData berkurang 1.
	   - Jika produk tidak ditemukan atau penghapusan dibatalkan, data tidak berubah.
	*/
	var kodeProduk string
	var productIndex, index int
	var jawab string

	fmt.Println()
	fmt.Println("+================ HAPUS PRODUK =================+")
	fmt.Print("Masukkan kode produk yang ingin dihapus: ")
	fmt.Scan(&kodeProduk)

	productIndex = cariProdukSequential(*dataSembako, *jumlahData, kodeProduk)

	if productIndex == -1 {
		fmt.Println("Produk tidak ditemukan.")
	} else {
		fmt.Println("Apakah anda yakin ingin menghapus produk dengan kode :", kodeProduk)
		fmt.Print("Yes/No : ")
		fmt.Scan(&jawab)

		if jawab == "No" {
			fmt.Println("Tidak jadi hapus data.")
		} else if jawab == "Yes" {

			for index = productIndex; index < *jumlahData-1; index++ {
				(*dataSembako)[index] = (*dataSembako)[index+1]
			}

			*jumlahData--

			fmt.Println("Produk berhasil dihapus.")
			fmt.Println("Silahkan ketik 2 untuk melihat data terbaru.")

		} else {
			fmt.Println("Pilihan hanya boleh Yes atau No.")
		}
	}
}

func sortNamaUntukPencarian(dataSembako *tabSembako, jumlahData int) {
	/* IS : dataSembako berisi sejumlah data produk dan jumlahData menyatakan banyaknya data yang tersimpan.
	   FS : Seluruh data produk terurut secara ascending berdasarkan field nama menggunakan metode Insertion Sort.
	*/
	var currentIndex, previousIndex int
	var tempProduct daftarSembako
	for currentIndex = 1; currentIndex < jumlahData; currentIndex++ {
		tempProduct = dataSembako[currentIndex]
		previousIndex = currentIndex - 1

		for previousIndex >= 0 && dataSembako[previousIndex].nama > tempProduct.nama {
			dataSembako[previousIndex+1] = dataSembako[previousIndex]
			previousIndex--
		}

		dataSembako[previousIndex+1] = tempProduct
	}
}

func selectionSortByStok(dataSembako *tabSembako, jumlahData int) {
	/* IS : dataSembako dan jumlahData terdefinisi, dengan jumlahData >= 0.
	   FS : Data produk pada dataSembako terurut menaik (ascending)berdasarkan jumlah stok menggunakan metode Selection Sort.
	*/
	var currentIndex, minimumIndex, nextIndex int
	var tempProduct daftarSembako
	for currentIndex = 0; currentIndex < jumlahData-1; currentIndex++ {
		minimumIndex = currentIndex

		for nextIndex = currentIndex + 1; nextIndex < jumlahData; nextIndex++ {
			if dataSembako[nextIndex].stok < dataSembako[minimumIndex].stok {
				minimumIndex = nextIndex
			}
		}

		tempProduct = dataSembako[currentIndex]
		dataSembako[currentIndex] = dataSembako[minimumIndex]
		dataSembako[minimumIndex] = tempProduct
	}
}

func insertionSortByHarga(dataSembako *tabSembako, jumlahData int) {
	/* IS : dataSembako dan jumlahData terdefinisi, dengan jumlahData >= 0.
	   FS : Data produk pada dataSembako terurut menaik (ascending)berdasarkan harga menggunakan metode Insertion Sort.
	*/
	var currentIndex, previousIndex int
	var tempProduct daftarSembako
	for currentIndex = 1; currentIndex < jumlahData; currentIndex++ {
		tempProduct = dataSembako[currentIndex]
		previousIndex = currentIndex - 1

		for previousIndex >= 0 && dataSembako[previousIndex].harga > tempProduct.harga {
			dataSembako[previousIndex+1] = dataSembako[previousIndex]
			previousIndex--
		}

		dataSembako[previousIndex+1] = tempProduct
	}
}

func urutkanProduk(dataSembako *tabSembako, jumlahData int) {
	/* IS : dataSembako dan jumlahData terdefinisi.
	   FS :
	   - Data produk terurut sesuai pilihan pengguna:
	     1. berdasarkan nama produk,
	     2. berdasarkan stok produk,
	     3. berdasarkan harga produk.
	   - Jika pilihan tidak valid, data tidak berubah.
	*/
	var pilihanSort int

	fmt.Println()
	fmt.Println("+================ URUTKAN PRODUK =================+")
	fmt.Println("1. Urutkan berdasarkan nama produk")
	fmt.Println("2. Urutkan berdasarkan stok produk")
	fmt.Println("3. Urutkan berdasarkan harga produk")
	fmt.Print("Pilih jenis pengurutan: ")
	fmt.Scan(&pilihanSort)

	if pilihanSort == 1 {
		sortNamaUntukPencarian(dataSembako, jumlahData)
		fmt.Println("Produk berhasil diurutkan berdasarkan nama")
		tampilProduk(*dataSembako, jumlahData)

	} else if pilihanSort == 2 {
		selectionSortByStok(dataSembako, jumlahData)
		fmt.Println("Produk berhasil diurutkan berdasarkan stok")
		tampilProduk(*dataSembako, jumlahData)

	} else if pilihanSort == 3 {
		insertionSortByHarga(dataSembako, jumlahData)
		fmt.Println("Produk berhasil diurutkan berdasarkan harga")
		tampilProduk(*dataSembako, jumlahData)

	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func laporanBarang(dataSembako tabSembako, jumlahData int) {
	/* IS : dataSembako dan jumlahData terdefinisi.
	   FS :
	   - Menampilkan total seluruh barang masuk.
	   - Menampilkan total seluruh barang keluar.
	   - Menampilkan seluruh data produk yang tersimpan.
	*/
	var totalBarangMasuk int
	var totalBarangKeluar int
	var productIndex int

	fmt.Println()
	fmt.Println("+================ LAPORAN BARANG MASUK DAN KELUAR =================+")

	for productIndex = 0; productIndex < jumlahData; productIndex++ {
		totalBarangMasuk += dataSembako[productIndex].barangMasuk
		totalBarangKeluar += dataSembako[productIndex].barangKeluar
	}

	fmt.Println("Total semua barang masuk :", totalBarangMasuk)
	fmt.Println("Total semua barang keluar:", totalBarangKeluar)
	fmt.Println()

	tampilProduk(dataSembako, jumlahData)
}

func cariStokEkstrem(dataSembako tabSembako, jumlahData int) {
	/* IS : dataSembako berisi sejumlah data produk.
	   FS :
	   - Menampilkan data produk yang memiliki stok maksimum
	     dan stok minimum.
	   - Jika tidak ada data produk, ditampilkan pesan yang sesuai.
	*/
	var maximumIndex, minimumIndex, productIndex int

	if jumlahData == 0 {
		fmt.Println("Belum ada data produk.")
	} else {
		maximumIndex = 0
		minimumIndex = 0

		for productIndex = 1; productIndex < jumlahData; productIndex++ {
			if dataSembako[productIndex].stok > dataSembako[maximumIndex].stok {
				maximumIndex = productIndex
			}

			if dataSembako[productIndex].stok < dataSembako[minimumIndex].stok {
				minimumIndex = productIndex
			}
		}

		fmt.Println()
		fmt.Println("+================ STOK TERBANYAK DAN TERSEDIKIT =================+")
		fmt.Println("Produk dengan stok terbanyak:")
		fmt.Println("Nama :", dataSembako[maximumIndex].nama)
		fmt.Println("Stok :", dataSembako[maximumIndex].stok)

		fmt.Println()
		fmt.Println("Produk dengan stok tersedikit:")
		fmt.Println("Nama :", dataSembako[minimumIndex].nama)
		fmt.Println("Stok :", dataSembako[minimumIndex].stok)
	}
}

func main() {
	/*Program utama yang menginisialisasi data awal, menampilkan menu utama, membaca pilihan pengguna,
	dan memanggil fungsi atau prosedur yang sesuai hingga pengguna memilih keluar dari program.*/
	var dataSembako tabSembako
	var jumlahData int
	var pilihanMenu int
	var selesai bool

	dataAwal(&dataSembako, &jumlahData)
	menuAwal()

	selesai = false

	for !selesai {
		menuUtama()
		fmt.Scan(&pilihanMenu)

		switch pilihanMenu {
		case 1:
			tambahProduk(&dataSembako, &jumlahData)
		case 2:
			tampilProduk(dataSembako, jumlahData)
		case 3:
			barangMasuk(&dataSembako, jumlahData)
		case 4:
			barangKeluar(&dataSembako, jumlahData)
		case 5:
			cariProduk(dataSembako, jumlahData)
		case 6:
			editProduk(&dataSembako, jumlahData)
		case 7:
			hapusProduk(&dataSembako, &jumlahData)
		case 8:
			urutkanProduk(&dataSembako, jumlahData)
		case 9:
			laporanBarang(dataSembako, jumlahData)
		case 10:
			cariStokEkstrem(dataSembako, jumlahData)
		case 0:
			fmt.Println("Terima kasih sudah menggunakan aplikasi.")
			selesai = true
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

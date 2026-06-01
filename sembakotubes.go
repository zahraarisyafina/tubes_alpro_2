package main 
import "fmt"


const NMAX int = 1500 //konstanta jumlah maksimal array
type daftarSembako struct{ //struct tipe alias untuk menyimpan data sembako  
	nama string
	nomor string
	harga float64
	stok int
	kategori string
}

type tabSembako [NMAX]daftarSembako //tipe data array untuk menyimpan daftar produk sembako 


var listKategori = [9]string{ //daftar kategori sembako 
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

func menuAwal(jenis string){ //untuk menampilkan menu awal pada program 
	fmt.Println()
	fmt.Println("+==========================================================+")
	fmt.Println("|                                                          |")
	fmt.Println("|              SELAMAT DATANG DI APLIKASI                  |")
	fmt.Println("|                 MANAJEMEN TOKO SEMBAKO                   |")
	fmt.Println("|                                                          |")
	fmt.Println("+==========================================================+")
	fmt.Println()
}

func menu(tipe string){ // Fungsi untuk menampilkan menu untuk pilih opsi yang tersedia
	fmt.Println("+==========================================================+")
	fmt.Println("|                    DAFTAR MENU                           |")
	fmt.Println("+==========================================================+")
	switch tipe{
	case "menuAwal":
		fmt.Println("| 1. Fitur Pengaturan Sembako                              |")
		//fmt.Printf("\n| %-40s |", "[2] Rekomendasi AI")
		//fmt.Printf("\n| %-40s |", "[0] Exit")
		fmt.Println("+==========================================================+")
		fmt.Print("Pilih (1/2/0)?")
		fmt.Println()
	case "Fitur Pengaturan Sembako":
		fmt.Println("| 1. Daftar Pengelolaan Sembako                            |")
		fmt.Println("| 2. Edit Data Produk                                      |")
		fmt.Println("| 3. Cari Produk                                           |")
		fmt.Println("| 4. Mengurutkan Produk                                    |")
		fmt.Println("| 5. Hapus Produk                                          |")
		fmt.Println("| 6. Daftar Pengelolaan Sembako                            |")
		fmt.Println("| 7. Home                                                  |")
		fmt.Println("+==========================================================+")
		fmt.Print("Pilih (1/2/3/4/5/6)?")
		fmt.Println()
	}
	case "Fitur Pengaturan Sembako":
		fmt.Println("| 1. Daftar Barang Masuk                                   |")
		fmt.Println("| 2. Daftar Barang Keluar                                  |")
		fmt.Println("| 3. Mengurutkan Produk                                    |")
		fmt.Println("| 4.                                         |")
		fmt.Println("| 5.                            |")
		fmt.Println("| 6. Home                                                  |")
		fmt.Println("+==========================================================+")
		fmt.Print("Pilih (1/2/3/4/5/6)?")
		fmt.Println()
	}

}

//func eksukusi()



func main() {
	var jenis string
	var pilih int

	menuAwal(jenis)

	menu("menuAwal")
	fmt.Scan(&pilih)

	if pilih == 1 {
		menu("Fitur Pengaturan Sembako")
		fmt.Scan(&pilih) // untuk memilih submenu nanti
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}
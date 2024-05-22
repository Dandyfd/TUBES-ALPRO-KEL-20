package main

import "fmt"

type Makanan struct {
	namaMakanan  string
	hargaMakanan int
}

type Minuman struct {
	namaMinuman  string
	hargaMinuman int
}

const MaxMeja int = 20
const NMAX int = 50

type MaxMejaPelanggan [MaxMeja]int
type ArrMakanan [NMAX]Makanan
type ArrMinuman [NMAX]Minuman

var jumlahMakanan = 7
var jumlahMinuman = 7

var daftarMakanan = ArrMakanan{
	{"Ayam geprek", 20000},
	{"Chicken Katsu", 23000},
	{"Kentang Goreng", 15000},
	{"Pisang bakar", 15000},
	{"Cireng bumbu rujak", 10000},
	{"Nasi Goreng", 15000},
	{"Steak", 30000},
}

var daftarMinuman = ArrMinuman{
	{"Americano", 14000},
	{"Latte", 15000},
	{"Espresso", 17000},
	{"Signature Chocolate", 14000},
	{"Matcha Latte", 13000},
	{"Es teh manis", 10000},
	{"Teh Tarik", 12000},
}

func main() {
	var login string

	fmt.Println("|| ------------------ ||")
	fmt.Println("||   Selamat Datang   ||")
	fmt.Println("|| ------------------ ||")
	fmt.Println("||       Login        ||")
	fmt.Println("||                    ||")
	fmt.Println("||       Admin        ||")
	fmt.Println("||                    ||")
	fmt.Println("||     Pelanggan      ||")
	fmt.Println("|| ------------------ ||")

	fmt.Print("Masuk: ")
	fmt.Scan(&login)
	if login == "Admin" {
		admin()
	} else if login == "Pelanggan" {
		pelanggan()
	} else {
		fmt.Println("Login tidak valid")
	}
}

func admin() {
	var action, login string

	fmt.Println("|| ------------------------ ||")
	fmt.Println("||   1. Tambah Item         ||")
	fmt.Println("||   2. Hapus Item          ||")
	fmt.Println("||   3. Cari Menu Makanan   ||")
	fmt.Println("||   4. Cari Menu Minuman   ||")
	fmt.Println("||   0. Exit                ||")
	fmt.Println("|| ------------------------ ||")
	fmt.Scan(&action)
	for action != "0" {
		fmt.Println("|| ------------------------ ||")
		fmt.Println("||   1. Tambah Item         ||")
		fmt.Println("||   2. Hapus Item          ||")
		fmt.Println("||   3. Cari Menu Makanan   ||")
		fmt.Println("||   4. Cari Menu Minuman   ||")
		fmt.Println("||   0. Exit                ||")
		fmt.Println("|| ------------------------ ||")
		if action == "1" {
			tambahItem()
		} else if action == "2" {
			hapusItem()
		} else if action == "3" {
			cariMenuMakanan()
		} else if action == "4" {
			cariMenuMinuman()
		}
		fmt.Scan(&action)
	}
	fmt.Println("|| ------------------ ||")
	fmt.Println("||   Selamat Datang   ||")
	fmt.Println("|| ------------------ ||")
	fmt.Println("||       Login        ||")
	fmt.Println("||                    ||")
	fmt.Println("||       Admin        ||")
	fmt.Println("||                    ||")
	fmt.Println("||     Pelanggan      ||")
	fmt.Println("|| ------------------ ||")

	fmt.Print("Masuk: ")
	fmt.Scan(&login)
	if login == "Admin" {
		admin()
	} else if login == "Pelanggan" {
		pelanggan()
	} else {
		fmt.Println("Login tidak valid")
	}
}

func tambahItem() {
	var jenis, nama string
	var harga int

	fmt.Print("Tambah item ke (makanan/minuman): ")
	fmt.Scan(&jenis)
	fmt.Print("Nama item: ")
	fmt.Scan(&nama)
	fmt.Print("Harga item: ")
	fmt.Scan(&harga)

	if jenis == "makanan" {
		if jumlahMakanan < NMAX {
			daftarMakanan[jumlahMakanan] = Makanan{nama, harga}
			jumlahMakanan++
			fmt.Println("Item makanan berhasil ditambahkan")
		} else {
			fmt.Println("Kapasitas daftar makanan penuh")
		}
	} else if jenis == "minuman" {
		if jumlahMinuman < NMAX {
			daftarMinuman[jumlahMinuman] = Minuman{nama, harga}
			jumlahMinuman++
			fmt.Println("Item minuman berhasil ditambahkan")
		} else {
			fmt.Println("Kapasitas daftar minuman penuh")
		}
	} else {
		fmt.Println("Jenis item tidak valid")
	}
}

func hapusItem() {
	var jenis, nama string

	fmt.Print("Hapus item dari (makanan/minuman): ")
	fmt.Scan(&jenis)
	fmt.Print("Nama item: ")
	fmt.Scan(&nama)

	if jenis == "makanan" {
		for i := 0; i < jumlahMakanan; i++ {
			if daftarMakanan[i].namaMakanan == nama {
				daftarMakanan[i] = daftarMakanan[jumlahMakanan-1]
				jumlahMakanan--
				fmt.Println("Item makanan berhasil dihapus")
				return
			}
		}
		fmt.Println("Item makanan tidak ditemukan")
	} else if jenis == "minuman" {
		for i := 0; i < jumlahMinuman; i++ {
			if daftarMinuman[i].namaMinuman == nama {
				daftarMinuman[i] = daftarMinuman[jumlahMinuman-1]
				jumlahMinuman--
				fmt.Println("Item minuman berhasil dihapus")
				return
			}
		}
		fmt.Println("Item minuman tidak ditemukan")
	} else {
		fmt.Println("Jenis item tidak valid")
	}
}

func cariMenuMakanan() {
	var nama string

	fmt.Print("Masukkan nama menu yang dicari: ")
	fmt.Scan(&nama)

	for i := 0; i < jumlahMakanan-1; i++ {
		for j := 0; j < jumlahMakanan-i-1; j++ {
			if daftarMakanan[j].hargaMakanan > daftarMakanan[j+1].hargaMakanan {
				daftarMakanan[j], daftarMakanan[j+1] = daftarMakanan[j+1], daftarMakanan[j]
			}
		}
	}
	for i := 0; i < jumlahMakanan; i++ {
		if daftarMakanan[i].namaMakanan == nama {
			fmt.Printf("Makanan ditemukan: %s - Rp%d\n", daftarMakanan[i].namaMakanan, daftarMakanan[i].hargaMakanan)
			return
		}
	}
	fmt.Println("Menu tidak ditemukan")
}

func cariMenuMinuman() {
	var nama string

	fmt.Print("Masukkan nama menu yang dicari: ")
	fmt.Scan(&nama)

	for i := 0; i < jumlahMinuman; i++ {
		if daftarMinuman[i].namaMinuman == nama {
			fmt.Printf("Minuman ditemukan: %s - Rp%d\n", daftarMinuman[i].namaMinuman, daftarMinuman[i].hargaMinuman)
			return
		}
	}
	fmt.Println("Menu tidak ditemukan")
}

func pelanggan() {
	var meja int

	fmt.Print("Masukkan nomor meja (1-20): ")
	fmt.Scan(&meja)
	if meja < 1 || meja > 20 {
		fmt.Println("Nomor meja tidak valid")
		pelanggan()
	}
	menupelanggan()
}

func menupelanggan() {
	var act int
	var pesanan []string
	var totalHarga int

	for {
		fmt.Println("|| ------------------------------ ||")
		fmt.Println("||            8'Caffe             ||")
		fmt.Println("|| ------------------------------ ||")
		fmt.Println("||   1. Makanan      2. Minuman   ||")
		fmt.Println("||   3. Cari menu makanan         ||")
		fmt.Println("||   4. Cari menu minuman         ||")
		fmt.Println("||   5. Lihat pesanan             ||")
		fmt.Println("||   6. Bayar                     ||")
		fmt.Println("|| ------------------------------ ||")

		fmt.Print("Pilih menu (1-6): ")
		fmt.Scan(&act)
		if act == 1 {
			totalHarga += pesanMakanan(&pesanan)
		} else if act == 2 {
			totalHarga += pesanMinuman(&pesanan)
		} else if act == 3 {
			cariMenuMakanan()
		} else if act == 4 {
			cariMenuMinuman()
		} else if act == 5 {
			lihatPesanan(pesanan, totalHarga)
		} else if act == 6 {
			selesaiDanBayar(totalHarga)
			break
		} else {
			fmt.Println("Aksi tidak valid")
		}
	}
}

func pesanMakanan(pesanan *[]string) int {
	var pilihan, jumlah int
	var totalHarga int
	var pilih string

	tampilkanDaftarMakanan()
	fmt.Print("Urutkan menu? (ya/tidak): ")
	fmt.Scan(&pilih)
	if pilih == "ya" {
		urutkanHargaMakanan()
	}
	fmt.Print("Masukkan nomor makanan yang ingin dipesan: ")
	fmt.Scan(&pilihan)
	fmt.Print("Masukkan jumlah: ")
	fmt.Scan(&jumlah)
	fmt.Println("Item ditambahkan")

	if pilihan < 1 || pilihan > jumlahMakanan {
		fmt.Println("Pilihan tidak valid")
		return 0
	}

	*pesanan = append(*pesanan, fmt.Sprintf("%s x%d", daftarMakanan[pilihan-1].namaMakanan, jumlah))
	totalHarga = daftarMakanan[pilihan-1].hargaMakanan * jumlah
	return totalHarga
}

func pesanMinuman(pesanan *[]string) int {
	var pilihan, jumlah int
	var totalHarga int
	var pilih string

	tampilkanDaftarMinuman()
	fmt.Print("Urutkan menu? (ya/tidak): ")
	fmt.Scan(&pilih)
	if pilih == "ya" {
		urutkanHargaMinuman()
	}
	fmt.Print("Masukkan nomor minuman yang ingin dipesan: ")
	fmt.Scan(&pilihan)
	fmt.Print("Masukkan jumlah: ")
	fmt.Scan(&jumlah)
	fmt.Println("Item ditambahkan")

	if pilihan < 1 || pilihan > jumlahMinuman {
		fmt.Println("Pilihan tidak valid")
		return 0
	}

	*pesanan = append(*pesanan, fmt.Sprintf("%s x%d", daftarMinuman[pilihan-1].namaMinuman, jumlah))
	totalHarga = daftarMinuman[pilihan-1].hargaMinuman * jumlah
	return totalHarga
}

func tampilkanDaftarMakanan() {
	fmt.Println("---------- Daftar Makanan ----------")
	for i := 0; i < jumlahMakanan; i++ {
		fmt.Printf(" %d. %s - Rp%d\n", i+1, daftarMakanan[i].namaMakanan, daftarMakanan[i].hargaMakanan)
	}
	fmt.Println("------------------------------------")
}

func tampilkanDaftarMinuman() {
	fmt.Println("---------- Daftar Minuman ----------")
	for i := 0; i < jumlahMinuman; i++ {
		fmt.Printf(" %d. %s - Rp%d\n", i+1, daftarMinuman[i].namaMinuman, daftarMinuman[i].hargaMinuman)
	}
	fmt.Println("------------------------------------")
}

func urutkanHargaMakanan() {
	for i := 0; i < jumlahMakanan-1; i++ {
		for j := 0; j < jumlahMakanan-i-1; j++ {
			if daftarMakanan[j].hargaMakanan > daftarMakanan[j+1].hargaMakanan {
				daftarMakanan[j], daftarMakanan[j+1] = daftarMakanan[j+1], daftarMakanan[j]
			}
		}
	}
	fmt.Println("------ Daftar Makanan Terurut ------")
	for i := 0; i < jumlahMakanan; i++ {
		fmt.Printf("%d. %s - Rp%d\n", i+1, daftarMakanan[i].namaMakanan, daftarMakanan[i].hargaMakanan)
	}
	fmt.Println("------------------------------------")
}

func urutkanHargaMinuman() {
	// Mengurutkan daftarMinuman menggunakan Insertion Sort
	for i := 1; i < jumlahMinuman; i++ {
		key := daftarMinuman[i]
		j := i - 1

		// Pindahkan elemen-elemen daftarMinuman[0..i-1], yang lebih besar dari key,
		// ke satu posisi di depan posisi saat ini
		for j >= 0 && daftarMinuman[j].hargaMinuman > key.hargaMinuman {
			daftarMinuman[j+1] = daftarMinuman[j]
			j = j - 1
		}
		daftarMinuman[j+1] = key
	}

	fmt.Println("------ Daftar Minuman Terurut ------")
	for i := 0; i < jumlahMinuman; i++ {
		fmt.Printf("%d. %s - Rp%d\n", i+1, daftarMinuman[i].namaMinuman, daftarMinuman[i].hargaMinuman)
	}
	fmt.Println("------------------------------------")
}
func lihatPesanan(pesanan []string, totalHarga int) {
	fmt.Println("---------- Pesanan Anda ----------")
	for _, p := range pesanan {
		fmt.Println(p)
	}
	fmt.Printf("Total harga: Rp%d\n", totalHarga)
}

func selesaiDanBayar(totalHarga int) {
	var bayar int

	fmt.Printf("Total harga: Rp.%d\n", totalHarga)
	fmt.Print("Masukkan jumlah uang yang dibayar: Rp.")
	fmt.Scan(&bayar)

	if bayar >= totalHarga {
		fmt.Printf("Kembalian: Rp%d\n", bayar-totalHarga)
	} else {
		fmt.Printf("Uang yang dibayar kurang Rp%d\n", totalHarga-bayar)
	}
	fmt.Println("|| ------------------- ||")
	fmt.Println("||     Terimakasih     ||")
	fmt.Println("|| ------------------- ||")
}

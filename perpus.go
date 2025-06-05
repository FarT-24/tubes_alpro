package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Buku struct { // Struct untuk menyimpan data buku
	ID      string
	Judul   string
	Penulis string
	Tahun   int
}

type Anggota struct { // Struct untuk menyimpan data anggota
	ID     string
	Nama   string
	Alamat string
}

type Peminjaman struct { // Struct untuk menyimpan data peminjaman
	IDBuku         string
	IDAnggota      string
	TanggalPinjam  string
	TanggalKembali string
	TanggalKembaliAktual string
	Denda          int
	StatusKembali string
}

var bukuList []Buku // Slice menyimpan daftar buku
var anggotaList []Anggota // Slice menyimpan daftar anggota
var peminjamanList []Peminjaman // Slice menyimpan daftar peminjaman

var reader = bufio.NewReader(os.Stdin)

func input(teks string) string { // Input fungsi untuk membaca input dari pengguna
	fmt.Print(teks)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func tambahBuku() { // Fungsi menambahkan buku baru
	id := input("ID Buku: ")
	judul := input("Judul: ")
	penulis := input("Penulis: ")
	var tahun int
	fmt.Print("Tahun Terbit: ")
	fmt.Scan(&tahun)
	reader.ReadString('\n')
	bukuList = append(bukuList, Buku{id, judul, penulis, tahun})
	fmt.Println("Buku berhasil ditambahkan!")
}

func tambahAnggota() { // Fungsi menambahkan anggota baru
	id := input("ID Anggota: ")
	nama := input("Nama: ")
	alamat := input("Alamat: ")
	anggotaList = append(anggotaList, Anggota{id, nama, alamat})
	fmt.Println("Anggota berhasil ditambahkan!")
}

func pinjamBuku() { // Fungsi mencatat peminjaman buku
	idBuku := input("ID Buku: ")
	idAnggota := input("ID Anggota: ")
	tglPinjam := input("Tanggal Pinjam (DD-MM-YYYY): ")
	tglKembali := input("Tanggal Kembali (DD-MM-YYYY): ")
	peminjamanList = append(peminjamanList, Peminjaman{
		IDBuku: idBuku,
		IDAnggota: idAnggota,
		TanggalPinjam: tglPinjam,
		TanggalKembali: tglKembali,
		TanggalKembaliAktual: "",
		Denda: 0, // Denda dihitung di func kembalikanBuku
		StatusKembali: "Belum Kembali",
	})
	fmt.Println("Peminjaman berhasil dicatat!")
}

func kembalikanBuku() { // Fungsi mencatat pengembalian buku
	idBuku := input("ID Buku: ")
	idAnggota := input("ID Anggota: ")
	tglKembali := input("Tanggal Kembali (DD-MM-YYYY): ")
	format := "02-01-2006"
	aktual, err := time.Parse(format, tglKembali)
	if err != nil {
		fmt.Println("Format tanggal tidak valid.")
		return
	}
	for i, p := range peminjamanList {
		if p.IDBuku == idBuku && p.IDAnggota == idAnggota && p.StatusKembali == "Belum Kembali" {
			rencana, err := time.Parse(format, p.TanggalKembali)
			if err != nil {
				fmt.Println("Data tanggal rencana kembali di sistem tidak valid.")
				return
			}
			terlambat := int(aktual.Sub(rencana).Hours() / 24)
			if terlambat < 0 {
				terlambat = 0
			}
			denda := terlambat * 1000
			peminjamanList[i].Denda = denda
			peminjamanList[i].TanggalKembaliAktual = tglKembali
			peminjamanList[i].StatusKembali = "Sudah"

			fmt.Printf("Pengembalian dicatat! Keterlambatan: %d hari, Denda: Rp.%d\n", terlambat, denda)
			return
		}
	}
	fmt.Println("Data peminjaman tidak ditemukan.")
}

func cariBukuByJudul(judul string) { // Pakai linear search
	found := false
	keyword := strings.ToLower(judul)
	for _, b := range bukuList {
		if strings.Contains(strings.ToLower(b.Judul), keyword) {
			fmt.Printf("ID: %s, Judul: %s, Penulis: %s, Tahun: %d\n", b.ID, b.Judul, b.Penulis, b.Tahun)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan kata tersebut di judulnya.")
	}
}
	
func cariBukuByID(id string) { // Pakai binary search
	left, right := 0, len(bukuList)-1
	for left <= right {
		mid := (left + right) / 2
		if bukuList[mid].ID == id {
			b := bukuList[mid]
			fmt.Printf("ID: %s, Judul: %s, Penulis: %s, Tahun: %d\n", b.ID, b.Judul, b.Penulis, b.Tahun)
			return
			} else if bukuList[mid].ID < id {
				left = mid + 1
			} else {
				right = mid - 1
			}
	}
	fmt.Println("Buku tidak ditemukan.")
}

func urutkanBukuByJudul() { // Pakai bubble sort
	n := len(bukuList)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if strings.ToLower(bukuList[j].Judul) > strings.ToLower(bukuList[j+1].Judul) {
				bukuList[j], bukuList[j+1] = bukuList[j+1], bukuList[j]
			}
		}
	}
	fmt.Println("Buku diurutkan berdasarkan Judul.")
}

func urutkanBukuByTahun() { // Pakai insertion sort
	for i := 1; i < len(bukuList); i++ {
		key := bukuList[i]
		j := i - 1
		// Geser elemen yang lebih besar ke kanan
		for j >= 0 && bukuList[j].Tahun > key.Tahun {
			bukuList[j+1] = bukuList[j]
			j--
		}
		bukuList[j+1] = key
	}
	fmt.Println("Buku diurutkan berdasarkan Tahun.")
}

func tampilkanBuku() { // Tampilkan semua buku
	for _, b := range bukuList {
		fmt.Printf("ID: %s, Judul: %s, Penulis: %s, Tahun: %d\n", b.ID, b.Judul, b.Penulis, b.Tahun)
	}
}

func tampilkanAnggota() { // Tampilkan semua anggota
	for _, a := range anggotaList {
		fmt.Printf("ID: %s, Nama: %s, Alamat: %s\n", a.ID, a.Nama, a.Alamat)
	}
}

func tampilkanPeminjaman() { // Tampilkan semua transaksi peminjaman
	format := "02-01-2006"
	for _, p := range peminjamanList {
		terlambat := 0
		tglAktual := p.TanggalKembaliAktual

		if p.StatusKembali == "Sudah" {
			rencana, err1 := time.Parse(format, p.TanggalKembali)
			aktual, err2 := time.Parse(format, p.TanggalKembaliAktual)
			if err1 == nil && err2 == nil {
				terlambat = int(aktual.Sub(rencana).Hours() / 24)
				if terlambat < 0 {
					terlambat = 0
				}
			}
		} else {
			tglAktual = "-"
		}
		fmt.Printf("ID Buku: %s, ID Anggota: %s, Pinjam: %s, Rencana Kembali: %s, 
			Aktual: %s, Keterlambatan: %d hari, Denda: Rp%d, Status: %s\n",
			p.IDBuku, p.IDAnggota, p.TanggalPinjam, p.TanggalKembali, tglAktual, 
			terlambat, p.Denda, p.StatusKembali)
	}
}

func main() { // Fungsi utama untuk menjalankan program
	for {
		fmt.Println("\n=== SISTEM INFORMASI PERPUSTAKAAN ===")
		fmt.Println("1. Tambah Buku")
		fmt.Println("2. Tambah Anggota")
		fmt.Println("3. Peminjaman Buku")
		fmt.Println("4. Kembalikan Buku")
		fmt.Println("5. Cari Buku (Judul)")
		fmt.Println("6. Cari Buku (ID)")
		fmt.Println("7. Urutkan Buku (Judul)")
		fmt.Println("8. Urutkan Buku (Tahun)")
		fmt.Println("9. Tampilkan Semua Buku")
		fmt.Println("10. Tampilkan Semua Anggota")
		fmt.Println("11. Tampilkan Semua Transaksi")
		fmt.Println("12. Keluar")
		fmt.Print("\nPilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)
		reader.ReadString('\n')

		switch pilihan {
		case 1:
			tambahBuku()
		case 2:
			tambahAnggota()
		case 3:
			pinjamBuku()
		case 4:
			kembalikanBuku()
		case 5:
			judul := input("Masukkan judul buku: ")
			cariBukuByJudul(judul)
		case 6:
			id := input("Masukkan ID buku: ")
			cariBukuByID(id)
		case 7:
			urutkanBukuByJudul()
		case 8:
			urutkanBukuByTahun()
		case 9:
			tampilkanBuku()
		case 10:
			tampilkanAnggota()
		case 11:
			tampilkanPeminjaman()
		case 12:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
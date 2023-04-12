package dipa

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoConn = atdb. MongoConnect(MongoInfo)


func TestInsertlisttamu(t *testing.T) {
	name := "GABYAZANA"
	notelp := "(+62)85231151994)"
	email := "gaby@gmail.com"
	kota := "BOGOR"
	hasil := Insertlisttamu(name, notelp, email, kota)
	fmt.Println(hasil)
}

func TestInsertUndanganRapat(t *testing.T) {
	lokasi := "Auditorium"
	agenda := "Rapat Umum"
	peserta := "seluruh anggota yayasan"
	hasil := InsertUndanganRapat(lokasi, agenda, peserta)
	fmt.Println(hasil)
}

func TestInsertpesertarapat(t *testing.T) {
	nama := "Gaby"
	instansi := "ULBI"
	status := "Aktif"
	hasil := Insertpesertarapat(nama, instansi, status)
	fmt.Println(hasil)
}

func TestInsertwakturapat(t *testing.T) {
	hal := "rapat"
	materi := "Generasi Muda"
	hasil := Insertwakturapat(hal, materi)
	fmt.Println(hasil)
}

func TestInsertrapatmulai(t *testing.T) {
	pembicara := "Jokowi"
	durasi := "satu jam"
	hasil := Insertrapatmulai(pembicara, durasi)
	fmt.Println(hasil)
}
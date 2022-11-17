package main

import (
	"fmt"
)

type PropertiMajalahStruct struct {
	nama_majalah string
	brand        string
	edisi        string
	halaman      string
}

type PropertiMajalahInterface interface {
	get_nama_majalah() string
	set_nama_majalah(nama string)

	get_brand() string
	set_brand(brand string)

	get_edisi() string
	set_edisi(edisi string)

	get_halaman() string
	set_halaman(halaman string)
}

type AbstractMajalah struct {
	Penulis    string
	Percetakan string
}

type AbstractMajalahImpl struct {
	AbstractMajalah
}

func (am *AbstractMajalahImpl) get_penulis() string {
	return am.Penulis
}

func (am *AbstractMajalah) set_penulis(penulis string) {
	am.Penulis = penulis
}

func (am *AbstractMajalahImpl) get_percetakan() string {
	return am.Percetakan
}

func (am *AbstractMajalah) set_percetakan(percetakan string) {
	am.Percetakan = percetakan
}

func (pm *PropertiMajalahStruct) get_brand() string {
	return pm.brand
}

func (pm *PropertiMajalahStruct) get_edisi() string {
	return pm.edisi
}

func (pm *PropertiMajalahStruct) get_halaman() string {
	return pm.halaman
}

func (pm *PropertiMajalahStruct) get_nama_majalah() string {
	return pm.nama_majalah
}

func (pm *PropertiMajalahStruct) set_brand(brand string) {
	pm.brand = brand
}

func (pm *PropertiMajalahStruct) set_edisi(edisi string) {
	pm.edisi = edisi
}

func (pm *PropertiMajalahStruct) set_halaman(halaman string) {
	pm.halaman = halaman
}

func (pm *PropertiMajalahStruct) set_nama_majalah(nama string) {
	pm.nama_majalah = nama
}

func NewClassMajalah(pm *PropertiMajalahStruct) PropertiMajalahInterface {
	return &PropertiMajalahStruct{
		nama_majalah: pm.nama_majalah,
		brand:        pm.brand,
		edisi:        pm.edisi,
		halaman:      pm.halaman,
	}
}

func main() {

	am := AbstractMajalahImpl{
		AbstractMajalah: AbstractMajalah{},
	}

	am.set_penulis("Abdul Kadir")

	fmt.Println("Penulis", am.get_penulis())

	pm := PropertiMajalahStruct{
		nama_majalah: "",
		brand:        "",
		edisi:        "",
		halaman:      "",
	}

	majalah := NewClassMajalah(&pm)

	majalah.set_nama_majalah("Bobo Kids")
	majalah.set_brand("Bobo")
	majalah.set_edisi("22 2022")
	majalah.set_halaman("17")

	fmt.Println("Nama majalah", majalah.get_nama_majalah())
	fmt.Println("Brand majalah", majalah.get_brand())
	fmt.Println("Edisi majalah", majalah.get_edisi())
	fmt.Println("Halaman majalah", majalah.get_halaman())

	pm = PropertiMajalahStruct{
		nama_majalah: "Bobo Junior",
		brand:        "Bobo",
		edisi:        "01 2022",
		halaman:      "18",
	}

	majalah = NewClassMajalah(&pm)

	fmt.Println("Nama majalah", majalah.get_nama_majalah())
	fmt.Println("Brand majalah", majalah.get_brand())
	fmt.Println("Edisi majalah", majalah.get_edisi())
	fmt.Println("Halaman majalah", majalah.get_halaman())
}

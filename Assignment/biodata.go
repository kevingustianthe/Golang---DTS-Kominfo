package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {

	var siswa = []Biodata{
		{nama: "Kevin Gustian The", alamat: "Jl. Pemuda", pekerjaan: "Mahasiswa", alasan: "Bahasa GO lagi populer"},
		{nama: "M. Iqbal Saputra", alamat: "Jl. Sudirman", pekerjaan: "Mahasiswa", alasan: "Tertarik belajar Backend"},
		{nama: "Ibra Hasan Suraya", alamat: "Jl. Fatmawati", pekerjaan: "Mahasiswa", alasan: "Menambah bahasa baru"},
		{nama: "Harianto Billy", alamat: "Jl. Merdeka", pekerjaan: "Mahasiwa", alasan: "Go memiliki sintaksis yang sederhana"},
		{nama: "Rafi Musthafa", alamat: "Jl. Melati", pekerjaan: "Mahasiswa", alasan: "Memiliki komunitas yang berkembang"},
	}

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: go run biodata.go <number>")
		return
	}

	number, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid argument:", args[0])
		return
	}

	if number < 1 || number > len(siswa) {
		fmt.Println("Invalid number. Please enter a number between 1 and", len(siswa))
		return
	}

	displayData(siswa[number-1])
}

func displayData(data Biodata) {
	fmt.Println("Nama\t\t\t\t: ", data.nama)
	fmt.Println("Alamat\t\t\t\t: ", data.alamat)
	fmt.Println("Pekerjaan\t\t\t: ", data.pekerjaan)
	fmt.Println("Alasan pilih kelas Golang\t: ", data.alasan)
}
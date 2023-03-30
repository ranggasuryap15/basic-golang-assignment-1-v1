package main

import (
	"fmt"
	"strings"

	"a21hc3NpZ25tZW50/helper"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

func Login(id string, name string) string {
	var statusLogin string
	var students[] string

	if len(id) < 5 {
		statusLogin = "ID must be 5 characters long!"
	}

	if id == "" || name == "" {
		statusLogin = "ID, Name or Major is undefined!"
	}

	students = strings.Split(Students, ", ")

	for i := 0; i < len(students); i++ {
		if strings.Split(students[i], "_")[0] != id {
			statusLogin = "Login gagal: data mahasiswa tidak ditemukan"
		} else {
			statusLogin = "Login berhasil: " + name + " (" + strings.Split(students[i], "_")[2] + ")"
			break
		}
	}
	return statusLogin // TODO: replace this
}



func Register(id string, name string, major string) string {
	var statusRegister string
	var students[] string
	// var studentDetails[] string
	// var studentID[] string

	if len(id) < 5 {
		statusRegister = "ID must be 5 characters long!"
	}

	if id == "" || name == "" || major == "" {
		statusRegister = "ID, Name or Major is undefined!"
	}

	students = strings.Split(Students, ", ")

	for i := 0; i < len(students); i++ {
		if strings.Split(students[i], "_")[0] == id {
			statusRegister = "Register gagal: id sudah digunakan"
			break
		} else {
			statusRegister = "Register berhasil: " + name + " (" + major + ")"
		}
	}
	
	return statusRegister // TODO: SELESAI REGISTER
}

func GetStudyProgram(code string) string {
	return "" // TODO: replace this
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		fmt.Println("Students: ", Students)
		fmt.Println("Student Study Programs: ", StudentStudyPrograms)

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)

			fmt.Println(Login(id, name))

			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scan(&jurusan)
			fmt.Println(Register(id, name, jurusan))

			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scan(&kode)

			fmt.Println(GetStudyProgram(kode))
			helper.Delay(5)
		case "4":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

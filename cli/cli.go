package cli

import (
	"fmt"

	"github.com/Srijan-Sengupta/Backup-Master/backend"
)

func Cli() {
	fmt.Print("Enter the folder to archive : ")
	fmt.Scanln(&backend.Inputf)
	fmt.Print("Enter the file name (The zip file with .zip extension): ")
	fmt.Scanln(&backend.Outputf)
	backend.StartTakingBackup(func(msg string) { fmt.Println(msg) }, func(p float64) { fmt.Println(fmt.Sprintf("Progress percent = %f", p)) })
}

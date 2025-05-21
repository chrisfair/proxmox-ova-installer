package main
import
	("fmt"
	"os"
"github.com/chrisfair/proxmox-ova-installer/internal/ova")



func main() {
	// Check if the user is root
	if os.Geteuid() != 0 {
		fmt.Println("This program must be run as root.")
		os.Exit(1)
	}

	// Check if the user is running on Proxmox
	// if !ova.IsProxmox() {
	// 	fmt.Println("This program must be run on Proxmox.")
	// 	os.Exit(1)
	// }
	//
	// // Check if the user has provided an OVA file
	// if len(os.Args) < 2 {
	// 	fmt.Println("Usage: importova <path to OVA file>")
	// 	os.Exit(1)
	// }
	//
	// // Import the OVA file
	// err := ova.ImportOVA(os.Args[1])
	// if err != nil {
	// 	fmt.Println("Error importing OVA file:", err)
	// 	os.Exit(1)
	// }

	newOVA := ova.OVA{}

	newOVA.Path = "./orion-vmware-10.6.0.ova"
	newOVA.Dir = "./"

	_, err := newOVA.Extract(newOVA.Path)

	if err == nil {
			fmt.Println("OVA file imported successfully.")
	}
}

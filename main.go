package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var err error
	// R Version urls goes here
	rversions := []string{"https://cran.r-project.org/bin/macosx/old/R-3.0.3.pkg", "https://cran.r-project.org/bin/macosx/el-capitan/base/R-3.4.4.pkg", "https://cran.r-project.org/bin/macosx/R-4.0.0.pkg"}
	fmt.Println("Starting R Installer for MacOS")
	for _, rversion := range rversions {
		urlParts := strings.Split(rversion, "/")
		pkg := urlParts[len(urlParts)-1]
		version := strings.ReplaceAll(pkg, ".pkg", "")
		fmt.Printf("Installing %v ...... \n", version)
		err = install(rversion)
		if err != nil {
			log.Fatalf("ERROR: %v", err)
		}
		fmt.Printf("Installation completed for %v ...... \n", version)
	}

	//Forget install paths
	fmt.Println("Forgetting paths")
	executeCommand("sudo", "pkgutil", "--forget", "org.r-project.R.el-capitan.fw.pkg", "--forget", "org.r-project.x86_64.tcltk.x11", "--forget", "org.r-project.x86_64.texinfo", "--forget", "org.r-project.R.el-capitan.GUI.pkg", "--forget", "org.r-project.R.mavericks.fw.pkg", "--forget", "org.r-project.R.mavericks.GUI.pkg", "--forget", "org.r-project.R.mavericks.GUI64.pkg")
	// if err != nil {
	// 	log.Printf("WARN: %v \n Can be ingnored if its first install", err)
	// }

	//R Switch url goes here
	rswitchURL := "http://groundhogr.com/rswitch/RSwitch-1.7.0.app.zip"
	fmt.Println("Installing RSwitch ......")
	err = installAppFromArchive(rswitchURL)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Println("Installation completed for RSwitch ......")

}

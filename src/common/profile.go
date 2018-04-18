package common

import (
	"os"
	"fmt"
)
/******************************************************************
 DIRECTORY
 ******************************************************************/
var BasePath string = os.Getenv("GOPATH")
var Staging string = os.Getenv("STAGING")

/******************************************************************
 LOG CONFIGURATION
 ******************************************************************/
const LogFileName string = "app.log"
var LogPath string = fmt.Sprintf("%v/var/log/%v", BasePath, LogFileName)

/******************************************************************
 SERVER CONFIGURATION
 ******************************************************************/
const HttpPort int = 80
const HttpsPort int = 443
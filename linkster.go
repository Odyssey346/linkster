package linkster

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func init() {
	fmt.Println(nil)
}

func OpenSite(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("Your platform can't use this module. Supported platforms: Linux, Windows and macOS.")
	}
	if err != nil {
		log.Fatal(err)
	}
}

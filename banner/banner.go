package banner

import (
	"fmt"
)

// prints the version message
const version = "v0.0.1"

func PrintVersion() {
	fmt.Printf("Current robotxt version %s\n", version)
}

// Prints the Colorful banner
func PrintBanner() {
	banner := `
    ____          __            __         __ 
   / __ \ ____   / /_   ____   / /_ _  __ / /_
  / /_/ // __ \ / __ \ / __ \ / __/| |/_// __/
 / _, _// /_/ // /_/ // /_/ // /_ _>  < / /_  
/_/ |_| \____//_.___/ \____/ \__//_/|_| \__/
`
	fmt.Printf("%s\n%50s\n\n", banner, "Current robotxt version "+version)
}

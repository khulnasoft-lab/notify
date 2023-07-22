package runner

import (
	"github.com/khulnasoft-labs/gologger"
	updateutils "github.com/khulnasoft-labs/utils/update"
)

const banner = `
             __  _ ___    
  ___  ___  / /_(_) _/_ __
 / _ \/ _ \/ __/ / _/ // /
/_//_/\___/\__/_/_/ \_, / v1.0.5
                   /___/  
`

// version is the current version
const version = `1.0.5`

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\tkhulnasoft-labs.io\n\n")
}

// GetUpdateCallback returns a callback function that updates notify
func GetUpdateCallback() func() {
	return func() {
		showBanner()
		updateutils.GetUpdateToolCallback("notify", version)()
	}
}

package upgrade

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/xzeu/tz/config"
)

var CmdUpgrade = &cobra.Command{
	Use:     "upgrade",
	Short:   "Upgrade the tz command.",
	Long:    "Upgrade the tz command.",
	Example: "tz upgrade",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("go install %s\n", config.TzCmd)
		cmd := exec.Command("go", "install", config.TzCmd)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("go install %s error\n", err)
		}
		fmt.Printf("\n🎉 Tz upgrade successfully!\n\n")
	},
}

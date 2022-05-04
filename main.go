package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	tz := "America/New_York"
	a, err := time.LoadLocation(tz)

	if err != nil {
		panic(err)
	}

	var rootCmd = &cobra.Command{
		Use:   "yadate",
		Short: "Yet another date command",
		Long:  `Yet another date command for providing more flexible ways for standard date command.`,
		Run: func(cmd *cobra.Command, args []string) {
			if bUtc, err := cmd.Flags().GetBool("utc"); err == nil {
				if bUtc {
					now := time.Now().In(a)
					fmt.Println(now)
				} else {
					now := time.Now()
					fmt.Println(now)
				}
			}
		},
	}

	rootCmd.Flags().BoolP(
		"utc", "u", false,
		"print or set Coordinated Universal Time (UTC)")

	rootCmd.Execute()
}

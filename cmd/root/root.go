package root

import (
	"time"

	"github.com/spf13/cobra"
)

type Config struct {
	URL             string
	RefreshInterval time.Duration
	Timeout         time.Duration
	ShouldFail      bool
	FollowRedirects bool
	SkipSSL         bool
	AssertText      string
	ReceiveAlert    bool
	Simple          bool
	NoFancy         bool
	Count           int
}

var AppConfig Config

var RootCmd = &cobra.Command{
	Use:   "updo",
	Short: "A simple website monitoring tool",
	Long: `Updo is a lightweight, easy-to-use website monitoring tool that checks
website availability and response time. It provides both a terminal UI
and a simple text-based output mode.`,
	Example: `  updo --url https://example.com
  updo --url https://example.com -r 10 -t 5
  updo --url https://example.com --simple -c 10
  updo --url https://example.com --simple --no-fancy
  updo --url https://example.com -a "Welcome"`,
}

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&AppConfig.URL, "url", "u", "", "URL or IP address to monitor")
	RootCmd.PersistentFlags().IntP("refresh", "r", 5, "Refresh interval in seconds")
	RootCmd.PersistentFlags().IntP("timeout", "t", 10, "HTTP request timeout in seconds")
	RootCmd.PersistentFlags().BoolVarP(&AppConfig.ShouldFail, "should-fail", "f", false, "Invert success code range")
	RootCmd.PersistentFlags().BoolVarP(&AppConfig.FollowRedirects, "follow-redirects", "l", true, "Follow redirects")
	RootCmd.PersistentFlags().BoolVarP(&AppConfig.SkipSSL, "skip-ssl", "s", false, "Skip SSL certificate verification")
	RootCmd.PersistentFlags().StringVarP(&AppConfig.AssertText, "assert-text", "a", "", "Text to assert in the response body")
	RootCmd.PersistentFlags().BoolVarP(&AppConfig.ReceiveAlert, "receive-alert", "n", true, "Enable alert notifications")
	RootCmd.PersistentFlags().BoolVar(&AppConfig.Simple, "simple", false, "Use simple output instead of TUI")
	RootCmd.PersistentFlags().BoolVar(&AppConfig.NoFancy, "no-fancy", false, "Disable fancy terminal formatting in simple mode")
	RootCmd.PersistentFlags().IntVarP(&AppConfig.Count, "count", "c", 0, "Number of checks to perform (0 = infinite)")

	RootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		refresh, _ := cmd.Flags().GetInt("refresh")
		timeout, _ := cmd.Flags().GetInt("timeout")

		AppConfig.RefreshInterval = time.Duration(refresh) * time.Second
		AppConfig.Timeout = time.Duration(timeout) * time.Second
	}
}

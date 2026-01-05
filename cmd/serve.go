package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var botToken string
var listenAddr string
var dataBase string
var dataRoot string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run bot",
	Long:  `Run the bot and start to serve http requests and telegram commands.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("In development")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVar(&botToken, "token", "", "Telergam bot token")
	serveCmd.Flags().StringVar(&listenAddr, "addr", "127.0.0.1:11090", "Listen addr for api")
	serveCmd.Flags().StringVar(&dataBase, "db", "", "Database address")
	serveCmd.Flags().StringVar(&dataRoot, "data", "", "Path to hold files")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

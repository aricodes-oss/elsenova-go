/*
Copyright © 2024 Aria Taylor <ari@aricodes.net>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"elsenova/bot"

	"github.com/rs/zerolog/log"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Runs the discord bot.",
	Long:  `Runs the discord bot.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Starting up!")
		token := viper.GetString("token")
		if token == "" {
			// cobra provides the RunE field which lets us return an `err` from Run
			// but we're using zerlog everywhere else for status messages
			log.Fatal().Msg("No token specified in elsenova.yml!")
		}

		b, err := bot.New(token)
		if err != nil {
			log.Fatal().Err(err).Msg("Error connecting to discord!")
		}

		err = b.Start()
		if err != nil {
			log.Fatal().Err(err).Msg("Error starting elsenova!")
		}
		log.Info().Msg("Ret-2-go!")

		b.Wait()
		b.Stop()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

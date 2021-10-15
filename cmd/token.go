package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Save token to config file",
	Long: `It's hack day!'`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("provide only one argument")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		dirname, err := os.UserHomeDir()
    		if err != nil {
        		panic( err )
    		}
		filename := path.Join(dirname, ".rcli.yml")
		 _, err = os.Stat(filename)

		if os.IsNotExist(err){
			file, err := os.Create(filename)
			if err != nil {
				fmt.Println("os.Create:", err)
				return
			}
		defer file.Close()
		fmt.Fprintf(file, "api: %s\n", token)
		fmt.Println("Token added in ", filename)

		} else {
			file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
			// the O_CREATE flag actually will create the file if it doesn't exist, which made me feel silly for writing the above, 
			// but the messaging to the user about what is happening is better with this approach I suppose : )
			if err!= nil{
				fmt.Println(err)
				return
			}
			defer file.Close()
			fmt.Println("Config file already exists. Overwriting token in ", filename)
			fmt.Fprintf(file, "api: %s\n", token)
		}

	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)
}

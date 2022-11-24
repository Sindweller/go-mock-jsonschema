package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"go-mock-jsonschema/pkg/mock"
	"io"
	"os"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "mock",
		Short: "generate mock data",
		Long:  `generate mock data`,
		Run: func(cmd *cobra.Command, args []string) {
			path, _ := cmd.Flags().GetString("input")
			file, err := os.Open(path)
			if err != nil {
				fmt.Println(err)
				return
			}

			buf := new(bytes.Buffer)
			_, err = io.Copy(buf, file)
			if err != nil {
				fmt.Println(err)
				return
			}
			// TODO 输出
			fmt.Println(mock.GetMock(buf.String()))
		},
	}
	jsonPath := "./example/input.json"
	rootCmd.PersistentFlags().StringVarP(&jsonPath, "input", "i", "", "input json file")
	return rootCmd
}

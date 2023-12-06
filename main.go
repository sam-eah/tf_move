package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var in, out, from_trim, from_add, to_trim, to_add string
	rootCmd.PersistentFlags().StringVarP(&in, "in", "i", "state.txt", "input file")
	rootCmd.PersistentFlags().StringVarP(&out, "out", "o", "moved.tf", "output file")
	rootCmd.PersistentFlags().StringVarP(&from_trim, "from-trim", "", "", "prefix to trim for from")
	rootCmd.PersistentFlags().StringVarP(&from_add, "from-add", "", "", "prefix to add for from")
	rootCmd.PersistentFlags().StringVarP(&to_trim, "to-trim", "", "", "prefix to trim for to")
	rootCmd.PersistentFlags().StringVarP(&to_add, "to-add", "", "", "prefix to add for to")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	file, err := os.Open(in)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	output, err := os.Create(out)
	if err != nil {
		fmt.Println(err)
	}
	defer output.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if !strings.HasPrefix(scanner.Text(), "data.") {
			str := fmt.Sprintf(`moved {
  from = %s%s
  to = %s%s
}
`,
				from_add,
				strings.TrimPrefix(scanner.Text(), from_trim),
				to_add,
				strings.TrimPrefix(scanner.Text(), to_trim))
			_, err := output.WriteString(str)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "tf_move",
	Short: "Move terraform resources",
	Long: `Complete documentation is available at github.com/sam-eah/tf_move`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

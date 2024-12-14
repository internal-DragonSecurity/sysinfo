/*
Copyright © 2024 Dragon Security

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
	"encoding/json"
	"fmt"
	"github.com/dragonsecurity/sysinfo/pkg/sysinfo"
	"log"

	"github.com/spf13/cobra"
)

// fullCmd represents the full command
var fullCmd = &cobra.Command{
	Use:   "full",
	Short: "full output of all sysinfo",
	Long:  "full output of all sysinfo",
	Run: func(cmd *cobra.Command, args []string) {
		var si sysinfo.SysInfo
		si.GetSysInfo()
		data, err := json.MarshalIndent(&si, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
	},
}

func init() {
	rootCmd.AddCommand(fullCmd)
}

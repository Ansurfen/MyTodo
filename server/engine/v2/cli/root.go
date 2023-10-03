package cli

import "github.com/spf13/cobra"

type BootOption struct {
	ProcessMode bool     `json:"processMode"`
	Filter      []string `json:"filter"`
	Port        []int    `json:"port"`
	Tmpl        string   `json:"tmpl"`
	Range       int      `json:"range"`
	Base        int      `json:"base"`
}

var todoCmd = &cobra.Command{
	Use: `todo`,
}

func Execute() {
	err := todoCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func init() {
	todoCmd.PersistentFlags().BoolVarP(&Option.ProcessMode, "process", "", false, "")
	todoCmd.PersistentFlags().StringSliceVarP(&Option.Filter, "filter", "f", nil, "filters service of given name or port from origin config")
	todoCmd.PersistentFlags().IntSliceVarP(&Option.Port, "port", "p", nil, "")
	todoCmd.PersistentFlags().StringVarP(&Option.Tmpl, "template", "t", "", "")
	todoCmd.PersistentFlags().IntVarP(&Option.Range, "range", "r", -1, "")
	todoCmd.PersistentFlags().StringSlice("rkset", nil, "")
}

var Option BootOption

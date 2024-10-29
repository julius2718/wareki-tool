/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type AlfredStruct struct {
	title    string
	subtitle string
	arg      string
}

var Alfred bool

// henkanCmd represents the henkan command
var henkanCmd = &cobra.Command{
	Use:   "henkan",
	Short: "Convert Japanese Year to Common Era Year",
	Long:  `Convert Japanese Year to Common Era Year`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		output := convert(input)

		result := fmt.Sprintf("%d年", output)

		if Alfred == true {
			alfred_item_dict := map[string]string{
				"title":    result,
				"subtitle": input,
				"arg":      result,
			}
			alfred_items := []map[string]string{alfred_item_dict}
			alfred_dict := map[string][]map[string]string{
				"items": alfred_items,
			}
			alfred_json, err := json.Marshal(alfred_dict)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(alfred_json))
		} else {
			fmt.Println(result)
		}
	},
}

func init() {
	rootCmd.AddCommand(henkanCmd)

	henkanCmd.PersistentFlags().BoolVarP(&Alfred, "alfred", "a", false, "Convert output into JSON for Alfred 5 (default: false)")
	viper.BindPFlag("alfred", rootCmd.PersistentFlags().Lookup("alfred"))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// henkanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// henkanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func convert(input string) int {
	era_year_str := extract_wyear(input)
	year := convert_wyear(era_year_str)

	return year
}

func extract_wyear(wyear string) [2]string {
	wyear_rune := []rune(wyear)
	wyear_len := len(wyear_rune)

	era_list := []string{"明治", "大正", "昭和", "平成", "令和"}

	if (slices.Contains(era_list, string(wyear_rune[:2]))) && (string(wyear_rune[wyear_len-1]) == "年") {
		era_rune := wyear_rune[:2]
		year_rune := wyear_rune[2 : wyear_len-1]
		era_str := string(era_rune)
		year_str := string(year_rune)
		era_year_str := [2]string{era_str, year_str}
		return era_year_str
	} else {
		panic("年号が正しくありません。「平成10年」のような形で入力してください。")
	}
}

func convert_wyear(era_year [2]string) int {
	era_str := era_year[0]
	wyear_str := era_year[1]

	if wyear_str == "元" {
		wyear_str = "1"
	}
	wyear_int, err := strconv.Atoi(wyear_str)
	if err != nil {
		panic(err)
	}

	year_int := 9999
	if era_str == "令和" {
		year_int = wyear_int + 2018
	} else if era_str == "平成" {
		year_int = wyear_int + 1988
	} else if era_str == "昭和" {
		year_int = wyear_int + 1925
	} else if era_str == "大正" {
		year_int = wyear_int + 1911
	} else if era_str == "明治" {
		year_int = wyear_int + 1867
	} else {
		panic("元号が正しくありません。")
	}

	return year_int
}

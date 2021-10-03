/*
Copyright © 2021 Muhammad Yarjuna Rohmat <myarjunar@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/myarjunar/pokcli/pkg/pokermachine"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pokcli <hand_1> <hand_2>",
	Short: "Find a winner from two poker hands.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("you need to provide two poker hands e.g AAAQQ")
		}

		firstCards := args[0]
		secondCards := args[1]

		hand1 := pokermachine.NewPokerHand(firstCards)
		hand2 := pokermachine.NewPokerHand(secondCards)

		winner, err := pokermachine.FindWinner(hand1, hand2)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(winner)
		}

		return nil
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

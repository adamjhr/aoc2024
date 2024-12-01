package main

import (
	"fmt";
	"bufio";
	"os";
	"strings";
	"strconv";
	"sort";
)



func main() {
	var diff int
	var err error
	team_1 := []int{}
	team_2 := []int{}

	fmt.Println("Day 1")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println("Scanning next pair")
		pair := s.Text()
		fmt.Printf("%s\n", pair)
		pair_slice := strings.Split(pair, "   ")
		
		var first_id int
		first_id, err = strconv.Atoi(pair_slice[0])
		if (err != nil) {
			fmt.Println("Error, exiting");
			return
		}

		var second_id int
		second_id, err = strconv.Atoi(pair_slice[1])
		if (err != nil) {
			fmt.Println("Error, exiting");
			return
		}

		fmt.Printf("%d %d\n", first_id, second_id)

		team_1 = append(team_1, first_id)
		team_2 = append(team_2, second_id)
	}
	fmt.Println("Finished scanning")
	sort.Slice(team_1, func(i, j int) bool {return team_1[i] < team_1[j]})
	sort.Slice(team_2, func(i, j int) bool {return team_2[i] < team_2[j]})

	for i, id_1 := range(team_1) {
		id_2 := team_2[i];
		fmt.Printf("%d %d\n", id_1, id_2)
		if (id_1 > id_2) {
			diff += id_1 - id_2 
		} else {
			diff += id_2 - id_1
		}
		fmt.Printf("Diff is is %d\n", diff)
	}

	fmt.Printf("Total diff is %d\n", diff)
}


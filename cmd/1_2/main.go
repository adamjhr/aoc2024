package main

import (
	"fmt";
	"bufio";
	"os";
	"strings";
	"strconv";
)



func main() {
	var sim int
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

	for _, id_1 := range(team_1) {

		for _, id_2 := range(team_2) {
			if (id_1 == id_2) {
				sim += id_1
			}	
		} 
		fmt.Printf("Sim is is %d\n", sim)
	}

	fmt.Printf("Total sim is %d\n", sim)
}


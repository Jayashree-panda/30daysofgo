package main

import (
	"encoding/csv"
	"strings"
	"flag"
	"fmt"
	"os"
	"time"
)

func main(){
	//setting the filename using command line arguments
	csvFilemane := flag.String("csv","problems.csv","The csv file contains a set of quiz questions and answers")
	timeLimit := flag.Int("limit", 30, "time limit of quiz game in seconds")
	flag.Parse()

	//reading the file
	file,err := os.Open(*csvFilemane) 
	if(err !=nil){
		exit(fmt.Sprintf("Failed to open the csv file %s\n", *csvFilemane))
	}

	//creating csvreader
	r := csv.NewReader(file)
	
	//read all the lines in csv
	lines,err := r.ReadAll()
	if(err != nil){
		exit("Failed to parse the csv file")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit)*time.Second)

	count := 0
	for i,p := range(problems){
		fmt.Printf("Problem #%d: %s = \n",i+1, p.question)
		answerChan := make(chan string)
		go func(){
			var answer string
			fmt.Scanf("%s \n", &answer)
			answerChan <- answer
		}()
		select {

		case <-timer.C:
			fmt.Printf("\n Your score is %d out of %d \n", count, len(problems) )
			return //terminate the program

		case answer := <- answerChan:
			if( answer == p.answer ){
				count++
				fmt.Printf("Correct! \n")
			}
		}
		
	}
}


//creating the set of questions from the csv data
func parseLines(lines [][] string) []problem {
	ret := make( [] problem, len(lines) )
	
	for i,line := range(lines){
		ret[i] = problem {
			question : line[0],
			answer : strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct{
	question string
	answer string
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}

	

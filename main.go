/*
the goal of this project to demonstrate setting up voting abstractions for instance:
*/
package main

import (
	"fmt"

	"github.com/goosedrunk/good/testpackage"
)

//type VotingResult interface {
//Implement()
//PrintResult() string
//}

//type VotingSystem interface {
//Results() VotingResult
//Vote() func Vote(name string, vote VoteType)
//}

func main() {
	fmt.Println("Please specify the type of voting system between (TODO specify options)")
	testpackage.SomeTestFunction()
	// TODO get user input
	// TODO process user input
	// ...
}

// TODO learn about adding different CLI options using https://github.com/spf13/cobra
// we should have the commands:
// - add voting system (many options, first past the post, preferential vote
//     - https://ottawacitizen.com/opinion/columnists/electoral-reform-five-different-voting-systems-at-a-glance
//     - create a general interface for "voting systems" and a bunch of structs for each voting system which fulfill the interface
//     - things like "execute" vote should be function variables
// - add a proposal in a voting system
// - cast your vote on a particular proposal
// - Tally up and execute the results of a vote (aka "Close" a vote - new votes casted in this proposal should be rejected)
//    - in this case, execution should just be printing a special message based on the results
//
// none of these commands should require verification, only adding your name
// (because this is an easy first program to write)
//
// All this information should be stored in JSON format (aka in a text file which JSON format) checkout https://gobyexample.com/json
//
// Your code should be organized into sub packages (we are in package main)

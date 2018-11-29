package commands

import "fmt"

type HelpCmd struct {
	Args struct {
		Topic string `positional-arg-name:"topic" description:"get information about available commands"`
	} `positional-args:"true"  required:"true"`
}

func (cmd *HelpCmd) Execute(args []string) error {

	if cmd.Args.Topic == "login" {
		fmt.Printf("%s\n", LoginCmdHelpMsg)
	}

	if cmd.Args.Topic == "authorizevote" {
		fmt.Printf("%s\n", AuthorizeVoteCmdHelpMsg)
	}

	if cmd.Args.Topic == "newuser" {
		fmt.Printf("%s\n", NewUserCmdHelpMsg)
	}
	
	if cmd.Args.Topic == "newproposal" {
		fmt.Printf("%s\n", NewProposalCmdHelpMsg)
	}

	return nil
}


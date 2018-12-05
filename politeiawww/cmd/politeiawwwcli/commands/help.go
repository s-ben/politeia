package commands

import "fmt"

type HelpCmd struct {
	Args struct {
		Topic string `positional-arg-name:"topic" description:"get information about available commands"`
	} `positional-args:"true"  required:"true"`
}

func (cmd *HelpCmd) Execute(args []string) error {

	switch cmd.Args.Topic {
	case "login":
		fmt.Printf("%s\n", LoginCmdHelpMsg)
	case "logout":
		fmt.Printf("%s\n", LogoutCmdHelpMsg)
	case "authorizevote":
		fmt.Printf("%s\n", AuthorizeVoteCmdHelpMsg)
	case "newuser":
		fmt.Printf("%s\n", NewUserCmdHelpMsg)
	case "newproposal":
		fmt.Printf("%s\n", NewProposalCmdHelpMsg)
	case "changepassword":
		fmt.Printf("%s\n", ChangePasswordCmdHelpMsg)
	case "changeusername":
		fmt.Printf("%s\n", ChangeUsernameCmdHelpMsg)
	case "faucet":
		fmt.Printf("%s\n", FaucetCmdHelpMsg)
	case "userdetails":
		fmt.Printf("%s\n", UserDetailsCmdHelpMsg)
	case "getproposal":
		fmt.Printf("%s\n", GetProposalCmdHelpMsg)
	default:
		fmt.Printf("invalid command\n")
	}

	return nil
}

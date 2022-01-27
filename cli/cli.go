package cli

import (
	"errors"
	"log"
	"os"
	"runtime"

	"github.com/Guilherme-De-Marchi/orwell-blockchain/blockchain"
)

type CommandLine struct {
	Chain *blockchain.Blockchain
}

func NewCommandLine(chain *blockchain.Blockchain) *CommandLine {
	return &CommandLine{
		Chain: chain,
	}
}

func (cli *CommandLine) Run() {
	argsLen := len(os.Args)
	if argsLen < 2 {
		cli.Error(errors.New("expected 2 or more arguments on command line"))
	}

	switch os.Args[1] {
	case "add":
		if argsLen != 3 {
			cli.PrintUsage()
			runtime.Goexit()
		}
		cli.AddBlock(os.Args[2])
	case "print":
		if argsLen != 2 {
			cli.PrintUsage()
			runtime.Goexit()
		}
		cli.PrintAllChain()
	default:
		cli.PrintUsage()
		runtime.Goexit()
	}
}

func (*CommandLine) Error(err error) {
	log.Println(err)
	runtime.Goexit()
}

func (cli *CommandLine) PrintUsage() {
	log.Println("Usage:")
	log.Println("add BLOCK_DATA `adds a block to the chain`")
}

func (cli *CommandLine) AddBlock(data string) {
	_, err := cli.Chain.AddNewBlock([]byte(data))
	if err != nil {
		cli.Error(err)
	}

	log.Println("Block added")
}

func (cli *CommandLine) PrintAllChain() {
	iter := cli.Chain.GetIterator()
	for iter.CurrentHash != nil {
		log.Println(iter.CurrentHash)
		b, err := iter.Next()
		log.Println("aqui0")
		if err != nil {
			cli.Error(err)
		}
		log.Println("aqui")
		b.Print()
		log.Println("aqui2")
	}
	log.Println("f", iter.CurrentHash)
}

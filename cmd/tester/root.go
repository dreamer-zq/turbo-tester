package cmd

import (
	"context"
	"math/big"
	"strconv"

	"github.com/dreamer-zq/turbo-tester/simple"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	flagURL     = "url"
	flagChainID = "chain-id"
)

// NewRootCmd returns a new instance of the cobra.Command struct.
//
// This function does not take any parameters.
// It returns a pointer to a cobra.Command struct.
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "tester",
		Short: "Turbo tester app command",
	}

	sampler := &simple.ETicketSampler{}

	rootCmd.AddCommand(DeployCmd(sampler))
	rootCmd.AddCommand(GentxCmd(sampler))
	rootCmd.AddCommand(StartCmd(sampler))

	rootCmd.PersistentFlags().String(flagURL, "", "turbo endpoint url")
	rootCmd.PersistentFlags().Int64(flagChainID, 0, "turbo chain-id")
	rootCmd.MarkFlagRequired(flagURL)
	rootCmd.MarkFlagRequired(flagChainID)
	return rootCmd
}

// GlobalConnfig represents a global config
type GlobalConnfig struct {
	chainID *big.Int
	url     string

	client *ethclient.Client
}

func loadGlobalFlags(cmd *cobra.Command) (*GlobalConnfig, error) {
	url, err := cmd.Flags().GetString(flagURL)
	if err != nil {
		return nil, err
	}

	chainIDInt, err := cmd.Flags().GetInt64(flagChainID)
	if err != nil {
		return nil, err
	}

	rpcClient, err := rpc.DialOptions(context.Background(), url, rpc.WithHeader("X-Chain", strconv.FormatInt(chainIDInt, 10)))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to the Ethereum client")
	}

	return &GlobalConnfig{
		chainID: big.NewInt(chainIDInt),
		url:     url,
		client:  ethclient.NewClient(rpcClient),
	}, nil
}

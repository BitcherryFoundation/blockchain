package keys

import (
	"bufio"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/hashrs/blockchain/framework/chain-app/client/flags"
	"github.com/hashrs/blockchain/framework/chain-app/client/input"
	"github.com/hashrs/blockchain/framework/chain-app/crypto/keys"
	sdk "github.com/hashrs/blockchain/framework/chain-app/types"
)

// ImportKeyCommand imports private keys from a keyfile.
func ImportKeyCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "import <name> <keyfile>",
		Short: "Import private keys into the local keybase",
		Long:  "Import a ASCII armored private key into the local keybase.",
		Args:  cobra.ExactArgs(2),
		RunE:  runImportCmd,
	}
}

func runImportCmd(cmd *cobra.Command, args []string) error {
	buf := bufio.NewReader(cmd.InOrStdin())
	kb, err := keys.NewKeyring(sdk.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), buf)
	if err != nil {
		return err
	}

	bz, err := ioutil.ReadFile(args[1])
	if err != nil {
		return err
	}

	passphrase, err := input.GetPassword("Enter passphrase to decrypt your key:", buf)
	if err != nil {
		return err
	}

	return kb.ImportPrivKey(args[0], string(bz), passphrase)
}

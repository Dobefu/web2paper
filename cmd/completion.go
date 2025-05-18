package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewCompletionCmd() (cmd *cobra.Command, err error) {
	cmd = &cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate a shell completion script",
		Long: `To load completions:

Bash:
  $ source <(./app completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ ./app completion bash > /etc/bash_completion.d/app
  # macOS:
  $ ./app completion bash > /usr/local/etc/bash_completion.d/app
  # or
  $ ./app completion bash > ~/.local/share/bash-completion/completions/app

Zsh:
  # If shell completion is not already enabled in your environment you will need
  # to enable it, see instructions below:
  # https://github.com/zsh-users/zsh-completions/blob/master/zsh-completions-howto.org#oh-my-zsh

  $ source <(./app completion zsh)

  # To load completions for each session, execute once:
  $ ./app completion zsh > "${fpath[1]}/_app"

  # You will need to start a new shell for this setup to take effect.

Fish:
  $ ./app completion fish | source

  # To load completions for each session, execute once:
  $ ./app completion fish > ~/.config/fish/completions/app.fish

  # You will need to start a new shell for this setup to take effect.
`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Annotations: map[string]string{
			"commandType": "main",
		},
		RunE: runCompletionCmd,
	}

	return cmd, nil
}
func init() {
	cmd, _ := NewCompletionCmd()

	rootCmd.AddCommand(cmd)
}

func runCompletionCmd(cmd *cobra.Command, args []string) error {
	switch args[0] {
	case "bash":
		return cmd.Root().GenBashCompletion(os.Stdout)
	case "zsh":
		return cmd.Root().GenZshCompletion(os.Stdout)
	case "fish":
		return cmd.Root().GenFishCompletion(os.Stdout, true)
	case "powershell":
		return cmd.Root().GenPowerShellCompletion(os.Stdout)
	default:
		return fmt.Errorf("unsupported shell type %q", args[0])
	}
}

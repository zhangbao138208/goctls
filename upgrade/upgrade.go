package upgrade

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zhangbao138208/goctls/rpc/execx"
)

// upgrade gets the latest goctl by
// go install github.com/zhangbao138208/goctls@latest
func upgrade(_ *cobra.Command, _ []string) error {
	cmd := `go install github.com/zhangbao138208/goctls@latest`
	info, err := execx.Run(cmd, "")
	if err != nil {
		return err
	}

	fmt.Print(info)
	return nil
}

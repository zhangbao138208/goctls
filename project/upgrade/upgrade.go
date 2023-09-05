package upgrade

import (
	"errors"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	conf "github.com/zhangbao138208/goctls/config"
	"github.com/zhangbao138208/goctls/rpc/execx"
	"os"
)

var (
	// VarBoolUpgradeMakefile describe whether to upgrade makefile
	VarBoolUpgradeMakefile bool
)

func UpgradeProject(_ *cobra.Command, _ []string) error {
	color.Green.Println("Start upgrading dependencies...")

	err := editMod(conf.DefaultGoZeroVersion, conf.DefaultToolVersion)
	if err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	err = upgradeDependencies(wd)
	if err != nil {
		return err
	}

	if VarBoolUpgradeMakefile {
		color.Green.Println("Start upgrading Makefile ...")
		_, err = execx.Run("goctls extra makefile", wd)
		if err != nil {
			return errors.New("failed to upgrade makefile")
		}
	}

	color.Green.Println("Done.")

	return nil
}

package args

import (
	"fmt"
	"os"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/version"
)

func Version(versionFiles []byte) error {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		info, err := version.Load(versionFiles)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		version.Print(info)
		return fmt.Errorf("version displayed")
	}
	return nil
}

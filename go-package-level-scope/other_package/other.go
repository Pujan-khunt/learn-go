package other

import (
	"fmt"

	aliasForTheImportedPackage "test-module/lib_package"
)

func useIt() {
	fmt.Println(aliasForTheImportedPackage.GloballyAccessibleVariableViaImports)
}

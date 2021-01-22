package gdbc

import "github.com/procyon-projects/procyon-core"

func init() {
	core.Register(newSimpleDatabaseConnectionProvider)
}

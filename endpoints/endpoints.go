package endpoints

import (
	"os"
)

// WebAppRoot ...
var WebAppRoot = os.Getenv("GOPHERFACE_APP_ROOT")

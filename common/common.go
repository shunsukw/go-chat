package common

import (
	"github.com/shunsukw/go-chat/common/datastore"
	"go.isomorphicgo.org/go/isokit"
)

// Env ...
type Env struct {
	DB          datastore.Datastore
	TemplateSet *isokit.TemplateSet
}

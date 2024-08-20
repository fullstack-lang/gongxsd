package reqif

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gongxsd-test-reqif/dist/ng-github.com-fullstack-lang-gongxsd-test-reqif
var NgDistNg embed.FS

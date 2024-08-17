package alt

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gongxsd-alt/dist/ng-github.com-fullstack-lang-gongxsd-alt
var NgDistNg embed.FS

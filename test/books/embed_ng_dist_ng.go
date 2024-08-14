package books

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gongxsd-test-books/dist/ng-github.com-fullstack-lang-gongxsd-test-books
var NgDistNg embed.FS

// do not modify, generated file
package fullstack

import (
	"github.com/fullstack-lang/gongxsd/go/controllers"
	"github.com/fullstack-lang/gongxsd/go/models"
	"github.com/fullstack-lang/gongxsd/go/orm"

	"github.com/gin-gonic/gin"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code
	_ "github.com/fullstack-lang/gongxsd/ng-github.com-fullstack-lang-gongxsd"
)

// NewStackInstance creates a new stack instance from the Stack Model
// and returns the backRepo of the stack instance (you can get the stage from backRepo.GetStage()
//
// - the stackPath is the unique identifier of the stack
// - the optional parameter filenames is for the name of the database filename
// if filenames is omitted, the database is persisted in memory
func NewStackInstance(
	r *gin.Engine,
	stackPath string,
	// filesnames is an optional parameter for the name of the database
	filenames ...string) (
	stage *models.Stage,
	backRepo *orm.BackRepoStruct) {

	stage = models.NewStage(stackPath)

	if len(filenames) == 0 {
		filenames = append(filenames, ":memory:")
	}

	backRepo = orm.NewBackRepo(stage, filenames[0])

	controllers.GetController().AddBackRepo(backRepo, stackPath)

	controllers.Register(r)

	// add orchestration
	// insertion point
	models.SetOrchestratorOnAfterUpdate[models.All](stage)
	models.SetOrchestratorOnAfterUpdate[models.Annotation](stage)
	models.SetOrchestratorOnAfterUpdate[models.Attribute](stage)
	models.SetOrchestratorOnAfterUpdate[models.AttributeGroup](stage)
	models.SetOrchestratorOnAfterUpdate[models.Choice](stage)
	models.SetOrchestratorOnAfterUpdate[models.ComplexContent](stage)
	models.SetOrchestratorOnAfterUpdate[models.ComplexType](stage)
	models.SetOrchestratorOnAfterUpdate[models.Documentation](stage)
	models.SetOrchestratorOnAfterUpdate[models.Element](stage)
	models.SetOrchestratorOnAfterUpdate[models.Enumeration](stage)
	models.SetOrchestratorOnAfterUpdate[models.Extension](stage)
	models.SetOrchestratorOnAfterUpdate[models.Group](stage)
	models.SetOrchestratorOnAfterUpdate[models.Length](stage)
	models.SetOrchestratorOnAfterUpdate[models.MaxInclusive](stage)
	models.SetOrchestratorOnAfterUpdate[models.MaxLength](stage)
	models.SetOrchestratorOnAfterUpdate[models.MinInclusive](stage)
	models.SetOrchestratorOnAfterUpdate[models.MinLength](stage)
	models.SetOrchestratorOnAfterUpdate[models.Pattern](stage)
	models.SetOrchestratorOnAfterUpdate[models.Restriction](stage)
	models.SetOrchestratorOnAfterUpdate[models.Schema](stage)
	models.SetOrchestratorOnAfterUpdate[models.Sequence](stage)
	models.SetOrchestratorOnAfterUpdate[models.SimpleContent](stage)
	models.SetOrchestratorOnAfterUpdate[models.SimpleType](stage)
	models.SetOrchestratorOnAfterUpdate[models.TotalDigit](stage)
	models.SetOrchestratorOnAfterUpdate[models.Union](stage)
	models.SetOrchestratorOnAfterUpdate[models.WhiteSpace](stage)

	return
}

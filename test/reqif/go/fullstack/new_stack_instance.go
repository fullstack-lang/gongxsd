// do not modify, generated file
package fullstack

import (
	"github.com/fullstack-lang/gongxsd/test/reqif/go/controllers"
	"github.com/fullstack-lang/gongxsd/test/reqif/go/models"
	"github.com/fullstack-lang/gongxsd/test/reqif/go/orm"

	"github.com/gin-gonic/gin"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code
	_ "github.com/fullstack-lang/gongxsd/test/reqif/ng-github.com-fullstack-lang-gongxsd-test-reqif/projects"
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
	stage *models.StageStruct,
	backRepo *orm.BackRepoStruct) {

	stage = models.NewStage(stackPath)

	if len(filenames) == 0 {
		filenames = append(filenames, ":memory:")
	}

	backRepo = orm.NewBackRepo(stage, filenames[0])

	if stackPath != "" {
		controllers.GetController().AddBackRepo(backRepo, stackPath)
	}

	controllers.Register(r)

	// add orchestration
	// insertion point
	models.SetOrchestratorOnAfterUpdate[models.ALTERNATIVE_ID](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_BOOLEAN](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_DATE](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_ENUMERATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_INTEGER](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_REAL](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_STRING](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_XHTML](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_BOOLEAN](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_DATE](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_ENUMERATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_INTEGER](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_REAL](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_STRING](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_XHTML](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_BOOLEAN](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_DATE](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_ENUMERATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_INTEGER](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_REAL](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_STRING](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_XHTML](stage)
	models.SetOrchestratorOnAfterUpdate[models.EMBEDDED_VALUE](stage)
	models.SetOrchestratorOnAfterUpdate[models.ENUM_VALUE](stage)
	models.SetOrchestratorOnAfterUpdate[models.RELATION_GROUP](stage)
	models.SetOrchestratorOnAfterUpdate[models.RELATION_GROUP_TYPE](stage)
	models.SetOrchestratorOnAfterUpdate[models.REQ_IF](stage)
	models.SetOrchestratorOnAfterUpdate[models.REQ_IF_CONTENT](stage)
	models.SetOrchestratorOnAfterUpdate[models.REQ_IF_HEADER](stage)
	models.SetOrchestratorOnAfterUpdate[models.REQ_IF_TOOL_EXTENSION](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECIFICATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECIFICATION_TYPE](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPEC_HIERARCHY](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPEC_OBJECT](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPEC_OBJECT_TYPE](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPEC_RELATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPEC_RELATION_TYPE](stage)
	models.SetOrchestratorOnAfterUpdate[models.XHTML_CONTENT](stage)

	return
}
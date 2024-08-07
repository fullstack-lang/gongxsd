// generated code - do not edit
package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fullstack-lang/gongxsd/go/orm"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// registerControllers register controllers
func registerControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gongxsd/go")
	{ // insertion point for registrations
		v1.GET("/v1/annotations", GetController().GetAnnotations)
		v1.GET("/v1/annotations/:id", GetController().GetAnnotation)
		v1.POST("/v1/annotations", GetController().PostAnnotation)
		v1.PATCH("/v1/annotations/:id", GetController().UpdateAnnotation)
		v1.PUT("/v1/annotations/:id", GetController().UpdateAnnotation)
		v1.DELETE("/v1/annotations/:id", GetController().DeleteAnnotation)

		v1.GET("/v1/complextypes", GetController().GetComplexTypes)
		v1.GET("/v1/complextypes/:id", GetController().GetComplexType)
		v1.POST("/v1/complextypes", GetController().PostComplexType)
		v1.PATCH("/v1/complextypes/:id", GetController().UpdateComplexType)
		v1.PUT("/v1/complextypes/:id", GetController().UpdateComplexType)
		v1.DELETE("/v1/complextypes/:id", GetController().DeleteComplexType)

		v1.GET("/v1/documentations", GetController().GetDocumentations)
		v1.GET("/v1/documentations/:id", GetController().GetDocumentation)
		v1.POST("/v1/documentations", GetController().PostDocumentation)
		v1.PATCH("/v1/documentations/:id", GetController().UpdateDocumentation)
		v1.PUT("/v1/documentations/:id", GetController().UpdateDocumentation)
		v1.DELETE("/v1/documentations/:id", GetController().DeleteDocumentation)

		v1.GET("/v1/elements", GetController().GetElements)
		v1.GET("/v1/elements/:id", GetController().GetElement)
		v1.POST("/v1/elements", GetController().PostElement)
		v1.PATCH("/v1/elements/:id", GetController().UpdateElement)
		v1.PUT("/v1/elements/:id", GetController().UpdateElement)
		v1.DELETE("/v1/elements/:id", GetController().DeleteElement)

		v1.GET("/v1/enumerations", GetController().GetEnumerations)
		v1.GET("/v1/enumerations/:id", GetController().GetEnumeration)
		v1.POST("/v1/enumerations", GetController().PostEnumeration)
		v1.PATCH("/v1/enumerations/:id", GetController().UpdateEnumeration)
		v1.PUT("/v1/enumerations/:id", GetController().UpdateEnumeration)
		v1.DELETE("/v1/enumerations/:id", GetController().DeleteEnumeration)

		v1.GET("/v1/maxinclusives", GetController().GetMaxInclusives)
		v1.GET("/v1/maxinclusives/:id", GetController().GetMaxInclusive)
		v1.POST("/v1/maxinclusives", GetController().PostMaxInclusive)
		v1.PATCH("/v1/maxinclusives/:id", GetController().UpdateMaxInclusive)
		v1.PUT("/v1/maxinclusives/:id", GetController().UpdateMaxInclusive)
		v1.DELETE("/v1/maxinclusives/:id", GetController().DeleteMaxInclusive)

		v1.GET("/v1/mininclusives", GetController().GetMinInclusives)
		v1.GET("/v1/mininclusives/:id", GetController().GetMinInclusive)
		v1.POST("/v1/mininclusives", GetController().PostMinInclusive)
		v1.PATCH("/v1/mininclusives/:id", GetController().UpdateMinInclusive)
		v1.PUT("/v1/mininclusives/:id", GetController().UpdateMinInclusive)
		v1.DELETE("/v1/mininclusives/:id", GetController().DeleteMinInclusive)

		v1.GET("/v1/patterns", GetController().GetPatterns)
		v1.GET("/v1/patterns/:id", GetController().GetPattern)
		v1.POST("/v1/patterns", GetController().PostPattern)
		v1.PATCH("/v1/patterns/:id", GetController().UpdatePattern)
		v1.PUT("/v1/patterns/:id", GetController().UpdatePattern)
		v1.DELETE("/v1/patterns/:id", GetController().DeletePattern)

		v1.GET("/v1/restrictions", GetController().GetRestrictions)
		v1.GET("/v1/restrictions/:id", GetController().GetRestriction)
		v1.POST("/v1/restrictions", GetController().PostRestriction)
		v1.PATCH("/v1/restrictions/:id", GetController().UpdateRestriction)
		v1.PUT("/v1/restrictions/:id", GetController().UpdateRestriction)
		v1.DELETE("/v1/restrictions/:id", GetController().DeleteRestriction)

		v1.GET("/v1/schemas", GetController().GetSchemas)
		v1.GET("/v1/schemas/:id", GetController().GetSchema)
		v1.POST("/v1/schemas", GetController().PostSchema)
		v1.PATCH("/v1/schemas/:id", GetController().UpdateSchema)
		v1.PUT("/v1/schemas/:id", GetController().UpdateSchema)
		v1.DELETE("/v1/schemas/:id", GetController().DeleteSchema)

		v1.GET("/v1/sequences", GetController().GetSequences)
		v1.GET("/v1/sequences/:id", GetController().GetSequence)
		v1.POST("/v1/sequences", GetController().PostSequence)
		v1.PATCH("/v1/sequences/:id", GetController().UpdateSequence)
		v1.PUT("/v1/sequences/:id", GetController().UpdateSequence)
		v1.DELETE("/v1/sequences/:id", GetController().DeleteSequence)

		v1.GET("/v1/simpletypes", GetController().GetSimpleTypes)
		v1.GET("/v1/simpletypes/:id", GetController().GetSimpleType)
		v1.POST("/v1/simpletypes", GetController().PostSimpleType)
		v1.PATCH("/v1/simpletypes/:id", GetController().UpdateSimpleType)
		v1.PUT("/v1/simpletypes/:id", GetController().UpdateSimpleType)
		v1.DELETE("/v1/simpletypes/:id", GetController().DeleteSimpleType)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)

		v1.GET("/v1/ws/commitfrombacknb", GetController().onWebSocketRequestForCommitFromBackNb)
		v1.GET("/v1/ws/stage", GetController().onWebSocketRequestForBackRepoContent)

		v1.GET("/v1/stacks", GetController().stacks)
	}
}

func (controller *Controller) stacks(c *gin.Context) {

	var res []string

	for k, _ := range controller.Map_BackRepos {
		res = append(res, k)
	}

	c.JSON(http.StatusOK, res)
}

// onWebSocketRequestForCommitFromBackNb is a function that is started each time
// a web socket request is received
//
// 1. upgrade the incomming web connection to a web socket
// 1. it subscribe to the backend commit number broadcaster
// 1. it stays live and pool for incomming backend commit number broadcast and forward
// them on the web socket connection
func (controller *Controller) onWebSocketRequestForCommitFromBackNb(c *gin.Context) {

	// log.Println("Stack github.com/fullstack-lang/gongxsd/go, onWebSocketRequestForCommitFromBackNb")

	// Upgrader specifies parameters for upgrading an HTTP connection to a
	// WebSocket connection.
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == "http://localhost:8080" || origin == "http://localhost:4200"
		},
	}

	wsConnection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wsConnection.Close()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go, Unkown stack", stackPath)
	}
	updateCommitBackRepoNbChannel := backRepo.SubscribeToCommitNb()

	for nbCommitBackRepo := range updateCommitBackRepoNbChannel {

		// Send elapsed time as a string over the WebSocket connection
		err = wsConnection.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%d", nbCommitBackRepo)))
		if err != nil {
			log.Println("client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
			fmt.Println(err)
			return
		}
	}
}

// onWebSocketRequestForBackRepoContent is a function that is started each time
// a web socket request is received
//
// 1. upgrade the incomming web connection to a web socket
// 1. it subscribe to the backend commit number broadcaster
// 1. it stays live and pool for incomming backend commit number broadcast and forward
// them on the web socket connection
func (controller *Controller) onWebSocketRequestForBackRepoContent(c *gin.Context) {

	// log.Println("Stack github.com/fullstack-lang/gongxsd/go, onWebSocketRequestForBackRepoContent")

	// Upgrader specifies parameters for upgrading an HTTP connection to a
	// WebSocket connection.
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == "http://localhost:8080" || origin == "http://localhost:4200"
		},
	}

	wsConnection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wsConnection.Close()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go, Unkown stack", stackPath)
	}
	updateCommitBackRepoNbChannel := backRepo.SubscribeToCommitNb()

	backRepoData := new(orm.BackRepoData)
	orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)

	err = wsConnection.WriteJSON(backRepoData)
	// log.Println("Stack github.com/fullstack-lang/gongxsd/go, onWebSocketRequestForBackRepoContent, first sent back repo of", stackPath)
	if err != nil {
		log.Println("client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
		fmt.Println(err)
		return
	}

	for nbCommitBackRepo := range updateCommitBackRepoNbChannel {
		_ = nbCommitBackRepo

		backRepoData := new(orm.BackRepoData)
		orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)

		// Send backRepo data
		err = wsConnection.WriteJSON(backRepoData)

		// log.Println("Stack github.com/fullstack-lang/gongxsd/go, onWebSocketRequestForBackRepoContent, sent back repo of", stackPath)

		if err != nil {
			log.Println("client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
			fmt.Println(err)
			return
		}
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func (controller *Controller) GetLastCommitFromBackNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func (controller *Controller) GetLastPushFromFrontNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastPushFromFrontNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}

// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongxsd/go/models"
	"github.com/fullstack-lang/gongxsd/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Documentation__dummysDeclaration__ models.Documentation
var __Documentation_time__dummyDeclaration time.Duration

var mutexDocumentation sync.Mutex

// An DocumentationID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDocumentation updateDocumentation deleteDocumentation
type DocumentationID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DocumentationInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDocumentation updateDocumentation
type DocumentationInput struct {
	// The Documentation to submit or modify
	// in: body
	Documentation *orm.DocumentationAPI
}

// GetDocumentations
//
// swagger:route GET /documentations documentations getDocumentations
//
// # Get all documentations
//
// Responses:
// default: genericError
//
//	200: documentationDBResponse
func (controller *Controller) GetDocumentations(c *gin.Context) {

	// source slice
	var documentationDBs []orm.DocumentationDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDocumentations", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDocumentation.GetDB()

	_, err := db.Find(&documentationDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	documentationAPIs := make([]orm.DocumentationAPI, 0)

	// for each documentation, update fields from the database nullable fields
	for idx := range documentationDBs {
		documentationDB := &documentationDBs[idx]
		_ = documentationDB
		var documentationAPI orm.DocumentationAPI

		// insertion point for updating fields
		documentationAPI.ID = documentationDB.ID
		documentationDB.CopyBasicFieldsToDocumentation_WOP(&documentationAPI.Documentation_WOP)
		documentationAPI.DocumentationPointersEncoding = documentationDB.DocumentationPointersEncoding
		documentationAPIs = append(documentationAPIs, documentationAPI)
	}

	c.JSON(http.StatusOK, documentationAPIs)
}

// PostDocumentation
//
// swagger:route POST /documentations documentations postDocumentation
//
// Creates a documentation
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDocumentation(c *gin.Context) {

	mutexDocumentation.Lock()
	defer mutexDocumentation.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDocumentations", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDocumentation.GetDB()

	// Validate input
	var input orm.DocumentationAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create documentation
	documentationDB := orm.DocumentationDB{}
	documentationDB.DocumentationPointersEncoding = input.DocumentationPointersEncoding
	documentationDB.CopyBasicFieldsFromDocumentation_WOP(&input.Documentation_WOP)

	_, err = db.Create(&documentationDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDocumentation.CheckoutPhaseOneInstance(&documentationDB)
	documentation := backRepo.BackRepoDocumentation.Map_DocumentationDBID_DocumentationPtr[documentationDB.ID]

	if documentation != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), documentation)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, documentationDB)
}

// GetDocumentation
//
// swagger:route GET /documentations/{ID} documentations getDocumentation
//
// Gets the details for a documentation.
//
// Responses:
// default: genericError
//
//	200: documentationDBResponse
func (controller *Controller) GetDocumentation(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDocumentation", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDocumentation.GetDB()

	// Get documentationDB in DB
	var documentationDB orm.DocumentationDB
	if _, err := db.First(&documentationDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var documentationAPI orm.DocumentationAPI
	documentationAPI.ID = documentationDB.ID
	documentationAPI.DocumentationPointersEncoding = documentationDB.DocumentationPointersEncoding
	documentationDB.CopyBasicFieldsToDocumentation_WOP(&documentationAPI.Documentation_WOP)

	c.JSON(http.StatusOK, documentationAPI)
}

// UpdateDocumentation
//
// swagger:route PATCH /documentations/{ID} documentations updateDocumentation
//
// # Update a documentation
//
// Responses:
// default: genericError
//
//	200: documentationDBResponse
func (controller *Controller) UpdateDocumentation(c *gin.Context) {

	mutexDocumentation.Lock()
	defer mutexDocumentation.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDocumentation", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDocumentation.GetDB()

	// Validate input
	var input orm.DocumentationAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var documentationDB orm.DocumentationDB

	// fetch the documentation
	_, err := db.First(&documentationDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	documentationDB.CopyBasicFieldsFromDocumentation_WOP(&input.Documentation_WOP)
	documentationDB.DocumentationPointersEncoding = input.DocumentationPointersEncoding

	db, _ = db.Model(&documentationDB)
	_, err = db.Updates(documentationDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	documentationNew := new(models.Documentation)
	documentationDB.CopyBasicFieldsToDocumentation(documentationNew)

	// redeem pointers
	documentationDB.DecodePointers(backRepo, documentationNew)

	// get stage instance from DB instance, and call callback function
	documentationOld := backRepo.BackRepoDocumentation.Map_DocumentationDBID_DocumentationPtr[documentationDB.ID]
	if documentationOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), documentationOld, documentationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the documentationDB
	c.JSON(http.StatusOK, documentationDB)
}

// DeleteDocumentation
//
// swagger:route DELETE /documentations/{ID} documentations deleteDocumentation
//
// # Delete a documentation
//
// default: genericError
//
//	200: documentationDBResponse
func (controller *Controller) DeleteDocumentation(c *gin.Context) {

	mutexDocumentation.Lock()
	defer mutexDocumentation.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDocumentation", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDocumentation.GetDB()

	// Get model if exist
	var documentationDB orm.DocumentationDB
	if _, err := db.First(&documentationDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&documentationDB)

	// get an instance (not staged) from DB instance, and call callback function
	documentationDeleted := new(models.Documentation)
	documentationDB.CopyBasicFieldsToDocumentation(documentationDeleted)

	// get stage instance from DB instance, and call callback function
	documentationStaged := backRepo.BackRepoDocumentation.Map_DocumentationDBID_DocumentationPtr[documentationDB.ID]
	if documentationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), documentationStaged, documentationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

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
var __ComplexContent__dummysDeclaration__ models.ComplexContent
var __ComplexContent_time__dummyDeclaration time.Duration

var mutexComplexContent sync.Mutex

// An ComplexContentID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getComplexContent updateComplexContent deleteComplexContent
type ComplexContentID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ComplexContentInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postComplexContent updateComplexContent
type ComplexContentInput struct {
	// The ComplexContent to submit or modify
	// in: body
	ComplexContent *orm.ComplexContentAPI
}

// GetComplexContents
//
// swagger:route GET /complexcontents complexcontents getComplexContents
//
// # Get all complexcontents
//
// Responses:
// default: genericError
//
//	200: complexcontentDBResponse
func (controller *Controller) GetComplexContents(c *gin.Context) {

	// source slice
	var complexcontentDBs []orm.ComplexContentDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetComplexContents", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoComplexContent.GetDB()

	_, err := db.Find(&complexcontentDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	complexcontentAPIs := make([]orm.ComplexContentAPI, 0)

	// for each complexcontent, update fields from the database nullable fields
	for idx := range complexcontentDBs {
		complexcontentDB := &complexcontentDBs[idx]
		_ = complexcontentDB
		var complexcontentAPI orm.ComplexContentAPI

		// insertion point for updating fields
		complexcontentAPI.ID = complexcontentDB.ID
		complexcontentDB.CopyBasicFieldsToComplexContent_WOP(&complexcontentAPI.ComplexContent_WOP)
		complexcontentAPI.ComplexContentPointersEncoding = complexcontentDB.ComplexContentPointersEncoding
		complexcontentAPIs = append(complexcontentAPIs, complexcontentAPI)
	}

	c.JSON(http.StatusOK, complexcontentAPIs)
}

// PostComplexContent
//
// swagger:route POST /complexcontents complexcontents postComplexContent
//
// Creates a complexcontent
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostComplexContent(c *gin.Context) {

	mutexComplexContent.Lock()
	defer mutexComplexContent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostComplexContents", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoComplexContent.GetDB()

	// Validate input
	var input orm.ComplexContentAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create complexcontent
	complexcontentDB := orm.ComplexContentDB{}
	complexcontentDB.ComplexContentPointersEncoding = input.ComplexContentPointersEncoding
	complexcontentDB.CopyBasicFieldsFromComplexContent_WOP(&input.ComplexContent_WOP)

	_, err = db.Create(&complexcontentDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoComplexContent.CheckoutPhaseOneInstance(&complexcontentDB)
	complexcontent := backRepo.BackRepoComplexContent.Map_ComplexContentDBID_ComplexContentPtr[complexcontentDB.ID]

	if complexcontent != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), complexcontent)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, complexcontentDB)
}

// GetComplexContent
//
// swagger:route GET /complexcontents/{ID} complexcontents getComplexContent
//
// Gets the details for a complexcontent.
//
// Responses:
// default: genericError
//
//	200: complexcontentDBResponse
func (controller *Controller) GetComplexContent(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetComplexContent", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoComplexContent.GetDB()

	// Get complexcontentDB in DB
	var complexcontentDB orm.ComplexContentDB
	if _, err := db.First(&complexcontentDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var complexcontentAPI orm.ComplexContentAPI
	complexcontentAPI.ID = complexcontentDB.ID
	complexcontentAPI.ComplexContentPointersEncoding = complexcontentDB.ComplexContentPointersEncoding
	complexcontentDB.CopyBasicFieldsToComplexContent_WOP(&complexcontentAPI.ComplexContent_WOP)

	c.JSON(http.StatusOK, complexcontentAPI)
}

// UpdateComplexContent
//
// swagger:route PATCH /complexcontents/{ID} complexcontents updateComplexContent
//
// # Update a complexcontent
//
// Responses:
// default: genericError
//
//	200: complexcontentDBResponse
func (controller *Controller) UpdateComplexContent(c *gin.Context) {

	mutexComplexContent.Lock()
	defer mutexComplexContent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateComplexContent", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoComplexContent.GetDB()

	// Validate input
	var input orm.ComplexContentAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var complexcontentDB orm.ComplexContentDB

	// fetch the complexcontent
	_, err := db.First(&complexcontentDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	complexcontentDB.CopyBasicFieldsFromComplexContent_WOP(&input.ComplexContent_WOP)
	complexcontentDB.ComplexContentPointersEncoding = input.ComplexContentPointersEncoding

	db, _ = db.Model(&complexcontentDB)
	_, err = db.Updates(complexcontentDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	complexcontentNew := new(models.ComplexContent)
	complexcontentDB.CopyBasicFieldsToComplexContent(complexcontentNew)

	// redeem pointers
	complexcontentDB.DecodePointers(backRepo, complexcontentNew)

	// get stage instance from DB instance, and call callback function
	complexcontentOld := backRepo.BackRepoComplexContent.Map_ComplexContentDBID_ComplexContentPtr[complexcontentDB.ID]
	if complexcontentOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), complexcontentOld, complexcontentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the complexcontentDB
	c.JSON(http.StatusOK, complexcontentDB)
}

// DeleteComplexContent
//
// swagger:route DELETE /complexcontents/{ID} complexcontents deleteComplexContent
//
// # Delete a complexcontent
//
// default: genericError
//
//	200: complexcontentDBResponse
func (controller *Controller) DeleteComplexContent(c *gin.Context) {

	mutexComplexContent.Lock()
	defer mutexComplexContent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteComplexContent", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoComplexContent.GetDB()

	// Get model if exist
	var complexcontentDB orm.ComplexContentDB
	if _, err := db.First(&complexcontentDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&complexcontentDB)

	// get an instance (not staged) from DB instance, and call callback function
	complexcontentDeleted := new(models.ComplexContent)
	complexcontentDB.CopyBasicFieldsToComplexContent(complexcontentDeleted)

	// get stage instance from DB instance, and call callback function
	complexcontentStaged := backRepo.BackRepoComplexContent.Map_ComplexContentDBID_ComplexContentPtr[complexcontentDB.ID]
	if complexcontentStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), complexcontentStaged, complexcontentDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

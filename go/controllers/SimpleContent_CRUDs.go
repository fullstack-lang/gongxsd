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
var __SimpleContent__dummysDeclaration__ models.SimpleContent
var __SimpleContent_time__dummyDeclaration time.Duration

var mutexSimpleContent sync.Mutex

// An SimpleContentID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSimpleContent updateSimpleContent deleteSimpleContent
type SimpleContentID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SimpleContentInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSimpleContent updateSimpleContent
type SimpleContentInput struct {
	// The SimpleContent to submit or modify
	// in: body
	SimpleContent *orm.SimpleContentAPI
}

// GetSimpleContents
//
// swagger:route GET /simplecontents simplecontents getSimpleContents
//
// # Get all simplecontents
//
// Responses:
// default: genericError
//
//	200: simplecontentDBResponse
func (controller *Controller) GetSimpleContents(c *gin.Context) {

	// source slice
	var simplecontentDBs []orm.SimpleContentDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSimpleContents", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSimpleContent.GetDB()

	_, err := db.Find(&simplecontentDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	simplecontentAPIs := make([]orm.SimpleContentAPI, 0)

	// for each simplecontent, update fields from the database nullable fields
	for idx := range simplecontentDBs {
		simplecontentDB := &simplecontentDBs[idx]
		_ = simplecontentDB
		var simplecontentAPI orm.SimpleContentAPI

		// insertion point for updating fields
		simplecontentAPI.ID = simplecontentDB.ID
		simplecontentDB.CopyBasicFieldsToSimpleContent_WOP(&simplecontentAPI.SimpleContent_WOP)
		simplecontentAPI.SimpleContentPointersEncoding = simplecontentDB.SimpleContentPointersEncoding
		simplecontentAPIs = append(simplecontentAPIs, simplecontentAPI)
	}

	c.JSON(http.StatusOK, simplecontentAPIs)
}

// PostSimpleContent
//
// swagger:route POST /simplecontents simplecontents postSimpleContent
//
// Creates a simplecontent
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSimpleContent(c *gin.Context) {

	mutexSimpleContent.Lock()
	defer mutexSimpleContent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSimpleContents", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSimpleContent.GetDB()

	// Validate input
	var input orm.SimpleContentAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create simplecontent
	simplecontentDB := orm.SimpleContentDB{}
	simplecontentDB.SimpleContentPointersEncoding = input.SimpleContentPointersEncoding
	simplecontentDB.CopyBasicFieldsFromSimpleContent_WOP(&input.SimpleContent_WOP)

	_, err = db.Create(&simplecontentDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSimpleContent.CheckoutPhaseOneInstance(&simplecontentDB)
	simplecontent := backRepo.BackRepoSimpleContent.Map_SimpleContentDBID_SimpleContentPtr[simplecontentDB.ID]

	if simplecontent != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), simplecontent)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, simplecontentDB)
}

// GetSimpleContent
//
// swagger:route GET /simplecontents/{ID} simplecontents getSimpleContent
//
// Gets the details for a simplecontent.
//
// Responses:
// default: genericError
//
//	200: simplecontentDBResponse
func (controller *Controller) GetSimpleContent(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSimpleContent", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSimpleContent.GetDB()

	// Get simplecontentDB in DB
	var simplecontentDB orm.SimpleContentDB
	if _, err := db.First(&simplecontentDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var simplecontentAPI orm.SimpleContentAPI
	simplecontentAPI.ID = simplecontentDB.ID
	simplecontentAPI.SimpleContentPointersEncoding = simplecontentDB.SimpleContentPointersEncoding
	simplecontentDB.CopyBasicFieldsToSimpleContent_WOP(&simplecontentAPI.SimpleContent_WOP)

	c.JSON(http.StatusOK, simplecontentAPI)
}

// UpdateSimpleContent
//
// swagger:route PATCH /simplecontents/{ID} simplecontents updateSimpleContent
//
// # Update a simplecontent
//
// Responses:
// default: genericError
//
//	200: simplecontentDBResponse
func (controller *Controller) UpdateSimpleContent(c *gin.Context) {

	mutexSimpleContent.Lock()
	defer mutexSimpleContent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSimpleContent", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSimpleContent.GetDB()

	// Validate input
	var input orm.SimpleContentAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var simplecontentDB orm.SimpleContentDB

	// fetch the simplecontent
	_, err := db.First(&simplecontentDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	simplecontentDB.CopyBasicFieldsFromSimpleContent_WOP(&input.SimpleContent_WOP)
	simplecontentDB.SimpleContentPointersEncoding = input.SimpleContentPointersEncoding

	db, _ = db.Model(&simplecontentDB)
	_, err = db.Updates(&simplecontentDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	simplecontentNew := new(models.SimpleContent)
	simplecontentDB.CopyBasicFieldsToSimpleContent(simplecontentNew)

	// redeem pointers
	simplecontentDB.DecodePointers(backRepo, simplecontentNew)

	// get stage instance from DB instance, and call callback function
	simplecontentOld := backRepo.BackRepoSimpleContent.Map_SimpleContentDBID_SimpleContentPtr[simplecontentDB.ID]
	if simplecontentOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), simplecontentOld, simplecontentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the simplecontentDB
	c.JSON(http.StatusOK, simplecontentDB)
}

// DeleteSimpleContent
//
// swagger:route DELETE /simplecontents/{ID} simplecontents deleteSimpleContent
//
// # Delete a simplecontent
//
// default: genericError
//
//	200: simplecontentDBResponse
func (controller *Controller) DeleteSimpleContent(c *gin.Context) {

	mutexSimpleContent.Lock()
	defer mutexSimpleContent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSimpleContent", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSimpleContent.GetDB()

	// Get model if exist
	var simplecontentDB orm.SimpleContentDB
	if _, err := db.First(&simplecontentDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&simplecontentDB)

	// get an instance (not staged) from DB instance, and call callback function
	simplecontentDeleted := new(models.SimpleContent)
	simplecontentDB.CopyBasicFieldsToSimpleContent(simplecontentDeleted)

	// get stage instance from DB instance, and call callback function
	simplecontentStaged := backRepo.BackRepoSimpleContent.Map_SimpleContentDBID_SimpleContentPtr[simplecontentDB.ID]
	if simplecontentStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), simplecontentStaged, simplecontentDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

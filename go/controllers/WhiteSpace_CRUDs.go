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
var __WhiteSpace__dummysDeclaration__ models.WhiteSpace
var __WhiteSpace_time__dummyDeclaration time.Duration

var mutexWhiteSpace sync.Mutex

// An WhiteSpaceID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getWhiteSpace updateWhiteSpace deleteWhiteSpace
type WhiteSpaceID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// WhiteSpaceInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postWhiteSpace updateWhiteSpace
type WhiteSpaceInput struct {
	// The WhiteSpace to submit or modify
	// in: body
	WhiteSpace *orm.WhiteSpaceAPI
}

// GetWhiteSpaces
//
// swagger:route GET /whitespaces whitespaces getWhiteSpaces
//
// # Get all whitespaces
//
// Responses:
// default: genericError
//
//	200: whitespaceDBResponse
func (controller *Controller) GetWhiteSpaces(c *gin.Context) {

	// source slice
	var whitespaceDBs []orm.WhiteSpaceDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetWhiteSpaces", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWhiteSpace.GetDB()

	query := db.Find(&whitespaceDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	whitespaceAPIs := make([]orm.WhiteSpaceAPI, 0)

	// for each whitespace, update fields from the database nullable fields
	for idx := range whitespaceDBs {
		whitespaceDB := &whitespaceDBs[idx]
		_ = whitespaceDB
		var whitespaceAPI orm.WhiteSpaceAPI

		// insertion point for updating fields
		whitespaceAPI.ID = whitespaceDB.ID
		whitespaceDB.CopyBasicFieldsToWhiteSpace_WOP(&whitespaceAPI.WhiteSpace_WOP)
		whitespaceAPI.WhiteSpacePointersEncoding = whitespaceDB.WhiteSpacePointersEncoding
		whitespaceAPIs = append(whitespaceAPIs, whitespaceAPI)
	}

	c.JSON(http.StatusOK, whitespaceAPIs)
}

// PostWhiteSpace
//
// swagger:route POST /whitespaces whitespaces postWhiteSpace
//
// Creates a whitespace
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostWhiteSpace(c *gin.Context) {

	mutexWhiteSpace.Lock()
	defer mutexWhiteSpace.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostWhiteSpaces", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWhiteSpace.GetDB()

	// Validate input
	var input orm.WhiteSpaceAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create whitespace
	whitespaceDB := orm.WhiteSpaceDB{}
	whitespaceDB.WhiteSpacePointersEncoding = input.WhiteSpacePointersEncoding
	whitespaceDB.CopyBasicFieldsFromWhiteSpace_WOP(&input.WhiteSpace_WOP)

	query := db.Create(&whitespaceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoWhiteSpace.CheckoutPhaseOneInstance(&whitespaceDB)
	whitespace := backRepo.BackRepoWhiteSpace.Map_WhiteSpaceDBID_WhiteSpacePtr[whitespaceDB.ID]

	if whitespace != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), whitespace)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, whitespaceDB)
}

// GetWhiteSpace
//
// swagger:route GET /whitespaces/{ID} whitespaces getWhiteSpace
//
// Gets the details for a whitespace.
//
// Responses:
// default: genericError
//
//	200: whitespaceDBResponse
func (controller *Controller) GetWhiteSpace(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetWhiteSpace", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWhiteSpace.GetDB()

	// Get whitespaceDB in DB
	var whitespaceDB orm.WhiteSpaceDB
	if err := db.First(&whitespaceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var whitespaceAPI orm.WhiteSpaceAPI
	whitespaceAPI.ID = whitespaceDB.ID
	whitespaceAPI.WhiteSpacePointersEncoding = whitespaceDB.WhiteSpacePointersEncoding
	whitespaceDB.CopyBasicFieldsToWhiteSpace_WOP(&whitespaceAPI.WhiteSpace_WOP)

	c.JSON(http.StatusOK, whitespaceAPI)
}

// UpdateWhiteSpace
//
// swagger:route PATCH /whitespaces/{ID} whitespaces updateWhiteSpace
//
// # Update a whitespace
//
// Responses:
// default: genericError
//
//	200: whitespaceDBResponse
func (controller *Controller) UpdateWhiteSpace(c *gin.Context) {

	mutexWhiteSpace.Lock()
	defer mutexWhiteSpace.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateWhiteSpace", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWhiteSpace.GetDB()

	// Validate input
	var input orm.WhiteSpaceAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var whitespaceDB orm.WhiteSpaceDB

	// fetch the whitespace
	query := db.First(&whitespaceDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	whitespaceDB.CopyBasicFieldsFromWhiteSpace_WOP(&input.WhiteSpace_WOP)
	whitespaceDB.WhiteSpacePointersEncoding = input.WhiteSpacePointersEncoding

	query = db.Model(&whitespaceDB).Updates(whitespaceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	whitespaceNew := new(models.WhiteSpace)
	whitespaceDB.CopyBasicFieldsToWhiteSpace(whitespaceNew)

	// redeem pointers
	whitespaceDB.DecodePointers(backRepo, whitespaceNew)

	// get stage instance from DB instance, and call callback function
	whitespaceOld := backRepo.BackRepoWhiteSpace.Map_WhiteSpaceDBID_WhiteSpacePtr[whitespaceDB.ID]
	if whitespaceOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), whitespaceOld, whitespaceNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the whitespaceDB
	c.JSON(http.StatusOK, whitespaceDB)
}

// DeleteWhiteSpace
//
// swagger:route DELETE /whitespaces/{ID} whitespaces deleteWhiteSpace
//
// # Delete a whitespace
//
// default: genericError
//
//	200: whitespaceDBResponse
func (controller *Controller) DeleteWhiteSpace(c *gin.Context) {

	mutexWhiteSpace.Lock()
	defer mutexWhiteSpace.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteWhiteSpace", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWhiteSpace.GetDB()

	// Get model if exist
	var whitespaceDB orm.WhiteSpaceDB
	if err := db.First(&whitespaceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&whitespaceDB)

	// get an instance (not staged) from DB instance, and call callback function
	whitespaceDeleted := new(models.WhiteSpace)
	whitespaceDB.CopyBasicFieldsToWhiteSpace(whitespaceDeleted)

	// get stage instance from DB instance, and call callback function
	whitespaceStaged := backRepo.BackRepoWhiteSpace.Map_WhiteSpaceDBID_WhiteSpacePtr[whitespaceDB.ID]
	if whitespaceStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), whitespaceStaged, whitespaceDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

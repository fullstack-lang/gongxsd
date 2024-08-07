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
var __Enumeration__dummysDeclaration__ models.Enumeration
var __Enumeration_time__dummyDeclaration time.Duration

var mutexEnumeration sync.Mutex

// An EnumerationID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEnumeration updateEnumeration deleteEnumeration
type EnumerationID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// EnumerationInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEnumeration updateEnumeration
type EnumerationInput struct {
	// The Enumeration to submit or modify
	// in: body
	Enumeration *orm.EnumerationAPI
}

// GetEnumerations
//
// swagger:route GET /enumerations enumerations getEnumerations
//
// # Get all enumerations
//
// Responses:
// default: genericError
//
//	200: enumerationDBResponse
func (controller *Controller) GetEnumerations(c *gin.Context) {

	// source slice
	var enumerationDBs []orm.EnumerationDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEnumerations", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEnumeration.GetDB()

	query := db.Find(&enumerationDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	enumerationAPIs := make([]orm.EnumerationAPI, 0)

	// for each enumeration, update fields from the database nullable fields
	for idx := range enumerationDBs {
		enumerationDB := &enumerationDBs[idx]
		_ = enumerationDB
		var enumerationAPI orm.EnumerationAPI

		// insertion point for updating fields
		enumerationAPI.ID = enumerationDB.ID
		enumerationDB.CopyBasicFieldsToEnumeration_WOP(&enumerationAPI.Enumeration_WOP)
		enumerationAPI.EnumerationPointersEncoding = enumerationDB.EnumerationPointersEncoding
		enumerationAPIs = append(enumerationAPIs, enumerationAPI)
	}

	c.JSON(http.StatusOK, enumerationAPIs)
}

// PostEnumeration
//
// swagger:route POST /enumerations enumerations postEnumeration
//
// Creates a enumeration
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEnumeration(c *gin.Context) {

	mutexEnumeration.Lock()
	defer mutexEnumeration.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEnumerations", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEnumeration.GetDB()

	// Validate input
	var input orm.EnumerationAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create enumeration
	enumerationDB := orm.EnumerationDB{}
	enumerationDB.EnumerationPointersEncoding = input.EnumerationPointersEncoding
	enumerationDB.CopyBasicFieldsFromEnumeration_WOP(&input.Enumeration_WOP)

	query := db.Create(&enumerationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEnumeration.CheckoutPhaseOneInstance(&enumerationDB)
	enumeration := backRepo.BackRepoEnumeration.Map_EnumerationDBID_EnumerationPtr[enumerationDB.ID]

	if enumeration != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), enumeration)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, enumerationDB)
}

// GetEnumeration
//
// swagger:route GET /enumerations/{ID} enumerations getEnumeration
//
// Gets the details for a enumeration.
//
// Responses:
// default: genericError
//
//	200: enumerationDBResponse
func (controller *Controller) GetEnumeration(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEnumeration", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEnumeration.GetDB()

	// Get enumerationDB in DB
	var enumerationDB orm.EnumerationDB
	if err := db.First(&enumerationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var enumerationAPI orm.EnumerationAPI
	enumerationAPI.ID = enumerationDB.ID
	enumerationAPI.EnumerationPointersEncoding = enumerationDB.EnumerationPointersEncoding
	enumerationDB.CopyBasicFieldsToEnumeration_WOP(&enumerationAPI.Enumeration_WOP)

	c.JSON(http.StatusOK, enumerationAPI)
}

// UpdateEnumeration
//
// swagger:route PATCH /enumerations/{ID} enumerations updateEnumeration
//
// # Update a enumeration
//
// Responses:
// default: genericError
//
//	200: enumerationDBResponse
func (controller *Controller) UpdateEnumeration(c *gin.Context) {

	mutexEnumeration.Lock()
	defer mutexEnumeration.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateEnumeration", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEnumeration.GetDB()

	// Validate input
	var input orm.EnumerationAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var enumerationDB orm.EnumerationDB

	// fetch the enumeration
	query := db.First(&enumerationDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	enumerationDB.CopyBasicFieldsFromEnumeration_WOP(&input.Enumeration_WOP)
	enumerationDB.EnumerationPointersEncoding = input.EnumerationPointersEncoding

	query = db.Model(&enumerationDB).Updates(enumerationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	enumerationNew := new(models.Enumeration)
	enumerationDB.CopyBasicFieldsToEnumeration(enumerationNew)

	// redeem pointers
	enumerationDB.DecodePointers(backRepo, enumerationNew)

	// get stage instance from DB instance, and call callback function
	enumerationOld := backRepo.BackRepoEnumeration.Map_EnumerationDBID_EnumerationPtr[enumerationDB.ID]
	if enumerationOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), enumerationOld, enumerationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the enumerationDB
	c.JSON(http.StatusOK, enumerationDB)
}

// DeleteEnumeration
//
// swagger:route DELETE /enumerations/{ID} enumerations deleteEnumeration
//
// # Delete a enumeration
//
// default: genericError
//
//	200: enumerationDBResponse
func (controller *Controller) DeleteEnumeration(c *gin.Context) {

	mutexEnumeration.Lock()
	defer mutexEnumeration.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEnumeration", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEnumeration.GetDB()

	// Get model if exist
	var enumerationDB orm.EnumerationDB
	if err := db.First(&enumerationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&enumerationDB)

	// get an instance (not staged) from DB instance, and call callback function
	enumerationDeleted := new(models.Enumeration)
	enumerationDB.CopyBasicFieldsToEnumeration(enumerationDeleted)

	// get stage instance from DB instance, and call callback function
	enumerationStaged := backRepo.BackRepoEnumeration.Map_EnumerationDBID_EnumerationPtr[enumerationDB.ID]
	if enumerationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), enumerationStaged, enumerationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

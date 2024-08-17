// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongxsd/alt/go/models"
	"github.com/fullstack-lang/gongxsd/alt/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __ComplexType__dummysDeclaration__ models.ComplexType
var __ComplexType_time__dummyDeclaration time.Duration

var mutexComplexType sync.Mutex

// An ComplexTypeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getComplexType updateComplexType deleteComplexType
type ComplexTypeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ComplexTypeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postComplexType updateComplexType
type ComplexTypeInput struct {
	// The ComplexType to submit or modify
	// in: body
	ComplexType *orm.ComplexTypeAPI
}

// GetComplexTypes
//
// swagger:route GET /complextypes complextypes getComplexTypes
//
// # Get all complextypes
//
// Responses:
// default: genericError
//
//	200: complextypeDBResponse
func (controller *Controller) GetComplexTypes(c *gin.Context) {

	// source slice
	var complextypeDBs []orm.ComplexTypeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetComplexTypes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/alt/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoComplexType.GetDB()

	query := db.Find(&complextypeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	complextypeAPIs := make([]orm.ComplexTypeAPI, 0)

	// for each complextype, update fields from the database nullable fields
	for idx := range complextypeDBs {
		complextypeDB := &complextypeDBs[idx]
		_ = complextypeDB
		var complextypeAPI orm.ComplexTypeAPI

		// insertion point for updating fields
		complextypeAPI.ID = complextypeDB.ID
		complextypeDB.CopyBasicFieldsToComplexType_WOP(&complextypeAPI.ComplexType_WOP)
		complextypeAPI.ComplexTypePointersEncoding = complextypeDB.ComplexTypePointersEncoding
		complextypeAPIs = append(complextypeAPIs, complextypeAPI)
	}

	c.JSON(http.StatusOK, complextypeAPIs)
}

// PostComplexType
//
// swagger:route POST /complextypes complextypes postComplexType
//
// Creates a complextype
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostComplexType(c *gin.Context) {

	mutexComplexType.Lock()
	defer mutexComplexType.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostComplexTypes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/alt/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoComplexType.GetDB()

	// Validate input
	var input orm.ComplexTypeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create complextype
	complextypeDB := orm.ComplexTypeDB{}
	complextypeDB.ComplexTypePointersEncoding = input.ComplexTypePointersEncoding
	complextypeDB.CopyBasicFieldsFromComplexType_WOP(&input.ComplexType_WOP)

	query := db.Create(&complextypeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoComplexType.CheckoutPhaseOneInstance(&complextypeDB)
	complextype := backRepo.BackRepoComplexType.Map_ComplexTypeDBID_ComplexTypePtr[complextypeDB.ID]

	if complextype != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), complextype)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, complextypeDB)
}

// GetComplexType
//
// swagger:route GET /complextypes/{ID} complextypes getComplexType
//
// Gets the details for a complextype.
//
// Responses:
// default: genericError
//
//	200: complextypeDBResponse
func (controller *Controller) GetComplexType(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetComplexType", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/alt/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoComplexType.GetDB()

	// Get complextypeDB in DB
	var complextypeDB orm.ComplexTypeDB
	if err := db.First(&complextypeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var complextypeAPI orm.ComplexTypeAPI
	complextypeAPI.ID = complextypeDB.ID
	complextypeAPI.ComplexTypePointersEncoding = complextypeDB.ComplexTypePointersEncoding
	complextypeDB.CopyBasicFieldsToComplexType_WOP(&complextypeAPI.ComplexType_WOP)

	c.JSON(http.StatusOK, complextypeAPI)
}

// UpdateComplexType
//
// swagger:route PATCH /complextypes/{ID} complextypes updateComplexType
//
// # Update a complextype
//
// Responses:
// default: genericError
//
//	200: complextypeDBResponse
func (controller *Controller) UpdateComplexType(c *gin.Context) {

	mutexComplexType.Lock()
	defer mutexComplexType.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateComplexType", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/alt/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoComplexType.GetDB()

	// Validate input
	var input orm.ComplexTypeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var complextypeDB orm.ComplexTypeDB

	// fetch the complextype
	query := db.First(&complextypeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	complextypeDB.CopyBasicFieldsFromComplexType_WOP(&input.ComplexType_WOP)
	complextypeDB.ComplexTypePointersEncoding = input.ComplexTypePointersEncoding

	query = db.Model(&complextypeDB).Updates(complextypeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	complextypeNew := new(models.ComplexType)
	complextypeDB.CopyBasicFieldsToComplexType(complextypeNew)

	// redeem pointers
	complextypeDB.DecodePointers(backRepo, complextypeNew)

	// get stage instance from DB instance, and call callback function
	complextypeOld := backRepo.BackRepoComplexType.Map_ComplexTypeDBID_ComplexTypePtr[complextypeDB.ID]
	if complextypeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), complextypeOld, complextypeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the complextypeDB
	c.JSON(http.StatusOK, complextypeDB)
}

// DeleteComplexType
//
// swagger:route DELETE /complextypes/{ID} complextypes deleteComplexType
//
// # Delete a complextype
//
// default: genericError
//
//	200: complextypeDBResponse
func (controller *Controller) DeleteComplexType(c *gin.Context) {

	mutexComplexType.Lock()
	defer mutexComplexType.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteComplexType", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/alt/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoComplexType.GetDB()

	// Get model if exist
	var complextypeDB orm.ComplexTypeDB
	if err := db.First(&complextypeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&complextypeDB)

	// get an instance (not staged) from DB instance, and call callback function
	complextypeDeleted := new(models.ComplexType)
	complextypeDB.CopyBasicFieldsToComplexType(complextypeDeleted)

	// get stage instance from DB instance, and call callback function
	complextypeStaged := backRepo.BackRepoComplexType.Map_ComplexTypeDBID_ComplexTypePtr[complextypeDB.ID]
	if complextypeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), complextypeStaged, complextypeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

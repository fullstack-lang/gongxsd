// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongxsd/test/reqif/go/models"
	"github.com/fullstack-lang/gongxsd/test/reqif/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __ATTRIBUTE_DEFINITION_ENUMERATION__dummysDeclaration__ models.ATTRIBUTE_DEFINITION_ENUMERATION
var __ATTRIBUTE_DEFINITION_ENUMERATION_time__dummyDeclaration time.Duration

var mutexATTRIBUTE_DEFINITION_ENUMERATION sync.Mutex

// An ATTRIBUTE_DEFINITION_ENUMERATIONID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTE_DEFINITION_ENUMERATION updateATTRIBUTE_DEFINITION_ENUMERATION deleteATTRIBUTE_DEFINITION_ENUMERATION
type ATTRIBUTE_DEFINITION_ENUMERATIONID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTE_DEFINITION_ENUMERATIONInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTE_DEFINITION_ENUMERATION updateATTRIBUTE_DEFINITION_ENUMERATION
type ATTRIBUTE_DEFINITION_ENUMERATIONInput struct {
	// The ATTRIBUTE_DEFINITION_ENUMERATION to submit or modify
	// in: body
	ATTRIBUTE_DEFINITION_ENUMERATION *orm.ATTRIBUTE_DEFINITION_ENUMERATIONAPI
}

// GetATTRIBUTE_DEFINITION_ENUMERATIONs
//
// swagger:route GET /attribute_definition_enumerations attribute_definition_enumerations getATTRIBUTE_DEFINITION_ENUMERATIONs
//
// # Get all attribute_definition_enumerations
//
// Responses:
// default: genericError
//
//	200: attribute_definition_enumerationDBResponse
func (controller *Controller) GetATTRIBUTE_DEFINITION_ENUMERATIONs(c *gin.Context) {

	// source slice
	var attribute_definition_enumerationDBs []orm.ATTRIBUTE_DEFINITION_ENUMERATIONDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_DEFINITION_ENUMERATIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.GetDB()

	query := db.Find(&attribute_definition_enumerationDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attribute_definition_enumerationAPIs := make([]orm.ATTRIBUTE_DEFINITION_ENUMERATIONAPI, 0)

	// for each attribute_definition_enumeration, update fields from the database nullable fields
	for idx := range attribute_definition_enumerationDBs {
		attribute_definition_enumerationDB := &attribute_definition_enumerationDBs[idx]
		_ = attribute_definition_enumerationDB
		var attribute_definition_enumerationAPI orm.ATTRIBUTE_DEFINITION_ENUMERATIONAPI

		// insertion point for updating fields
		attribute_definition_enumerationAPI.ID = attribute_definition_enumerationDB.ID
		attribute_definition_enumerationDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_ENUMERATION_WOP(&attribute_definition_enumerationAPI.ATTRIBUTE_DEFINITION_ENUMERATION_WOP)
		attribute_definition_enumerationAPI.ATTRIBUTE_DEFINITION_ENUMERATIONPointersEncoding = attribute_definition_enumerationDB.ATTRIBUTE_DEFINITION_ENUMERATIONPointersEncoding
		attribute_definition_enumerationAPIs = append(attribute_definition_enumerationAPIs, attribute_definition_enumerationAPI)
	}

	c.JSON(http.StatusOK, attribute_definition_enumerationAPIs)
}

// PostATTRIBUTE_DEFINITION_ENUMERATION
//
// swagger:route POST /attribute_definition_enumerations attribute_definition_enumerations postATTRIBUTE_DEFINITION_ENUMERATION
//
// Creates a attribute_definition_enumeration
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTE_DEFINITION_ENUMERATION(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_ENUMERATION.Lock()
	defer mutexATTRIBUTE_DEFINITION_ENUMERATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTE_DEFINITION_ENUMERATIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_DEFINITION_ENUMERATIONAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attribute_definition_enumeration
	attribute_definition_enumerationDB := orm.ATTRIBUTE_DEFINITION_ENUMERATIONDB{}
	attribute_definition_enumerationDB.ATTRIBUTE_DEFINITION_ENUMERATIONPointersEncoding = input.ATTRIBUTE_DEFINITION_ENUMERATIONPointersEncoding
	attribute_definition_enumerationDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_ENUMERATION_WOP(&input.ATTRIBUTE_DEFINITION_ENUMERATION_WOP)

	query := db.Create(&attribute_definition_enumerationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.CheckoutPhaseOneInstance(&attribute_definition_enumerationDB)
	attribute_definition_enumeration := backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.Map_ATTRIBUTE_DEFINITION_ENUMERATIONDBID_ATTRIBUTE_DEFINITION_ENUMERATIONPtr[attribute_definition_enumerationDB.ID]

	if attribute_definition_enumeration != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attribute_definition_enumeration)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attribute_definition_enumerationDB)
}

// GetATTRIBUTE_DEFINITION_ENUMERATION
//
// swagger:route GET /attribute_definition_enumerations/{ID} attribute_definition_enumerations getATTRIBUTE_DEFINITION_ENUMERATION
//
// Gets the details for a attribute_definition_enumeration.
//
// Responses:
// default: genericError
//
//	200: attribute_definition_enumerationDBResponse
func (controller *Controller) GetATTRIBUTE_DEFINITION_ENUMERATION(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_DEFINITION_ENUMERATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.GetDB()

	// Get attribute_definition_enumerationDB in DB
	var attribute_definition_enumerationDB orm.ATTRIBUTE_DEFINITION_ENUMERATIONDB
	if err := db.First(&attribute_definition_enumerationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attribute_definition_enumerationAPI orm.ATTRIBUTE_DEFINITION_ENUMERATIONAPI
	attribute_definition_enumerationAPI.ID = attribute_definition_enumerationDB.ID
	attribute_definition_enumerationAPI.ATTRIBUTE_DEFINITION_ENUMERATIONPointersEncoding = attribute_definition_enumerationDB.ATTRIBUTE_DEFINITION_ENUMERATIONPointersEncoding
	attribute_definition_enumerationDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_ENUMERATION_WOP(&attribute_definition_enumerationAPI.ATTRIBUTE_DEFINITION_ENUMERATION_WOP)

	c.JSON(http.StatusOK, attribute_definition_enumerationAPI)
}

// UpdateATTRIBUTE_DEFINITION_ENUMERATION
//
// swagger:route PATCH /attribute_definition_enumerations/{ID} attribute_definition_enumerations updateATTRIBUTE_DEFINITION_ENUMERATION
//
// # Update a attribute_definition_enumeration
//
// Responses:
// default: genericError
//
//	200: attribute_definition_enumerationDBResponse
func (controller *Controller) UpdateATTRIBUTE_DEFINITION_ENUMERATION(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_ENUMERATION.Lock()
	defer mutexATTRIBUTE_DEFINITION_ENUMERATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTE_DEFINITION_ENUMERATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_DEFINITION_ENUMERATIONAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attribute_definition_enumerationDB orm.ATTRIBUTE_DEFINITION_ENUMERATIONDB

	// fetch the attribute_definition_enumeration
	query := db.First(&attribute_definition_enumerationDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attribute_definition_enumerationDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_ENUMERATION_WOP(&input.ATTRIBUTE_DEFINITION_ENUMERATION_WOP)
	attribute_definition_enumerationDB.ATTRIBUTE_DEFINITION_ENUMERATIONPointersEncoding = input.ATTRIBUTE_DEFINITION_ENUMERATIONPointersEncoding

	query = db.Model(&attribute_definition_enumerationDB).Updates(attribute_definition_enumerationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attribute_definition_enumerationNew := new(models.ATTRIBUTE_DEFINITION_ENUMERATION)
	attribute_definition_enumerationDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumerationNew)

	// redeem pointers
	attribute_definition_enumerationDB.DecodePointers(backRepo, attribute_definition_enumerationNew)

	// get stage instance from DB instance, and call callback function
	attribute_definition_enumerationOld := backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.Map_ATTRIBUTE_DEFINITION_ENUMERATIONDBID_ATTRIBUTE_DEFINITION_ENUMERATIONPtr[attribute_definition_enumerationDB.ID]
	if attribute_definition_enumerationOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attribute_definition_enumerationOld, attribute_definition_enumerationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attribute_definition_enumerationDB
	c.JSON(http.StatusOK, attribute_definition_enumerationDB)
}

// DeleteATTRIBUTE_DEFINITION_ENUMERATION
//
// swagger:route DELETE /attribute_definition_enumerations/{ID} attribute_definition_enumerations deleteATTRIBUTE_DEFINITION_ENUMERATION
//
// # Delete a attribute_definition_enumeration
//
// default: genericError
//
//	200: attribute_definition_enumerationDBResponse
func (controller *Controller) DeleteATTRIBUTE_DEFINITION_ENUMERATION(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_ENUMERATION.Lock()
	defer mutexATTRIBUTE_DEFINITION_ENUMERATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTE_DEFINITION_ENUMERATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.GetDB()

	// Get model if exist
	var attribute_definition_enumerationDB orm.ATTRIBUTE_DEFINITION_ENUMERATIONDB
	if err := db.First(&attribute_definition_enumerationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&attribute_definition_enumerationDB)

	// get an instance (not staged) from DB instance, and call callback function
	attribute_definition_enumerationDeleted := new(models.ATTRIBUTE_DEFINITION_ENUMERATION)
	attribute_definition_enumerationDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumerationDeleted)

	// get stage instance from DB instance, and call callback function
	attribute_definition_enumerationStaged := backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.Map_ATTRIBUTE_DEFINITION_ENUMERATIONDBID_ATTRIBUTE_DEFINITION_ENUMERATIONPtr[attribute_definition_enumerationDB.ID]
	if attribute_definition_enumerationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attribute_definition_enumerationStaged, attribute_definition_enumerationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

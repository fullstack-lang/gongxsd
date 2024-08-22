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
var __A_DEFINITION__dummysDeclaration__ models.A_DEFINITION
var __A_DEFINITION_time__dummyDeclaration time.Duration

var mutexA_DEFINITION sync.Mutex

// An A_DEFINITIONID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_DEFINITION updateA_DEFINITION deleteA_DEFINITION
type A_DEFINITIONID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_DEFINITIONInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_DEFINITION updateA_DEFINITION
type A_DEFINITIONInput struct {
	// The A_DEFINITION to submit or modify
	// in: body
	A_DEFINITION *orm.A_DEFINITIONAPI
}

// GetA_DEFINITIONs
//
// swagger:route GET /a_definitions a_definitions getA_DEFINITIONs
//
// # Get all a_definitions
//
// Responses:
// default: genericError
//
//	200: a_definitionDBResponse
func (controller *Controller) GetA_DEFINITIONs(c *gin.Context) {

	// source slice
	var a_definitionDBs []orm.A_DEFINITIONDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFINITIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION.GetDB()

	query := db.Find(&a_definitionDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_definitionAPIs := make([]orm.A_DEFINITIONAPI, 0)

	// for each a_definition, update fields from the database nullable fields
	for idx := range a_definitionDBs {
		a_definitionDB := &a_definitionDBs[idx]
		_ = a_definitionDB
		var a_definitionAPI orm.A_DEFINITIONAPI

		// insertion point for updating fields
		a_definitionAPI.ID = a_definitionDB.ID
		a_definitionDB.CopyBasicFieldsToA_DEFINITION_WOP(&a_definitionAPI.A_DEFINITION_WOP)
		a_definitionAPI.A_DEFINITIONPointersEncoding = a_definitionDB.A_DEFINITIONPointersEncoding
		a_definitionAPIs = append(a_definitionAPIs, a_definitionAPI)
	}

	c.JSON(http.StatusOK, a_definitionAPIs)
}

// PostA_DEFINITION
//
// swagger:route POST /a_definitions a_definitions postA_DEFINITION
//
// Creates a a_definition
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_DEFINITION(c *gin.Context) {

	mutexA_DEFINITION.Lock()
	defer mutexA_DEFINITION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_DEFINITIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION.GetDB()

	// Validate input
	var input orm.A_DEFINITIONAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_definition
	a_definitionDB := orm.A_DEFINITIONDB{}
	a_definitionDB.A_DEFINITIONPointersEncoding = input.A_DEFINITIONPointersEncoding
	a_definitionDB.CopyBasicFieldsFromA_DEFINITION_WOP(&input.A_DEFINITION_WOP)

	query := db.Create(&a_definitionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_DEFINITION.CheckoutPhaseOneInstance(&a_definitionDB)
	a_definition := backRepo.BackRepoA_DEFINITION.Map_A_DEFINITIONDBID_A_DEFINITIONPtr[a_definitionDB.ID]

	if a_definition != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_definition)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_definitionDB)
}

// GetA_DEFINITION
//
// swagger:route GET /a_definitions/{ID} a_definitions getA_DEFINITION
//
// Gets the details for a a_definition.
//
// Responses:
// default: genericError
//
//	200: a_definitionDBResponse
func (controller *Controller) GetA_DEFINITION(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFINITION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION.GetDB()

	// Get a_definitionDB in DB
	var a_definitionDB orm.A_DEFINITIONDB
	if err := db.First(&a_definitionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_definitionAPI orm.A_DEFINITIONAPI
	a_definitionAPI.ID = a_definitionDB.ID
	a_definitionAPI.A_DEFINITIONPointersEncoding = a_definitionDB.A_DEFINITIONPointersEncoding
	a_definitionDB.CopyBasicFieldsToA_DEFINITION_WOP(&a_definitionAPI.A_DEFINITION_WOP)

	c.JSON(http.StatusOK, a_definitionAPI)
}

// UpdateA_DEFINITION
//
// swagger:route PATCH /a_definitions/{ID} a_definitions updateA_DEFINITION
//
// # Update a a_definition
//
// Responses:
// default: genericError
//
//	200: a_definitionDBResponse
func (controller *Controller) UpdateA_DEFINITION(c *gin.Context) {

	mutexA_DEFINITION.Lock()
	defer mutexA_DEFINITION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_DEFINITION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION.GetDB()

	// Validate input
	var input orm.A_DEFINITIONAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_definitionDB orm.A_DEFINITIONDB

	// fetch the a_definition
	query := db.First(&a_definitionDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_definitionDB.CopyBasicFieldsFromA_DEFINITION_WOP(&input.A_DEFINITION_WOP)
	a_definitionDB.A_DEFINITIONPointersEncoding = input.A_DEFINITIONPointersEncoding

	query = db.Model(&a_definitionDB).Updates(a_definitionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_definitionNew := new(models.A_DEFINITION)
	a_definitionDB.CopyBasicFieldsToA_DEFINITION(a_definitionNew)

	// redeem pointers
	a_definitionDB.DecodePointers(backRepo, a_definitionNew)

	// get stage instance from DB instance, and call callback function
	a_definitionOld := backRepo.BackRepoA_DEFINITION.Map_A_DEFINITIONDBID_A_DEFINITIONPtr[a_definitionDB.ID]
	if a_definitionOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_definitionOld, a_definitionNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_definitionDB
	c.JSON(http.StatusOK, a_definitionDB)
}

// DeleteA_DEFINITION
//
// swagger:route DELETE /a_definitions/{ID} a_definitions deleteA_DEFINITION
//
// # Delete a a_definition
//
// default: genericError
//
//	200: a_definitionDBResponse
func (controller *Controller) DeleteA_DEFINITION(c *gin.Context) {

	mutexA_DEFINITION.Lock()
	defer mutexA_DEFINITION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_DEFINITION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION.GetDB()

	// Get model if exist
	var a_definitionDB orm.A_DEFINITIONDB
	if err := db.First(&a_definitionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_definitionDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_definitionDeleted := new(models.A_DEFINITION)
	a_definitionDB.CopyBasicFieldsToA_DEFINITION(a_definitionDeleted)

	// get stage instance from DB instance, and call callback function
	a_definitionStaged := backRepo.BackRepoA_DEFINITION.Map_A_DEFINITIONDBID_A_DEFINITIONPtr[a_definitionDB.ID]
	if a_definitionStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_definitionStaged, a_definitionDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

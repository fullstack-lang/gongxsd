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
var __A_SOURCE_SPECIFICATION__dummysDeclaration__ models.A_SOURCE_SPECIFICATION
var __A_SOURCE_SPECIFICATION_time__dummyDeclaration time.Duration

var mutexA_SOURCE_SPECIFICATION sync.Mutex

// An A_SOURCE_SPECIFICATIONID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_SOURCE_SPECIFICATION updateA_SOURCE_SPECIFICATION deleteA_SOURCE_SPECIFICATION
type A_SOURCE_SPECIFICATIONID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_SOURCE_SPECIFICATIONInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_SOURCE_SPECIFICATION updateA_SOURCE_SPECIFICATION
type A_SOURCE_SPECIFICATIONInput struct {
	// The A_SOURCE_SPECIFICATION to submit or modify
	// in: body
	A_SOURCE_SPECIFICATION *orm.A_SOURCE_SPECIFICATIONAPI
}

// GetA_SOURCE_SPECIFICATIONs
//
// swagger:route GET /a_source_specifications a_source_specifications getA_SOURCE_SPECIFICATIONs
//
// # Get all a_source_specifications
//
// Responses:
// default: genericError
//
//	200: a_source_specificationDBResponse
func (controller *Controller) GetA_SOURCE_SPECIFICATIONs(c *gin.Context) {

	// source slice
	var a_source_specificationDBs []orm.A_SOURCE_SPECIFICATIONDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SOURCE_SPECIFICATIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_SOURCE_SPECIFICATION.GetDB()

	query := db.Find(&a_source_specificationDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_source_specificationAPIs := make([]orm.A_SOURCE_SPECIFICATIONAPI, 0)

	// for each a_source_specification, update fields from the database nullable fields
	for idx := range a_source_specificationDBs {
		a_source_specificationDB := &a_source_specificationDBs[idx]
		_ = a_source_specificationDB
		var a_source_specificationAPI orm.A_SOURCE_SPECIFICATIONAPI

		// insertion point for updating fields
		a_source_specificationAPI.ID = a_source_specificationDB.ID
		a_source_specificationDB.CopyBasicFieldsToA_SOURCE_SPECIFICATION_WOP(&a_source_specificationAPI.A_SOURCE_SPECIFICATION_WOP)
		a_source_specificationAPI.A_SOURCE_SPECIFICATIONPointersEncoding = a_source_specificationDB.A_SOURCE_SPECIFICATIONPointersEncoding
		a_source_specificationAPIs = append(a_source_specificationAPIs, a_source_specificationAPI)
	}

	c.JSON(http.StatusOK, a_source_specificationAPIs)
}

// PostA_SOURCE_SPECIFICATION
//
// swagger:route POST /a_source_specifications a_source_specifications postA_SOURCE_SPECIFICATION
//
// Creates a a_source_specification
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_SOURCE_SPECIFICATION(c *gin.Context) {

	mutexA_SOURCE_SPECIFICATION.Lock()
	defer mutexA_SOURCE_SPECIFICATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_SOURCE_SPECIFICATIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_SOURCE_SPECIFICATION.GetDB()

	// Validate input
	var input orm.A_SOURCE_SPECIFICATIONAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_source_specification
	a_source_specificationDB := orm.A_SOURCE_SPECIFICATIONDB{}
	a_source_specificationDB.A_SOURCE_SPECIFICATIONPointersEncoding = input.A_SOURCE_SPECIFICATIONPointersEncoding
	a_source_specificationDB.CopyBasicFieldsFromA_SOURCE_SPECIFICATION_WOP(&input.A_SOURCE_SPECIFICATION_WOP)

	query := db.Create(&a_source_specificationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_SOURCE_SPECIFICATION.CheckoutPhaseOneInstance(&a_source_specificationDB)
	a_source_specification := backRepo.BackRepoA_SOURCE_SPECIFICATION.Map_A_SOURCE_SPECIFICATIONDBID_A_SOURCE_SPECIFICATIONPtr[a_source_specificationDB.ID]

	if a_source_specification != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_source_specification)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_source_specificationDB)
}

// GetA_SOURCE_SPECIFICATION
//
// swagger:route GET /a_source_specifications/{ID} a_source_specifications getA_SOURCE_SPECIFICATION
//
// Gets the details for a a_source_specification.
//
// Responses:
// default: genericError
//
//	200: a_source_specificationDBResponse
func (controller *Controller) GetA_SOURCE_SPECIFICATION(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SOURCE_SPECIFICATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_SOURCE_SPECIFICATION.GetDB()

	// Get a_source_specificationDB in DB
	var a_source_specificationDB orm.A_SOURCE_SPECIFICATIONDB
	if err := db.First(&a_source_specificationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_source_specificationAPI orm.A_SOURCE_SPECIFICATIONAPI
	a_source_specificationAPI.ID = a_source_specificationDB.ID
	a_source_specificationAPI.A_SOURCE_SPECIFICATIONPointersEncoding = a_source_specificationDB.A_SOURCE_SPECIFICATIONPointersEncoding
	a_source_specificationDB.CopyBasicFieldsToA_SOURCE_SPECIFICATION_WOP(&a_source_specificationAPI.A_SOURCE_SPECIFICATION_WOP)

	c.JSON(http.StatusOK, a_source_specificationAPI)
}

// UpdateA_SOURCE_SPECIFICATION
//
// swagger:route PATCH /a_source_specifications/{ID} a_source_specifications updateA_SOURCE_SPECIFICATION
//
// # Update a a_source_specification
//
// Responses:
// default: genericError
//
//	200: a_source_specificationDBResponse
func (controller *Controller) UpdateA_SOURCE_SPECIFICATION(c *gin.Context) {

	mutexA_SOURCE_SPECIFICATION.Lock()
	defer mutexA_SOURCE_SPECIFICATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_SOURCE_SPECIFICATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_SOURCE_SPECIFICATION.GetDB()

	// Validate input
	var input orm.A_SOURCE_SPECIFICATIONAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_source_specificationDB orm.A_SOURCE_SPECIFICATIONDB

	// fetch the a_source_specification
	query := db.First(&a_source_specificationDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_source_specificationDB.CopyBasicFieldsFromA_SOURCE_SPECIFICATION_WOP(&input.A_SOURCE_SPECIFICATION_WOP)
	a_source_specificationDB.A_SOURCE_SPECIFICATIONPointersEncoding = input.A_SOURCE_SPECIFICATIONPointersEncoding

	query = db.Model(&a_source_specificationDB).Updates(a_source_specificationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_source_specificationNew := new(models.A_SOURCE_SPECIFICATION)
	a_source_specificationDB.CopyBasicFieldsToA_SOURCE_SPECIFICATION(a_source_specificationNew)

	// redeem pointers
	a_source_specificationDB.DecodePointers(backRepo, a_source_specificationNew)

	// get stage instance from DB instance, and call callback function
	a_source_specificationOld := backRepo.BackRepoA_SOURCE_SPECIFICATION.Map_A_SOURCE_SPECIFICATIONDBID_A_SOURCE_SPECIFICATIONPtr[a_source_specificationDB.ID]
	if a_source_specificationOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_source_specificationOld, a_source_specificationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_source_specificationDB
	c.JSON(http.StatusOK, a_source_specificationDB)
}

// DeleteA_SOURCE_SPECIFICATION
//
// swagger:route DELETE /a_source_specifications/{ID} a_source_specifications deleteA_SOURCE_SPECIFICATION
//
// # Delete a a_source_specification
//
// default: genericError
//
//	200: a_source_specificationDBResponse
func (controller *Controller) DeleteA_SOURCE_SPECIFICATION(c *gin.Context) {

	mutexA_SOURCE_SPECIFICATION.Lock()
	defer mutexA_SOURCE_SPECIFICATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_SOURCE_SPECIFICATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_SOURCE_SPECIFICATION.GetDB()

	// Get model if exist
	var a_source_specificationDB orm.A_SOURCE_SPECIFICATIONDB
	if err := db.First(&a_source_specificationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_source_specificationDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_source_specificationDeleted := new(models.A_SOURCE_SPECIFICATION)
	a_source_specificationDB.CopyBasicFieldsToA_SOURCE_SPECIFICATION(a_source_specificationDeleted)

	// get stage instance from DB instance, and call callback function
	a_source_specificationStaged := backRepo.BackRepoA_SOURCE_SPECIFICATION.Map_A_SOURCE_SPECIFICATIONDBID_A_SOURCE_SPECIFICATIONPtr[a_source_specificationDB.ID]
	if a_source_specificationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_source_specificationStaged, a_source_specificationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

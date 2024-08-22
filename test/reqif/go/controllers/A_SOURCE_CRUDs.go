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
var __A_SOURCE__dummysDeclaration__ models.A_SOURCE
var __A_SOURCE_time__dummyDeclaration time.Duration

var mutexA_SOURCE sync.Mutex

// An A_SOURCEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_SOURCE updateA_SOURCE deleteA_SOURCE
type A_SOURCEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_SOURCEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_SOURCE updateA_SOURCE
type A_SOURCEInput struct {
	// The A_SOURCE to submit or modify
	// in: body
	A_SOURCE *orm.A_SOURCEAPI
}

// GetA_SOURCEs
//
// swagger:route GET /a_sources a_sources getA_SOURCEs
//
// # Get all a_sources
//
// Responses:
// default: genericError
//
//	200: a_sourceDBResponse
func (controller *Controller) GetA_SOURCEs(c *gin.Context) {

	// source slice
	var a_sourceDBs []orm.A_SOURCEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SOURCEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_SOURCE.GetDB()

	query := db.Find(&a_sourceDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_sourceAPIs := make([]orm.A_SOURCEAPI, 0)

	// for each a_source, update fields from the database nullable fields
	for idx := range a_sourceDBs {
		a_sourceDB := &a_sourceDBs[idx]
		_ = a_sourceDB
		var a_sourceAPI orm.A_SOURCEAPI

		// insertion point for updating fields
		a_sourceAPI.ID = a_sourceDB.ID
		a_sourceDB.CopyBasicFieldsToA_SOURCE_WOP(&a_sourceAPI.A_SOURCE_WOP)
		a_sourceAPI.A_SOURCEPointersEncoding = a_sourceDB.A_SOURCEPointersEncoding
		a_sourceAPIs = append(a_sourceAPIs, a_sourceAPI)
	}

	c.JSON(http.StatusOK, a_sourceAPIs)
}

// PostA_SOURCE
//
// swagger:route POST /a_sources a_sources postA_SOURCE
//
// Creates a a_source
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_SOURCE(c *gin.Context) {

	mutexA_SOURCE.Lock()
	defer mutexA_SOURCE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_SOURCEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_SOURCE.GetDB()

	// Validate input
	var input orm.A_SOURCEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_source
	a_sourceDB := orm.A_SOURCEDB{}
	a_sourceDB.A_SOURCEPointersEncoding = input.A_SOURCEPointersEncoding
	a_sourceDB.CopyBasicFieldsFromA_SOURCE_WOP(&input.A_SOURCE_WOP)

	query := db.Create(&a_sourceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_SOURCE.CheckoutPhaseOneInstance(&a_sourceDB)
	a_source := backRepo.BackRepoA_SOURCE.Map_A_SOURCEDBID_A_SOURCEPtr[a_sourceDB.ID]

	if a_source != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_source)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_sourceDB)
}

// GetA_SOURCE
//
// swagger:route GET /a_sources/{ID} a_sources getA_SOURCE
//
// Gets the details for a a_source.
//
// Responses:
// default: genericError
//
//	200: a_sourceDBResponse
func (controller *Controller) GetA_SOURCE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SOURCE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_SOURCE.GetDB()

	// Get a_sourceDB in DB
	var a_sourceDB orm.A_SOURCEDB
	if err := db.First(&a_sourceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_sourceAPI orm.A_SOURCEAPI
	a_sourceAPI.ID = a_sourceDB.ID
	a_sourceAPI.A_SOURCEPointersEncoding = a_sourceDB.A_SOURCEPointersEncoding
	a_sourceDB.CopyBasicFieldsToA_SOURCE_WOP(&a_sourceAPI.A_SOURCE_WOP)

	c.JSON(http.StatusOK, a_sourceAPI)
}

// UpdateA_SOURCE
//
// swagger:route PATCH /a_sources/{ID} a_sources updateA_SOURCE
//
// # Update a a_source
//
// Responses:
// default: genericError
//
//	200: a_sourceDBResponse
func (controller *Controller) UpdateA_SOURCE(c *gin.Context) {

	mutexA_SOURCE.Lock()
	defer mutexA_SOURCE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_SOURCE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_SOURCE.GetDB()

	// Validate input
	var input orm.A_SOURCEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_sourceDB orm.A_SOURCEDB

	// fetch the a_source
	query := db.First(&a_sourceDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_sourceDB.CopyBasicFieldsFromA_SOURCE_WOP(&input.A_SOURCE_WOP)
	a_sourceDB.A_SOURCEPointersEncoding = input.A_SOURCEPointersEncoding

	query = db.Model(&a_sourceDB).Updates(a_sourceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_sourceNew := new(models.A_SOURCE)
	a_sourceDB.CopyBasicFieldsToA_SOURCE(a_sourceNew)

	// redeem pointers
	a_sourceDB.DecodePointers(backRepo, a_sourceNew)

	// get stage instance from DB instance, and call callback function
	a_sourceOld := backRepo.BackRepoA_SOURCE.Map_A_SOURCEDBID_A_SOURCEPtr[a_sourceDB.ID]
	if a_sourceOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_sourceOld, a_sourceNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_sourceDB
	c.JSON(http.StatusOK, a_sourceDB)
}

// DeleteA_SOURCE
//
// swagger:route DELETE /a_sources/{ID} a_sources deleteA_SOURCE
//
// # Delete a a_source
//
// default: genericError
//
//	200: a_sourceDBResponse
func (controller *Controller) DeleteA_SOURCE(c *gin.Context) {

	mutexA_SOURCE.Lock()
	defer mutexA_SOURCE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_SOURCE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_SOURCE.GetDB()

	// Get model if exist
	var a_sourceDB orm.A_SOURCEDB
	if err := db.First(&a_sourceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_sourceDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_sourceDeleted := new(models.A_SOURCE)
	a_sourceDB.CopyBasicFieldsToA_SOURCE(a_sourceDeleted)

	// get stage instance from DB instance, and call callback function
	a_sourceStaged := backRepo.BackRepoA_SOURCE.Map_A_SOURCEDBID_A_SOURCEPtr[a_sourceDB.ID]
	if a_sourceStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_sourceStaged, a_sourceDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

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
var __Schema__dummysDeclaration__ models.Schema
var __Schema_time__dummyDeclaration time.Duration

var mutexSchema sync.Mutex

// An SchemaID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSchema updateSchema deleteSchema
type SchemaID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SchemaInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSchema updateSchema
type SchemaInput struct {
	// The Schema to submit or modify
	// in: body
	Schema *orm.SchemaAPI
}

// GetSchemas
//
// swagger:route GET /schemas schemas getSchemas
//
// # Get all schemas
//
// Responses:
// default: genericError
//
//	200: schemaDBResponse
func (controller *Controller) GetSchemas(c *gin.Context) {

	// source slice
	var schemaDBs []orm.SchemaDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSchemas", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSchema.GetDB()

	_, err := db.Find(&schemaDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	schemaAPIs := make([]orm.SchemaAPI, 0)

	// for each schema, update fields from the database nullable fields
	for idx := range schemaDBs {
		schemaDB := &schemaDBs[idx]
		_ = schemaDB
		var schemaAPI orm.SchemaAPI

		// insertion point for updating fields
		schemaAPI.ID = schemaDB.ID
		schemaDB.CopyBasicFieldsToSchema_WOP(&schemaAPI.Schema_WOP)
		schemaAPI.SchemaPointersEncoding = schemaDB.SchemaPointersEncoding
		schemaAPIs = append(schemaAPIs, schemaAPI)
	}

	c.JSON(http.StatusOK, schemaAPIs)
}

// PostSchema
//
// swagger:route POST /schemas schemas postSchema
//
// Creates a schema
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSchema(c *gin.Context) {

	mutexSchema.Lock()
	defer mutexSchema.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSchemas", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSchema.GetDB()

	// Validate input
	var input orm.SchemaAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create schema
	schemaDB := orm.SchemaDB{}
	schemaDB.SchemaPointersEncoding = input.SchemaPointersEncoding
	schemaDB.CopyBasicFieldsFromSchema_WOP(&input.Schema_WOP)

	_, err = db.Create(&schemaDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSchema.CheckoutPhaseOneInstance(&schemaDB)
	schema := backRepo.BackRepoSchema.Map_SchemaDBID_SchemaPtr[schemaDB.ID]

	if schema != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), schema)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, schemaDB)
}

// GetSchema
//
// swagger:route GET /schemas/{ID} schemas getSchema
//
// Gets the details for a schema.
//
// Responses:
// default: genericError
//
//	200: schemaDBResponse
func (controller *Controller) GetSchema(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSchema", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSchema.GetDB()

	// Get schemaDB in DB
	var schemaDB orm.SchemaDB
	if _, err := db.First(&schemaDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var schemaAPI orm.SchemaAPI
	schemaAPI.ID = schemaDB.ID
	schemaAPI.SchemaPointersEncoding = schemaDB.SchemaPointersEncoding
	schemaDB.CopyBasicFieldsToSchema_WOP(&schemaAPI.Schema_WOP)

	c.JSON(http.StatusOK, schemaAPI)
}

// UpdateSchema
//
// swagger:route PATCH /schemas/{ID} schemas updateSchema
//
// # Update a schema
//
// Responses:
// default: genericError
//
//	200: schemaDBResponse
func (controller *Controller) UpdateSchema(c *gin.Context) {

	mutexSchema.Lock()
	defer mutexSchema.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSchema", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSchema.GetDB()

	// Validate input
	var input orm.SchemaAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var schemaDB orm.SchemaDB

	// fetch the schema
	_, err := db.First(&schemaDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	schemaDB.CopyBasicFieldsFromSchema_WOP(&input.Schema_WOP)
	schemaDB.SchemaPointersEncoding = input.SchemaPointersEncoding

	db, _ = db.Model(&schemaDB)
	_, err = db.Updates(schemaDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	schemaNew := new(models.Schema)
	schemaDB.CopyBasicFieldsToSchema(schemaNew)

	// redeem pointers
	schemaDB.DecodePointers(backRepo, schemaNew)

	// get stage instance from DB instance, and call callback function
	schemaOld := backRepo.BackRepoSchema.Map_SchemaDBID_SchemaPtr[schemaDB.ID]
	if schemaOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), schemaOld, schemaNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the schemaDB
	c.JSON(http.StatusOK, schemaDB)
}

// DeleteSchema
//
// swagger:route DELETE /schemas/{ID} schemas deleteSchema
//
// # Delete a schema
//
// default: genericError
//
//	200: schemaDBResponse
func (controller *Controller) DeleteSchema(c *gin.Context) {

	mutexSchema.Lock()
	defer mutexSchema.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSchema", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSchema.GetDB()

	// Get model if exist
	var schemaDB orm.SchemaDB
	if _, err := db.First(&schemaDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&schemaDB)

	// get an instance (not staged) from DB instance, and call callback function
	schemaDeleted := new(models.Schema)
	schemaDB.CopyBasicFieldsToSchema(schemaDeleted)

	// get stage instance from DB instance, and call callback function
	schemaStaged := backRepo.BackRepoSchema.Map_SchemaDBID_SchemaPtr[schemaDB.ID]
	if schemaStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), schemaStaged, schemaDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

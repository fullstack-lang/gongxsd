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
var __A_DEFINITION_2__dummysDeclaration__ models.A_DEFINITION_2
var __A_DEFINITION_2_time__dummyDeclaration time.Duration

var mutexA_DEFINITION_2 sync.Mutex

// An A_DEFINITION_2ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_DEFINITION_2 updateA_DEFINITION_2 deleteA_DEFINITION_2
type A_DEFINITION_2ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_DEFINITION_2Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_DEFINITION_2 updateA_DEFINITION_2
type A_DEFINITION_2Input struct {
	// The A_DEFINITION_2 to submit or modify
	// in: body
	A_DEFINITION_2 *orm.A_DEFINITION_2API
}

// GetA_DEFINITION_2s
//
// swagger:route GET /a_definition_2s a_definition_2s getA_DEFINITION_2s
//
// # Get all a_definition_2s
//
// Responses:
// default: genericError
//
//	200: a_definition_2DBResponse
func (controller *Controller) GetA_DEFINITION_2s(c *gin.Context) {

	// source slice
	var a_definition_2DBs []orm.A_DEFINITION_2DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFINITION_2s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION_2.GetDB()

	query := db.Find(&a_definition_2DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_definition_2APIs := make([]orm.A_DEFINITION_2API, 0)

	// for each a_definition_2, update fields from the database nullable fields
	for idx := range a_definition_2DBs {
		a_definition_2DB := &a_definition_2DBs[idx]
		_ = a_definition_2DB
		var a_definition_2API orm.A_DEFINITION_2API

		// insertion point for updating fields
		a_definition_2API.ID = a_definition_2DB.ID
		a_definition_2DB.CopyBasicFieldsToA_DEFINITION_2_WOP(&a_definition_2API.A_DEFINITION_2_WOP)
		a_definition_2API.A_DEFINITION_2PointersEncoding = a_definition_2DB.A_DEFINITION_2PointersEncoding
		a_definition_2APIs = append(a_definition_2APIs, a_definition_2API)
	}

	c.JSON(http.StatusOK, a_definition_2APIs)
}

// PostA_DEFINITION_2
//
// swagger:route POST /a_definition_2s a_definition_2s postA_DEFINITION_2
//
// Creates a a_definition_2
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_DEFINITION_2(c *gin.Context) {

	mutexA_DEFINITION_2.Lock()
	defer mutexA_DEFINITION_2.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_DEFINITION_2s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION_2.GetDB()

	// Validate input
	var input orm.A_DEFINITION_2API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_definition_2
	a_definition_2DB := orm.A_DEFINITION_2DB{}
	a_definition_2DB.A_DEFINITION_2PointersEncoding = input.A_DEFINITION_2PointersEncoding
	a_definition_2DB.CopyBasicFieldsFromA_DEFINITION_2_WOP(&input.A_DEFINITION_2_WOP)

	query := db.Create(&a_definition_2DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_DEFINITION_2.CheckoutPhaseOneInstance(&a_definition_2DB)
	a_definition_2 := backRepo.BackRepoA_DEFINITION_2.Map_A_DEFINITION_2DBID_A_DEFINITION_2Ptr[a_definition_2DB.ID]

	if a_definition_2 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_definition_2)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_definition_2DB)
}

// GetA_DEFINITION_2
//
// swagger:route GET /a_definition_2s/{ID} a_definition_2s getA_DEFINITION_2
//
// Gets the details for a a_definition_2.
//
// Responses:
// default: genericError
//
//	200: a_definition_2DBResponse
func (controller *Controller) GetA_DEFINITION_2(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFINITION_2", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION_2.GetDB()

	// Get a_definition_2DB in DB
	var a_definition_2DB orm.A_DEFINITION_2DB
	if err := db.First(&a_definition_2DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_definition_2API orm.A_DEFINITION_2API
	a_definition_2API.ID = a_definition_2DB.ID
	a_definition_2API.A_DEFINITION_2PointersEncoding = a_definition_2DB.A_DEFINITION_2PointersEncoding
	a_definition_2DB.CopyBasicFieldsToA_DEFINITION_2_WOP(&a_definition_2API.A_DEFINITION_2_WOP)

	c.JSON(http.StatusOK, a_definition_2API)
}

// UpdateA_DEFINITION_2
//
// swagger:route PATCH /a_definition_2s/{ID} a_definition_2s updateA_DEFINITION_2
//
// # Update a a_definition_2
//
// Responses:
// default: genericError
//
//	200: a_definition_2DBResponse
func (controller *Controller) UpdateA_DEFINITION_2(c *gin.Context) {

	mutexA_DEFINITION_2.Lock()
	defer mutexA_DEFINITION_2.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_DEFINITION_2", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION_2.GetDB()

	// Validate input
	var input orm.A_DEFINITION_2API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_definition_2DB orm.A_DEFINITION_2DB

	// fetch the a_definition_2
	query := db.First(&a_definition_2DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_definition_2DB.CopyBasicFieldsFromA_DEFINITION_2_WOP(&input.A_DEFINITION_2_WOP)
	a_definition_2DB.A_DEFINITION_2PointersEncoding = input.A_DEFINITION_2PointersEncoding

	query = db.Model(&a_definition_2DB).Updates(a_definition_2DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_definition_2New := new(models.A_DEFINITION_2)
	a_definition_2DB.CopyBasicFieldsToA_DEFINITION_2(a_definition_2New)

	// redeem pointers
	a_definition_2DB.DecodePointers(backRepo, a_definition_2New)

	// get stage instance from DB instance, and call callback function
	a_definition_2Old := backRepo.BackRepoA_DEFINITION_2.Map_A_DEFINITION_2DBID_A_DEFINITION_2Ptr[a_definition_2DB.ID]
	if a_definition_2Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_definition_2Old, a_definition_2New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_definition_2DB
	c.JSON(http.StatusOK, a_definition_2DB)
}

// DeleteA_DEFINITION_2
//
// swagger:route DELETE /a_definition_2s/{ID} a_definition_2s deleteA_DEFINITION_2
//
// # Delete a a_definition_2
//
// default: genericError
//
//	200: a_definition_2DBResponse
func (controller *Controller) DeleteA_DEFINITION_2(c *gin.Context) {

	mutexA_DEFINITION_2.Lock()
	defer mutexA_DEFINITION_2.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_DEFINITION_2", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION_2.GetDB()

	// Get model if exist
	var a_definition_2DB orm.A_DEFINITION_2DB
	if err := db.First(&a_definition_2DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_definition_2DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_definition_2Deleted := new(models.A_DEFINITION_2)
	a_definition_2DB.CopyBasicFieldsToA_DEFINITION_2(a_definition_2Deleted)

	// get stage instance from DB instance, and call callback function
	a_definition_2Staged := backRepo.BackRepoA_DEFINITION_2.Map_A_DEFINITION_2DBID_A_DEFINITION_2Ptr[a_definition_2DB.ID]
	if a_definition_2Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_definition_2Staged, a_definition_2Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

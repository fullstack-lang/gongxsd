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
var __A_SPEC_RELATIONS_1__dummysDeclaration__ models.A_SPEC_RELATIONS_1
var __A_SPEC_RELATIONS_1_time__dummyDeclaration time.Duration

var mutexA_SPEC_RELATIONS_1 sync.Mutex

// An A_SPEC_RELATIONS_1ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_SPEC_RELATIONS_1 updateA_SPEC_RELATIONS_1 deleteA_SPEC_RELATIONS_1
type A_SPEC_RELATIONS_1ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_SPEC_RELATIONS_1Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_SPEC_RELATIONS_1 updateA_SPEC_RELATIONS_1
type A_SPEC_RELATIONS_1Input struct {
	// The A_SPEC_RELATIONS_1 to submit or modify
	// in: body
	A_SPEC_RELATIONS_1 *orm.A_SPEC_RELATIONS_1API
}

// GetA_SPEC_RELATIONS_1s
//
// swagger:route GET /a_spec_relations_1s a_spec_relations_1s getA_SPEC_RELATIONS_1s
//
// # Get all a_spec_relations_1s
//
// Responses:
// default: genericError
//
//	200: a_spec_relations_1DBResponse
func (controller *Controller) GetA_SPEC_RELATIONS_1s(c *gin.Context) {

	// source slice
	var a_spec_relations_1DBs []orm.A_SPEC_RELATIONS_1DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SPEC_RELATIONS_1s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_SPEC_RELATIONS_1.GetDB()

	query := db.Find(&a_spec_relations_1DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_spec_relations_1APIs := make([]orm.A_SPEC_RELATIONS_1API, 0)

	// for each a_spec_relations_1, update fields from the database nullable fields
	for idx := range a_spec_relations_1DBs {
		a_spec_relations_1DB := &a_spec_relations_1DBs[idx]
		_ = a_spec_relations_1DB
		var a_spec_relations_1API orm.A_SPEC_RELATIONS_1API

		// insertion point for updating fields
		a_spec_relations_1API.ID = a_spec_relations_1DB.ID
		a_spec_relations_1DB.CopyBasicFieldsToA_SPEC_RELATIONS_1_WOP(&a_spec_relations_1API.A_SPEC_RELATIONS_1_WOP)
		a_spec_relations_1API.A_SPEC_RELATIONS_1PointersEncoding = a_spec_relations_1DB.A_SPEC_RELATIONS_1PointersEncoding
		a_spec_relations_1APIs = append(a_spec_relations_1APIs, a_spec_relations_1API)
	}

	c.JSON(http.StatusOK, a_spec_relations_1APIs)
}

// PostA_SPEC_RELATIONS_1
//
// swagger:route POST /a_spec_relations_1s a_spec_relations_1s postA_SPEC_RELATIONS_1
//
// Creates a a_spec_relations_1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_SPEC_RELATIONS_1(c *gin.Context) {

	mutexA_SPEC_RELATIONS_1.Lock()
	defer mutexA_SPEC_RELATIONS_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_SPEC_RELATIONS_1s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_SPEC_RELATIONS_1.GetDB()

	// Validate input
	var input orm.A_SPEC_RELATIONS_1API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_spec_relations_1
	a_spec_relations_1DB := orm.A_SPEC_RELATIONS_1DB{}
	a_spec_relations_1DB.A_SPEC_RELATIONS_1PointersEncoding = input.A_SPEC_RELATIONS_1PointersEncoding
	a_spec_relations_1DB.CopyBasicFieldsFromA_SPEC_RELATIONS_1_WOP(&input.A_SPEC_RELATIONS_1_WOP)

	query := db.Create(&a_spec_relations_1DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_SPEC_RELATIONS_1.CheckoutPhaseOneInstance(&a_spec_relations_1DB)
	a_spec_relations_1 := backRepo.BackRepoA_SPEC_RELATIONS_1.Map_A_SPEC_RELATIONS_1DBID_A_SPEC_RELATIONS_1Ptr[a_spec_relations_1DB.ID]

	if a_spec_relations_1 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_spec_relations_1)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_spec_relations_1DB)
}

// GetA_SPEC_RELATIONS_1
//
// swagger:route GET /a_spec_relations_1s/{ID} a_spec_relations_1s getA_SPEC_RELATIONS_1
//
// Gets the details for a a_spec_relations_1.
//
// Responses:
// default: genericError
//
//	200: a_spec_relations_1DBResponse
func (controller *Controller) GetA_SPEC_RELATIONS_1(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SPEC_RELATIONS_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_SPEC_RELATIONS_1.GetDB()

	// Get a_spec_relations_1DB in DB
	var a_spec_relations_1DB orm.A_SPEC_RELATIONS_1DB
	if err := db.First(&a_spec_relations_1DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_spec_relations_1API orm.A_SPEC_RELATIONS_1API
	a_spec_relations_1API.ID = a_spec_relations_1DB.ID
	a_spec_relations_1API.A_SPEC_RELATIONS_1PointersEncoding = a_spec_relations_1DB.A_SPEC_RELATIONS_1PointersEncoding
	a_spec_relations_1DB.CopyBasicFieldsToA_SPEC_RELATIONS_1_WOP(&a_spec_relations_1API.A_SPEC_RELATIONS_1_WOP)

	c.JSON(http.StatusOK, a_spec_relations_1API)
}

// UpdateA_SPEC_RELATIONS_1
//
// swagger:route PATCH /a_spec_relations_1s/{ID} a_spec_relations_1s updateA_SPEC_RELATIONS_1
//
// # Update a a_spec_relations_1
//
// Responses:
// default: genericError
//
//	200: a_spec_relations_1DBResponse
func (controller *Controller) UpdateA_SPEC_RELATIONS_1(c *gin.Context) {

	mutexA_SPEC_RELATIONS_1.Lock()
	defer mutexA_SPEC_RELATIONS_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_SPEC_RELATIONS_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_SPEC_RELATIONS_1.GetDB()

	// Validate input
	var input orm.A_SPEC_RELATIONS_1API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_spec_relations_1DB orm.A_SPEC_RELATIONS_1DB

	// fetch the a_spec_relations_1
	query := db.First(&a_spec_relations_1DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_spec_relations_1DB.CopyBasicFieldsFromA_SPEC_RELATIONS_1_WOP(&input.A_SPEC_RELATIONS_1_WOP)
	a_spec_relations_1DB.A_SPEC_RELATIONS_1PointersEncoding = input.A_SPEC_RELATIONS_1PointersEncoding

	query = db.Model(&a_spec_relations_1DB).Updates(a_spec_relations_1DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_spec_relations_1New := new(models.A_SPEC_RELATIONS_1)
	a_spec_relations_1DB.CopyBasicFieldsToA_SPEC_RELATIONS_1(a_spec_relations_1New)

	// redeem pointers
	a_spec_relations_1DB.DecodePointers(backRepo, a_spec_relations_1New)

	// get stage instance from DB instance, and call callback function
	a_spec_relations_1Old := backRepo.BackRepoA_SPEC_RELATIONS_1.Map_A_SPEC_RELATIONS_1DBID_A_SPEC_RELATIONS_1Ptr[a_spec_relations_1DB.ID]
	if a_spec_relations_1Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_spec_relations_1Old, a_spec_relations_1New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_spec_relations_1DB
	c.JSON(http.StatusOK, a_spec_relations_1DB)
}

// DeleteA_SPEC_RELATIONS_1
//
// swagger:route DELETE /a_spec_relations_1s/{ID} a_spec_relations_1s deleteA_SPEC_RELATIONS_1
//
// # Delete a a_spec_relations_1
//
// default: genericError
//
//	200: a_spec_relations_1DBResponse
func (controller *Controller) DeleteA_SPEC_RELATIONS_1(c *gin.Context) {

	mutexA_SPEC_RELATIONS_1.Lock()
	defer mutexA_SPEC_RELATIONS_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_SPEC_RELATIONS_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_SPEC_RELATIONS_1.GetDB()

	// Get model if exist
	var a_spec_relations_1DB orm.A_SPEC_RELATIONS_1DB
	if err := db.First(&a_spec_relations_1DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_spec_relations_1DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_spec_relations_1Deleted := new(models.A_SPEC_RELATIONS_1)
	a_spec_relations_1DB.CopyBasicFieldsToA_SPEC_RELATIONS_1(a_spec_relations_1Deleted)

	// get stage instance from DB instance, and call callback function
	a_spec_relations_1Staged := backRepo.BackRepoA_SPEC_RELATIONS_1.Map_A_SPEC_RELATIONS_1DBID_A_SPEC_RELATIONS_1Ptr[a_spec_relations_1DB.ID]
	if a_spec_relations_1Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_spec_relations_1Staged, a_spec_relations_1Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

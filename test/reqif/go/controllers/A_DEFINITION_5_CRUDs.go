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
var __A_DEFINITION_5__dummysDeclaration__ models.A_DEFINITION_5
var __A_DEFINITION_5_time__dummyDeclaration time.Duration

var mutexA_DEFINITION_5 sync.Mutex

// An A_DEFINITION_5ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_DEFINITION_5 updateA_DEFINITION_5 deleteA_DEFINITION_5
type A_DEFINITION_5ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_DEFINITION_5Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_DEFINITION_5 updateA_DEFINITION_5
type A_DEFINITION_5Input struct {
	// The A_DEFINITION_5 to submit or modify
	// in: body
	A_DEFINITION_5 *orm.A_DEFINITION_5API
}

// GetA_DEFINITION_5s
//
// swagger:route GET /a_definition_5s a_definition_5s getA_DEFINITION_5s
//
// # Get all a_definition_5s
//
// Responses:
// default: genericError
//
//	200: a_definition_5DBResponse
func (controller *Controller) GetA_DEFINITION_5s(c *gin.Context) {

	// source slice
	var a_definition_5DBs []orm.A_DEFINITION_5DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFINITION_5s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION_5.GetDB()

	query := db.Find(&a_definition_5DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_definition_5APIs := make([]orm.A_DEFINITION_5API, 0)

	// for each a_definition_5, update fields from the database nullable fields
	for idx := range a_definition_5DBs {
		a_definition_5DB := &a_definition_5DBs[idx]
		_ = a_definition_5DB
		var a_definition_5API orm.A_DEFINITION_5API

		// insertion point for updating fields
		a_definition_5API.ID = a_definition_5DB.ID
		a_definition_5DB.CopyBasicFieldsToA_DEFINITION_5_WOP(&a_definition_5API.A_DEFINITION_5_WOP)
		a_definition_5API.A_DEFINITION_5PointersEncoding = a_definition_5DB.A_DEFINITION_5PointersEncoding
		a_definition_5APIs = append(a_definition_5APIs, a_definition_5API)
	}

	c.JSON(http.StatusOK, a_definition_5APIs)
}

// PostA_DEFINITION_5
//
// swagger:route POST /a_definition_5s a_definition_5s postA_DEFINITION_5
//
// Creates a a_definition_5
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_DEFINITION_5(c *gin.Context) {

	mutexA_DEFINITION_5.Lock()
	defer mutexA_DEFINITION_5.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_DEFINITION_5s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION_5.GetDB()

	// Validate input
	var input orm.A_DEFINITION_5API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_definition_5
	a_definition_5DB := orm.A_DEFINITION_5DB{}
	a_definition_5DB.A_DEFINITION_5PointersEncoding = input.A_DEFINITION_5PointersEncoding
	a_definition_5DB.CopyBasicFieldsFromA_DEFINITION_5_WOP(&input.A_DEFINITION_5_WOP)

	query := db.Create(&a_definition_5DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_DEFINITION_5.CheckoutPhaseOneInstance(&a_definition_5DB)
	a_definition_5 := backRepo.BackRepoA_DEFINITION_5.Map_A_DEFINITION_5DBID_A_DEFINITION_5Ptr[a_definition_5DB.ID]

	if a_definition_5 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_definition_5)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_definition_5DB)
}

// GetA_DEFINITION_5
//
// swagger:route GET /a_definition_5s/{ID} a_definition_5s getA_DEFINITION_5
//
// Gets the details for a a_definition_5.
//
// Responses:
// default: genericError
//
//	200: a_definition_5DBResponse
func (controller *Controller) GetA_DEFINITION_5(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFINITION_5", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION_5.GetDB()

	// Get a_definition_5DB in DB
	var a_definition_5DB orm.A_DEFINITION_5DB
	if err := db.First(&a_definition_5DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_definition_5API orm.A_DEFINITION_5API
	a_definition_5API.ID = a_definition_5DB.ID
	a_definition_5API.A_DEFINITION_5PointersEncoding = a_definition_5DB.A_DEFINITION_5PointersEncoding
	a_definition_5DB.CopyBasicFieldsToA_DEFINITION_5_WOP(&a_definition_5API.A_DEFINITION_5_WOP)

	c.JSON(http.StatusOK, a_definition_5API)
}

// UpdateA_DEFINITION_5
//
// swagger:route PATCH /a_definition_5s/{ID} a_definition_5s updateA_DEFINITION_5
//
// # Update a a_definition_5
//
// Responses:
// default: genericError
//
//	200: a_definition_5DBResponse
func (controller *Controller) UpdateA_DEFINITION_5(c *gin.Context) {

	mutexA_DEFINITION_5.Lock()
	defer mutexA_DEFINITION_5.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_DEFINITION_5", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION_5.GetDB()

	// Validate input
	var input orm.A_DEFINITION_5API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_definition_5DB orm.A_DEFINITION_5DB

	// fetch the a_definition_5
	query := db.First(&a_definition_5DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_definition_5DB.CopyBasicFieldsFromA_DEFINITION_5_WOP(&input.A_DEFINITION_5_WOP)
	a_definition_5DB.A_DEFINITION_5PointersEncoding = input.A_DEFINITION_5PointersEncoding

	query = db.Model(&a_definition_5DB).Updates(a_definition_5DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_definition_5New := new(models.A_DEFINITION_5)
	a_definition_5DB.CopyBasicFieldsToA_DEFINITION_5(a_definition_5New)

	// redeem pointers
	a_definition_5DB.DecodePointers(backRepo, a_definition_5New)

	// get stage instance from DB instance, and call callback function
	a_definition_5Old := backRepo.BackRepoA_DEFINITION_5.Map_A_DEFINITION_5DBID_A_DEFINITION_5Ptr[a_definition_5DB.ID]
	if a_definition_5Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_definition_5Old, a_definition_5New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_definition_5DB
	c.JSON(http.StatusOK, a_definition_5DB)
}

// DeleteA_DEFINITION_5
//
// swagger:route DELETE /a_definition_5s/{ID} a_definition_5s deleteA_DEFINITION_5
//
// # Delete a a_definition_5
//
// default: genericError
//
//	200: a_definition_5DBResponse
func (controller *Controller) DeleteA_DEFINITION_5(c *gin.Context) {

	mutexA_DEFINITION_5.Lock()
	defer mutexA_DEFINITION_5.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_DEFINITION_5", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION_5.GetDB()

	// Get model if exist
	var a_definition_5DB orm.A_DEFINITION_5DB
	if err := db.First(&a_definition_5DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_definition_5DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_definition_5Deleted := new(models.A_DEFINITION_5)
	a_definition_5DB.CopyBasicFieldsToA_DEFINITION_5(a_definition_5Deleted)

	// get stage instance from DB instance, and call callback function
	a_definition_5Staged := backRepo.BackRepoA_DEFINITION_5.Map_A_DEFINITION_5DBID_A_DEFINITION_5Ptr[a_definition_5DB.ID]
	if a_definition_5Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_definition_5Staged, a_definition_5Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

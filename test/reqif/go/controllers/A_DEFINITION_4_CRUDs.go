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
var __A_DEFINITION_4__dummysDeclaration__ models.A_DEFINITION_4
var __A_DEFINITION_4_time__dummyDeclaration time.Duration

var mutexA_DEFINITION_4 sync.Mutex

// An A_DEFINITION_4ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_DEFINITION_4 updateA_DEFINITION_4 deleteA_DEFINITION_4
type A_DEFINITION_4ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_DEFINITION_4Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_DEFINITION_4 updateA_DEFINITION_4
type A_DEFINITION_4Input struct {
	// The A_DEFINITION_4 to submit or modify
	// in: body
	A_DEFINITION_4 *orm.A_DEFINITION_4API
}

// GetA_DEFINITION_4s
//
// swagger:route GET /a_definition_4s a_definition_4s getA_DEFINITION_4s
//
// # Get all a_definition_4s
//
// Responses:
// default: genericError
//
//	200: a_definition_4DBResponse
func (controller *Controller) GetA_DEFINITION_4s(c *gin.Context) {

	// source slice
	var a_definition_4DBs []orm.A_DEFINITION_4DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFINITION_4s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION_4.GetDB()

	query := db.Find(&a_definition_4DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_definition_4APIs := make([]orm.A_DEFINITION_4API, 0)

	// for each a_definition_4, update fields from the database nullable fields
	for idx := range a_definition_4DBs {
		a_definition_4DB := &a_definition_4DBs[idx]
		_ = a_definition_4DB
		var a_definition_4API orm.A_DEFINITION_4API

		// insertion point for updating fields
		a_definition_4API.ID = a_definition_4DB.ID
		a_definition_4DB.CopyBasicFieldsToA_DEFINITION_4_WOP(&a_definition_4API.A_DEFINITION_4_WOP)
		a_definition_4API.A_DEFINITION_4PointersEncoding = a_definition_4DB.A_DEFINITION_4PointersEncoding
		a_definition_4APIs = append(a_definition_4APIs, a_definition_4API)
	}

	c.JSON(http.StatusOK, a_definition_4APIs)
}

// PostA_DEFINITION_4
//
// swagger:route POST /a_definition_4s a_definition_4s postA_DEFINITION_4
//
// Creates a a_definition_4
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_DEFINITION_4(c *gin.Context) {

	mutexA_DEFINITION_4.Lock()
	defer mutexA_DEFINITION_4.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_DEFINITION_4s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION_4.GetDB()

	// Validate input
	var input orm.A_DEFINITION_4API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_definition_4
	a_definition_4DB := orm.A_DEFINITION_4DB{}
	a_definition_4DB.A_DEFINITION_4PointersEncoding = input.A_DEFINITION_4PointersEncoding
	a_definition_4DB.CopyBasicFieldsFromA_DEFINITION_4_WOP(&input.A_DEFINITION_4_WOP)

	query := db.Create(&a_definition_4DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_DEFINITION_4.CheckoutPhaseOneInstance(&a_definition_4DB)
	a_definition_4 := backRepo.BackRepoA_DEFINITION_4.Map_A_DEFINITION_4DBID_A_DEFINITION_4Ptr[a_definition_4DB.ID]

	if a_definition_4 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_definition_4)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_definition_4DB)
}

// GetA_DEFINITION_4
//
// swagger:route GET /a_definition_4s/{ID} a_definition_4s getA_DEFINITION_4
//
// Gets the details for a a_definition_4.
//
// Responses:
// default: genericError
//
//	200: a_definition_4DBResponse
func (controller *Controller) GetA_DEFINITION_4(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFINITION_4", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION_4.GetDB()

	// Get a_definition_4DB in DB
	var a_definition_4DB orm.A_DEFINITION_4DB
	if err := db.First(&a_definition_4DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_definition_4API orm.A_DEFINITION_4API
	a_definition_4API.ID = a_definition_4DB.ID
	a_definition_4API.A_DEFINITION_4PointersEncoding = a_definition_4DB.A_DEFINITION_4PointersEncoding
	a_definition_4DB.CopyBasicFieldsToA_DEFINITION_4_WOP(&a_definition_4API.A_DEFINITION_4_WOP)

	c.JSON(http.StatusOK, a_definition_4API)
}

// UpdateA_DEFINITION_4
//
// swagger:route PATCH /a_definition_4s/{ID} a_definition_4s updateA_DEFINITION_4
//
// # Update a a_definition_4
//
// Responses:
// default: genericError
//
//	200: a_definition_4DBResponse
func (controller *Controller) UpdateA_DEFINITION_4(c *gin.Context) {

	mutexA_DEFINITION_4.Lock()
	defer mutexA_DEFINITION_4.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_DEFINITION_4", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION_4.GetDB()

	// Validate input
	var input orm.A_DEFINITION_4API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_definition_4DB orm.A_DEFINITION_4DB

	// fetch the a_definition_4
	query := db.First(&a_definition_4DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_definition_4DB.CopyBasicFieldsFromA_DEFINITION_4_WOP(&input.A_DEFINITION_4_WOP)
	a_definition_4DB.A_DEFINITION_4PointersEncoding = input.A_DEFINITION_4PointersEncoding

	query = db.Model(&a_definition_4DB).Updates(a_definition_4DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_definition_4New := new(models.A_DEFINITION_4)
	a_definition_4DB.CopyBasicFieldsToA_DEFINITION_4(a_definition_4New)

	// redeem pointers
	a_definition_4DB.DecodePointers(backRepo, a_definition_4New)

	// get stage instance from DB instance, and call callback function
	a_definition_4Old := backRepo.BackRepoA_DEFINITION_4.Map_A_DEFINITION_4DBID_A_DEFINITION_4Ptr[a_definition_4DB.ID]
	if a_definition_4Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_definition_4Old, a_definition_4New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_definition_4DB
	c.JSON(http.StatusOK, a_definition_4DB)
}

// DeleteA_DEFINITION_4
//
// swagger:route DELETE /a_definition_4s/{ID} a_definition_4s deleteA_DEFINITION_4
//
// # Delete a a_definition_4
//
// default: genericError
//
//	200: a_definition_4DBResponse
func (controller *Controller) DeleteA_DEFINITION_4(c *gin.Context) {

	mutexA_DEFINITION_4.Lock()
	defer mutexA_DEFINITION_4.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_DEFINITION_4", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION_4.GetDB()

	// Get model if exist
	var a_definition_4DB orm.A_DEFINITION_4DB
	if err := db.First(&a_definition_4DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_definition_4DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_definition_4Deleted := new(models.A_DEFINITION_4)
	a_definition_4DB.CopyBasicFieldsToA_DEFINITION_4(a_definition_4Deleted)

	// get stage instance from DB instance, and call callback function
	a_definition_4Staged := backRepo.BackRepoA_DEFINITION_4.Map_A_DEFINITION_4DBID_A_DEFINITION_4Ptr[a_definition_4DB.ID]
	if a_definition_4Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_definition_4Staged, a_definition_4Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

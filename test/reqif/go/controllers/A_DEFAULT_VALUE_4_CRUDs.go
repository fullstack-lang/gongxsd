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
var __A_DEFAULT_VALUE_4__dummysDeclaration__ models.A_DEFAULT_VALUE_4
var __A_DEFAULT_VALUE_4_time__dummyDeclaration time.Duration

var mutexA_DEFAULT_VALUE_4 sync.Mutex

// An A_DEFAULT_VALUE_4ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_DEFAULT_VALUE_4 updateA_DEFAULT_VALUE_4 deleteA_DEFAULT_VALUE_4
type A_DEFAULT_VALUE_4ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_DEFAULT_VALUE_4Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_DEFAULT_VALUE_4 updateA_DEFAULT_VALUE_4
type A_DEFAULT_VALUE_4Input struct {
	// The A_DEFAULT_VALUE_4 to submit or modify
	// in: body
	A_DEFAULT_VALUE_4 *orm.A_DEFAULT_VALUE_4API
}

// GetA_DEFAULT_VALUE_4s
//
// swagger:route GET /a_default_value_4s a_default_value_4s getA_DEFAULT_VALUE_4s
//
// # Get all a_default_value_4s
//
// Responses:
// default: genericError
//
//	200: a_default_value_4DBResponse
func (controller *Controller) GetA_DEFAULT_VALUE_4s(c *gin.Context) {

	// source slice
	var a_default_value_4DBs []orm.A_DEFAULT_VALUE_4DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFAULT_VALUE_4s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_4.GetDB()

	query := db.Find(&a_default_value_4DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_default_value_4APIs := make([]orm.A_DEFAULT_VALUE_4API, 0)

	// for each a_default_value_4, update fields from the database nullable fields
	for idx := range a_default_value_4DBs {
		a_default_value_4DB := &a_default_value_4DBs[idx]
		_ = a_default_value_4DB
		var a_default_value_4API orm.A_DEFAULT_VALUE_4API

		// insertion point for updating fields
		a_default_value_4API.ID = a_default_value_4DB.ID
		a_default_value_4DB.CopyBasicFieldsToA_DEFAULT_VALUE_4_WOP(&a_default_value_4API.A_DEFAULT_VALUE_4_WOP)
		a_default_value_4API.A_DEFAULT_VALUE_4PointersEncoding = a_default_value_4DB.A_DEFAULT_VALUE_4PointersEncoding
		a_default_value_4APIs = append(a_default_value_4APIs, a_default_value_4API)
	}

	c.JSON(http.StatusOK, a_default_value_4APIs)
}

// PostA_DEFAULT_VALUE_4
//
// swagger:route POST /a_default_value_4s a_default_value_4s postA_DEFAULT_VALUE_4
//
// Creates a a_default_value_4
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_DEFAULT_VALUE_4(c *gin.Context) {

	mutexA_DEFAULT_VALUE_4.Lock()
	defer mutexA_DEFAULT_VALUE_4.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_DEFAULT_VALUE_4s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_4.GetDB()

	// Validate input
	var input orm.A_DEFAULT_VALUE_4API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_default_value_4
	a_default_value_4DB := orm.A_DEFAULT_VALUE_4DB{}
	a_default_value_4DB.A_DEFAULT_VALUE_4PointersEncoding = input.A_DEFAULT_VALUE_4PointersEncoding
	a_default_value_4DB.CopyBasicFieldsFromA_DEFAULT_VALUE_4_WOP(&input.A_DEFAULT_VALUE_4_WOP)

	query := db.Create(&a_default_value_4DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_DEFAULT_VALUE_4.CheckoutPhaseOneInstance(&a_default_value_4DB)
	a_default_value_4 := backRepo.BackRepoA_DEFAULT_VALUE_4.Map_A_DEFAULT_VALUE_4DBID_A_DEFAULT_VALUE_4Ptr[a_default_value_4DB.ID]

	if a_default_value_4 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_default_value_4)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_default_value_4DB)
}

// GetA_DEFAULT_VALUE_4
//
// swagger:route GET /a_default_value_4s/{ID} a_default_value_4s getA_DEFAULT_VALUE_4
//
// Gets the details for a a_default_value_4.
//
// Responses:
// default: genericError
//
//	200: a_default_value_4DBResponse
func (controller *Controller) GetA_DEFAULT_VALUE_4(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFAULT_VALUE_4", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_4.GetDB()

	// Get a_default_value_4DB in DB
	var a_default_value_4DB orm.A_DEFAULT_VALUE_4DB
	if err := db.First(&a_default_value_4DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_default_value_4API orm.A_DEFAULT_VALUE_4API
	a_default_value_4API.ID = a_default_value_4DB.ID
	a_default_value_4API.A_DEFAULT_VALUE_4PointersEncoding = a_default_value_4DB.A_DEFAULT_VALUE_4PointersEncoding
	a_default_value_4DB.CopyBasicFieldsToA_DEFAULT_VALUE_4_WOP(&a_default_value_4API.A_DEFAULT_VALUE_4_WOP)

	c.JSON(http.StatusOK, a_default_value_4API)
}

// UpdateA_DEFAULT_VALUE_4
//
// swagger:route PATCH /a_default_value_4s/{ID} a_default_value_4s updateA_DEFAULT_VALUE_4
//
// # Update a a_default_value_4
//
// Responses:
// default: genericError
//
//	200: a_default_value_4DBResponse
func (controller *Controller) UpdateA_DEFAULT_VALUE_4(c *gin.Context) {

	mutexA_DEFAULT_VALUE_4.Lock()
	defer mutexA_DEFAULT_VALUE_4.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_DEFAULT_VALUE_4", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_4.GetDB()

	// Validate input
	var input orm.A_DEFAULT_VALUE_4API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_default_value_4DB orm.A_DEFAULT_VALUE_4DB

	// fetch the a_default_value_4
	query := db.First(&a_default_value_4DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_default_value_4DB.CopyBasicFieldsFromA_DEFAULT_VALUE_4_WOP(&input.A_DEFAULT_VALUE_4_WOP)
	a_default_value_4DB.A_DEFAULT_VALUE_4PointersEncoding = input.A_DEFAULT_VALUE_4PointersEncoding

	query = db.Model(&a_default_value_4DB).Updates(a_default_value_4DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_default_value_4New := new(models.A_DEFAULT_VALUE_4)
	a_default_value_4DB.CopyBasicFieldsToA_DEFAULT_VALUE_4(a_default_value_4New)

	// redeem pointers
	a_default_value_4DB.DecodePointers(backRepo, a_default_value_4New)

	// get stage instance from DB instance, and call callback function
	a_default_value_4Old := backRepo.BackRepoA_DEFAULT_VALUE_4.Map_A_DEFAULT_VALUE_4DBID_A_DEFAULT_VALUE_4Ptr[a_default_value_4DB.ID]
	if a_default_value_4Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_default_value_4Old, a_default_value_4New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_default_value_4DB
	c.JSON(http.StatusOK, a_default_value_4DB)
}

// DeleteA_DEFAULT_VALUE_4
//
// swagger:route DELETE /a_default_value_4s/{ID} a_default_value_4s deleteA_DEFAULT_VALUE_4
//
// # Delete a a_default_value_4
//
// default: genericError
//
//	200: a_default_value_4DBResponse
func (controller *Controller) DeleteA_DEFAULT_VALUE_4(c *gin.Context) {

	mutexA_DEFAULT_VALUE_4.Lock()
	defer mutexA_DEFAULT_VALUE_4.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_DEFAULT_VALUE_4", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_4.GetDB()

	// Get model if exist
	var a_default_value_4DB orm.A_DEFAULT_VALUE_4DB
	if err := db.First(&a_default_value_4DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_default_value_4DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_default_value_4Deleted := new(models.A_DEFAULT_VALUE_4)
	a_default_value_4DB.CopyBasicFieldsToA_DEFAULT_VALUE_4(a_default_value_4Deleted)

	// get stage instance from DB instance, and call callback function
	a_default_value_4Staged := backRepo.BackRepoA_DEFAULT_VALUE_4.Map_A_DEFAULT_VALUE_4DBID_A_DEFAULT_VALUE_4Ptr[a_default_value_4DB.ID]
	if a_default_value_4Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_default_value_4Staged, a_default_value_4Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

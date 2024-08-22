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
var __A_TYPE_10__dummysDeclaration__ models.A_TYPE_10
var __A_TYPE_10_time__dummyDeclaration time.Duration

var mutexA_TYPE_10 sync.Mutex

// An A_TYPE_10ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_TYPE_10 updateA_TYPE_10 deleteA_TYPE_10
type A_TYPE_10ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_TYPE_10Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_TYPE_10 updateA_TYPE_10
type A_TYPE_10Input struct {
	// The A_TYPE_10 to submit or modify
	// in: body
	A_TYPE_10 *orm.A_TYPE_10API
}

// GetA_TYPE_10s
//
// swagger:route GET /a_type_10s a_type_10s getA_TYPE_10s
//
// # Get all a_type_10s
//
// Responses:
// default: genericError
//
//	200: a_type_10DBResponse
func (controller *Controller) GetA_TYPE_10s(c *gin.Context) {

	// source slice
	var a_type_10DBs []orm.A_TYPE_10DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE_10s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_10.GetDB()

	query := db.Find(&a_type_10DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_type_10APIs := make([]orm.A_TYPE_10API, 0)

	// for each a_type_10, update fields from the database nullable fields
	for idx := range a_type_10DBs {
		a_type_10DB := &a_type_10DBs[idx]
		_ = a_type_10DB
		var a_type_10API orm.A_TYPE_10API

		// insertion point for updating fields
		a_type_10API.ID = a_type_10DB.ID
		a_type_10DB.CopyBasicFieldsToA_TYPE_10_WOP(&a_type_10API.A_TYPE_10_WOP)
		a_type_10API.A_TYPE_10PointersEncoding = a_type_10DB.A_TYPE_10PointersEncoding
		a_type_10APIs = append(a_type_10APIs, a_type_10API)
	}

	c.JSON(http.StatusOK, a_type_10APIs)
}

// PostA_TYPE_10
//
// swagger:route POST /a_type_10s a_type_10s postA_TYPE_10
//
// Creates a a_type_10
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_TYPE_10(c *gin.Context) {

	mutexA_TYPE_10.Lock()
	defer mutexA_TYPE_10.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_TYPE_10s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_10.GetDB()

	// Validate input
	var input orm.A_TYPE_10API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_type_10
	a_type_10DB := orm.A_TYPE_10DB{}
	a_type_10DB.A_TYPE_10PointersEncoding = input.A_TYPE_10PointersEncoding
	a_type_10DB.CopyBasicFieldsFromA_TYPE_10_WOP(&input.A_TYPE_10_WOP)

	query := db.Create(&a_type_10DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_TYPE_10.CheckoutPhaseOneInstance(&a_type_10DB)
	a_type_10 := backRepo.BackRepoA_TYPE_10.Map_A_TYPE_10DBID_A_TYPE_10Ptr[a_type_10DB.ID]

	if a_type_10 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_type_10)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_type_10DB)
}

// GetA_TYPE_10
//
// swagger:route GET /a_type_10s/{ID} a_type_10s getA_TYPE_10
//
// Gets the details for a a_type_10.
//
// Responses:
// default: genericError
//
//	200: a_type_10DBResponse
func (controller *Controller) GetA_TYPE_10(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE_10", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_10.GetDB()

	// Get a_type_10DB in DB
	var a_type_10DB orm.A_TYPE_10DB
	if err := db.First(&a_type_10DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_type_10API orm.A_TYPE_10API
	a_type_10API.ID = a_type_10DB.ID
	a_type_10API.A_TYPE_10PointersEncoding = a_type_10DB.A_TYPE_10PointersEncoding
	a_type_10DB.CopyBasicFieldsToA_TYPE_10_WOP(&a_type_10API.A_TYPE_10_WOP)

	c.JSON(http.StatusOK, a_type_10API)
}

// UpdateA_TYPE_10
//
// swagger:route PATCH /a_type_10s/{ID} a_type_10s updateA_TYPE_10
//
// # Update a a_type_10
//
// Responses:
// default: genericError
//
//	200: a_type_10DBResponse
func (controller *Controller) UpdateA_TYPE_10(c *gin.Context) {

	mutexA_TYPE_10.Lock()
	defer mutexA_TYPE_10.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_TYPE_10", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_10.GetDB()

	// Validate input
	var input orm.A_TYPE_10API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_type_10DB orm.A_TYPE_10DB

	// fetch the a_type_10
	query := db.First(&a_type_10DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_type_10DB.CopyBasicFieldsFromA_TYPE_10_WOP(&input.A_TYPE_10_WOP)
	a_type_10DB.A_TYPE_10PointersEncoding = input.A_TYPE_10PointersEncoding

	query = db.Model(&a_type_10DB).Updates(a_type_10DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_type_10New := new(models.A_TYPE_10)
	a_type_10DB.CopyBasicFieldsToA_TYPE_10(a_type_10New)

	// redeem pointers
	a_type_10DB.DecodePointers(backRepo, a_type_10New)

	// get stage instance from DB instance, and call callback function
	a_type_10Old := backRepo.BackRepoA_TYPE_10.Map_A_TYPE_10DBID_A_TYPE_10Ptr[a_type_10DB.ID]
	if a_type_10Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_type_10Old, a_type_10New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_type_10DB
	c.JSON(http.StatusOK, a_type_10DB)
}

// DeleteA_TYPE_10
//
// swagger:route DELETE /a_type_10s/{ID} a_type_10s deleteA_TYPE_10
//
// # Delete a a_type_10
//
// default: genericError
//
//	200: a_type_10DBResponse
func (controller *Controller) DeleteA_TYPE_10(c *gin.Context) {

	mutexA_TYPE_10.Lock()
	defer mutexA_TYPE_10.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_TYPE_10", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_10.GetDB()

	// Get model if exist
	var a_type_10DB orm.A_TYPE_10DB
	if err := db.First(&a_type_10DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_type_10DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_type_10Deleted := new(models.A_TYPE_10)
	a_type_10DB.CopyBasicFieldsToA_TYPE_10(a_type_10Deleted)

	// get stage instance from DB instance, and call callback function
	a_type_10Staged := backRepo.BackRepoA_TYPE_10.Map_A_TYPE_10DBID_A_TYPE_10Ptr[a_type_10DB.ID]
	if a_type_10Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_type_10Staged, a_type_10Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

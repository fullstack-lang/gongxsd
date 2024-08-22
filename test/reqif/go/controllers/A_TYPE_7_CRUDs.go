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
var __A_TYPE_7__dummysDeclaration__ models.A_TYPE_7
var __A_TYPE_7_time__dummyDeclaration time.Duration

var mutexA_TYPE_7 sync.Mutex

// An A_TYPE_7ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_TYPE_7 updateA_TYPE_7 deleteA_TYPE_7
type A_TYPE_7ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_TYPE_7Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_TYPE_7 updateA_TYPE_7
type A_TYPE_7Input struct {
	// The A_TYPE_7 to submit or modify
	// in: body
	A_TYPE_7 *orm.A_TYPE_7API
}

// GetA_TYPE_7s
//
// swagger:route GET /a_type_7s a_type_7s getA_TYPE_7s
//
// # Get all a_type_7s
//
// Responses:
// default: genericError
//
//	200: a_type_7DBResponse
func (controller *Controller) GetA_TYPE_7s(c *gin.Context) {

	// source slice
	var a_type_7DBs []orm.A_TYPE_7DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE_7s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_7.GetDB()

	query := db.Find(&a_type_7DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_type_7APIs := make([]orm.A_TYPE_7API, 0)

	// for each a_type_7, update fields from the database nullable fields
	for idx := range a_type_7DBs {
		a_type_7DB := &a_type_7DBs[idx]
		_ = a_type_7DB
		var a_type_7API orm.A_TYPE_7API

		// insertion point for updating fields
		a_type_7API.ID = a_type_7DB.ID
		a_type_7DB.CopyBasicFieldsToA_TYPE_7_WOP(&a_type_7API.A_TYPE_7_WOP)
		a_type_7API.A_TYPE_7PointersEncoding = a_type_7DB.A_TYPE_7PointersEncoding
		a_type_7APIs = append(a_type_7APIs, a_type_7API)
	}

	c.JSON(http.StatusOK, a_type_7APIs)
}

// PostA_TYPE_7
//
// swagger:route POST /a_type_7s a_type_7s postA_TYPE_7
//
// Creates a a_type_7
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_TYPE_7(c *gin.Context) {

	mutexA_TYPE_7.Lock()
	defer mutexA_TYPE_7.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_TYPE_7s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_7.GetDB()

	// Validate input
	var input orm.A_TYPE_7API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_type_7
	a_type_7DB := orm.A_TYPE_7DB{}
	a_type_7DB.A_TYPE_7PointersEncoding = input.A_TYPE_7PointersEncoding
	a_type_7DB.CopyBasicFieldsFromA_TYPE_7_WOP(&input.A_TYPE_7_WOP)

	query := db.Create(&a_type_7DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_TYPE_7.CheckoutPhaseOneInstance(&a_type_7DB)
	a_type_7 := backRepo.BackRepoA_TYPE_7.Map_A_TYPE_7DBID_A_TYPE_7Ptr[a_type_7DB.ID]

	if a_type_7 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_type_7)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_type_7DB)
}

// GetA_TYPE_7
//
// swagger:route GET /a_type_7s/{ID} a_type_7s getA_TYPE_7
//
// Gets the details for a a_type_7.
//
// Responses:
// default: genericError
//
//	200: a_type_7DBResponse
func (controller *Controller) GetA_TYPE_7(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE_7", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_7.GetDB()

	// Get a_type_7DB in DB
	var a_type_7DB orm.A_TYPE_7DB
	if err := db.First(&a_type_7DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_type_7API orm.A_TYPE_7API
	a_type_7API.ID = a_type_7DB.ID
	a_type_7API.A_TYPE_7PointersEncoding = a_type_7DB.A_TYPE_7PointersEncoding
	a_type_7DB.CopyBasicFieldsToA_TYPE_7_WOP(&a_type_7API.A_TYPE_7_WOP)

	c.JSON(http.StatusOK, a_type_7API)
}

// UpdateA_TYPE_7
//
// swagger:route PATCH /a_type_7s/{ID} a_type_7s updateA_TYPE_7
//
// # Update a a_type_7
//
// Responses:
// default: genericError
//
//	200: a_type_7DBResponse
func (controller *Controller) UpdateA_TYPE_7(c *gin.Context) {

	mutexA_TYPE_7.Lock()
	defer mutexA_TYPE_7.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_TYPE_7", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_7.GetDB()

	// Validate input
	var input orm.A_TYPE_7API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_type_7DB orm.A_TYPE_7DB

	// fetch the a_type_7
	query := db.First(&a_type_7DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_type_7DB.CopyBasicFieldsFromA_TYPE_7_WOP(&input.A_TYPE_7_WOP)
	a_type_7DB.A_TYPE_7PointersEncoding = input.A_TYPE_7PointersEncoding

	query = db.Model(&a_type_7DB).Updates(a_type_7DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_type_7New := new(models.A_TYPE_7)
	a_type_7DB.CopyBasicFieldsToA_TYPE_7(a_type_7New)

	// redeem pointers
	a_type_7DB.DecodePointers(backRepo, a_type_7New)

	// get stage instance from DB instance, and call callback function
	a_type_7Old := backRepo.BackRepoA_TYPE_7.Map_A_TYPE_7DBID_A_TYPE_7Ptr[a_type_7DB.ID]
	if a_type_7Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_type_7Old, a_type_7New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_type_7DB
	c.JSON(http.StatusOK, a_type_7DB)
}

// DeleteA_TYPE_7
//
// swagger:route DELETE /a_type_7s/{ID} a_type_7s deleteA_TYPE_7
//
// # Delete a a_type_7
//
// default: genericError
//
//	200: a_type_7DBResponse
func (controller *Controller) DeleteA_TYPE_7(c *gin.Context) {

	mutexA_TYPE_7.Lock()
	defer mutexA_TYPE_7.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_TYPE_7", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_7.GetDB()

	// Get model if exist
	var a_type_7DB orm.A_TYPE_7DB
	if err := db.First(&a_type_7DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_type_7DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_type_7Deleted := new(models.A_TYPE_7)
	a_type_7DB.CopyBasicFieldsToA_TYPE_7(a_type_7Deleted)

	// get stage instance from DB instance, and call callback function
	a_type_7Staged := backRepo.BackRepoA_TYPE_7.Map_A_TYPE_7DBID_A_TYPE_7Ptr[a_type_7DB.ID]
	if a_type_7Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_type_7Staged, a_type_7Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

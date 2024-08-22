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
var __A_TYPE_5__dummysDeclaration__ models.A_TYPE_5
var __A_TYPE_5_time__dummyDeclaration time.Duration

var mutexA_TYPE_5 sync.Mutex

// An A_TYPE_5ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_TYPE_5 updateA_TYPE_5 deleteA_TYPE_5
type A_TYPE_5ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_TYPE_5Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_TYPE_5 updateA_TYPE_5
type A_TYPE_5Input struct {
	// The A_TYPE_5 to submit or modify
	// in: body
	A_TYPE_5 *orm.A_TYPE_5API
}

// GetA_TYPE_5s
//
// swagger:route GET /a_type_5s a_type_5s getA_TYPE_5s
//
// # Get all a_type_5s
//
// Responses:
// default: genericError
//
//	200: a_type_5DBResponse
func (controller *Controller) GetA_TYPE_5s(c *gin.Context) {

	// source slice
	var a_type_5DBs []orm.A_TYPE_5DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE_5s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_5.GetDB()

	query := db.Find(&a_type_5DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_type_5APIs := make([]orm.A_TYPE_5API, 0)

	// for each a_type_5, update fields from the database nullable fields
	for idx := range a_type_5DBs {
		a_type_5DB := &a_type_5DBs[idx]
		_ = a_type_5DB
		var a_type_5API orm.A_TYPE_5API

		// insertion point for updating fields
		a_type_5API.ID = a_type_5DB.ID
		a_type_5DB.CopyBasicFieldsToA_TYPE_5_WOP(&a_type_5API.A_TYPE_5_WOP)
		a_type_5API.A_TYPE_5PointersEncoding = a_type_5DB.A_TYPE_5PointersEncoding
		a_type_5APIs = append(a_type_5APIs, a_type_5API)
	}

	c.JSON(http.StatusOK, a_type_5APIs)
}

// PostA_TYPE_5
//
// swagger:route POST /a_type_5s a_type_5s postA_TYPE_5
//
// Creates a a_type_5
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_TYPE_5(c *gin.Context) {

	mutexA_TYPE_5.Lock()
	defer mutexA_TYPE_5.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_TYPE_5s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_5.GetDB()

	// Validate input
	var input orm.A_TYPE_5API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_type_5
	a_type_5DB := orm.A_TYPE_5DB{}
	a_type_5DB.A_TYPE_5PointersEncoding = input.A_TYPE_5PointersEncoding
	a_type_5DB.CopyBasicFieldsFromA_TYPE_5_WOP(&input.A_TYPE_5_WOP)

	query := db.Create(&a_type_5DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_TYPE_5.CheckoutPhaseOneInstance(&a_type_5DB)
	a_type_5 := backRepo.BackRepoA_TYPE_5.Map_A_TYPE_5DBID_A_TYPE_5Ptr[a_type_5DB.ID]

	if a_type_5 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_type_5)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_type_5DB)
}

// GetA_TYPE_5
//
// swagger:route GET /a_type_5s/{ID} a_type_5s getA_TYPE_5
//
// Gets the details for a a_type_5.
//
// Responses:
// default: genericError
//
//	200: a_type_5DBResponse
func (controller *Controller) GetA_TYPE_5(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE_5", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_5.GetDB()

	// Get a_type_5DB in DB
	var a_type_5DB orm.A_TYPE_5DB
	if err := db.First(&a_type_5DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_type_5API orm.A_TYPE_5API
	a_type_5API.ID = a_type_5DB.ID
	a_type_5API.A_TYPE_5PointersEncoding = a_type_5DB.A_TYPE_5PointersEncoding
	a_type_5DB.CopyBasicFieldsToA_TYPE_5_WOP(&a_type_5API.A_TYPE_5_WOP)

	c.JSON(http.StatusOK, a_type_5API)
}

// UpdateA_TYPE_5
//
// swagger:route PATCH /a_type_5s/{ID} a_type_5s updateA_TYPE_5
//
// # Update a a_type_5
//
// Responses:
// default: genericError
//
//	200: a_type_5DBResponse
func (controller *Controller) UpdateA_TYPE_5(c *gin.Context) {

	mutexA_TYPE_5.Lock()
	defer mutexA_TYPE_5.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_TYPE_5", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_5.GetDB()

	// Validate input
	var input orm.A_TYPE_5API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_type_5DB orm.A_TYPE_5DB

	// fetch the a_type_5
	query := db.First(&a_type_5DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_type_5DB.CopyBasicFieldsFromA_TYPE_5_WOP(&input.A_TYPE_5_WOP)
	a_type_5DB.A_TYPE_5PointersEncoding = input.A_TYPE_5PointersEncoding

	query = db.Model(&a_type_5DB).Updates(a_type_5DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_type_5New := new(models.A_TYPE_5)
	a_type_5DB.CopyBasicFieldsToA_TYPE_5(a_type_5New)

	// redeem pointers
	a_type_5DB.DecodePointers(backRepo, a_type_5New)

	// get stage instance from DB instance, and call callback function
	a_type_5Old := backRepo.BackRepoA_TYPE_5.Map_A_TYPE_5DBID_A_TYPE_5Ptr[a_type_5DB.ID]
	if a_type_5Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_type_5Old, a_type_5New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_type_5DB
	c.JSON(http.StatusOK, a_type_5DB)
}

// DeleteA_TYPE_5
//
// swagger:route DELETE /a_type_5s/{ID} a_type_5s deleteA_TYPE_5
//
// # Delete a a_type_5
//
// default: genericError
//
//	200: a_type_5DBResponse
func (controller *Controller) DeleteA_TYPE_5(c *gin.Context) {

	mutexA_TYPE_5.Lock()
	defer mutexA_TYPE_5.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_TYPE_5", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_5.GetDB()

	// Get model if exist
	var a_type_5DB orm.A_TYPE_5DB
	if err := db.First(&a_type_5DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_type_5DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_type_5Deleted := new(models.A_TYPE_5)
	a_type_5DB.CopyBasicFieldsToA_TYPE_5(a_type_5Deleted)

	// get stage instance from DB instance, and call callback function
	a_type_5Staged := backRepo.BackRepoA_TYPE_5.Map_A_TYPE_5DBID_A_TYPE_5Ptr[a_type_5DB.ID]
	if a_type_5Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_type_5Staged, a_type_5Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

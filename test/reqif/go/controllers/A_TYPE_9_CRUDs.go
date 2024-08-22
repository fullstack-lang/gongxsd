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
var __A_TYPE_9__dummysDeclaration__ models.A_TYPE_9
var __A_TYPE_9_time__dummyDeclaration time.Duration

var mutexA_TYPE_9 sync.Mutex

// An A_TYPE_9ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_TYPE_9 updateA_TYPE_9 deleteA_TYPE_9
type A_TYPE_9ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_TYPE_9Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_TYPE_9 updateA_TYPE_9
type A_TYPE_9Input struct {
	// The A_TYPE_9 to submit or modify
	// in: body
	A_TYPE_9 *orm.A_TYPE_9API
}

// GetA_TYPE_9s
//
// swagger:route GET /a_type_9s a_type_9s getA_TYPE_9s
//
// # Get all a_type_9s
//
// Responses:
// default: genericError
//
//	200: a_type_9DBResponse
func (controller *Controller) GetA_TYPE_9s(c *gin.Context) {

	// source slice
	var a_type_9DBs []orm.A_TYPE_9DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE_9s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_9.GetDB()

	query := db.Find(&a_type_9DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_type_9APIs := make([]orm.A_TYPE_9API, 0)

	// for each a_type_9, update fields from the database nullable fields
	for idx := range a_type_9DBs {
		a_type_9DB := &a_type_9DBs[idx]
		_ = a_type_9DB
		var a_type_9API orm.A_TYPE_9API

		// insertion point for updating fields
		a_type_9API.ID = a_type_9DB.ID
		a_type_9DB.CopyBasicFieldsToA_TYPE_9_WOP(&a_type_9API.A_TYPE_9_WOP)
		a_type_9API.A_TYPE_9PointersEncoding = a_type_9DB.A_TYPE_9PointersEncoding
		a_type_9APIs = append(a_type_9APIs, a_type_9API)
	}

	c.JSON(http.StatusOK, a_type_9APIs)
}

// PostA_TYPE_9
//
// swagger:route POST /a_type_9s a_type_9s postA_TYPE_9
//
// Creates a a_type_9
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_TYPE_9(c *gin.Context) {

	mutexA_TYPE_9.Lock()
	defer mutexA_TYPE_9.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_TYPE_9s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_9.GetDB()

	// Validate input
	var input orm.A_TYPE_9API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_type_9
	a_type_9DB := orm.A_TYPE_9DB{}
	a_type_9DB.A_TYPE_9PointersEncoding = input.A_TYPE_9PointersEncoding
	a_type_9DB.CopyBasicFieldsFromA_TYPE_9_WOP(&input.A_TYPE_9_WOP)

	query := db.Create(&a_type_9DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_TYPE_9.CheckoutPhaseOneInstance(&a_type_9DB)
	a_type_9 := backRepo.BackRepoA_TYPE_9.Map_A_TYPE_9DBID_A_TYPE_9Ptr[a_type_9DB.ID]

	if a_type_9 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_type_9)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_type_9DB)
}

// GetA_TYPE_9
//
// swagger:route GET /a_type_9s/{ID} a_type_9s getA_TYPE_9
//
// Gets the details for a a_type_9.
//
// Responses:
// default: genericError
//
//	200: a_type_9DBResponse
func (controller *Controller) GetA_TYPE_9(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE_9", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_9.GetDB()

	// Get a_type_9DB in DB
	var a_type_9DB orm.A_TYPE_9DB
	if err := db.First(&a_type_9DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_type_9API orm.A_TYPE_9API
	a_type_9API.ID = a_type_9DB.ID
	a_type_9API.A_TYPE_9PointersEncoding = a_type_9DB.A_TYPE_9PointersEncoding
	a_type_9DB.CopyBasicFieldsToA_TYPE_9_WOP(&a_type_9API.A_TYPE_9_WOP)

	c.JSON(http.StatusOK, a_type_9API)
}

// UpdateA_TYPE_9
//
// swagger:route PATCH /a_type_9s/{ID} a_type_9s updateA_TYPE_9
//
// # Update a a_type_9
//
// Responses:
// default: genericError
//
//	200: a_type_9DBResponse
func (controller *Controller) UpdateA_TYPE_9(c *gin.Context) {

	mutexA_TYPE_9.Lock()
	defer mutexA_TYPE_9.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_TYPE_9", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_9.GetDB()

	// Validate input
	var input orm.A_TYPE_9API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_type_9DB orm.A_TYPE_9DB

	// fetch the a_type_9
	query := db.First(&a_type_9DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_type_9DB.CopyBasicFieldsFromA_TYPE_9_WOP(&input.A_TYPE_9_WOP)
	a_type_9DB.A_TYPE_9PointersEncoding = input.A_TYPE_9PointersEncoding

	query = db.Model(&a_type_9DB).Updates(a_type_9DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_type_9New := new(models.A_TYPE_9)
	a_type_9DB.CopyBasicFieldsToA_TYPE_9(a_type_9New)

	// redeem pointers
	a_type_9DB.DecodePointers(backRepo, a_type_9New)

	// get stage instance from DB instance, and call callback function
	a_type_9Old := backRepo.BackRepoA_TYPE_9.Map_A_TYPE_9DBID_A_TYPE_9Ptr[a_type_9DB.ID]
	if a_type_9Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_type_9Old, a_type_9New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_type_9DB
	c.JSON(http.StatusOK, a_type_9DB)
}

// DeleteA_TYPE_9
//
// swagger:route DELETE /a_type_9s/{ID} a_type_9s deleteA_TYPE_9
//
// # Delete a a_type_9
//
// default: genericError
//
//	200: a_type_9DBResponse
func (controller *Controller) DeleteA_TYPE_9(c *gin.Context) {

	mutexA_TYPE_9.Lock()
	defer mutexA_TYPE_9.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_TYPE_9", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_9.GetDB()

	// Get model if exist
	var a_type_9DB orm.A_TYPE_9DB
	if err := db.First(&a_type_9DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_type_9DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_type_9Deleted := new(models.A_TYPE_9)
	a_type_9DB.CopyBasicFieldsToA_TYPE_9(a_type_9Deleted)

	// get stage instance from DB instance, and call callback function
	a_type_9Staged := backRepo.BackRepoA_TYPE_9.Map_A_TYPE_9DBID_A_TYPE_9Ptr[a_type_9DB.ID]
	if a_type_9Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_type_9Staged, a_type_9Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

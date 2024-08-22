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
var __A_TYPE_6__dummysDeclaration__ models.A_TYPE_6
var __A_TYPE_6_time__dummyDeclaration time.Duration

var mutexA_TYPE_6 sync.Mutex

// An A_TYPE_6ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_TYPE_6 updateA_TYPE_6 deleteA_TYPE_6
type A_TYPE_6ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_TYPE_6Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_TYPE_6 updateA_TYPE_6
type A_TYPE_6Input struct {
	// The A_TYPE_6 to submit or modify
	// in: body
	A_TYPE_6 *orm.A_TYPE_6API
}

// GetA_TYPE_6s
//
// swagger:route GET /a_type_6s a_type_6s getA_TYPE_6s
//
// # Get all a_type_6s
//
// Responses:
// default: genericError
//
//	200: a_type_6DBResponse
func (controller *Controller) GetA_TYPE_6s(c *gin.Context) {

	// source slice
	var a_type_6DBs []orm.A_TYPE_6DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE_6s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_6.GetDB()

	query := db.Find(&a_type_6DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_type_6APIs := make([]orm.A_TYPE_6API, 0)

	// for each a_type_6, update fields from the database nullable fields
	for idx := range a_type_6DBs {
		a_type_6DB := &a_type_6DBs[idx]
		_ = a_type_6DB
		var a_type_6API orm.A_TYPE_6API

		// insertion point for updating fields
		a_type_6API.ID = a_type_6DB.ID
		a_type_6DB.CopyBasicFieldsToA_TYPE_6_WOP(&a_type_6API.A_TYPE_6_WOP)
		a_type_6API.A_TYPE_6PointersEncoding = a_type_6DB.A_TYPE_6PointersEncoding
		a_type_6APIs = append(a_type_6APIs, a_type_6API)
	}

	c.JSON(http.StatusOK, a_type_6APIs)
}

// PostA_TYPE_6
//
// swagger:route POST /a_type_6s a_type_6s postA_TYPE_6
//
// Creates a a_type_6
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_TYPE_6(c *gin.Context) {

	mutexA_TYPE_6.Lock()
	defer mutexA_TYPE_6.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_TYPE_6s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_6.GetDB()

	// Validate input
	var input orm.A_TYPE_6API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_type_6
	a_type_6DB := orm.A_TYPE_6DB{}
	a_type_6DB.A_TYPE_6PointersEncoding = input.A_TYPE_6PointersEncoding
	a_type_6DB.CopyBasicFieldsFromA_TYPE_6_WOP(&input.A_TYPE_6_WOP)

	query := db.Create(&a_type_6DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_TYPE_6.CheckoutPhaseOneInstance(&a_type_6DB)
	a_type_6 := backRepo.BackRepoA_TYPE_6.Map_A_TYPE_6DBID_A_TYPE_6Ptr[a_type_6DB.ID]

	if a_type_6 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_type_6)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_type_6DB)
}

// GetA_TYPE_6
//
// swagger:route GET /a_type_6s/{ID} a_type_6s getA_TYPE_6
//
// Gets the details for a a_type_6.
//
// Responses:
// default: genericError
//
//	200: a_type_6DBResponse
func (controller *Controller) GetA_TYPE_6(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE_6", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_6.GetDB()

	// Get a_type_6DB in DB
	var a_type_6DB orm.A_TYPE_6DB
	if err := db.First(&a_type_6DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_type_6API orm.A_TYPE_6API
	a_type_6API.ID = a_type_6DB.ID
	a_type_6API.A_TYPE_6PointersEncoding = a_type_6DB.A_TYPE_6PointersEncoding
	a_type_6DB.CopyBasicFieldsToA_TYPE_6_WOP(&a_type_6API.A_TYPE_6_WOP)

	c.JSON(http.StatusOK, a_type_6API)
}

// UpdateA_TYPE_6
//
// swagger:route PATCH /a_type_6s/{ID} a_type_6s updateA_TYPE_6
//
// # Update a a_type_6
//
// Responses:
// default: genericError
//
//	200: a_type_6DBResponse
func (controller *Controller) UpdateA_TYPE_6(c *gin.Context) {

	mutexA_TYPE_6.Lock()
	defer mutexA_TYPE_6.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_TYPE_6", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_6.GetDB()

	// Validate input
	var input orm.A_TYPE_6API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_type_6DB orm.A_TYPE_6DB

	// fetch the a_type_6
	query := db.First(&a_type_6DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_type_6DB.CopyBasicFieldsFromA_TYPE_6_WOP(&input.A_TYPE_6_WOP)
	a_type_6DB.A_TYPE_6PointersEncoding = input.A_TYPE_6PointersEncoding

	query = db.Model(&a_type_6DB).Updates(a_type_6DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_type_6New := new(models.A_TYPE_6)
	a_type_6DB.CopyBasicFieldsToA_TYPE_6(a_type_6New)

	// redeem pointers
	a_type_6DB.DecodePointers(backRepo, a_type_6New)

	// get stage instance from DB instance, and call callback function
	a_type_6Old := backRepo.BackRepoA_TYPE_6.Map_A_TYPE_6DBID_A_TYPE_6Ptr[a_type_6DB.ID]
	if a_type_6Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_type_6Old, a_type_6New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_type_6DB
	c.JSON(http.StatusOK, a_type_6DB)
}

// DeleteA_TYPE_6
//
// swagger:route DELETE /a_type_6s/{ID} a_type_6s deleteA_TYPE_6
//
// # Delete a a_type_6
//
// default: genericError
//
//	200: a_type_6DBResponse
func (controller *Controller) DeleteA_TYPE_6(c *gin.Context) {

	mutexA_TYPE_6.Lock()
	defer mutexA_TYPE_6.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_TYPE_6", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_6.GetDB()

	// Get model if exist
	var a_type_6DB orm.A_TYPE_6DB
	if err := db.First(&a_type_6DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_type_6DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_type_6Deleted := new(models.A_TYPE_6)
	a_type_6DB.CopyBasicFieldsToA_TYPE_6(a_type_6Deleted)

	// get stage instance from DB instance, and call callback function
	a_type_6Staged := backRepo.BackRepoA_TYPE_6.Map_A_TYPE_6DBID_A_TYPE_6Ptr[a_type_6DB.ID]
	if a_type_6Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_type_6Staged, a_type_6Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

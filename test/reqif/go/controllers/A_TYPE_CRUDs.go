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
var __A_TYPE__dummysDeclaration__ models.A_TYPE
var __A_TYPE_time__dummyDeclaration time.Duration

var mutexA_TYPE sync.Mutex

// An A_TYPEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_TYPE updateA_TYPE deleteA_TYPE
type A_TYPEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_TYPEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_TYPE updateA_TYPE
type A_TYPEInput struct {
	// The A_TYPE to submit or modify
	// in: body
	A_TYPE *orm.A_TYPEAPI
}

// GetA_TYPEs
//
// swagger:route GET /a_types a_types getA_TYPEs
//
// # Get all a_types
//
// Responses:
// default: genericError
//
//	200: a_typeDBResponse
func (controller *Controller) GetA_TYPEs(c *gin.Context) {

	// source slice
	var a_typeDBs []orm.A_TYPEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE.GetDB()

	query := db.Find(&a_typeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_typeAPIs := make([]orm.A_TYPEAPI, 0)

	// for each a_type, update fields from the database nullable fields
	for idx := range a_typeDBs {
		a_typeDB := &a_typeDBs[idx]
		_ = a_typeDB
		var a_typeAPI orm.A_TYPEAPI

		// insertion point for updating fields
		a_typeAPI.ID = a_typeDB.ID
		a_typeDB.CopyBasicFieldsToA_TYPE_WOP(&a_typeAPI.A_TYPE_WOP)
		a_typeAPI.A_TYPEPointersEncoding = a_typeDB.A_TYPEPointersEncoding
		a_typeAPIs = append(a_typeAPIs, a_typeAPI)
	}

	c.JSON(http.StatusOK, a_typeAPIs)
}

// PostA_TYPE
//
// swagger:route POST /a_types a_types postA_TYPE
//
// Creates a a_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_TYPE(c *gin.Context) {

	mutexA_TYPE.Lock()
	defer mutexA_TYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_TYPEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE.GetDB()

	// Validate input
	var input orm.A_TYPEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_type
	a_typeDB := orm.A_TYPEDB{}
	a_typeDB.A_TYPEPointersEncoding = input.A_TYPEPointersEncoding
	a_typeDB.CopyBasicFieldsFromA_TYPE_WOP(&input.A_TYPE_WOP)

	query := db.Create(&a_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_TYPE.CheckoutPhaseOneInstance(&a_typeDB)
	a_type := backRepo.BackRepoA_TYPE.Map_A_TYPEDBID_A_TYPEPtr[a_typeDB.ID]

	if a_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_typeDB)
}

// GetA_TYPE
//
// swagger:route GET /a_types/{ID} a_types getA_TYPE
//
// Gets the details for a a_type.
//
// Responses:
// default: genericError
//
//	200: a_typeDBResponse
func (controller *Controller) GetA_TYPE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE.GetDB()

	// Get a_typeDB in DB
	var a_typeDB orm.A_TYPEDB
	if err := db.First(&a_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_typeAPI orm.A_TYPEAPI
	a_typeAPI.ID = a_typeDB.ID
	a_typeAPI.A_TYPEPointersEncoding = a_typeDB.A_TYPEPointersEncoding
	a_typeDB.CopyBasicFieldsToA_TYPE_WOP(&a_typeAPI.A_TYPE_WOP)

	c.JSON(http.StatusOK, a_typeAPI)
}

// UpdateA_TYPE
//
// swagger:route PATCH /a_types/{ID} a_types updateA_TYPE
//
// # Update a a_type
//
// Responses:
// default: genericError
//
//	200: a_typeDBResponse
func (controller *Controller) UpdateA_TYPE(c *gin.Context) {

	mutexA_TYPE.Lock()
	defer mutexA_TYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_TYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE.GetDB()

	// Validate input
	var input orm.A_TYPEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_typeDB orm.A_TYPEDB

	// fetch the a_type
	query := db.First(&a_typeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_typeDB.CopyBasicFieldsFromA_TYPE_WOP(&input.A_TYPE_WOP)
	a_typeDB.A_TYPEPointersEncoding = input.A_TYPEPointersEncoding

	query = db.Model(&a_typeDB).Updates(a_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_typeNew := new(models.A_TYPE)
	a_typeDB.CopyBasicFieldsToA_TYPE(a_typeNew)

	// redeem pointers
	a_typeDB.DecodePointers(backRepo, a_typeNew)

	// get stage instance from DB instance, and call callback function
	a_typeOld := backRepo.BackRepoA_TYPE.Map_A_TYPEDBID_A_TYPEPtr[a_typeDB.ID]
	if a_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_typeOld, a_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_typeDB
	c.JSON(http.StatusOK, a_typeDB)
}

// DeleteA_TYPE
//
// swagger:route DELETE /a_types/{ID} a_types deleteA_TYPE
//
// # Delete a a_type
//
// default: genericError
//
//	200: a_typeDBResponse
func (controller *Controller) DeleteA_TYPE(c *gin.Context) {

	mutexA_TYPE.Lock()
	defer mutexA_TYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_TYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE.GetDB()

	// Get model if exist
	var a_typeDB orm.A_TYPEDB
	if err := db.First(&a_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_typeDeleted := new(models.A_TYPE)
	a_typeDB.CopyBasicFieldsToA_TYPE(a_typeDeleted)

	// get stage instance from DB instance, and call callback function
	a_typeStaged := backRepo.BackRepoA_TYPE.Map_A_TYPEDBID_A_TYPEPtr[a_typeDB.ID]
	if a_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_typeStaged, a_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

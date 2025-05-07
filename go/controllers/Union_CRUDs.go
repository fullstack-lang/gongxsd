// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongxsd/go/models"
	"github.com/fullstack-lang/gongxsd/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Union__dummysDeclaration__ models.Union
var __Union_time__dummyDeclaration time.Duration

var mutexUnion sync.Mutex

// An UnionID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getUnion updateUnion deleteUnion
type UnionID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// UnionInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postUnion updateUnion
type UnionInput struct {
	// The Union to submit or modify
	// in: body
	Union *orm.UnionAPI
}

// GetUnions
//
// swagger:route GET /unions unions getUnions
//
// # Get all unions
//
// Responses:
// default: genericError
//
//	200: unionDBResponse
func (controller *Controller) GetUnions(c *gin.Context) {

	// source slice
	var unionDBs []orm.UnionDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetUnions", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gongxsd/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoUnion.GetDB()

	_, err := db.Find(&unionDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	unionAPIs := make([]orm.UnionAPI, 0)

	// for each union, update fields from the database nullable fields
	for idx := range unionDBs {
		unionDB := &unionDBs[idx]
		_ = unionDB
		var unionAPI orm.UnionAPI

		// insertion point for updating fields
		unionAPI.ID = unionDB.ID
		unionDB.CopyBasicFieldsToUnion_WOP(&unionAPI.Union_WOP)
		unionAPI.UnionPointersEncoding = unionDB.UnionPointersEncoding
		unionAPIs = append(unionAPIs, unionAPI)
	}

	c.JSON(http.StatusOK, unionAPIs)
}

// PostUnion
//
// swagger:route POST /unions unions postUnion
//
// Creates a union
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostUnion(c *gin.Context) {

	mutexUnion.Lock()
	defer mutexUnion.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostUnions", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gongxsd/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoUnion.GetDB()

	// Validate input
	var input orm.UnionAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create union
	unionDB := orm.UnionDB{}
	unionDB.UnionPointersEncoding = input.UnionPointersEncoding
	unionDB.CopyBasicFieldsFromUnion_WOP(&input.Union_WOP)

	_, err = db.Create(&unionDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoUnion.CheckoutPhaseOneInstance(&unionDB)
	union := backRepo.BackRepoUnion.Map_UnionDBID_UnionPtr[unionDB.ID]

	if union != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), union)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, unionDB)
}

// GetUnion
//
// swagger:route GET /unions/{ID} unions getUnion
//
// Gets the details for a union.
//
// Responses:
// default: genericError
//
//	200: unionDBResponse
func (controller *Controller) GetUnion(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetUnion", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gongxsd/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoUnion.GetDB()

	// Get unionDB in DB
	var unionDB orm.UnionDB
	if _, err := db.First(&unionDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var unionAPI orm.UnionAPI
	unionAPI.ID = unionDB.ID
	unionAPI.UnionPointersEncoding = unionDB.UnionPointersEncoding
	unionDB.CopyBasicFieldsToUnion_WOP(&unionAPI.Union_WOP)

	c.JSON(http.StatusOK, unionAPI)
}

// UpdateUnion
//
// swagger:route PATCH /unions/{ID} unions updateUnion
//
// # Update a union
//
// Responses:
// default: genericError
//
//	200: unionDBResponse
func (controller *Controller) UpdateUnion(c *gin.Context) {

	mutexUnion.Lock()
	defer mutexUnion.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateUnion", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gongxsd/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoUnion.GetDB()

	// Validate input
	var input orm.UnionAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var unionDB orm.UnionDB

	// fetch the union
	_, err := db.First(&unionDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	unionDB.CopyBasicFieldsFromUnion_WOP(&input.Union_WOP)
	unionDB.UnionPointersEncoding = input.UnionPointersEncoding

	db, _ = db.Model(&unionDB)
	_, err = db.Updates(&unionDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	unionNew := new(models.Union)
	unionDB.CopyBasicFieldsToUnion(unionNew)

	// redeem pointers
	unionDB.DecodePointers(backRepo, unionNew)

	// get stage instance from DB instance, and call callback function
	unionOld := backRepo.BackRepoUnion.Map_UnionDBID_UnionPtr[unionDB.ID]
	if unionOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), unionOld, unionNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the unionDB
	c.JSON(http.StatusOK, unionDB)
}

// DeleteUnion
//
// swagger:route DELETE /unions/{ID} unions deleteUnion
//
// # Delete a union
//
// default: genericError
//
//	200: unionDBResponse
func (controller *Controller) DeleteUnion(c *gin.Context) {

	mutexUnion.Lock()
	defer mutexUnion.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteUnion", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gongxsd/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoUnion.GetDB()

	// Get model if exist
	var unionDB orm.UnionDB
	if _, err := db.First(&unionDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&unionDB)

	// get an instance (not staged) from DB instance, and call callback function
	unionDeleted := new(models.Union)
	unionDB.CopyBasicFieldsToUnion(unionDeleted)

	// get stage instance from DB instance, and call callback function
	unionStaged := backRepo.BackRepoUnion.Map_UnionDBID_UnionPtr[unionDB.ID]
	if unionStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), unionStaged, unionDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

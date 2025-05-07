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
var __SimpleType__dummysDeclaration__ models.SimpleType
var __SimpleType_time__dummyDeclaration time.Duration

var mutexSimpleType sync.Mutex

// An SimpleTypeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSimpleType updateSimpleType deleteSimpleType
type SimpleTypeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SimpleTypeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSimpleType updateSimpleType
type SimpleTypeInput struct {
	// The SimpleType to submit or modify
	// in: body
	SimpleType *orm.SimpleTypeAPI
}

// GetSimpleTypes
//
// swagger:route GET /simpletypes simpletypes getSimpleTypes
//
// # Get all simpletypes
//
// Responses:
// default: genericError
//
//	200: simpletypeDBResponse
func (controller *Controller) GetSimpleTypes(c *gin.Context) {

	// source slice
	var simpletypeDBs []orm.SimpleTypeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSimpleTypes", "Name", stackPath)
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
	db := backRepo.BackRepoSimpleType.GetDB()

	_, err := db.Find(&simpletypeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	simpletypeAPIs := make([]orm.SimpleTypeAPI, 0)

	// for each simpletype, update fields from the database nullable fields
	for idx := range simpletypeDBs {
		simpletypeDB := &simpletypeDBs[idx]
		_ = simpletypeDB
		var simpletypeAPI orm.SimpleTypeAPI

		// insertion point for updating fields
		simpletypeAPI.ID = simpletypeDB.ID
		simpletypeDB.CopyBasicFieldsToSimpleType_WOP(&simpletypeAPI.SimpleType_WOP)
		simpletypeAPI.SimpleTypePointersEncoding = simpletypeDB.SimpleTypePointersEncoding
		simpletypeAPIs = append(simpletypeAPIs, simpletypeAPI)
	}

	c.JSON(http.StatusOK, simpletypeAPIs)
}

// PostSimpleType
//
// swagger:route POST /simpletypes simpletypes postSimpleType
//
// Creates a simpletype
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSimpleType(c *gin.Context) {

	mutexSimpleType.Lock()
	defer mutexSimpleType.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSimpleTypes", "Name", stackPath)
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
	db := backRepo.BackRepoSimpleType.GetDB()

	// Validate input
	var input orm.SimpleTypeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create simpletype
	simpletypeDB := orm.SimpleTypeDB{}
	simpletypeDB.SimpleTypePointersEncoding = input.SimpleTypePointersEncoding
	simpletypeDB.CopyBasicFieldsFromSimpleType_WOP(&input.SimpleType_WOP)

	_, err = db.Create(&simpletypeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSimpleType.CheckoutPhaseOneInstance(&simpletypeDB)
	simpletype := backRepo.BackRepoSimpleType.Map_SimpleTypeDBID_SimpleTypePtr[simpletypeDB.ID]

	if simpletype != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), simpletype)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, simpletypeDB)
}

// GetSimpleType
//
// swagger:route GET /simpletypes/{ID} simpletypes getSimpleType
//
// Gets the details for a simpletype.
//
// Responses:
// default: genericError
//
//	200: simpletypeDBResponse
func (controller *Controller) GetSimpleType(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSimpleType", "Name", stackPath)
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
	db := backRepo.BackRepoSimpleType.GetDB()

	// Get simpletypeDB in DB
	var simpletypeDB orm.SimpleTypeDB
	if _, err := db.First(&simpletypeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var simpletypeAPI orm.SimpleTypeAPI
	simpletypeAPI.ID = simpletypeDB.ID
	simpletypeAPI.SimpleTypePointersEncoding = simpletypeDB.SimpleTypePointersEncoding
	simpletypeDB.CopyBasicFieldsToSimpleType_WOP(&simpletypeAPI.SimpleType_WOP)

	c.JSON(http.StatusOK, simpletypeAPI)
}

// UpdateSimpleType
//
// swagger:route PATCH /simpletypes/{ID} simpletypes updateSimpleType
//
// # Update a simpletype
//
// Responses:
// default: genericError
//
//	200: simpletypeDBResponse
func (controller *Controller) UpdateSimpleType(c *gin.Context) {

	mutexSimpleType.Lock()
	defer mutexSimpleType.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSimpleType", "Name", stackPath)
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
	db := backRepo.BackRepoSimpleType.GetDB()

	// Validate input
	var input orm.SimpleTypeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var simpletypeDB orm.SimpleTypeDB

	// fetch the simpletype
	_, err := db.First(&simpletypeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	simpletypeDB.CopyBasicFieldsFromSimpleType_WOP(&input.SimpleType_WOP)
	simpletypeDB.SimpleTypePointersEncoding = input.SimpleTypePointersEncoding

	db, _ = db.Model(&simpletypeDB)
	_, err = db.Updates(&simpletypeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	simpletypeNew := new(models.SimpleType)
	simpletypeDB.CopyBasicFieldsToSimpleType(simpletypeNew)

	// redeem pointers
	simpletypeDB.DecodePointers(backRepo, simpletypeNew)

	// get stage instance from DB instance, and call callback function
	simpletypeOld := backRepo.BackRepoSimpleType.Map_SimpleTypeDBID_SimpleTypePtr[simpletypeDB.ID]
	if simpletypeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), simpletypeOld, simpletypeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the simpletypeDB
	c.JSON(http.StatusOK, simpletypeDB)
}

// DeleteSimpleType
//
// swagger:route DELETE /simpletypes/{ID} simpletypes deleteSimpleType
//
// # Delete a simpletype
//
// default: genericError
//
//	200: simpletypeDBResponse
func (controller *Controller) DeleteSimpleType(c *gin.Context) {

	mutexSimpleType.Lock()
	defer mutexSimpleType.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSimpleType", "Name", stackPath)
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
	db := backRepo.BackRepoSimpleType.GetDB()

	// Get model if exist
	var simpletypeDB orm.SimpleTypeDB
	if _, err := db.First(&simpletypeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&simpletypeDB)

	// get an instance (not staged) from DB instance, and call callback function
	simpletypeDeleted := new(models.SimpleType)
	simpletypeDB.CopyBasicFieldsToSimpleType(simpletypeDeleted)

	// get stage instance from DB instance, and call callback function
	simpletypeStaged := backRepo.BackRepoSimpleType.Map_SimpleTypeDBID_SimpleTypePtr[simpletypeDB.ID]
	if simpletypeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), simpletypeStaged, simpletypeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

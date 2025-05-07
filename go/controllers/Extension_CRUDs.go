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
var __Extension__dummysDeclaration__ models.Extension
var __Extension_time__dummyDeclaration time.Duration

var mutexExtension sync.Mutex

// An ExtensionID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getExtension updateExtension deleteExtension
type ExtensionID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ExtensionInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postExtension updateExtension
type ExtensionInput struct {
	// The Extension to submit or modify
	// in: body
	Extension *orm.ExtensionAPI
}

// GetExtensions
//
// swagger:route GET /extensions extensions getExtensions
//
// # Get all extensions
//
// Responses:
// default: genericError
//
//	200: extensionDBResponse
func (controller *Controller) GetExtensions(c *gin.Context) {

	// source slice
	var extensionDBs []orm.ExtensionDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetExtensions", "Name", stackPath)
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
	db := backRepo.BackRepoExtension.GetDB()

	_, err := db.Find(&extensionDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	extensionAPIs := make([]orm.ExtensionAPI, 0)

	// for each extension, update fields from the database nullable fields
	for idx := range extensionDBs {
		extensionDB := &extensionDBs[idx]
		_ = extensionDB
		var extensionAPI orm.ExtensionAPI

		// insertion point for updating fields
		extensionAPI.ID = extensionDB.ID
		extensionDB.CopyBasicFieldsToExtension_WOP(&extensionAPI.Extension_WOP)
		extensionAPI.ExtensionPointersEncoding = extensionDB.ExtensionPointersEncoding
		extensionAPIs = append(extensionAPIs, extensionAPI)
	}

	c.JSON(http.StatusOK, extensionAPIs)
}

// PostExtension
//
// swagger:route POST /extensions extensions postExtension
//
// Creates a extension
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostExtension(c *gin.Context) {

	mutexExtension.Lock()
	defer mutexExtension.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostExtensions", "Name", stackPath)
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
	db := backRepo.BackRepoExtension.GetDB()

	// Validate input
	var input orm.ExtensionAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create extension
	extensionDB := orm.ExtensionDB{}
	extensionDB.ExtensionPointersEncoding = input.ExtensionPointersEncoding
	extensionDB.CopyBasicFieldsFromExtension_WOP(&input.Extension_WOP)

	_, err = db.Create(&extensionDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoExtension.CheckoutPhaseOneInstance(&extensionDB)
	extension := backRepo.BackRepoExtension.Map_ExtensionDBID_ExtensionPtr[extensionDB.ID]

	if extension != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), extension)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, extensionDB)
}

// GetExtension
//
// swagger:route GET /extensions/{ID} extensions getExtension
//
// Gets the details for a extension.
//
// Responses:
// default: genericError
//
//	200: extensionDBResponse
func (controller *Controller) GetExtension(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetExtension", "Name", stackPath)
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
	db := backRepo.BackRepoExtension.GetDB()

	// Get extensionDB in DB
	var extensionDB orm.ExtensionDB
	if _, err := db.First(&extensionDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var extensionAPI orm.ExtensionAPI
	extensionAPI.ID = extensionDB.ID
	extensionAPI.ExtensionPointersEncoding = extensionDB.ExtensionPointersEncoding
	extensionDB.CopyBasicFieldsToExtension_WOP(&extensionAPI.Extension_WOP)

	c.JSON(http.StatusOK, extensionAPI)
}

// UpdateExtension
//
// swagger:route PATCH /extensions/{ID} extensions updateExtension
//
// # Update a extension
//
// Responses:
// default: genericError
//
//	200: extensionDBResponse
func (controller *Controller) UpdateExtension(c *gin.Context) {

	mutexExtension.Lock()
	defer mutexExtension.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateExtension", "Name", stackPath)
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
	db := backRepo.BackRepoExtension.GetDB()

	// Validate input
	var input orm.ExtensionAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var extensionDB orm.ExtensionDB

	// fetch the extension
	_, err := db.First(&extensionDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	extensionDB.CopyBasicFieldsFromExtension_WOP(&input.Extension_WOP)
	extensionDB.ExtensionPointersEncoding = input.ExtensionPointersEncoding

	db, _ = db.Model(&extensionDB)
	_, err = db.Updates(&extensionDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	extensionNew := new(models.Extension)
	extensionDB.CopyBasicFieldsToExtension(extensionNew)

	// redeem pointers
	extensionDB.DecodePointers(backRepo, extensionNew)

	// get stage instance from DB instance, and call callback function
	extensionOld := backRepo.BackRepoExtension.Map_ExtensionDBID_ExtensionPtr[extensionDB.ID]
	if extensionOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), extensionOld, extensionNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the extensionDB
	c.JSON(http.StatusOK, extensionDB)
}

// DeleteExtension
//
// swagger:route DELETE /extensions/{ID} extensions deleteExtension
//
// # Delete a extension
//
// default: genericError
//
//	200: extensionDBResponse
func (controller *Controller) DeleteExtension(c *gin.Context) {

	mutexExtension.Lock()
	defer mutexExtension.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteExtension", "Name", stackPath)
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
	db := backRepo.BackRepoExtension.GetDB()

	// Get model if exist
	var extensionDB orm.ExtensionDB
	if _, err := db.First(&extensionDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&extensionDB)

	// get an instance (not staged) from DB instance, and call callback function
	extensionDeleted := new(models.Extension)
	extensionDB.CopyBasicFieldsToExtension(extensionDeleted)

	// get stage instance from DB instance, and call callback function
	extensionStaged := backRepo.BackRepoExtension.Map_ExtensionDBID_ExtensionPtr[extensionDB.ID]
	if extensionStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), extensionStaged, extensionDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

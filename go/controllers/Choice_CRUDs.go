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
var __Choice__dummysDeclaration__ models.Choice
var __Choice_time__dummyDeclaration time.Duration

var mutexChoice sync.Mutex

// An ChoiceID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getChoice updateChoice deleteChoice
type ChoiceID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ChoiceInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postChoice updateChoice
type ChoiceInput struct {
	// The Choice to submit or modify
	// in: body
	Choice *orm.ChoiceAPI
}

// GetChoices
//
// swagger:route GET /choices choices getChoices
//
// # Get all choices
//
// Responses:
// default: genericError
//
//	200: choiceDBResponse
func (controller *Controller) GetChoices(c *gin.Context) {

	// source slice
	var choiceDBs []orm.ChoiceDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetChoices", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoChoice.GetDB()

	_, err := db.Find(&choiceDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	choiceAPIs := make([]orm.ChoiceAPI, 0)

	// for each choice, update fields from the database nullable fields
	for idx := range choiceDBs {
		choiceDB := &choiceDBs[idx]
		_ = choiceDB
		var choiceAPI orm.ChoiceAPI

		// insertion point for updating fields
		choiceAPI.ID = choiceDB.ID
		choiceDB.CopyBasicFieldsToChoice_WOP(&choiceAPI.Choice_WOP)
		choiceAPI.ChoicePointersEncoding = choiceDB.ChoicePointersEncoding
		choiceAPIs = append(choiceAPIs, choiceAPI)
	}

	c.JSON(http.StatusOK, choiceAPIs)
}

// PostChoice
//
// swagger:route POST /choices choices postChoice
//
// Creates a choice
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostChoice(c *gin.Context) {

	mutexChoice.Lock()
	defer mutexChoice.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostChoices", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoChoice.GetDB()

	// Validate input
	var input orm.ChoiceAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create choice
	choiceDB := orm.ChoiceDB{}
	choiceDB.ChoicePointersEncoding = input.ChoicePointersEncoding
	choiceDB.CopyBasicFieldsFromChoice_WOP(&input.Choice_WOP)

	_, err = db.Create(&choiceDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoChoice.CheckoutPhaseOneInstance(&choiceDB)
	choice := backRepo.BackRepoChoice.Map_ChoiceDBID_ChoicePtr[choiceDB.ID]

	if choice != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), choice)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, choiceDB)
}

// GetChoice
//
// swagger:route GET /choices/{ID} choices getChoice
//
// Gets the details for a choice.
//
// Responses:
// default: genericError
//
//	200: choiceDBResponse
func (controller *Controller) GetChoice(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetChoice", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoChoice.GetDB()

	// Get choiceDB in DB
	var choiceDB orm.ChoiceDB
	if _, err := db.First(&choiceDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var choiceAPI orm.ChoiceAPI
	choiceAPI.ID = choiceDB.ID
	choiceAPI.ChoicePointersEncoding = choiceDB.ChoicePointersEncoding
	choiceDB.CopyBasicFieldsToChoice_WOP(&choiceAPI.Choice_WOP)

	c.JSON(http.StatusOK, choiceAPI)
}

// UpdateChoice
//
// swagger:route PATCH /choices/{ID} choices updateChoice
//
// # Update a choice
//
// Responses:
// default: genericError
//
//	200: choiceDBResponse
func (controller *Controller) UpdateChoice(c *gin.Context) {

	mutexChoice.Lock()
	defer mutexChoice.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateChoice", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoChoice.GetDB()

	// Validate input
	var input orm.ChoiceAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var choiceDB orm.ChoiceDB

	// fetch the choice
	_, err := db.First(&choiceDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	choiceDB.CopyBasicFieldsFromChoice_WOP(&input.Choice_WOP)
	choiceDB.ChoicePointersEncoding = input.ChoicePointersEncoding

	db, _ = db.Model(&choiceDB)
	_, err = db.Updates(choiceDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	choiceNew := new(models.Choice)
	choiceDB.CopyBasicFieldsToChoice(choiceNew)

	// redeem pointers
	choiceDB.DecodePointers(backRepo, choiceNew)

	// get stage instance from DB instance, and call callback function
	choiceOld := backRepo.BackRepoChoice.Map_ChoiceDBID_ChoicePtr[choiceDB.ID]
	if choiceOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), choiceOld, choiceNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the choiceDB
	c.JSON(http.StatusOK, choiceDB)
}

// DeleteChoice
//
// swagger:route DELETE /choices/{ID} choices deleteChoice
//
// # Delete a choice
//
// default: genericError
//
//	200: choiceDBResponse
func (controller *Controller) DeleteChoice(c *gin.Context) {

	mutexChoice.Lock()
	defer mutexChoice.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteChoice", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoChoice.GetDB()

	// Get model if exist
	var choiceDB orm.ChoiceDB
	if _, err := db.First(&choiceDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&choiceDB)

	// get an instance (not staged) from DB instance, and call callback function
	choiceDeleted := new(models.Choice)
	choiceDB.CopyBasicFieldsToChoice(choiceDeleted)

	// get stage instance from DB instance, and call callback function
	choiceStaged := backRepo.BackRepoChoice.Map_ChoiceDBID_ChoicePtr[choiceDB.ID]
	if choiceStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), choiceStaged, choiceDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

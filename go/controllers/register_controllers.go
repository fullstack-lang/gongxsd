// generated code - do not edit
package controllers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/fullstack-lang/gongxsd/go/orm"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// registerControllers register controllers
func registerControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gongxsd/go")
	{ // insertion point for registrations
		v1.GET("/v1/alls", GetController().GetAlls)
		v1.GET("/v1/alls/:id", GetController().GetAll)
		v1.POST("/v1/alls", GetController().PostAll)
		v1.PATCH("/v1/alls/:id", GetController().UpdateAll)
		v1.PUT("/v1/alls/:id", GetController().UpdateAll)
		v1.DELETE("/v1/alls/:id", GetController().DeleteAll)

		v1.GET("/v1/annotations", GetController().GetAnnotations)
		v1.GET("/v1/annotations/:id", GetController().GetAnnotation)
		v1.POST("/v1/annotations", GetController().PostAnnotation)
		v1.PATCH("/v1/annotations/:id", GetController().UpdateAnnotation)
		v1.PUT("/v1/annotations/:id", GetController().UpdateAnnotation)
		v1.DELETE("/v1/annotations/:id", GetController().DeleteAnnotation)

		v1.GET("/v1/attributes", GetController().GetAttributes)
		v1.GET("/v1/attributes/:id", GetController().GetAttribute)
		v1.POST("/v1/attributes", GetController().PostAttribute)
		v1.PATCH("/v1/attributes/:id", GetController().UpdateAttribute)
		v1.PUT("/v1/attributes/:id", GetController().UpdateAttribute)
		v1.DELETE("/v1/attributes/:id", GetController().DeleteAttribute)

		v1.GET("/v1/attributegroups", GetController().GetAttributeGroups)
		v1.GET("/v1/attributegroups/:id", GetController().GetAttributeGroup)
		v1.POST("/v1/attributegroups", GetController().PostAttributeGroup)
		v1.PATCH("/v1/attributegroups/:id", GetController().UpdateAttributeGroup)
		v1.PUT("/v1/attributegroups/:id", GetController().UpdateAttributeGroup)
		v1.DELETE("/v1/attributegroups/:id", GetController().DeleteAttributeGroup)

		v1.GET("/v1/choices", GetController().GetChoices)
		v1.GET("/v1/choices/:id", GetController().GetChoice)
		v1.POST("/v1/choices", GetController().PostChoice)
		v1.PATCH("/v1/choices/:id", GetController().UpdateChoice)
		v1.PUT("/v1/choices/:id", GetController().UpdateChoice)
		v1.DELETE("/v1/choices/:id", GetController().DeleteChoice)

		v1.GET("/v1/complexcontents", GetController().GetComplexContents)
		v1.GET("/v1/complexcontents/:id", GetController().GetComplexContent)
		v1.POST("/v1/complexcontents", GetController().PostComplexContent)
		v1.PATCH("/v1/complexcontents/:id", GetController().UpdateComplexContent)
		v1.PUT("/v1/complexcontents/:id", GetController().UpdateComplexContent)
		v1.DELETE("/v1/complexcontents/:id", GetController().DeleteComplexContent)

		v1.GET("/v1/complextypes", GetController().GetComplexTypes)
		v1.GET("/v1/complextypes/:id", GetController().GetComplexType)
		v1.POST("/v1/complextypes", GetController().PostComplexType)
		v1.PATCH("/v1/complextypes/:id", GetController().UpdateComplexType)
		v1.PUT("/v1/complextypes/:id", GetController().UpdateComplexType)
		v1.DELETE("/v1/complextypes/:id", GetController().DeleteComplexType)

		v1.GET("/v1/documentations", GetController().GetDocumentations)
		v1.GET("/v1/documentations/:id", GetController().GetDocumentation)
		v1.POST("/v1/documentations", GetController().PostDocumentation)
		v1.PATCH("/v1/documentations/:id", GetController().UpdateDocumentation)
		v1.PUT("/v1/documentations/:id", GetController().UpdateDocumentation)
		v1.DELETE("/v1/documentations/:id", GetController().DeleteDocumentation)

		v1.GET("/v1/elements", GetController().GetElements)
		v1.GET("/v1/elements/:id", GetController().GetElement)
		v1.POST("/v1/elements", GetController().PostElement)
		v1.PATCH("/v1/elements/:id", GetController().UpdateElement)
		v1.PUT("/v1/elements/:id", GetController().UpdateElement)
		v1.DELETE("/v1/elements/:id", GetController().DeleteElement)

		v1.GET("/v1/enumerations", GetController().GetEnumerations)
		v1.GET("/v1/enumerations/:id", GetController().GetEnumeration)
		v1.POST("/v1/enumerations", GetController().PostEnumeration)
		v1.PATCH("/v1/enumerations/:id", GetController().UpdateEnumeration)
		v1.PUT("/v1/enumerations/:id", GetController().UpdateEnumeration)
		v1.DELETE("/v1/enumerations/:id", GetController().DeleteEnumeration)

		v1.GET("/v1/extensions", GetController().GetExtensions)
		v1.GET("/v1/extensions/:id", GetController().GetExtension)
		v1.POST("/v1/extensions", GetController().PostExtension)
		v1.PATCH("/v1/extensions/:id", GetController().UpdateExtension)
		v1.PUT("/v1/extensions/:id", GetController().UpdateExtension)
		v1.DELETE("/v1/extensions/:id", GetController().DeleteExtension)

		v1.GET("/v1/groups", GetController().GetGroups)
		v1.GET("/v1/groups/:id", GetController().GetGroup)
		v1.POST("/v1/groups", GetController().PostGroup)
		v1.PATCH("/v1/groups/:id", GetController().UpdateGroup)
		v1.PUT("/v1/groups/:id", GetController().UpdateGroup)
		v1.DELETE("/v1/groups/:id", GetController().DeleteGroup)

		v1.GET("/v1/lengths", GetController().GetLengths)
		v1.GET("/v1/lengths/:id", GetController().GetLength)
		v1.POST("/v1/lengths", GetController().PostLength)
		v1.PATCH("/v1/lengths/:id", GetController().UpdateLength)
		v1.PUT("/v1/lengths/:id", GetController().UpdateLength)
		v1.DELETE("/v1/lengths/:id", GetController().DeleteLength)

		v1.GET("/v1/maxinclusives", GetController().GetMaxInclusives)
		v1.GET("/v1/maxinclusives/:id", GetController().GetMaxInclusive)
		v1.POST("/v1/maxinclusives", GetController().PostMaxInclusive)
		v1.PATCH("/v1/maxinclusives/:id", GetController().UpdateMaxInclusive)
		v1.PUT("/v1/maxinclusives/:id", GetController().UpdateMaxInclusive)
		v1.DELETE("/v1/maxinclusives/:id", GetController().DeleteMaxInclusive)

		v1.GET("/v1/maxlengths", GetController().GetMaxLengths)
		v1.GET("/v1/maxlengths/:id", GetController().GetMaxLength)
		v1.POST("/v1/maxlengths", GetController().PostMaxLength)
		v1.PATCH("/v1/maxlengths/:id", GetController().UpdateMaxLength)
		v1.PUT("/v1/maxlengths/:id", GetController().UpdateMaxLength)
		v1.DELETE("/v1/maxlengths/:id", GetController().DeleteMaxLength)

		v1.GET("/v1/mininclusives", GetController().GetMinInclusives)
		v1.GET("/v1/mininclusives/:id", GetController().GetMinInclusive)
		v1.POST("/v1/mininclusives", GetController().PostMinInclusive)
		v1.PATCH("/v1/mininclusives/:id", GetController().UpdateMinInclusive)
		v1.PUT("/v1/mininclusives/:id", GetController().UpdateMinInclusive)
		v1.DELETE("/v1/mininclusives/:id", GetController().DeleteMinInclusive)

		v1.GET("/v1/minlengths", GetController().GetMinLengths)
		v1.GET("/v1/minlengths/:id", GetController().GetMinLength)
		v1.POST("/v1/minlengths", GetController().PostMinLength)
		v1.PATCH("/v1/minlengths/:id", GetController().UpdateMinLength)
		v1.PUT("/v1/minlengths/:id", GetController().UpdateMinLength)
		v1.DELETE("/v1/minlengths/:id", GetController().DeleteMinLength)

		v1.GET("/v1/patterns", GetController().GetPatterns)
		v1.GET("/v1/patterns/:id", GetController().GetPattern)
		v1.POST("/v1/patterns", GetController().PostPattern)
		v1.PATCH("/v1/patterns/:id", GetController().UpdatePattern)
		v1.PUT("/v1/patterns/:id", GetController().UpdatePattern)
		v1.DELETE("/v1/patterns/:id", GetController().DeletePattern)

		v1.GET("/v1/restrictions", GetController().GetRestrictions)
		v1.GET("/v1/restrictions/:id", GetController().GetRestriction)
		v1.POST("/v1/restrictions", GetController().PostRestriction)
		v1.PATCH("/v1/restrictions/:id", GetController().UpdateRestriction)
		v1.PUT("/v1/restrictions/:id", GetController().UpdateRestriction)
		v1.DELETE("/v1/restrictions/:id", GetController().DeleteRestriction)

		v1.GET("/v1/schemas", GetController().GetSchemas)
		v1.GET("/v1/schemas/:id", GetController().GetSchema)
		v1.POST("/v1/schemas", GetController().PostSchema)
		v1.PATCH("/v1/schemas/:id", GetController().UpdateSchema)
		v1.PUT("/v1/schemas/:id", GetController().UpdateSchema)
		v1.DELETE("/v1/schemas/:id", GetController().DeleteSchema)

		v1.GET("/v1/sequences", GetController().GetSequences)
		v1.GET("/v1/sequences/:id", GetController().GetSequence)
		v1.POST("/v1/sequences", GetController().PostSequence)
		v1.PATCH("/v1/sequences/:id", GetController().UpdateSequence)
		v1.PUT("/v1/sequences/:id", GetController().UpdateSequence)
		v1.DELETE("/v1/sequences/:id", GetController().DeleteSequence)

		v1.GET("/v1/simplecontents", GetController().GetSimpleContents)
		v1.GET("/v1/simplecontents/:id", GetController().GetSimpleContent)
		v1.POST("/v1/simplecontents", GetController().PostSimpleContent)
		v1.PATCH("/v1/simplecontents/:id", GetController().UpdateSimpleContent)
		v1.PUT("/v1/simplecontents/:id", GetController().UpdateSimpleContent)
		v1.DELETE("/v1/simplecontents/:id", GetController().DeleteSimpleContent)

		v1.GET("/v1/simpletypes", GetController().GetSimpleTypes)
		v1.GET("/v1/simpletypes/:id", GetController().GetSimpleType)
		v1.POST("/v1/simpletypes", GetController().PostSimpleType)
		v1.PATCH("/v1/simpletypes/:id", GetController().UpdateSimpleType)
		v1.PUT("/v1/simpletypes/:id", GetController().UpdateSimpleType)
		v1.DELETE("/v1/simpletypes/:id", GetController().DeleteSimpleType)

		v1.GET("/v1/totaldigits", GetController().GetTotalDigits)
		v1.GET("/v1/totaldigits/:id", GetController().GetTotalDigit)
		v1.POST("/v1/totaldigits", GetController().PostTotalDigit)
		v1.PATCH("/v1/totaldigits/:id", GetController().UpdateTotalDigit)
		v1.PUT("/v1/totaldigits/:id", GetController().UpdateTotalDigit)
		v1.DELETE("/v1/totaldigits/:id", GetController().DeleteTotalDigit)

		v1.GET("/v1/unions", GetController().GetUnions)
		v1.GET("/v1/unions/:id", GetController().GetUnion)
		v1.POST("/v1/unions", GetController().PostUnion)
		v1.PATCH("/v1/unions/:id", GetController().UpdateUnion)
		v1.PUT("/v1/unions/:id", GetController().UpdateUnion)
		v1.DELETE("/v1/unions/:id", GetController().DeleteUnion)

		v1.GET("/v1/whitespaces", GetController().GetWhiteSpaces)
		v1.GET("/v1/whitespaces/:id", GetController().GetWhiteSpace)
		v1.POST("/v1/whitespaces", GetController().PostWhiteSpace)
		v1.PATCH("/v1/whitespaces/:id", GetController().UpdateWhiteSpace)
		v1.PUT("/v1/whitespaces/:id", GetController().UpdateWhiteSpace)
		v1.DELETE("/v1/whitespaces/:id", GetController().DeleteWhiteSpace)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)

		v1.GET("/v1/ws/stage", GetController().onWebSocketRequestForBackRepoContent)

		v1.GET("/v1/stacks", GetController().stacks)
	}
}

func (controller *Controller) stacks(c *gin.Context) {

	var res []string

	for k := range controller.Map_BackRepos {
		res = append(res, k)
	}

	c.JSON(http.StatusOK, res)
}

// onWebSocketRequestForBackRepoContent is a function that is started each time
// a web socket request is received
//
// 1. upgrade the incomming web connection to a web socket
// 1. it subscribe to the backend commit number broadcaster
// 1. it stays live and pool for incomming backend commit number broadcast and forward
// them on the web socket connection
func (controller *Controller) onWebSocketRequestForBackRepoContent(c *gin.Context) {

	// log.Println("Stack github.com/fullstack-lang/gongxsd/go, onWebSocketRequestForBackRepoContent")

	// Upgrader specifies parameters for upgrading an HTTP connection to a
	// WebSocket connection.

	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			if origin == "" {
				log.Printf("CheckOrigin: Rejected - Origin header is empty. Request from: %s", r.RemoteAddr)
				return false // Or handle as per your security policy
			}

			u, err := url.Parse(origin)
			if err != nil {
				log.Printf("CheckOrigin: Rejected - Invalid Origin URL '%s'. Error: %v", origin, err)
				return false // Invalid URL
			}

			portStr := u.Port()

			if portStr == "" {
				// If no port is specified, it might be using default HTTP/HTTPS ports.
				// For this specific request, we'll assume a port must be present.
				log.Printf("CheckOrigin: Rejected - No port specified in Origin URL '%s'", origin)
				return false
			}

			port, err := strconv.Atoi(portStr)
			if err != nil {
				log.Printf("CheckOrigin: Rejected - Port '%s' in Origin URL '%s' is not a valid number. Error: %v", portStr, origin, err)
				return false // Port is not a valid number
			}

			// Check if the port is 4200 OR in the range 8000-9000
			allowed := port == 4200 || (port >= 8000 && port <= 9000)
			if !allowed {
				log.Printf("CheckOrigin: Rejected - Port %d from Origin '%s' is not in the allowed list (4200 or 8000-9000)", port, origin)
				return false
			}

			// log.Printf("CheckOrigin: Accepted - Origin '%s' with port %d", origin, port)
			return true
		},
	}

	wsConnection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wsConnection.Close()

	// Create a context that is canceled when the connection is closed
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "Name", stackPath)
		}
	}

	index := controller.listenerIndex
	controller.listenerIndex++

	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "WebSocket Connect, Stack github.com/fullstack-lang/gongxsd/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	updateCommitBackRepoNbChannel := backRepo.SubscribeToCommitNb(ctx)

	// Start a goroutine to read from the WebSocket to detect disconnection
	go func() {
		for {
			// ReadMessage is used to detect client disconnection
			_, _, err := wsConnection.ReadMessage()
			if err != nil {
				log.Println("github.com/fullstack-lang/gongxsd/go", stackPath, "WS client disconnected:", err)
				cancel() // Cancel the context
				return
			}
		}
	}()

	backRepoData := new(orm.BackRepoData)
	orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)
	backRepoData.GONG__Index = index

	refresh := 0
	// Marshal the data to JSON first to be able to get its size
	jsonData, err := json.Marshal(backRepoData)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		return
	}

	// Get the size of the JSON data in bytes
	jsonSize := len(jsonData)

    // Calculate the full SHA-256 hash
    fullHash := sha256.Sum256(jsonData)

    // Use the first 12 characters for a shorter, yet highly unique, signature
    shortHash := hex.EncodeToString(fullHash[:])[0:12]

	// Use WriteMessage to send the pre-marshaled JSON data.
	// websocket.TextMessage is typically what WriteJSON uses.
	err = wsConnection.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		log.Println("github.com/fullstack-lang/gongxsd/go:\n",
			"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
		fmt.Println(err)
		return
	} else {
		// 1. Extract the component name from the long path for cleaner logs
		// For example, "github.com/fullstack-lang/gong/lib/table/go" becomes "table"
		parts := strings.Split("github.com/fullstack-lang/gongxsd/go", "/") // Assuming goFilePath holds the path
		component := "unknown"
		if len(parts) > 2 {
			component = parts[len(parts)-2]
		}

		// 2. Use a single, formatted log line
		log.Printf(
			"%-12s | %-85s | Idx: %d | Size: %-9s | Hash: %s",
			component,
			stackPath,
			index,
			formatBytes(jsonSize),
			shortHash,
		)
	}
	for {
		select {
		case <-ctx.Done():
			// Context canceled, exit the loop
			return
		default:
			for nbCommitBackRepo := range updateCommitBackRepoNbChannel {
				_ = nbCommitBackRepo
				refresh += 1

				backRepoData := new(orm.BackRepoData)
				orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)
				backRepoData.GONG__Index = index

				// Set write deadline to prevent blocking indefinitely
				wsConnection.SetWriteDeadline(time.Now().Add(10 * time.Second))

				// Send backRepo data
				// Marshal the data to JSON first to be able to get its size
				jsonData, err := json.Marshal(backRepoData)
				if err != nil {
					log.Printf("Error marshaling JSON: %v", err)
					return
				}

				// Get the size of the JSON data in bytes
				jsonSize := len(jsonData)

				// Calculate the full SHA-256 hash
				fullHash := sha256.Sum256(jsonData)

				// Use the first 12 characters for a shorter, yet highly unique, signature
				shortHash := hex.EncodeToString(fullHash[:])[0:12]

				// Use WriteMessage to send the pre-marshaled JSON data.
				// websocket.TextMessage is typically what WriteJSON uses.
				err = wsConnection.WriteMessage(websocket.TextMessage, jsonData)
				if err != nil {
					log.Println("github.com/fullstack-lang/gongxsd/go:\n",
						"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
					fmt.Println(err)
					return
				} else {
					// 1. Extract the component name from the long path for cleaner logs
					// For example, "github.com/fullstack-lang/gong/lib/table/go" becomes "table"
					parts := strings.Split("github.com/fullstack-lang/gongxsd/go", "/") // Assuming goFilePath holds the path
					component := "unknown"
					if len(parts) > 2 {
						component = parts[len(parts)-2]
					}

					// 2. Use a single, formatted log line
					log.Printf(
						"%-12s | %-85s | Idx: %d | Size: %-9s | Hash: %s",
						component,
						stackPath,
						index,
						formatBytes(jsonSize),
						shortHash,
					)
				}
			}
		}
	}
}

// formatBytes converts a size in bytes to a human-readable string (KB, MB, GB).
func formatBytes(size int) string {
    if size < 1024 {
        return fmt.Sprintf("%d B", size)
    }
    sizeInKB := float64(size) / 1024.0
    if sizeInKB < 1024.0 {
        // For KB, show one decimal place if it's not a whole number
        if math.Mod(sizeInKB, 1.0) == 0 {
            return fmt.Sprintf("%.0f KB", sizeInKB)
        }
        return fmt.Sprintf("%.1f KB", sizeInKB)
    }
    sizeInMB := sizeInKB / 1024.0
    return fmt.Sprintf("%.2f MB", sizeInMB)
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func (controller *Controller) GetLastCommitFromBackNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "Name", stackPath)
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
	res := backRepo.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func (controller *Controller) GetLastPushFromFrontNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastPushFromFrontNb", "Name", stackPath)
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
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}

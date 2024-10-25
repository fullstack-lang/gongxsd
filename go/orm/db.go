// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gongxsd/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	allDBs map[uint]*AllDB

	nextIDAllDB uint

	annotationDBs map[uint]*AnnotationDB

	nextIDAnnotationDB uint

	attributeDBs map[uint]*AttributeDB

	nextIDAttributeDB uint

	attributegroupDBs map[uint]*AttributeGroupDB

	nextIDAttributeGroupDB uint

	choiceDBs map[uint]*ChoiceDB

	nextIDChoiceDB uint

	complexcontentDBs map[uint]*ComplexContentDB

	nextIDComplexContentDB uint

	complextypeDBs map[uint]*ComplexTypeDB

	nextIDComplexTypeDB uint

	documentationDBs map[uint]*DocumentationDB

	nextIDDocumentationDB uint

	elementDBs map[uint]*ElementDB

	nextIDElementDB uint

	enumerationDBs map[uint]*EnumerationDB

	nextIDEnumerationDB uint

	extensionDBs map[uint]*ExtensionDB

	nextIDExtensionDB uint

	groupDBs map[uint]*GroupDB

	nextIDGroupDB uint

	lengthDBs map[uint]*LengthDB

	nextIDLengthDB uint

	maxinclusiveDBs map[uint]*MaxInclusiveDB

	nextIDMaxInclusiveDB uint

	maxlengthDBs map[uint]*MaxLengthDB

	nextIDMaxLengthDB uint

	mininclusiveDBs map[uint]*MinInclusiveDB

	nextIDMinInclusiveDB uint

	minlengthDBs map[uint]*MinLengthDB

	nextIDMinLengthDB uint

	patternDBs map[uint]*PatternDB

	nextIDPatternDB uint

	restrictionDBs map[uint]*RestrictionDB

	nextIDRestrictionDB uint

	schemaDBs map[uint]*SchemaDB

	nextIDSchemaDB uint

	sequenceDBs map[uint]*SequenceDB

	nextIDSequenceDB uint

	simplecontentDBs map[uint]*SimpleContentDB

	nextIDSimpleContentDB uint

	simpletypeDBs map[uint]*SimpleTypeDB

	nextIDSimpleTypeDB uint

	totaldigitDBs map[uint]*TotalDigitDB

	nextIDTotalDigitDB uint

	unionDBs map[uint]*UnionDB

	nextIDUnionDB uint

	whitespaceDBs map[uint]*WhiteSpaceDB

	nextIDWhiteSpaceDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		allDBs: make(map[uint]*AllDB),

		annotationDBs: make(map[uint]*AnnotationDB),

		attributeDBs: make(map[uint]*AttributeDB),

		attributegroupDBs: make(map[uint]*AttributeGroupDB),

		choiceDBs: make(map[uint]*ChoiceDB),

		complexcontentDBs: make(map[uint]*ComplexContentDB),

		complextypeDBs: make(map[uint]*ComplexTypeDB),

		documentationDBs: make(map[uint]*DocumentationDB),

		elementDBs: make(map[uint]*ElementDB),

		enumerationDBs: make(map[uint]*EnumerationDB),

		extensionDBs: make(map[uint]*ExtensionDB),

		groupDBs: make(map[uint]*GroupDB),

		lengthDBs: make(map[uint]*LengthDB),

		maxinclusiveDBs: make(map[uint]*MaxInclusiveDB),

		maxlengthDBs: make(map[uint]*MaxLengthDB),

		mininclusiveDBs: make(map[uint]*MinInclusiveDB),

		minlengthDBs: make(map[uint]*MinLengthDB),

		patternDBs: make(map[uint]*PatternDB),

		restrictionDBs: make(map[uint]*RestrictionDB),

		schemaDBs: make(map[uint]*SchemaDB),

		sequenceDBs: make(map[uint]*SequenceDB),

		simplecontentDBs: make(map[uint]*SimpleContentDB),

		simpletypeDBs: make(map[uint]*SimpleTypeDB),

		totaldigitDBs: make(map[uint]*TotalDigitDB),

		unionDBs: make(map[uint]*UnionDB),

		whitespaceDBs: make(map[uint]*WhiteSpaceDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongxsd/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *AllDB:
		db.nextIDAllDB++
		v.ID = db.nextIDAllDB
		db.allDBs[v.ID] = v
	case *AnnotationDB:
		db.nextIDAnnotationDB++
		v.ID = db.nextIDAnnotationDB
		db.annotationDBs[v.ID] = v
	case *AttributeDB:
		db.nextIDAttributeDB++
		v.ID = db.nextIDAttributeDB
		db.attributeDBs[v.ID] = v
	case *AttributeGroupDB:
		db.nextIDAttributeGroupDB++
		v.ID = db.nextIDAttributeGroupDB
		db.attributegroupDBs[v.ID] = v
	case *ChoiceDB:
		db.nextIDChoiceDB++
		v.ID = db.nextIDChoiceDB
		db.choiceDBs[v.ID] = v
	case *ComplexContentDB:
		db.nextIDComplexContentDB++
		v.ID = db.nextIDComplexContentDB
		db.complexcontentDBs[v.ID] = v
	case *ComplexTypeDB:
		db.nextIDComplexTypeDB++
		v.ID = db.nextIDComplexTypeDB
		db.complextypeDBs[v.ID] = v
	case *DocumentationDB:
		db.nextIDDocumentationDB++
		v.ID = db.nextIDDocumentationDB
		db.documentationDBs[v.ID] = v
	case *ElementDB:
		db.nextIDElementDB++
		v.ID = db.nextIDElementDB
		db.elementDBs[v.ID] = v
	case *EnumerationDB:
		db.nextIDEnumerationDB++
		v.ID = db.nextIDEnumerationDB
		db.enumerationDBs[v.ID] = v
	case *ExtensionDB:
		db.nextIDExtensionDB++
		v.ID = db.nextIDExtensionDB
		db.extensionDBs[v.ID] = v
	case *GroupDB:
		db.nextIDGroupDB++
		v.ID = db.nextIDGroupDB
		db.groupDBs[v.ID] = v
	case *LengthDB:
		db.nextIDLengthDB++
		v.ID = db.nextIDLengthDB
		db.lengthDBs[v.ID] = v
	case *MaxInclusiveDB:
		db.nextIDMaxInclusiveDB++
		v.ID = db.nextIDMaxInclusiveDB
		db.maxinclusiveDBs[v.ID] = v
	case *MaxLengthDB:
		db.nextIDMaxLengthDB++
		v.ID = db.nextIDMaxLengthDB
		db.maxlengthDBs[v.ID] = v
	case *MinInclusiveDB:
		db.nextIDMinInclusiveDB++
		v.ID = db.nextIDMinInclusiveDB
		db.mininclusiveDBs[v.ID] = v
	case *MinLengthDB:
		db.nextIDMinLengthDB++
		v.ID = db.nextIDMinLengthDB
		db.minlengthDBs[v.ID] = v
	case *PatternDB:
		db.nextIDPatternDB++
		v.ID = db.nextIDPatternDB
		db.patternDBs[v.ID] = v
	case *RestrictionDB:
		db.nextIDRestrictionDB++
		v.ID = db.nextIDRestrictionDB
		db.restrictionDBs[v.ID] = v
	case *SchemaDB:
		db.nextIDSchemaDB++
		v.ID = db.nextIDSchemaDB
		db.schemaDBs[v.ID] = v
	case *SequenceDB:
		db.nextIDSequenceDB++
		v.ID = db.nextIDSequenceDB
		db.sequenceDBs[v.ID] = v
	case *SimpleContentDB:
		db.nextIDSimpleContentDB++
		v.ID = db.nextIDSimpleContentDB
		db.simplecontentDBs[v.ID] = v
	case *SimpleTypeDB:
		db.nextIDSimpleTypeDB++
		v.ID = db.nextIDSimpleTypeDB
		db.simpletypeDBs[v.ID] = v
	case *TotalDigitDB:
		db.nextIDTotalDigitDB++
		v.ID = db.nextIDTotalDigitDB
		db.totaldigitDBs[v.ID] = v
	case *UnionDB:
		db.nextIDUnionDB++
		v.ID = db.nextIDUnionDB
		db.unionDBs[v.ID] = v
	case *WhiteSpaceDB:
		db.nextIDWhiteSpaceDB++
		v.ID = db.nextIDWhiteSpaceDB
		db.whitespaceDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gongxsd/go, unsupported type in Create")
	}
	return db, nil
}

// Unscoped sets the unscoped flag for soft-deletes (not used in this implementation)
func (db *DBLite) Unscoped() (db.DBInterface, error) {
	return db, nil
}

// Model is a placeholder in this implementation
func (db *DBLite) Model(instanceDB any) (db.DBInterface, error) {
	// Not implemented as types are handled directly
	return db, nil
}

// Delete removes a record from the database
func (db *DBLite) Delete(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongxsd/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AllDB:
		delete(db.allDBs, v.ID)
	case *AnnotationDB:
		delete(db.annotationDBs, v.ID)
	case *AttributeDB:
		delete(db.attributeDBs, v.ID)
	case *AttributeGroupDB:
		delete(db.attributegroupDBs, v.ID)
	case *ChoiceDB:
		delete(db.choiceDBs, v.ID)
	case *ComplexContentDB:
		delete(db.complexcontentDBs, v.ID)
	case *ComplexTypeDB:
		delete(db.complextypeDBs, v.ID)
	case *DocumentationDB:
		delete(db.documentationDBs, v.ID)
	case *ElementDB:
		delete(db.elementDBs, v.ID)
	case *EnumerationDB:
		delete(db.enumerationDBs, v.ID)
	case *ExtensionDB:
		delete(db.extensionDBs, v.ID)
	case *GroupDB:
		delete(db.groupDBs, v.ID)
	case *LengthDB:
		delete(db.lengthDBs, v.ID)
	case *MaxInclusiveDB:
		delete(db.maxinclusiveDBs, v.ID)
	case *MaxLengthDB:
		delete(db.maxlengthDBs, v.ID)
	case *MinInclusiveDB:
		delete(db.mininclusiveDBs, v.ID)
	case *MinLengthDB:
		delete(db.minlengthDBs, v.ID)
	case *PatternDB:
		delete(db.patternDBs, v.ID)
	case *RestrictionDB:
		delete(db.restrictionDBs, v.ID)
	case *SchemaDB:
		delete(db.schemaDBs, v.ID)
	case *SequenceDB:
		delete(db.sequenceDBs, v.ID)
	case *SimpleContentDB:
		delete(db.simplecontentDBs, v.ID)
	case *SimpleTypeDB:
		delete(db.simpletypeDBs, v.ID)
	case *TotalDigitDB:
		delete(db.totaldigitDBs, v.ID)
	case *UnionDB:
		delete(db.unionDBs, v.ID)
	case *WhiteSpaceDB:
		delete(db.whitespaceDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongxsd/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongxsd/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AllDB:
		db.allDBs[v.ID] = v
		return db, nil
	case *AnnotationDB:
		db.annotationDBs[v.ID] = v
		return db, nil
	case *AttributeDB:
		db.attributeDBs[v.ID] = v
		return db, nil
	case *AttributeGroupDB:
		db.attributegroupDBs[v.ID] = v
		return db, nil
	case *ChoiceDB:
		db.choiceDBs[v.ID] = v
		return db, nil
	case *ComplexContentDB:
		db.complexcontentDBs[v.ID] = v
		return db, nil
	case *ComplexTypeDB:
		db.complextypeDBs[v.ID] = v
		return db, nil
	case *DocumentationDB:
		db.documentationDBs[v.ID] = v
		return db, nil
	case *ElementDB:
		db.elementDBs[v.ID] = v
		return db, nil
	case *EnumerationDB:
		db.enumerationDBs[v.ID] = v
		return db, nil
	case *ExtensionDB:
		db.extensionDBs[v.ID] = v
		return db, nil
	case *GroupDB:
		db.groupDBs[v.ID] = v
		return db, nil
	case *LengthDB:
		db.lengthDBs[v.ID] = v
		return db, nil
	case *MaxInclusiveDB:
		db.maxinclusiveDBs[v.ID] = v
		return db, nil
	case *MaxLengthDB:
		db.maxlengthDBs[v.ID] = v
		return db, nil
	case *MinInclusiveDB:
		db.mininclusiveDBs[v.ID] = v
		return db, nil
	case *MinLengthDB:
		db.minlengthDBs[v.ID] = v
		return db, nil
	case *PatternDB:
		db.patternDBs[v.ID] = v
		return db, nil
	case *RestrictionDB:
		db.restrictionDBs[v.ID] = v
		return db, nil
	case *SchemaDB:
		db.schemaDBs[v.ID] = v
		return db, nil
	case *SequenceDB:
		db.sequenceDBs[v.ID] = v
		return db, nil
	case *SimpleContentDB:
		db.simplecontentDBs[v.ID] = v
		return db, nil
	case *SimpleTypeDB:
		db.simpletypeDBs[v.ID] = v
		return db, nil
	case *TotalDigitDB:
		db.totaldigitDBs[v.ID] = v
		return db, nil
	case *UnionDB:
		db.unionDBs[v.ID] = v
		return db, nil
	case *WhiteSpaceDB:
		db.whitespaceDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongxsd/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongxsd/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AllDB:
		if existing, ok := db.allDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *AnnotationDB:
		if existing, ok := db.annotationDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *AttributeDB:
		if existing, ok := db.attributeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *AttributeGroupDB:
		if existing, ok := db.attributegroupDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *ChoiceDB:
		if existing, ok := db.choiceDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *ComplexContentDB:
		if existing, ok := db.complexcontentDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *ComplexTypeDB:
		if existing, ok := db.complextypeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *DocumentationDB:
		if existing, ok := db.documentationDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *ElementDB:
		if existing, ok := db.elementDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *EnumerationDB:
		if existing, ok := db.enumerationDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *ExtensionDB:
		if existing, ok := db.extensionDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *GroupDB:
		if existing, ok := db.groupDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *LengthDB:
		if existing, ok := db.lengthDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *MaxInclusiveDB:
		if existing, ok := db.maxinclusiveDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *MaxLengthDB:
		if existing, ok := db.maxlengthDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *MinInclusiveDB:
		if existing, ok := db.mininclusiveDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *MinLengthDB:
		if existing, ok := db.minlengthDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *PatternDB:
		if existing, ok := db.patternDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *RestrictionDB:
		if existing, ok := db.restrictionDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *SchemaDB:
		if existing, ok := db.schemaDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *SequenceDB:
		if existing, ok := db.sequenceDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *SimpleContentDB:
		if existing, ok := db.simplecontentDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *SimpleTypeDB:
		if existing, ok := db.simpletypeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *TotalDigitDB:
		if existing, ok := db.totaldigitDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *UnionDB:
		if existing, ok := db.unionDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	case *WhiteSpaceDB:
		if existing, ok := db.whitespaceDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gongxsd/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]AllDB:
        *ptr = make([]AllDB, 0, len(db.allDBs))
        for _, v := range db.allDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]AnnotationDB:
        *ptr = make([]AnnotationDB, 0, len(db.annotationDBs))
        for _, v := range db.annotationDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]AttributeDB:
        *ptr = make([]AttributeDB, 0, len(db.attributeDBs))
        for _, v := range db.attributeDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]AttributeGroupDB:
        *ptr = make([]AttributeGroupDB, 0, len(db.attributegroupDBs))
        for _, v := range db.attributegroupDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]ChoiceDB:
        *ptr = make([]ChoiceDB, 0, len(db.choiceDBs))
        for _, v := range db.choiceDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]ComplexContentDB:
        *ptr = make([]ComplexContentDB, 0, len(db.complexcontentDBs))
        for _, v := range db.complexcontentDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]ComplexTypeDB:
        *ptr = make([]ComplexTypeDB, 0, len(db.complextypeDBs))
        for _, v := range db.complextypeDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]DocumentationDB:
        *ptr = make([]DocumentationDB, 0, len(db.documentationDBs))
        for _, v := range db.documentationDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]ElementDB:
        *ptr = make([]ElementDB, 0, len(db.elementDBs))
        for _, v := range db.elementDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]EnumerationDB:
        *ptr = make([]EnumerationDB, 0, len(db.enumerationDBs))
        for _, v := range db.enumerationDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]ExtensionDB:
        *ptr = make([]ExtensionDB, 0, len(db.extensionDBs))
        for _, v := range db.extensionDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]GroupDB:
        *ptr = make([]GroupDB, 0, len(db.groupDBs))
        for _, v := range db.groupDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]LengthDB:
        *ptr = make([]LengthDB, 0, len(db.lengthDBs))
        for _, v := range db.lengthDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]MaxInclusiveDB:
        *ptr = make([]MaxInclusiveDB, 0, len(db.maxinclusiveDBs))
        for _, v := range db.maxinclusiveDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]MaxLengthDB:
        *ptr = make([]MaxLengthDB, 0, len(db.maxlengthDBs))
        for _, v := range db.maxlengthDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]MinInclusiveDB:
        *ptr = make([]MinInclusiveDB, 0, len(db.mininclusiveDBs))
        for _, v := range db.mininclusiveDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]MinLengthDB:
        *ptr = make([]MinLengthDB, 0, len(db.minlengthDBs))
        for _, v := range db.minlengthDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]PatternDB:
        *ptr = make([]PatternDB, 0, len(db.patternDBs))
        for _, v := range db.patternDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]RestrictionDB:
        *ptr = make([]RestrictionDB, 0, len(db.restrictionDBs))
        for _, v := range db.restrictionDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]SchemaDB:
        *ptr = make([]SchemaDB, 0, len(db.schemaDBs))
        for _, v := range db.schemaDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]SequenceDB:
        *ptr = make([]SequenceDB, 0, len(db.sequenceDBs))
        for _, v := range db.sequenceDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]SimpleContentDB:
        *ptr = make([]SimpleContentDB, 0, len(db.simplecontentDBs))
        for _, v := range db.simplecontentDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]SimpleTypeDB:
        *ptr = make([]SimpleTypeDB, 0, len(db.simpletypeDBs))
        for _, v := range db.simpletypeDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]TotalDigitDB:
        *ptr = make([]TotalDigitDB, 0, len(db.totaldigitDBs))
        for _, v := range db.totaldigitDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]UnionDB:
        *ptr = make([]UnionDB, 0, len(db.unionDBs))
        for _, v := range db.unionDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]WhiteSpaceDB:
        *ptr = make([]WhiteSpaceDB, 0, len(db.whitespaceDBs))
        for _, v := range db.whitespaceDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
    default:
        return nil, errors.New("github.com/fullstack-lang/gongxsd/go, Find: unsupported type")
    }
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gongxsd/go, Do not process when conds is not a single parameter")
	}

	str, ok := conds[0].(string)

	if !ok {
		return nil, errors.New("github.com/fullstack-lang/gongxsd/go, conds[0] is not a string")
	}

	i, err := strconv.ParseUint(str, 10, 32) // Base 10, 32-bit unsigned int
	if err != nil {
		return nil, errors.New("github.com/fullstack-lang/gongxsd/go, conds[0] is not a string number")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *AllDB:
		tmp, ok := db.allDBs[uint(i)]

		allDB, _ := instanceDB.(*AllDB)
		*allDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *AnnotationDB:
		tmp, ok := db.annotationDBs[uint(i)]

		annotationDB, _ := instanceDB.(*AnnotationDB)
		*annotationDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *AttributeDB:
		tmp, ok := db.attributeDBs[uint(i)]

		attributeDB, _ := instanceDB.(*AttributeDB)
		*attributeDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *AttributeGroupDB:
		tmp, ok := db.attributegroupDBs[uint(i)]

		attributegroupDB, _ := instanceDB.(*AttributeGroupDB)
		*attributegroupDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *ChoiceDB:
		tmp, ok := db.choiceDBs[uint(i)]

		choiceDB, _ := instanceDB.(*ChoiceDB)
		*choiceDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *ComplexContentDB:
		tmp, ok := db.complexcontentDBs[uint(i)]

		complexcontentDB, _ := instanceDB.(*ComplexContentDB)
		*complexcontentDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *ComplexTypeDB:
		tmp, ok := db.complextypeDBs[uint(i)]

		complextypeDB, _ := instanceDB.(*ComplexTypeDB)
		*complextypeDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *DocumentationDB:
		tmp, ok := db.documentationDBs[uint(i)]

		documentationDB, _ := instanceDB.(*DocumentationDB)
		*documentationDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *ElementDB:
		tmp, ok := db.elementDBs[uint(i)]

		elementDB, _ := instanceDB.(*ElementDB)
		*elementDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *EnumerationDB:
		tmp, ok := db.enumerationDBs[uint(i)]

		enumerationDB, _ := instanceDB.(*EnumerationDB)
		*enumerationDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *ExtensionDB:
		tmp, ok := db.extensionDBs[uint(i)]

		extensionDB, _ := instanceDB.(*ExtensionDB)
		*extensionDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *GroupDB:
		tmp, ok := db.groupDBs[uint(i)]

		groupDB, _ := instanceDB.(*GroupDB)
		*groupDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *LengthDB:
		tmp, ok := db.lengthDBs[uint(i)]

		lengthDB, _ := instanceDB.(*LengthDB)
		*lengthDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *MaxInclusiveDB:
		tmp, ok := db.maxinclusiveDBs[uint(i)]

		maxinclusiveDB, _ := instanceDB.(*MaxInclusiveDB)
		*maxinclusiveDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *MaxLengthDB:
		tmp, ok := db.maxlengthDBs[uint(i)]

		maxlengthDB, _ := instanceDB.(*MaxLengthDB)
		*maxlengthDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *MinInclusiveDB:
		tmp, ok := db.mininclusiveDBs[uint(i)]

		mininclusiveDB, _ := instanceDB.(*MinInclusiveDB)
		*mininclusiveDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *MinLengthDB:
		tmp, ok := db.minlengthDBs[uint(i)]

		minlengthDB, _ := instanceDB.(*MinLengthDB)
		*minlengthDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *PatternDB:
		tmp, ok := db.patternDBs[uint(i)]

		patternDB, _ := instanceDB.(*PatternDB)
		*patternDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *RestrictionDB:
		tmp, ok := db.restrictionDBs[uint(i)]

		restrictionDB, _ := instanceDB.(*RestrictionDB)
		*restrictionDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *SchemaDB:
		tmp, ok := db.schemaDBs[uint(i)]

		schemaDB, _ := instanceDB.(*SchemaDB)
		*schemaDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *SequenceDB:
		tmp, ok := db.sequenceDBs[uint(i)]

		sequenceDB, _ := instanceDB.(*SequenceDB)
		*sequenceDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *SimpleContentDB:
		tmp, ok := db.simplecontentDBs[uint(i)]

		simplecontentDB, _ := instanceDB.(*SimpleContentDB)
		*simplecontentDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *SimpleTypeDB:
		tmp, ok := db.simpletypeDBs[uint(i)]

		simpletypeDB, _ := instanceDB.(*SimpleTypeDB)
		*simpletypeDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *TotalDigitDB:
		tmp, ok := db.totaldigitDBs[uint(i)]

		totaldigitDB, _ := instanceDB.(*TotalDigitDB)
		*totaldigitDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *UnionDB:
		tmp, ok := db.unionDBs[uint(i)]

		unionDB, _ := instanceDB.(*UnionDB)
		*unionDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *WhiteSpaceDB:
		tmp, ok := db.whitespaceDBs[uint(i)]

		whitespaceDB, _ := instanceDB.(*WhiteSpaceDB)
		*whitespaceDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gongxsd/go, Unkown type")
	}
	
	return db, nil
}


// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gongxsd/test/books/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	booktypeDBs map[uint]*BookTypeDB

	nextIDBookTypeDB uint

	booksDBs map[uint]*BooksDB

	nextIDBooksDB uint

	creditDBs map[uint]*CreditDB

	nextIDCreditDB uint

	linkDBs map[uint]*LinkDB

	nextIDLinkDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		booktypeDBs: make(map[uint]*BookTypeDB),

		booksDBs: make(map[uint]*BooksDB),

		creditDBs: make(map[uint]*CreditDB),

		linkDBs: make(map[uint]*LinkDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongxsd/test/books/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *BookTypeDB:
		db.nextIDBookTypeDB++
		v.ID = db.nextIDBookTypeDB
		db.booktypeDBs[v.ID] = v
	case *BooksDB:
		db.nextIDBooksDB++
		v.ID = db.nextIDBooksDB
		db.booksDBs[v.ID] = v
	case *CreditDB:
		db.nextIDCreditDB++
		v.ID = db.nextIDCreditDB
		db.creditDBs[v.ID] = v
	case *LinkDB:
		db.nextIDLinkDB++
		v.ID = db.nextIDLinkDB
		db.linkDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gongxsd/test/books/go, unsupported type in Create")
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
		return nil, errors.New("github.com/fullstack-lang/gongxsd/test/books/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *BookTypeDB:
		delete(db.booktypeDBs, v.ID)
	case *BooksDB:
		delete(db.booksDBs, v.ID)
	case *CreditDB:
		delete(db.creditDBs, v.ID)
	case *LinkDB:
		delete(db.linkDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongxsd/test/books/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongxsd/test/books/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *BookTypeDB:
		db.booktypeDBs[v.ID] = v
		return db, nil
	case *BooksDB:
		db.booksDBs[v.ID] = v
		return db, nil
	case *CreditDB:
		db.creditDBs[v.ID] = v
		return db, nil
	case *LinkDB:
		db.linkDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongxsd/test/books/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongxsd/test/books/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *BookTypeDB:
		if existing, ok := db.booktypeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db BookType github.com/fullstack-lang/gongxsd/test/books/go, record not found")
		}
	case *BooksDB:
		if existing, ok := db.booksDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Books github.com/fullstack-lang/gongxsd/test/books/go, record not found")
		}
	case *CreditDB:
		if existing, ok := db.creditDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Credit github.com/fullstack-lang/gongxsd/test/books/go, record not found")
		}
	case *LinkDB:
		if existing, ok := db.linkDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Link github.com/fullstack-lang/gongxsd/test/books/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gongxsd/test/books/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]BookTypeDB:
		*ptr = make([]BookTypeDB, 0, len(db.booktypeDBs))
		for _, v := range db.booktypeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]BooksDB:
		*ptr = make([]BooksDB, 0, len(db.booksDBs))
		for _, v := range db.booksDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]CreditDB:
		*ptr = make([]CreditDB, 0, len(db.creditDBs))
		for _, v := range db.creditDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]LinkDB:
		*ptr = make([]LinkDB, 0, len(db.linkDBs))
		for _, v := range db.linkDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongxsd/test/books/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gongxsd/test/books/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gongxsd/test/books/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongxsd/test/books/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *BookTypeDB:
		tmp, ok := db.booktypeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First BookType Unkown entry %d", i))
		}

		booktypeDB, _ := instanceDB.(*BookTypeDB)
		*booktypeDB = *tmp
		
	case *BooksDB:
		tmp, ok := db.booksDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Books Unkown entry %d", i))
		}

		booksDB, _ := instanceDB.(*BooksDB)
		*booksDB = *tmp
		
	case *CreditDB:
		tmp, ok := db.creditDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Credit Unkown entry %d", i))
		}

		creditDB, _ := instanceDB.(*CreditDB)
		*creditDB = *tmp
		
	case *LinkDB:
		tmp, ok := db.linkDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Link Unkown entry %d", i))
		}

		linkDB, _ := instanceDB.(*LinkDB)
		*linkDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gongxsd/test/books/go, Unkown type")
	}
	
	return db, nil
}


package server

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	// OK when returned by MongoDB is really a float (0.0 = false, 1.0 = true)
	OK = float64(1)
)

// CurrentOps is returned by db.currentOp() and contains
// a list of all operations currently in-progress. The db.currentOp()
// operation will itself be an element of InProg[].
type CurrentOps struct {
	InProg []CurrentOp `bson:"inprog" json:"inprog"`
	Info   string      `bson:"info,omitempty" json:"info,omitempty"`
	Ok     float64     `bson:"ok" json:"ok"`
}

// CurrentOp is a database operation currently in-progress.
type CurrentOp struct {
	Active           bool   `bson:"active" json:"active"`
	OpID             uint32 `bson:"opid" json:"opid"`
	SecsRunning      uint32 `bson:"secs_running" json:"secs_running"`
	MicrosecsRunning uint64 `bson:"microsecs_running" json:"microsecs_running"`
	OpType           string `bson:"op" json:"op"`
	Namespace        string `bson:"ns" json:"ns"`
	KillPending      bool   `bson:"killPending" json:"killPending"`
	Query            bson.D `bson:"query" json:"query"`
}

// Reply is a response from a MongoDB command that doesn't return any results.
type Reply struct {
	Info string  `bson:"info,omitempty" json:"info,omitempty"`
	Ok   float64 `bson:"ok" json:"ok"`
}

// killLongRunningOps is intended to be run as a separate goroutine, off of
// the main server thread. killLongRunningOps periodically checks the admin
// database for long-running client-initiated operations (e.g. a slow pipeline)
// and kills those operations after the set Config.DatabaseOpTimeout.
//
// This is a common approach, similarly applied here:
// 1. https://blog.mlab.com/2014/02/mongodb-currentop-killop
// 2. https://dzone.com/articles/finding-and-terminating-long

// TODO: disabled as requires high-grade permissions. Remove completely?
func killLongRunningOps(ticker *time.Ticker, connectionString string, dbname string, config Config) {
	logKLRO(nil, fmt.Sprintf("Monitoring databases %s for long-running operations", config.DatabaseSuffix))

	session, err := mgo.Dial(connectionString)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	adminDB := session.DB(dbname)

	for now := range ticker.C {
		var err error
		ops := CurrentOps{}
		t := &now

		// This will return a set of client-initiated currentOps ONLY. There are numerous
		// more server operations that are returned when passed {"$all": true}.
		// see: https://docs.mongodb.com/manual/reference/command/currentOp/
		err = adminDB.Run("currentOp", &ops)

		if err != nil {
			logKLRO(t, err.Error())
		}

		if ops.Ok != OK {
			if ops.Info != "" {
				logKLRO(t, "!OK: "+ops.Info)
			} else {
				logKLRO(t, "!OK: No additional information")
			}
			continue
		}

		for _, op := range ops.InProg {

			// Only evaluate active operations.
			if !op.Active {
				continue
			}

			// Don't retry kills.
			if op.KillPending {
				continue
			}

			// Only interfere with operations on our database (e.g. "fhir").
			if !strings.HasSuffix(op.Namespace, config.DatabaseSuffix) {
				continue
			}

			// Check the current runtime.
			if float64(op.SecsRunning) < config.DatabaseOpTimeout.Seconds() {
				continue
			}

			// Operations that get here meet the following criteria:
			// 1. Have a runtime exceeding the current config.DatabaseOpTimeout
			// 2. Are in the config.DatabaseName namespace.
			switch op.OpType {
			// To protect data integrity, only kill these types of operations.
			// For a full list of command types, see:
			// https://docs.mongodb.com/manual/reference/command/currentOp/#currentOp.op
			case "command", "query", "getMore":
				if len(op.Query) == 0 {
					continue
				}

				queryDoc := op.Query[0]
				err = killOp(adminDB, op.OpID)
				if err != nil {
					logKLRO(t, err.Error())
					continue
				}

				// Successfully killed the operation.
				msg := fmt.Sprintf("killed op[%d] %s %s", op.OpID, queryDoc.Name, op.Namespace)
				logKLRO(t, msg)
			}
		}
	}
}

func killOp(adminDB *mgo.Database, opID uint32) error {
	var err error
	reply := Reply{}
	// see: https://docs.mongodb.com/manual/reference/command/killOp/
	err = adminDB.Run(bson.D{{Name: "killOp", Value: 1}, {Name: "op", Value: opID}}, &reply)
	if reply.Ok != OK {
		if reply.Info != "" {
			return errors.New(reply.Info)
		}
		return fmt.Errorf("Failed to kill op[%d]", opID)
	}
	return err
}

func logKLRO(t *time.Time, msg string) {
	if t != nil {
		log.Printf("%v KillLongRunningOps: %s\n", t, msg)
		return
	}
	log.Printf("KillLongRunningOps: %s\n", msg)
}

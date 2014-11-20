package jresp

// @author Robin Verlangen

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type JResp struct {
	data   map[string]interface{}
	tStart time.Time
}

func NewJsonResp() *JResp {
	return &JResp{
		data:   make(map[string]interface{}),
		tStart: time.Now(),
	}
}

func (jr *JResp) OK() {
	jr.set("status", "OK")
}

func (jr *JResp) Error(msg string) {
	jr.set("status", "ERROR")
	jr.set("error", msg)
}

func (jr *JResp) Set(k string, v interface{}) {
	if k == "status" || k == "error" || k == "parsetime" {
		log.Fatal(fmt.Sprintf("JSON key '%s' not allowed", k))
	}
	jr.set(k, v)
}

func (jr *JResp) set(k string, v interface{}) {
	jr.data[k] = v
}

func (jr *JResp) ToString() string {
	// Ensure we have a status
	if jr.data["status"] == nil {
		log.Fatal("Did not set status")
	}

	// Error logging
	if jr.data["status"] == "ERROR" {
		log.Println(fmt.Sprintf("Request error: %s", jr.data["error"]))
	}

	// Stop execution timer
	elapsed := time.Since(jr.tStart)
	jr.set("parsetime", fmt.Sprintf("%s", elapsed))

	// To JSON
	b, err := json.Marshal(jr.data)
	if err != nil {
		log.Println(fmt.Sprintf("error:", err))
		return ""
	}
	return string(b)
}


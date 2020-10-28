package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"go101.org/gold/internal/util"
)

const DurationToUpdate = time.Hour * 24 * 64

const (
	UpdateTip_Nothing = iota
	UpdateTip_ToUpdate
	UpdateTip_Updating
	UpdateTip_Updated
)

var UpdateTip2DivID = []string{
	UpdateTip_Nothing:  "",
	UpdateTip_ToUpdate: "to-update",
	UpdateTip_Updating: "updating",
	UpdateTip_Updated:  "updated",
}

// Must be called when locking.
func (ds *docServer) confirmUpdateTip() {
	if ds.updateTip == UpdateTip_Updating {
		return
	}

	d := time.Now().Sub(ds.roughBuildTime())
	needCheckUpdate := d > DurationToUpdate
	if needCheckUpdate {
		ds.updateTip = UpdateTip_ToUpdate
		ds.newerVersionInstalled = false
	} else if ds.newerVersionInstalled {
		ds.updateTip = UpdateTip_Updated
	} else {
		ds.updateTip = UpdateTip_Nothing
	}
}

// update page.
func (ds *docServer) startUpdatingGold() {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	ds.confirmUpdateTip()

	if ds.updateTip == UpdateTip_ToUpdate {
		ds.updateTip = UpdateTip_Updating
		go ds.updateGold()
	}
}

// api:update
// - GET: get current update info.
// - POST: do update
func (ds *docServer) updateAPI(w http.ResponseWriter, r *http.Request) {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	ds.confirmUpdateTip()

	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"updateStatus": "%s"}`, UpdateTip2DivID[ds.updateTip])
		return
	}

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "application/json")
		if ds.updateTip == UpdateTip_ToUpdate {
			ds.updateTip = UpdateTip_Updating
			go ds.updateGold()
		}
		fmt.Fprintf(w, `{"updateStatus": "%s"}`, UpdateTip2DivID[ds.updateTip])
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (ds *docServer) onUpdateDone(succeeded bool) {
	var now = time.Now()

	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	ds.roughBuildTime = func() time.Time {
		return now
	}
	ds.newerVersionInstalled = succeeded
	ds.updateTip = UpdateTip_Nothing
}

func (ds *docServer) updateGold() {
	if err := func() error {
		dir, err := ioutil.TempDir("", "*")
		if err != nil {
			return err
		}

		appPkgPath := "go101.org/golds"
		switch appName := filepath.Base(os.Args[0]); appName {
		case "godoge", "gocore", "gold":
			appPkgPath += "/" + appName
		case "golds":
		default:
			// ToDo: should update, then move the name
		}

		ds.updateLogger.Printf("Run: go get -u %s", appPkgPath)
		output, err := util.RunShellCommand(time.Minute*2, dir, []string{"GO111MODULE=on"}, "go", "get", "-u", appPkgPath)
		if len(output) > 0 {
			ds.updateLogger.Printf("\n%s\n", output)
		}
		if err != nil {
			return err
		}

		return nil
	}(); err != nil {
		ds.onUpdateDone(false)
		ds.updateLogger.Println("Update Gold error:", err)
	} else {
		ds.onUpdateDone(true)
		ds.updateLogger.Println("Update Gold succeeded.")
	}
}

package glusterfs

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	
	"github.com/boltdb/bolt"
	"github.com/heketi/heketi/pkg/glusterfs/api"
	"github.com/heketi/heketi/pkg/utils"
)

func (a *App) SnapshotDestroy(w http.ResponseWriter, r *http.Request) {
	type snapshotDestroyRequest struct {
		SnapshotId string `json:"snapshot_id"`
		VolumeId string `json:"volume_id"`
	}
	
	var msg snapshotDestroyRequest
	err := utils.GetJsonFromRequest(r, &msg)
	if err != nil {
		http.Error(w, "request unable to be parsed", http.StatusUnprocessableEntity)
		return
	}
	
	sshHost := ""
	var info *api.VolumeInfoResponse
	err = a.db.View(func(tx *bolt.Tx) error {
		entry, err := NewVolumeEntryFromId(tx, msg.VolumeId)
		if err == ErrNotFound || !entry.Visible() {
			// treat an invisible entry like it doesn't exist
			http.Error(w, "Id not found", http.StatusNotFound)
			return ErrNotFound
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
		
		info, err = entry.NewInfoResponse(tx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
		if sshHost == "" {
			sshHost = info.Mount.GlusterFS.Hosts[0]
		}
		
		return nil
	})
	
	if err != nil {
		return
	}
	op := NewSnapshotDestroyOperation(a.db, msg.SnapshotId, sshHost)
	if err := AsyncHttpOperation(a, w, r, op); err != nil {
		OperationHttpErrorf(w, err,
			"Failed destroy snapshot %v: %v", msg.SnapshotId, err)
		return
	}
}

func (a *App) SnapshotRestore(w http.ResponseWriter, r *http.Request) {
	type snapshotDestroyRequest struct {
		SnapshotId string `json:"snapshot_id"`
		VolumeId string `json:"volume_id"`
	}
	
	var msg snapshotDestroyRequest
	err := utils.GetJsonFromRequest(r, &msg)
	if err != nil {
		http.Error(w, "request unable to be parsed", http.StatusUnprocessableEntity)
		return
	}
	
	sshHost := ""
	var info *api.VolumeInfoResponse
	err = a.db.View(func(tx *bolt.Tx) error {
		entry, err := NewVolumeEntryFromId(tx, msg.VolumeId)
		if err == ErrNotFound || !entry.Visible() {
			// treat an invisible entry like it doesn't exist
			http.Error(w, "Id not found", http.StatusNotFound)
			return ErrNotFound
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
		
		info, err = entry.NewInfoResponse(tx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
		if sshHost == "" {
			sshHost = info.Mount.GlusterFS.Hosts[0]
		}
		
		return nil
	})
	
	if err != nil {
		return
	}
	volumeId := fmt.Sprintf("vol_%s", msg.VolumeId)
	op := NewSnapshotRestoreOperation(a.db, msg.SnapshotId, sshHost, volumeId)
	if err := AsyncHttpOperation(a, w, r, op); err != nil {
		OperationHttpErrorf(w, err,
			"Failed restore snapshot %v: %v", msg.SnapshotId, err)
		return
	}
}

func (a *App) SnapshotCreate(w http.ResponseWriter, r *http.Request) {
	
	var msg api.SnapshotRequest
	err := utils.GetJsonFromRequest(r, &msg)
	if err != nil {
		http.Error(w, "request unable to be parsed", http.StatusUnprocessableEntity)
		return
	}
	
	sshHost := ""
	var info *api.VolumeInfoResponse
	err = a.db.View(func(tx *bolt.Tx) error {
		entry, err := NewVolumeEntryFromId(tx, msg.VolumeId)
		if err == ErrNotFound || !entry.Visible() {
			// treat an invisible entry like it doesn't exist
			http.Error(w, "Id not found", http.StatusNotFound)
			return ErrNotFound
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
		
		info, err = entry.NewInfoResponse(tx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
		if sshHost == "" {
			sshHost = info.Mount.GlusterFS.Hosts[0]
		}
		
		return nil
	})
	
	if err != nil {
		return
	}
	volumeId := fmt.Sprintf("vol_%s", msg.VolumeId)
	op := NewSnapshotCreateOperation(a.db, msg.SnapshotId, sshHost, volumeId)
	if err := AsyncHttpOperation(a, w, r, op); err != nil {
		OperationHttpErrorf(w, err,
			"Failed create snapshot %v: %v", msg.SnapshotId, err)
		return
	}
}

func (a *App) SnapshotInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volumeId := vars["volume_id"]
	snapshotId := vars["snapshot_id"]
	
	sshHost := ""
	var info *api.VolumeInfoResponse
	err := a.db.View(func(tx *bolt.Tx) error {
		entry, err := NewVolumeEntryFromId(tx, volumeId)
		if err == ErrNotFound || !entry.Visible() {
			// treat an invisible entry like it doesn't exist
			http.Error(w, "Id not found", http.StatusNotFound)
			return ErrNotFound
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
		
		info, err = entry.NewInfoResponse(tx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
		if sshHost == "" {
			sshHost = info.Mount.GlusterFS.Hosts[0]
		}
		
		return nil
	})
	
	if err != nil {
		return
	}
	snap, err := a.executor.SnapshotInfo(sshHost, snapshotId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(snap); err != nil {
		panic(err)
	}
}


func (a *App) VolumeInfoDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volumeId := vars["volume_id"]
	
	sshHost := ""
	var info *api.VolumeInfoResponse
	err := a.db.View(func(tx *bolt.Tx) error {
		entry, err := NewVolumeEntryFromId(tx, volumeId)
		if err == ErrNotFound || !entry.Visible() {
			// treat an invisible entry like it doesn't exist
			http.Error(w, "Id not found", http.StatusNotFound)
			return ErrNotFound
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
		
		info, err = entry.NewInfoResponse(tx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
		if sshHost == "" {
			sshHost = info.Mount.GlusterFS.Hosts[0]
		}
		
		return nil
	})
	
	if err != nil {
		return
	}
	
	volumeInfo, err := a.executor.VolumeInfo(sshHost, volumeId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(volumeInfo); err != nil {
		panic(err)
	}
}

func (a *App) BrickInfoDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volumeId := vars["volume_id"]
	brickId := vars["brick_id"]
	
	sshHost := ""
	var info *api.VolumeInfoResponse
	err := a.db.View(func(tx *bolt.Tx) error {
		entry, err := NewVolumeEntryFromId(tx, volumeId)
		if err == ErrNotFound || !entry.Visible() {
			// treat an invisible entry like it doesn't exist
			http.Error(w, "Id not found", http.StatusNotFound)
			return ErrNotFound
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
		
		info, err = entry.NewInfoResponse(tx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
		if sshHost == "" {
			sshHost = info.Mount.GlusterFS.Hosts[0]
		}
		
		return nil
	})
	
	if err != nil {
		return
	}
	volume := fmt.Sprintf("vol_%s", volumeId)
	brick, err := a.executor.BrickInfo(sshHost, volume, brickId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(brick); err != nil {
		panic(err)
	}
}
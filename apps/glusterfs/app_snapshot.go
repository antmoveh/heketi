package glusterfs

import (
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
	
	op := NewSnapshotRestoreOperation(a.db, msg.SnapshotId, sshHost)
	if err := AsyncHttpOperation(a, w, r, op); err != nil {
		OperationHttpErrorf(w, err,
			"Failed restore snapshot %v: %v", msg.SnapshotId, err)
		return
	}
}
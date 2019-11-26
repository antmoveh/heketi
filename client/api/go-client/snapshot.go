
package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
	
	"github.com/heketi/heketi/pkg/glusterfs/api"
	"github.com/heketi/heketi/pkg/utils"
)

func (c *Client) SnapshotDestroy(request *api.SnapshotRequest) error {
	
	// Marshal request to JSON
	buffer, err := json.Marshal(request)
	if err != nil {
		return err
	}
	
	// Create a request
	req, err := http.NewRequest("Delete",
		c.host+"/snapshot/destroy",
		bytes.NewBuffer(buffer))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	
	// Set token
	err = c.setToken(req)
	if err != nil {
		return err
	}
	
	// Send request
	r, err := c.do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusAccepted {
		return utils.GetErrorFromResponse(r)
	}
	
	// Wait for response
	r, err = c.waitForResponseWithTimer(r, time.Second)
	if err != nil {
		return err
	}
	if r.StatusCode != http.StatusOK {
		return utils.GetErrorFromResponse(r)
	}
	
	// Read JSON response
	// var volume api.VolumeInfoResponse
	// err = utils.GetJsonFromResponse(r, &volume)
	// if err != nil {
	// 	return  err
	// }
	
	return  nil
	
}

func (c *Client) SnapshotRestore(request *api.SnapshotRequest) error {
	
	// Marshal request to JSON
	buffer, err := json.Marshal(request)
	if err != nil {
		return err
	}
	
	// Create a request
	req, err := http.NewRequest("PUT",
		c.host+"/snapshot/restore",
		bytes.NewBuffer(buffer))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	
	// Set token
	err = c.setToken(req)
	if err != nil {
		return err
	}
	
	// Send request
	r, err := c.do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusAccepted {
		return utils.GetErrorFromResponse(r)
	}
	
	// Wait for response
	r, err = c.waitForResponseWithTimer(r, time.Second)
	if err != nil {
		return  err
	}
	if r.StatusCode != http.StatusOK {
		return utils.GetErrorFromResponse(r)
	}
	
	// Read JSON response
	// var volume api.VolumeInfoResponse
	// err = utils.GetJsonFromResponse(r, &volume)
	// if err != nil {
	// 	return err
	// }
	
	return nil
}

func (c *Client) SnapshotCreate(request *api.SnapshotRequest) error {
	
	// Marshal request to JSON
	buffer, err := json.Marshal(request)
	if err != nil {
		return err
	}
	
	// Create a request
	req, err := http.NewRequest("POST",
		c.host+"/snapshot/create",
		bytes.NewBuffer(buffer))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	
	// Set token
	err = c.setToken(req)
	if err != nil {
		return err
	}
	
	// Send request
	r, err := c.do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusAccepted {
		return utils.GetErrorFromResponse(r)
	}
	
	// Wait for response
	r, err = c.waitForResponseWithTimer(r, time.Second)
	if err != nil {
		return  err
	}
	if r.StatusCode != http.StatusOK {
		return utils.GetErrorFromResponse(r)
	}
	
	// Read JSON response
	// var volume api.VolumeInfoResponse
	// err = utils.GetJsonFromResponse(r, &volume)
	// if err != nil {
	// 	return err
	// }
	
	return nil
}

func (c *Client) SnapshotInfo(request *api.SnapshotRequest) (*api.SnapshotInfoResponse, error) {
	
	// Create a request
	req, err := http.NewRequest("GET",
		c.host+"/snapshot/info/"+request.VolumeId+"/"+request.SnapshotId, nil)
	if err != nil {
		return nil, err
	}
	
	// Set token
	err = c.setToken(req)
	if err != nil {
		return nil, err
	}
	
	// Send request
	r, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return nil, utils.GetErrorFromResponse(r)
	}
	
	// Read JSON response
	var snap api.SnapshotInfoResponse
	err = utils.GetJsonFromResponse(r, &snap)
	if err != nil {
		return nil, err
	}
	
	return &snap, nil
}
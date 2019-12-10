
package client

import (
	"bytes"
	"encoding/json"
	"github.com/heketi/heketi/executors"
	"github.com/heketi/heketi/pkg/glusterfs/api"
	"github.com/heketi/heketi/pkg/utils"
	"net/http"
	"time"
)

func (c *Client) SnapshotDestroy(request *api.SnapshotRequest) (*api.SnapshotInfoResponse, error) {
	
	// Marshal request to JSON
	buffer, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	
	// Create a request
	req, err := http.NewRequest("DELETE",
		c.host+"/snapshot/destroy",
		bytes.NewBuffer(buffer))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	
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
	if r.StatusCode != http.StatusAccepted {
		return nil, utils.GetErrorFromResponse(r)
	}
	
	// Wait for response
	_, err = c.waitForResponseWithTimer(r, time.Second)
	if err != nil {
		return nil, err
	}
	
	return c.SnapshotInfo(request)
	
}

func (c *Client) SnapshotRestore(request *api.SnapshotRequest) (*api.SnapshotInfoResponse, error) {
	
	// Marshal request to JSON
	buffer, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	
	// Create a request
	req, err := http.NewRequest("PUT",
		c.host+"/snapshot/restore",
		bytes.NewBuffer(buffer))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	
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
	if r.StatusCode != http.StatusAccepted {
		return nil, utils.GetErrorFromResponse(r)
	}
	
	// Wait for response
	_, err = c.waitForResponseWithTimer(r, time.Second)
	if err != nil {
		return nil, err
	}
	
	return c.SnapshotInfo(request)
}

func (c *Client) SnapshotCreate(request *api.SnapshotRequest) (*api.SnapshotInfoResponse, error) {
	
	// Marshal request to JSON
	buffer, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	
	// Create a request
	req, err := http.NewRequest("POST",
		c.host+"/snapshot/create",
		bytes.NewBuffer(buffer))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	
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
	if r.StatusCode != http.StatusAccepted {
		return nil, utils.GetErrorFromResponse(r)
	}
	
	// Wait for response
	_, err = c.waitForResponseWithTimer(r, time.Second)
	if err != nil {
		return nil, err
	}
	
	return c.SnapshotInfo(request)
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

func (c *Client) VolumeInfoDetail(request *api.SnapshotRequest) (*executors.Volume, error) {
	
	// Create a request
	req, err := http.NewRequest("GET",
		c.host+"/volume/info/detail/"+request.VolumeId, nil)
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
	var volume executors.Volume
	err = utils.GetJsonFromResponse(r, &volume)
	if err != nil {
		return nil, err
	}
	
	return &volume, nil
}

func (c *Client) BrickInfoDetail(request *api.SnapshotRequest) (*executors.BrickDetailInfo, error) {
	
	// Marshal request to JSON
	buffer, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	
	// Create a request
	req, err := http.NewRequest("POST",
		c.host+"/brick/info/detail",
		bytes.NewBuffer(buffer))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	
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
	var brick executors.BrickDetailInfo
	err = utils.GetJsonFromResponse(r, &brick)
	if err != nil {
		return nil, err
	}
	
	return &brick, nil
}
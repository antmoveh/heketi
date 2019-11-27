//
// Copyright (c) 2015 The heketi Authors
//
// This file is licensed to you under your choice of the GNU Lesser
// General Public License, version 3 or any later version (LGPLv3 or
// later), or the GNU General Public License, version 2 (GPLv2), in all
// cases as published by the Free Software Foundation.
//

package executors

import (
	"encoding/xml"
)

type Executor interface {
	GlusterdCheck(host string) error
	PeerProbe(exec_host, newnode string) error
	PeerDetach(exec_host, detachnode string) error
	DeviceSetup(host, device, vgid string, destroy bool) (*DeviceInfo, error)
	GetDeviceInfo(host, device, vgid string) (*DeviceInfo, error)
	DeviceTeardown(host, device, vgid string) error
	BrickCreate(host string, brick *BrickRequest) (*BrickInfo, error)
	BrickDestroy(host string, brick *BrickRequest) (bool, error)
	BrickInfo(host string, volume, brick string) (*BrickDetailInfo, error)
	VolumeCreate(host string, volume *VolumeRequest) (*Volume, error)
	VolumeDestroy(host string, volume string) error
	VolumeDestroyCheck(host, volume string) error
	VolumeExpand(host string, volume *VolumeRequest) (*Volume, error)
	VolumeReplaceBrick(host string, volume string, oldBrick *BrickInfo, newBrick *BrickInfo) error
	VolumeInfo(host string, volume string) (*Volume, error)
	VolumeClone(host string, vsr *VolumeCloneRequest) (*Volume, error)
	VolumeSnapshot(host string, vsr *VolumeSnapshotRequest) (*Snapshot, error)
	SnapshotCloneVolume(host string, scr *SnapshotCloneRequest) (*Volume, error)
	SnapshotCloneBlockVolume(host string, scr *SnapshotCloneRequest) (*BlockVolumeInfo, error)
	SnapshotDestroy(host string, snapshot string) error
	HealInfo(host string, volume string) (*HealInfo, error)
	SetLogLevel(level string)
	BlockVolumeCreate(host string, blockVolume *BlockVolumeRequest) (*BlockVolumeInfo, error)
	BlockVolumeDestroy(host string, blockHostingVolumeName string, blockVolumeName string) error
	SnapshotRestore(host string, snapshot string, volumeId string) error
	SnapshotInfo(host, snapshotId string) (*SnapshotDetail, error)
}

// Enumerate durability types
type DurabilityType int

const (
	DurabilityNone DurabilityType = iota
	DurabilityReplica
	DurabilityDispersion
)

// Returns the size of the device
type DeviceInfo struct {
	// Size in KB
	TotalSize  uint64
	FreeSize   uint64
	UsedSize   uint64
	ExtentSize uint64
}

type BrickFormatType int

const (
	NormalFormat BrickFormatType = iota
	ArbiterFormat
)

// Brick description
type BrickRequest struct {
	VgId             string
	Name             string
	TpSize           uint64
	Size             uint64
	PoolMetadataSize uint64
	Gid              int64
	// Path is the brick mountpoint (named Path for symmetry with BrickInfo)
	Path string
	// lvm names
	TpName string
	LvName string
	Format BrickFormatType
}

// Returns information about the location of the brick
type BrickInfo struct {
	Path string
	Host string
}

type VolumeRequest struct {
	Bricks               []BrickInfo
	Name                 string
	Type                 DurabilityType
	GlusterVolumeOptions []string

	// Dispersion
	Data       int
	Redundancy int

	// Replica
	Replica int
	Arbiter bool
}

type VolumeCloneRequest struct {
	Volume string
	Clone  string
}

type VolumeSnapshotRequest struct {
	Volume      string
	Snapshot    string
	Description string
}

type SnapshotCloneRequest struct {
	Volume   string
	Snapshot string
}

type Snapshot struct {
	XMLName xml.Name `xml:"snapshot"`
	Name    string   `xml:"name"`
	UUID    string   `xml:"uuid"`
}

type SnapCreate struct {
	XMLName  xml.Name `xml:"snapCreate"`
	Snapshot Snapshot
}

type SnapClone struct {
	XMLName xml.Name `xml:"CloneCreate"`
	Volume  VolumeClone
}

type VolumeClone struct {
	XMLName xml.Name `xml:"volume"`
	Name    string   `xml:"name"`
	UUID    string   `xml:"uuid"`
}

type SnapDelete struct {
	XMLName   xml.Name         `xml:"snapDelete"`
	Snapshots []SnapshotStatus `xml:"snapshots"`
}

type SnapshotStatus struct {
	XMLName xml.Name `xml:"snapshot"`
	Status  string   `xml:"status"`
	Name    string   `xml:"name"`
	UUID    string   `xml:"uuid"`
}

type SnapActivate struct {
	XMLName  xml.Name `xml:"snapActivate"`
	Snapshot Snapshot
}

type SnapDeactivate struct {
	XMLName  xml.Name `xml:"snapDeactivate"`
	Snapshot Snapshot
}

type Brick struct {
	UUID      string `xml:"uuid,attr"`
	Name      string `xml:"name"`
	HostUUID  string `xml:"hostUuid"`
	IsArbiter int    `xml:"isArbiter"`
}

type Bricks struct {
	XMLName   xml.Name `xml:"bricks"`
	BrickList []Brick  `xml:"brick"`
}

type BrickHealStatus struct {
	HostUUID        string `xml:"hostUuid,attr"`
	Name            string `xml:"name"`
	Status          string `xml:"status"`
	NumberOfEntries string `xml:"numberOfEntries"`
}

type Option struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

type Options struct {
	XMLName    xml.Name `xml:"options"`
	OptionList []Option `xml:"option"`
}

type Volume struct {
	XMLName         xml.Name `xml:"volume"`
	VolumeName      string   `xml:"name"`
	ID              string   `xml:"id"`
	Status          int      `xml:"status"`
	StatusStr       string   `xml:"statusStr"`
	BrickCount      int      `xml:"brickCount"`
	DistCount       int      `xml:"distCount"`
	StripeCount     int      `xml:"stripeCount"`
	ReplicaCount    int      `xml:"replicaCount"`
	ArbiterCount    int      `xml:"arbiterCount"`
	DisperseCount   int      `xml:"disperseCount"`
	RedundancyCount int      `xml:"redundancyCount"`
	Type            int      `xml:"type"`
	TypeStr         string   `xml:"typeStr"`
	Transport       int      `xml:"transport"`
	Bricks          Bricks
	OptCount        int `xml:"optCount"`
	Options         Options
}

type Volumes struct {
	XMLName    xml.Name `xml:"volumes"`
	Count      int      `xml:"count"`
	VolumeList []Volume `xml:"volume"`
}

type VolInfo struct {
	XMLName xml.Name `xml:"volInfo"`
	Volumes Volumes  `xml:"volumes"`
}

type BrickDetailInfo struct {
	Volumes BrickVolumesInfo `xml:"volumes"`
}

type BrickVolumesInfo struct {
	Volumes BrickVolInfo `xml:"volume"`
}

type BrickVolInfo struct {
	VolName    string         `xml:"volName"`
	NodeCount  int            `xml:"nodeCount"`
	DetailInfo NodeDetailInfo `xml:"node"`
}

type NodeDetailInfo struct {
	Hostname    string `xml:"hostname"`
	Path        string `xml:"path"`
	PeerId      string `xml:"peerid"`
	Status      int    `xml:"status"`
	Port        int    `xml:"port"`
	Pid         int    `xml:"pid"`
	SizeTotal   uint64 `xml:"sizeTotal"`
	SizeFree    uint64 `xml:"sizeFree"`
	Device      string `xml:"device"`
	BlockSize   uint64 `xml:"blockSize"`
	FSName      string `xml:"fsName"`
	INodeSize   string `xml:"inodeSize"`
	INodesTotal uint64 `xml:"inodesTotal"`
	INodesFree  uint64 `xml:"inodesFree"`
}

type HealInfoBricks struct {
	BrickList []BrickHealStatus `xml:"brick"`
}

type HealInfo struct {
	XMLName xml.Name       `xml:"healInfo"`
	Bricks  HealInfoBricks `xml:"bricks"`
}

type BlockVolumeRequest struct {
	Name              string
	Size              int
	GlusterVolumeName string
	GlusterNode       string
	Hacount           int
	BlockHosts        []string
	Auth              bool
}

type BlockVolumeInfo struct {
	Name              string
	Size              int
	GlusterVolumeName string
	GlusterNode       string
	Hacount           int
	BlockHosts        []string
	Iqn               string
	Username          string
	Password          string
}

type VolumeDoesNotExistErr struct {
	Name string
}

func (dne *VolumeDoesNotExistErr) Error() string {
	return "Volume Does Not Exist: " + dne.Name
}

type SnapshotInfo struct {
	XMLName xml.Name `xml:"snapInfo"`
	Count int `xml:"count"`
	Snapshots Snapshots `xml:"snapshots"`
}

type Snapshots struct {
	XMLName xml.Name `xml:"snapshots"`
	SnapshotList []SnapshotDetail `xml:"snapshot"`
}

type SnapshotDetail struct {
	XMLName xml.Name `xml:"snapshot"`
	Name string `xml:"name"`
	Uuid string `json:"uuid"`
	CreateTime string `xml:"createTime"`
	VolCount int `xml:"volCount"`
	Error string `xml:"error"`
}
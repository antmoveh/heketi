[![Stories in Ready](https://badge.waffle.io/heketi/heketi.png?label=in%20progress&title=In%20Progress)](https://waffle.io/heketi/heketi)
[![Build Status](https://travis-ci.org/heketi/heketi.svg?branch=master)](https://travis-ci.org/heketi/heketi)
[![Coverage Status](https://coveralls.io/repos/heketi/heketi/badge.svg)](https://coveralls.io/r/heketi/heketi)
[![Go Report Card](https://goreportcard.com/badge/github.com/heketi/heketi)](https://goreportcard.com/report/github.com/heketi/heketi)

# Heketi
Heketi provides a RESTful management interface which can be used to manage the life cycle of GlusterFS volumes.  With Heketi, cloud services like OpenStack Manila, Kubernetes, and OpenShift can dynamically provision GlusterFS volumes with any of the supported durability types.  Heketi will automatically determine the location for bricks across the cluster, making sure to place bricks and its replicas across different failure domains.  Heketi also supports any number of GlusterFS clusters, allowing cloud services to provide network file storage without being limited to a single GlusterFS cluster.

# Workflow
When a request is received to create a volume, Heketi will first allocate the appropriate storage in a cluster, making sure to place brick replicas across failure domains.  It will then format, then mount the storage to create bricks for the volume requested.  Once all bricks have been automatically created, Heketi will finally satisfy the request by creating, then starting the newly created GlusterFS volume.

# Downloads

Heketi source code can be obtained via the
[project's releases page](https://github.com/heketi/heketi/releases)
or by cloning this repository.

# Documentation

Heketi's official documentation is located in the
[docs/ directory](https://github.com/heketi/heketi/tree/master/docs/)
within the repo.

# Demo
Please visit [Vagrant-Heketi](https://github.com/heketi/vagrant-heketi) to try out the demo.

# Community

* Mailing list: [Join our mailing list](http://lists.gluster.org/mailman/listinfo/heketi-devel)
* IRC: #heketi on Freenode

# Talks

* DevNation 2016

[![image](https://img.youtube.com/vi/gmEUnOmDziQ/3.jpg)](https://youtu.be/gmEUnOmDziQ)
[Slides](http://bit.ly/29avBJX)

* Devconf.cz 2016:

[![image](https://img.youtube.com/vi/jpkG4wciy4U/3.jpg)](https://www.youtube.com/watch?v=jpkG4wciy4U) [Slides](https://github.com/lpabon/go-slides)

# Updates

Please fetch branch `releses/8` for new supported features.
- Supported more fields to Brick for returning data usage in "volume info" API.
```shell
#HTTP REQUEST
curl -v http://127.0.0.1:8090/volumes/d56209edae84fda92500644532c20fb5

#HTTP RESPONSE
{
    "size": 2,
    "name": "vol_d56209edae84fda92500644532c20fb5",
    "durability": {
        "type": "replicate",
        "replicate": {
            "replica": 3
        },
        "disperse": {}
    },
    "snapshot": {
        "enable": false,
        "factor": 1
    },
    "id": "d56209edae84fda92500644532c20fb5",
    "cluster": "693d0f20edf4b10e789670ed11f8e702",
    "mount": {
        "glusterfs": {
            "hosts": [
                "10.203.40.99",
                "10.203.40.98",
                "10.203.40.97"
            ],
            "device": "10.203.40.99:vol_d56209edae84fda92500644532c20fb5",
            "options": {
                "backup-volfile-servers": "10.203.40.98,10.203.40.97"
            }
        }
    },
    "blockinfo": {},
    "bricks": [
        {
            "id": "7588e89959eb1f94bb696f7e3c5b73f7",
            "path": "/var/lib/heketi/mounts/vg_0adfc528cc4f9383cb0e7b3a672593b8/brick_7588e89959eb1f94bb696f7e3c5b73f7/brick",
            "device": "0adfc528cc4f9383cb0e7b3a672593b8",
            "node": "a239f268c58362d9ff3205f8b1de24b6",
            "volume": "d56209edae84fda92500644532c20fb5",
            "size_total": 2086912,      #new supported field
            "size_free": 1783008,       #new supported field
            "block_size": 4096,         #new supported field
            "inodes_total": 1048576,    #new supported field
            "inodes_free": 1048530,
            "status": 1,                #new supported field
            "host": "10.203.40.98",     #new supported field
            "size": 2097152
        },
        {
            "id": "bd8ffbab86122403cdcf9a2ca8e79a25",
            "path": "/var/lib/heketi/mounts/vg_d2b815bfb723ee885817f021c7f33f26/brick_bd8ffbab86122403cdcf9a2ca8e79a25/brick",
            "device": "d2b815bfb723ee885817f021c7f33f26",
            "node": "ee216dc39e70dbefffa454b090312ef8",
            "volume": "d56209edae84fda92500644532c20fb5",
            "size_total": 2086912,
            "size_free": 1783008,
            "block_size": 4096,
            "inodes_total": 1048576,
            "inodes_free": 1048530,
            "status": 1,
            "host": "10.203.40.97",
            "size": 2097152
        },
        {
            "id": "eec203b041262332ab604a036d705e01",
            "path": "/var/lib/heketi/mounts/vg_3e43e405ef3036799af859311bbbd1ed/brick_eec203b041262332ab604a036d705e01/brick",
            "device": "3e43e405ef3036799af859311bbbd1ed",
            "node": "3e6ea6254c1ec329e0dde5f46a4f0b3a",
            "volume": "d56209edae84fda92500644532c20fb5",
            "size_total": 0,
            "size_free": 0,
            "block_size": 0,
            "inodes_total": 0,
            "inodes_free": 0,
            "status": 0,
            "host": "10.203.40.99",
            "size": 2097152
        }
    ]
}
```


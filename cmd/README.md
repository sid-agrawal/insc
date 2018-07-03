# Commands

This package contains the commands which needs to be run as per the design document.



## host
This command is to be run on the host. It spawns the following different goroutines:
* aggregator: [github.com/sid-agrawal/insc/host/aggregator](github.com/sid-agrawal/insc/host/aggregator)

* db writer: [github.com/sid-agrawal/insc/host/db](github.com/sid-agrawal/insc/host/db)
* db reader: [github.com/sid-agrawal/insc/host/db](github.com/sid-agrawal/insc/host/db)
* publisher: [github.com/sid-agrawal/insc/host/publisher](github.com/sid-agrawal/insc/host/publisher)
* bluetooth server [github.com/sid-agrawal/insc/host/bluetoothserver](github.com/sid-agrawal/insc/host/bluetoothserver)


## remoteclient
This demonstrates the use of the API expoted by [github.com/sid-agrawal/insc/datasinks/remoteclient](github.com/sid-agrawal/insc/datasinks/remoteclient)


## blesensor
This cmd is meant to be run on a bluetooth enable client device and demonstrates the use of the API expoted by [github.com/sid-agrawal/insc/datasources/blesensor](github.com/sid-agrawal/insc/datasources/blesensor)


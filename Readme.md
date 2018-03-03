# Redis operator as dependency

This examples shows how you can use the redis-failover kubernetes CR client that is in the [Redis-operator][redis-operator].

## Usage

`make`

This will grab the correct dependencies and run the main program that only lists the redis failovers on the cluster.

**Note: this example uses your local kubeconfig configuration, this has been made on purpose for testing.**

[redis-operator]: https://github.com/spotahome/redis-operator
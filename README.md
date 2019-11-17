# Lugo - Arena

[![GoDoc](https://godoc.org/github.com/lugobots/arena?status.svg)](https://godoc.org/github.com/lugobots/arena)
[![Go Report Card](https://goreportcard.com/badge/github.com/lugobots/arena)](https://goreportcard.com/report/github.com/lugobots/arena)

Lugo - Arena is a [Go](http://golang.org/) module that provides some shareable features between the
game server and the clients of [Lugo](https://lugobots.dev/) game. This module is meant to be used by the  [Client Player](https://github.com/lugobots/client-player-go)
implemented in Go. However, you may implement another client and use this module as well.

If you wish to use part of this lib for any other project, please let me know if you find bugs, I will fix as soon as I can.   


Notes:

1. Most part of this library code is not tested at the current version (1.1.0). And there is no plans to 
   improve its tests. 
2. This module uses Lugo version 1.* constant values (distance, time, speed). Please, read the game documentation 
at the [Official website](https://lugobots.dev) for further information about all units. 

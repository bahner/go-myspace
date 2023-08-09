package app

import (
	"github.com/bahner/go-myspace/config"

	"github.com/ergo-services/ergo/node"
)

var (
	log             = config.Log
	nodeCookie      = config.NodeCookie
	nodeName        = config.NodeName
	myspaceNodeName = config.MyspaceNodeName
	appName         = config.AppName
	version         = config.Version
	description     = config.Description

	n node.Node
)

package main

import (
	"fmt"
	"os"
)

func (server *Server) getInitPeerId() string{
	var peerId string
	if peerId = os.Getenv("GO_FASTDFS_PEER_ID"); peerId == "" {
		peerId = fmt.Sprintf("%d", server.util.RandInt(0, 9))
	}
	return peerId
}

func (server *Server) getInitPort() string{
	var port string
	if port = os.Getenv("GO_FASTDFS_PORT"); port == "" {
		port = "8080"
	}
	return port
}

func (server *Server) getInitHost() string{
	var ip string
	if ip = os.Getenv("GO_FASTDFS_IP"); ip == "" {
		ip = server.util.GetPulicIP()
	}
	peer := "http://" + ip + ":" + server.getInitPort()
	return peer;
}

func (server *Server) getInitPeers() string{
	var peers string
	if peers = os.Getenv("GO_FASTDFS_PEERS"); peers == ""{
		peers = server.getInitHost()
	}
	return peers
}

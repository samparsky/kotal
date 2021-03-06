package controllers

import "os"

const (
	// PathConfig is the genesis file path
	PathConfig = "/mnt/config"
	// PathBlockchainData is the blockchain data path
	PathBlockchainData = "/mnt/data"
	// PathSecrets is the secrets (private keys, password ... etc) path
	PathSecrets = "/mnt/secrets"
)

// Images
const (
	// DefaultBesuImage is hyperledger besu image
	DefaultBesuImage = "hyperledger/besu:1.5.3"
	// DefaultGethImage is go-ethereum image
	DefaultGethImage = "ethereum/client-go:v1.9.20"
	// DefaultParityImage is parity image
	DefaultParityImage = "openethereum/openethereum:v3.0.1"
)

// Node settings
const (
	// DefaultHost is the host address used by rpc, ws and graphql server
	// rpcHost, wsHost, graphqlHost has been removed from node spec because
	// in geth v1.9.19 it has been removed along with graphqlPort
	// https://github.com/ethereum/go-ethereum/releases/tag/v1.9.19
	// in most use cases the value 0.0.0.0 makes most sense.
	DefaultHost = "0.0.0.0"
)

const (
	// EnvBesuImage is the environment variable used for hyperledger besu image
	EnvBesuImage = "BESU_IMAGE"
	// EnvGethImage is the environment variable used for go ethereum image
	EnvGethImage = "GETH_IMAGE"
	// EnvParityImage is the environment variable used for parity (OpenEthereum)
	EnvParityImage = "PARITY_IMAGE"
)

// GethImage returns geth docker image
func GethImage() string {
	if os.Getenv(EnvGethImage) == "" {
		return DefaultGethImage
	}
	return os.Getenv(EnvGethImage)
}

// BesuImage returns besu docker image
func BesuImage() string {
	if os.Getenv(EnvBesuImage) == "" {
		return DefaultBesuImage
	}
	return os.Getenv(EnvBesuImage)
}

// ParityImage returns parity docker image
func ParityImage() string {
	if os.Getenv(EnvParityImage) == "" {
		return DefaultParityImage
	}
	return os.Getenv(EnvParityImage)
}

// Hyperledger Besu client arguments
const (
	// BesuLogging is the argument used for logging verbosity level
	BesuLogging = "--logging"
	// BesuNetworkID is the argument used for network id
	BesuNetworkID = "--network-id"
	// BesuNatMethod is the argument used for nat method
	BesuNatMethod = "--nat-method"
	// BesuNodePrivateKey is the argument used for node private key
	BesuNodePrivateKey = "--node-private-key-file"
	// BesuGenesisFile is the argument used for genesis file
	BesuGenesisFile = "--genesis-file"
	// BesuDataPath is the argument used for data path
	BesuDataPath = "--data-path"
	// BesuNetwork is the argument used for selecting network
	BesuNetwork = "--network"
	// BesuP2PPort is the argument used for p2p port
	BesuP2PPort = "--p2p-port"
	// BesuBootnodes is the argument used for bootnodes
	BesuBootnodes = "--bootnodes"
	// BesuSyncMode is the argument used for sync mode
	BesuSyncMode = "--sync-mode"
	// BesuMinerEnabled is the argument used for turning on mining
	BesuMinerEnabled = "--miner-enabled"
	// BesuMinerCoinbase is the argument used for setting coinbase account
	BesuMinerCoinbase = "--miner-coinbase"
	// BesuRPCHTTPCorsOrigins is the argument used for setting rpc HTTP cors origins
	BesuRPCHTTPCorsOrigins = "--rpc-http-cors-origins"
	// BesuRPCHTTPEnabled is the argument used to enable RPC over HTTP
	BesuRPCHTTPEnabled = "--rpc-http-enabled"
	// BesuRPCHTTPPort is the argument used for RPC HTTP port
	BesuRPCHTTPPort = "--rpc-http-port"
	// BesuRPCHTTPHost is the argument used for RPC HTTP Host
	BesuRPCHTTPHost = "--rpc-http-host"
	// BesuRPCHTTPAPI is the argument used for RPC HTTP APIs
	BesuRPCHTTPAPI = "--rpc-http-api"
	// BesuRPCWSEnabled is the argument used to enable RPC WS
	BesuRPCWSEnabled = "--rpc-ws-enabled"
	// BesuRPCWSPort is the argument used for RPC WS port
	BesuRPCWSPort = "--rpc-ws-port"
	// BesuRPCWSHost is the argument used for RPC WS host
	BesuRPCWSHost = "--rpc-ws-host"
	// BesuRPCWSAPI is the argument used for RPC WS APIs
	BesuRPCWSAPI = "--rpc-ws-api"
	// BesuGraphQLHTTPEnabled is the argument used to enable QraphQL HTTP server
	BesuGraphQLHTTPEnabled = "--graphql-http-enabled"
	// BesuGraphQLHTTPPort is the argument used for GraphQL HTTP port
	BesuGraphQLHTTPPort = "--graphql-http-port"
	// BesuGraphQLHTTPHost is the argument used for GraphQL HTTP host
	BesuGraphQLHTTPHost = "--graphql-http-host"
	// BesuGraphQLHTTPCorsOrigins is the argument used for GraphQL HTTP Cors origins
	BesuGraphQLHTTPCorsOrigins = "--graphql-http-cors-origins"
	// BesuHostWhitelist is the argument used for whitelisting hosts
	BesuHostWhitelist = "--host-whitelist"
)

// Go ethereum client arguments
const (
	// GethLogging is the argument used for logging verbosity level
	GethLogging = "--verbosity"
	// GethNetworkID is the argument used for network id
	GethNetworkID = "--networkid"
	// GethNodeKey is the argument used for node private key
	GethNodeKey = "--nodekey"
	// GethDataDir is the argument used for data path
	GethDataDir = "--datadir"
	// GethP2PPort is the argument used for p2p port
	GethP2PPort = "--port"
	// GethBootnodes is the argument used for bootnodes
	GethBootnodes = "--bootnodes"
	// GethSyncMode is the argument used for sync mode
	GethSyncMode = "--syncmode"

	// GethMinerEnabled is the argument used for turning on mining
	GethMinerEnabled = "--mine"
	// GethMinerCoinbase is the argument used for setting coinbase account
	GethMinerCoinbase = "--miner.etherbase"

	// GethRPCHTTPCorsOrigins is the argument used for setting rpc HTTP cors origins
	GethRPCHTTPCorsOrigins = "--http.corsdomain"
	// GethRPCHTTPEnabled is the argument used to enable RPC over HTTP
	GethRPCHTTPEnabled = "--http"
	// GethRPCHTTPPort is the argument used for RPC HTTP port
	GethRPCHTTPPort = "--http.port"
	// GethRPCHTTPHost is the argument used for RPC HTTP Host
	GethRPCHTTPHost = "--http.addr"
	// GethRPCHTTPAPI is the argument used for RPC HTTP APIs
	GethRPCHTTPAPI = "--http.api"
	// GethRPCHostWhitelist is the argument used for whitelisting hosts
	GethRPCHostWhitelist = "--http.vhosts"

	// GethRPCWSEnabled is the argument used to enable RPC WS
	GethRPCWSEnabled = "--ws"
	// GethRPCWSPort is the argument used for RPC WS port
	GethRPCWSPort = "--ws.port"
	// GethRPCWSHost is the argument used for RPC WS host
	GethRPCWSHost = "--ws.addr"
	// GethRPCWSAPI is the argument used for RPC WS APIs
	GethRPCWSAPI = "--ws.api"

	// GethGraphQLHTTPEnabled is the argument used to enable QraphQL HTTP server
	GethGraphQLHTTPEnabled = "--graphql"
	// GethGraphQLHTTPCorsOrigins is the argument used for GraphQL HTTP Cors origins
	GethGraphQLHTTPCorsOrigins = "--graphql.corsdomain"
	// GethGraphQLHostWhitelist is the argument used for whitelisting hosts
	GethGraphQLHostWhitelist = "--graphql.vhosts"
	// GethUnlock is the argument used for unlocking imported ethereum account
	GethUnlock = "--unlock"
	// GethPassword is the argument used for locking imported ethereum address
	GethPassword = "--password"
)

// Parity client arguments
const (
	// ParityLogging is the argument used for logging verbosity level
	ParityLogging = "--logging"
	// ParityNetworkID is the argument used for network id
	ParityNetworkID = "--network-id"
	// ParityNodeKey is the argument used for node key
	ParityNodeKey = "--node-key"
	// ParityDataDir is the argument used for data path
	ParityDataDir = "--base-path"
	// ParityNetwork is the argument used for selecting network
	ParityNetwork = "--chain"
	// ParityP2PPort is the argument used for p2p port
	ParityP2PPort = "--port"
	// ParityBootnodes is the argument used for bootnodes
	ParityBootnodes = "--bootnodes"
	// ParitySyncMode is the argument used for sync mode
	ParitySyncMode = "--pruning"
	// ParityMinerCoinbase is the argument used for setting coinbase account
	ParityMinerCoinbase = "--author"
	// ParityEngineSigner is the argument used for engine singer
	ParityEngineSigner = "--engine-signer"

	// ParityRPCHTTPCorsOrigins is the argument used for setting rpc HTTP cors origins
	ParityRPCHTTPCorsOrigins = "--jsonrpc-cors"
	// ParityRPCHTTPPort is the argument used for RPC HTTP port
	ParityRPCHTTPPort = "--jsonrpc-port"
	// ParityRPCHTTPHost is the argument used for RPC HTTP Host
	ParityRPCHTTPHost = "--jsonrpc-interface"
	// ParityRPCHTTPAPI is the argument used for RPC HTTP APIs
	ParityRPCHTTPAPI = "--jsonrpc-apis"
	// ParityRPCHostWhitelist is the argument used for whitelisting hosts
	ParityRPCHostWhitelist = "--jsonrpc-hosts"

	// ParityRPCWSPort is the argument used for RPC WS port
	ParityRPCWSPort = "--ws-port"
	// ParityRPCWSHost is the argument used for RPC WS host
	ParityRPCWSHost = "--ws-interface"
	// ParityRPCWSAPI is the argument used for RPC WS APIs
	ParityRPCWSAPI = "--ws-apis"

	// ParityUnlock is the argument used for unlocking imported ethereum account
	ParityUnlock = "--unlock"
	// ParityPassword is the argument used for locking imported ethereum address
	ParityPassword = "--password"
)

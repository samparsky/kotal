package controllers

import (
	"encoding/json"
	"fmt"
	ethereumv1alpha1 "github.com/kotalco/kotal/apis/ethereum/v1alpha1"
	"strings"
)

// BesuClient is Hyperledger Besu client
type BesuClient struct{}

// GetArgs returns command line arguments required for client run
func (b *BesuClient) GetArgs(node *ethereumv1alpha1.Node, network *ethereumv1alpha1.Network, bootnodes []string) (args []string) {

	// appendArg appends argument with optional value to the arguments array
	appendArg := func(arg ...string) {
		args = append(args, arg...)
	}

	appendArg(BesuNatMethod, "KUBERNETES")

	if network.Spec.ID != 0 {
		appendArg(BesuNetworkID, fmt.Sprintf("%d", network.Spec.ID))
	}

	if node.WithNodekey() {
		appendArg(BesuNodePrivateKey, fmt.Sprintf("%s/nodekey", PathNodekey))
	}

	if network.Spec.Genesis != nil {
		appendArg(BesuGenesisFile, fmt.Sprintf("%s/genesis.json", PathGenesisFile))
	}

	appendArg(BesuDataPath, PathBlockchainData)

	if network.Spec.Join != "" {
		appendArg(BesuNetwork, network.Spec.Join)
	}

	if node.P2PPort != 0 {
		appendArg(BesuP2PPort, fmt.Sprintf("%d", node.P2PPort))
	}

	if len(bootnodes) != 0 {
		commaSeperatedBootnodes := strings.Join(bootnodes, ",")
		appendArg(BesuBootnodes, commaSeperatedBootnodes)
	}

	if node.SyncMode != "" {
		appendArg(BesuSyncMode, string(node.SyncMode))
	}

	if node.Miner {
		appendArg(BesuMinerEnabled)
	}

	if node.Coinbase != "" {
		appendArg(BesuMinerCoinbase, string(node.Coinbase))
	}

	if node.RPC {
		appendArg(BesuRPCHTTPEnabled)
	}

	if node.RPCPort != 0 {
		appendArg(BesuRPCHTTPPort, fmt.Sprintf("%d", node.RPCPort))
	}

	if node.RPCHost != "" {
		appendArg(BesuRPCHTTPHost, node.RPCHost)
	}

	if len(node.RPCAPI) != 0 {
		apis := []string{}
		for _, api := range node.RPCAPI {
			apis = append(apis, string(api))
		}
		commaSeperatedAPIs := strings.Join(apis, ",")
		appendArg(BesuRPCHTTPAPI, commaSeperatedAPIs)
	}

	if node.WS {
		appendArg(BesuRPCWSEnabled)
	}

	if node.WSPort != 0 {
		appendArg(BesuRPCWSPort, fmt.Sprintf("%d", node.WSPort))
	}

	if node.WSHost != "" {
		appendArg(BesuRPCWSHost, node.WSHost)
	}

	if len(node.WSAPI) != 0 {
		apis := []string{}
		for _, api := range node.WSAPI {
			apis = append(apis, string(api))
		}
		commaSeperatedAPIs := strings.Join(apis, ",")
		appendArg(BesuRPCWSAPI, commaSeperatedAPIs)
	}

	if node.GraphQL {
		appendArg(BesuGraphQLHTTPEnabled)
	}

	if node.GraphQLPort != 0 {
		appendArg(BesuGraphQLHTTPPort, fmt.Sprintf("%d", node.GraphQLPort))
	}

	if node.GraphQLHost != "" {
		appendArg(BesuGraphQLHTTPHost, node.GraphQLHost)
	}

	if len(node.Hosts) != 0 {
		commaSeperatedHosts := strings.Join(node.Hosts, ",")
		appendArg(BesuHostWhitelist, commaSeperatedHosts)
	}

	if len(node.CORSDomains) != 0 {
		commaSeperatedDomains := strings.Join(node.CORSDomains, ",")
		if node.RPC {
			appendArg(BesuRPCHTTPCorsOrigins, commaSeperatedDomains)
		}
		if node.GraphQL {
			appendArg(BesuGraphQLHTTPCorsOrigins, commaSeperatedDomains)
		}
	}

	return args
}

// GetGenesisFile returns genesis config parameter
func (b *BesuClient) GetGenesisFile(genesis *ethereumv1alpha1.Genesis, consensus ethereumv1alpha1.ConsensusAlgorithm) (content string, err error) {
	mixHash := genesis.MixHash
	nonce := genesis.Nonce
	difficulty := genesis.Difficulty
	result := map[string]interface{}{}
	var consensusConfig map[string]uint
	var extraData, engine string

	// ethash PoW settings
	if consensus == ethereumv1alpha1.ProofOfWork {
		consensusConfig = map[string]uint{}

		if genesis.Ethash.FixedDifficulty != nil {
			consensusConfig["fixeddifficulty"] = *genesis.Ethash.FixedDifficulty
		}

		engine = "ethash"
	}

	// clique PoA settings
	if consensus == ethereumv1alpha1.ProofOfAuthority {
		consensusConfig = map[string]uint{
			"blockperiodseconds": genesis.Clique.BlockPeriod,
			"epochlength":        genesis.Clique.EpochLength,
		}
		engine = "clique"
		extraData = createExtraDataFromSigners(genesis.Clique.Signers)
	}

	// clique ibft2 settings
	if consensus == ethereumv1alpha1.IstanbulBFT {

		consensusConfig = map[string]uint{
			"blockperiodseconds":        genesis.IBFT2.BlockPeriod,
			"epochlength":               genesis.IBFT2.EpochLength,
			"requesttimeoutseconds":     genesis.IBFT2.RequestTimeout,
			"messageQueueLimit":         genesis.IBFT2.MessageQueueLimit,
			"duplicateMessageLimit":     genesis.IBFT2.DuplicateMessageLimit,
			"futureMessagesLimit":       genesis.IBFT2.FutureMessagesLimit,
			"futureMessagesMaxDistance": genesis.IBFT2.FutureMessagesMaxDistance,
		}
		engine = "ibft2"
		mixHash = "0x63746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365"
		nonce = "0x0"
		difficulty = "0x1"
		extraData, err = createExtraDataFromValidators(genesis.IBFT2.Validators)
		if err != nil {
			return
		}
	}

	config := map[string]interface{}{
		"chainId":             genesis.ChainID,
		"homesteadBlock":      genesis.Forks.Homestead,
		"eip150Block":         genesis.Forks.EIP150,
		"eip150Hash":          genesis.Forks.EIP150Hash,
		"eip155Block":         genesis.Forks.EIP155,
		"eip158Block":         genesis.Forks.EIP158,
		"byzantiumBlock":      genesis.Forks.Byzantium,
		"constantinopleBlock": genesis.Forks.Constantinople,
		"petersburgBlock":     genesis.Forks.Petersburg,
		"istanbulBlock":       genesis.Forks.Istanbul,
		"muirGlacierBlock":    genesis.Forks.MuirGlacier,
		engine:                consensusConfig,
	}

	if genesis.Forks.DAO != nil {
		config["daoForkBlock"] = genesis.Forks.DAO
	}

	result["config"] = config
	result["nonce"] = nonce
	result["timestamp"] = genesis.Timestamp
	result["gasLimit"] = genesis.GasLimit
	result["difficulty"] = difficulty
	result["coinbase"] = genesis.Coinbase
	result["mixHash"] = mixHash
	result["extraData"] = extraData

	alloc := map[ethereumv1alpha1.EthereumAddress]interface{}{}
	for _, account := range genesis.Accounts {
		m := map[string]interface{}{
			"balance": account.Balance,
		}

		if account.Code != "" {
			m["code"] = account.Code
		}

		if account.Storage != nil {
			m["storage"] = account.Storage
		}

		alloc[account.Address] = m
	}

	result["alloc"] = alloc

	data, err := json.Marshal(result)
	if err != nil {
		return
	}

	content = string(data)

	return
}

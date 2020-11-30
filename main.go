package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

var tags = []string{"kindynos", "erc20", "ethereum"}

type ListVersion struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

type ListToken struct {
	ChainID  int      `json:"chainId"`
	Address  string   `json:"address"`
	Symbol   string   `json:"symbol"`
	Name     string   `json:"name"`
	Decimals int      `json:"decimals"`
	LogoURI  string   `json:"logoURI"`
	Tags     []string `json:"tags"`
}

type ListTag struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type List struct {
	Name      string             `json:"name"`
	LogoURI   string             `json:"logoURI"`
	Keywords  []string           `json:"keywords"`
	Timestamp time.Time          `json:"timestamp"`
	Tokens    []ListToken        `json:"tokens"`
	Version   ListVersion        `json:"version"`
	Tags      map[string]ListTag `json:"tags"`
}

var ListData = List{
	Name:      "Kindynos Token List",
	LogoURI:   "https://raw.githubusercontent.com/grupokindynos/kindynos-branding/master/kindynos/kindynos-isotype.png",
	Keywords:  []string{"defi", "tokens", "erc20", "kindynos"},
	Timestamp: time.Now(),
	Tags: map[string]ListTag{
		"stablecoin": {
			"Stablecoin",
			"Tokens that are fixed to an external asset, e.g. the US dollar",
		},
		"yieldfarm": {
			"YieldFarming",
			"Tokens that earn rewards based on liquidity providing",
		},
	},
	Tokens: []ListToken{
		{
			ChainID:  1,
			Address:  "0x622f2962ae78e8686ecc1e30cf2f9a6e5ac35626",
			Symbol:   "WPOLIS",
			Name:     "WrappedPolis",
			Decimals: 18,
			LogoURI:  "https://raw.githubusercontent.com/grupokindynos/kindynos-branding/master/polis/polis-isotype-background-round.png",
			Tags: tags,
		},
		{
			ChainID:  1,
			Address:  "0xf56842af3b56fd72d17cb103f92d027bba912e89",
			Symbol:   "BAMBOO",
			Name:     "BambooDeFi",
			Decimals: 18,
			LogoURI:  "https://raw.githubusercontent.com/bamboo-defi/frontend/main/src/assets/bamboo-head.png",
			Tags: tags,
		},
		{
			Name:     "Wrapped Ether",
			Address:  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
			Symbol:   "WETH",
			Decimals: 18,
			ChainID:  1,
			LogoURI:  "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2/logo.png",
			Tags: tags,
		},
		{
			Name: "Dai Stablecoin",
			Address: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
			Symbol: "DAI",
			Decimals: 18,
			ChainID: 1,
			LogoURI: "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0x6B175474E89094C44Da98b954EedeAC495271d0F/logo.png",
			Tags: tags,
		},
		{
			Name: "USDCoin",
			Address: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			Symbol: "USDC",
			Decimals: 6,
			ChainID: 1,
			LogoURI: "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48/logo.png",
			Tags: tags,
		},
		{
			Name: "Tether USD",
			Address: "0xdAC17F958D2ee523a2206206994597C13D831ec7",
			Symbol: "USDT",
			Decimals: 6,
			ChainID: 1,
			LogoURI: "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0xdAC17F958D2ee523a2206206994597C13D831ec7/logo.png",
			Tags: tags,
		},
		{
			Name: "HUSD",
			Address: "0xdF574c24545E5FfEcb9a659c229253D4111d87e1",
			Symbol: "HUSD",
			Decimals: 8,
			ChainID: 1,
			LogoURI: "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0xdF574c24545E5FfEcb9a659c229253D4111d87e1/logo.png",
			Tags: tags,
		},
		{
			Name: "Synthetix Network Token",
			Address: "0xC011a73ee8576Fb46F5E1c5751cA3B9Fe0af2a6F",
			Symbol: "SNX",
			Decimals: 18,
			ChainID: 1,
			LogoURI: "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0xC011a73ee8576Fb46F5E1c5751cA3B9Fe0af2a6F/logo.png",
			Tags: tags,
		},
		{
			Name: "ChainLink Token",
			Address: "0x514910771AF9Ca656af840dff83E8264EcF986CA",
			Symbol: "LINK",
			Decimals: 18,
			ChainID: 1,
			LogoURI: "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0x514910771AF9Ca656af840dff83E8264EcF986CA/logo.png",
			Tags: tags,
		},
		{
			Name: "Compound",
			Address: "0xc00e94Cb662C3520282E6f5717214004A7f26888",
			Symbol: "COMP",
			Decimals: 18,
			ChainID: 1,
			LogoURI: "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0xc00e94Cb662C3520282E6f5717214004A7f26888/logo.png",
			Tags: tags,
		},
		{
			Name: "Uniswap",
			Address: "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984",
			Symbol: "UNI",
			Decimals: 18,
			ChainID: 1,
			LogoURI: "ipfs://QmXttGpZrECX5qCyXbBQiqgQNytVGeZW5Anewvh2jc4psg",
			Tags: tags,
		},
		{
			Name: "EthLend Token",
			Address: "0x80fB784B7eD66730e8b1DBd9820aFD29931aab03",
			Symbol: "LEND",
			Decimals: 18,
			ChainID: 1,
			LogoURI: "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0x80fB784B7eD66730e8b1DBd9820aFD29931aab03/logo.png",
			Tags: tags,
		},
		{
			Name: "Wrapped BTC",
			Address: "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599",
			Symbol: "WBTC",
			Decimals: 8,
			ChainID: 1,
			LogoURI: "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599/logo.png",
			Tags: tags,
		},
	},
	Version: ListVersion{Major: 1, Minor: 0, Patch: 2},
}

func main() {
	raw, err := json.Marshal(ListData)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./kindynos.tokens.json", raw, 0700)
	if err != nil {
		panic(err)
	}
}

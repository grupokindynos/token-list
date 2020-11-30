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
			Tags:     tags,
		},
		{
			ChainID:  1,
			Address:  "0xf56842af3b56fd72d17cb103f92d027bba912e89",
			Symbol:   "BAMBOO",
			Name:     "BambooDeFi",
			Decimals: 18,
			LogoURI:  "https://raw.githubusercontent.com/bamboo-defi/frontend/main/src/assets/bamboo-head.png",
			Tags:     tags,
		},
		{
			Name:     "Wrapped Ether",
			Address:  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
			Symbol:   "WETH",
			Decimals: 18,
			ChainID:  1,
			LogoURI:  "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2/logo.png",
			Tags:     tags,
		},
		{
			Name:     "Dai Stablecoin",
			Address:  "0x6B175474E89094C44Da98b954EedeAC495271d0F",
			Symbol:   "DAI",
			Decimals: 18,
			ChainID:  1,
			LogoURI:  "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0x6B175474E89094C44Da98b954EedeAC495271d0F/logo.png",
			Tags:     tags,
		},
		{
			Name:     "USDCoin",
			Address:  "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			Symbol:   "USDC",
			Decimals: 6,
			ChainID:  1,
			LogoURI:  "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48/logo.png",
			Tags:     tags,
		},
		{
			Name:     "Tether USD",
			Address:  "0xdAC17F958D2ee523a2206206994597C13D831ec7",
			Symbol:   "USDT",
			Decimals: 6,
			ChainID:  1,
			LogoURI:  "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0xdAC17F958D2ee523a2206206994597C13D831ec7/logo.png",
			Tags:     tags,
		},
		{
			Name:     "HUSD",
			Address:  "0xdF574c24545E5FfEcb9a659c229253D4111d87e1",
			Symbol:   "HUSD",
			Decimals: 8,
			ChainID:  1,
			LogoURI:  "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0xdF574c24545E5FfEcb9a659c229253D4111d87e1/logo.png",
			Tags:     tags,
		},
		{
			Name:     "Synthetix Network Token",
			Address:  "0xC011a73ee8576Fb46F5E1c5751cA3B9Fe0af2a6F",
			Symbol:   "SNX",
			Decimals: 18,
			ChainID:  1,
			LogoURI:  "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0xC011a73ee8576Fb46F5E1c5751cA3B9Fe0af2a6F/logo.png",
			Tags:     tags,
		},
		{
			Name:     "ChainLink Token",
			Address:  "0x514910771AF9Ca656af840dff83E8264EcF986CA",
			Symbol:   "LINK",
			Decimals: 18,
			ChainID:  1,
			LogoURI:  "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0x514910771AF9Ca656af840dff83E8264EcF986CA/logo.png",
			Tags:     tags,
		},
		{
			Name:     "Compound",
			Address:  "0xc00e94Cb662C3520282E6f5717214004A7f26888",
			Symbol:   "COMP",
			Decimals: 18,
			ChainID:  1,
			LogoURI:  "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0xc00e94Cb662C3520282E6f5717214004A7f26888/logo.png",
			Tags:     tags,
		},
		{
			Name:     "Uniswap",
			Address:  "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984",
			Symbol:   "UNI",
			Decimals: 18,
			ChainID:  1,
			LogoURI:  "ipfs://QmXttGpZrECX5qCyXbBQiqgQNytVGeZW5Anewvh2jc4psg",
			Tags:     tags,
		},
		{
			Name:     "EthLend Token",
			Address:  "0x80fB784B7eD66730e8b1DBd9820aFD29931aab03",
			Symbol:   "LEND",
			Decimals: 18,
			ChainID:  1,
			LogoURI:  "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0x80fB784B7eD66730e8b1DBd9820aFD29931aab03/logo.png",
			Tags:     tags,
		},
		{
			Name:     "Wrapped BTC",
			Address:  "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599",
			Symbol:   "WBTC",
			Decimals: 8,
			ChainID:  1,
			LogoURI:  "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/ethereum/assets/0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599/logo.png",
			Tags:     tags,
		},
		{
			Name:     "TrueUSD",
			Address:  "0x0000000000085d4780b73119b644ae5ecd22b376",
			Symbol:   "TUSD",
			Decimals: 18,
			ChainID:  1,
			LogoURI:  "https://assets.coingecko.com/coins/images/3449/thumb/TUSD.png?1559172762",
			Tags:     tags,
		},
		{
			Name:     "Crypto.com Coin",
			Address:  "0xa0b73e1ff0b80914ab6fe0444e65848c4c34450b",
			Symbol:   "CRO",
			Decimals: 8,
			ChainID:  1,
			LogoURI:  "https://assets.coingecko.com/coins/images/7310/thumb/cypto.png?1547043960",
			Tags:     tags,
		},
		{
			Name:     "yearn finance",
			Address:  "0x0bc529c00c6401aef6d220be8c6ea1667f6ad93e",
			Symbol:   "YFI",
			Decimals: 18,
			ChainID:  1,
			LogoURI:  "https://assets.coingecko.com/coins/images/11849/thumb/yfi-192x192.png?1598325330",
			Tags:     tags,
		},
		{
			Name:     "Harvest Finance",
			Address:  "0xa0246c9032bc3a600820415ae600c6388619a14d",
			Symbol:   "FARM",
			Decimals: 18,
			ChainID:  1,
			LogoURI:  "https://assets.coingecko.com/coins/images/12304/thumb/Harvest.png?1599007988",
			Tags:     tags,
		},
		{
			Name:     "OKB",
			Address:  "0x75231f58b43240c9718dd58b4967c5114342a86c",
			Symbol:   "OKB",
			Decimals: 18,
			ChainID:  1,
			LogoURI:  "https://assets.coingecko.com/coins/images/4463/thumb/okb_token.png?1548386209",
			Tags:     tags,
		},
		{
			Name:     "Maker",
			Address:  "0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2",
			Symbol:   "MKR",
			Decimals: 18,
			ChainID:  1,
			LogoURI:  "https://assets.coingecko.com/coins/images/1364/thumb/Mark_Maker.png?1585191826",
			Tags:     tags,
		},
		{
			Name:     "Sushi",
			Address:  "0x6b3595068778dd592e39a122f4f5a5cf09c90fe2",
			Symbol:   "SUSHI",
			Decimals: 18,
			ChainID:  1,
			LogoURI:  "https://assets.coingecko.com/coins/images/12271/thumb/sushi.jpg?1598623048",
			Tags:     tags,
		},
		{
			Name:     "FTX Token",
			Address:  "0x50d1c9771902476076ecfc8b2a83ad6b9355a4c9",
			Symbol:   "FTT",
			Decimals: 18,
			ChainID:  1,
			LogoURI:  "https://assets.coingecko.com/coins/images/9026/thumb/ftt.png?1563776835",
			Tags:     tags,
		},
		{
			Name:     "Wrapped ZCASH",
			Address:  "0x4a64515e5e1d1073e83f30cb97bed20400b66e10",
			Symbol:   "WZEC",
			Decimals: 18,
			ChainID:  1,
			LogoURI:  "https://assets.coingecko.com/coins/images/13239/thumb/WZEC-icon.png?1606630700",
			Tags:     tags,
		},
		{
			Name:     "Probit Token",
			Address:  "0xfb559ce67ff522ec0b9ba7f5dc9dc7ef6c139803",
			Symbol:   "PROB",
			Decimals: 18,
			ChainID:  1,
			LogoURI:  "https://s2.coinmarketcap.com/static/img/coins/128x128/4586.png",
			Tags:     tags,
		},
		{
			Name:     "Eterbase Utility Token",
			Address:  "0x5bdC00B6676579b301B276198Db1ea9AffB94329",
			Symbol:   "XBASE",
			Decimals: 18,
			ChainID:  1,
			LogoURI:  "https://s2.coinmarketcap.com/static/img/coins/128x128/3815.png",
			Tags:     tags,
		},
		{
			Name:     "KuCoin Shares",
			Address:  "0x039b5649a59967e3e936d7471f9c3700100ee1ab",
			Symbol:   "KCS",
			Decimals: 6,
			ChainID:  1,
			LogoURI:  "https://s2.coinmarketcap.com/static/img/coins/128x128/2087.png",
			Tags:     tags,
		},
		{
			Name:     "Huobi Token",
			Address:  "0x6f259637dcd74c767781e37bc6133cd6a68aa161",
			Symbol:   "HT",
			Decimals: 18,
			ChainID:  1,
			LogoURI:  "https://s2.coinmarketcap.com/static/img/coins/128x128/2502.png",
			Tags:     tags,
		},
	},
	Version: ListVersion{Major: 1, Minor: 0, Patch: 4},
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

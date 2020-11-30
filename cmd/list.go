package cmd

import "time"

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

type List struct {
	Name      string      `json:"name"`
	LogoURI   string      `json:"logoURI"`
	Keywords  []string    `json:"keywords"`
	Timestamp time.Time   `json:"timestamp"`
	Tokens    []ListToken `json:"tokens"`
	Version   ListVersion `json:"version"`
}

var ListData = List{
	Name:     "Kindynos Token List",
	LogoURI:  "https://raw.githubusercontent.com/grupokindynos/kindynos-branding/master/kindynos/kindynos-isotype.png",
	Keywords: []string{"defi", "tokens", "erc20", "kindynos"},
	Timestamp: time.Now(),
	Tokens: []ListToken{
		{
			ChainID: 1,
			Address: "0x622f2962ae78e8686ecc1e30cf2f9a6e5ac35626",
			Symbol: "WPOLIS",
			Name:"WrappedPolis",
			Decimals: 18,
			LogoURI: "https://raw.githubusercontent.com/grupokindynos/kindynos-branding/master/polis/polis-isotype-background-round.png",
			Tags: []string{"polis", "erc20", "defi"},
		},
	},
	Version: ListVersion{Major: 1, Minor: 0, Patch: 0},
}

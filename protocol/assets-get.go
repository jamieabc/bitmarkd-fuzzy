package protocol

import "local/random"

type AssetsGetRpc struct {
	ID     string      `json:"id,omitempty"`
	Method string      `json:"method,omitempty"`
	Params []AssetsGet `json:"params,omitempty"`
}

type AssetsGet struct {
	FingerPrints []string `json:"fingerprints"`
}

func (rpc *AssetsGetRpc) JustifyData() {
	rpc.Method = "Assets.Get"
	rpc.ID = "1"
}

// genRandomData generates random data fits specific interface
func (rpc *AssetsGetRpc) GenRandomData() {
	r := random.New()
	r.Fuzz(rpc)
}

// sampleData generates correct data
func (rpc *AssetsGetRpc) SampleData() {
	rpc.ID = "1"
	rpc.Method = "Assets.Get"
	rpc.Params = []AssetsGet{
		AssetsGet{
			FingerPrints: []string{
				"015b9c9fb3e993bf64500977844104fc0b70ef5ca99141e9f56e2e837ce668fbe787643c34a0d51a32a82408eb36a6e93f7badbc5af50de29d9401b5affe564440",
				"e12571eff187f9a88e8a516c639fc51ef6ff9472fd39fd326cec33e266ae9ce74863f428f1e153f724b19b4b1d26df586f1ea3b794a5ca617b37129d315e3918"},
		},
	}
}

//go:generate mockgen -package httpclient -source=httpclient.go -destination httpclient_mock.go

package httpclient

type HexHTTPClient struct {

}

func (hc *HexHTTPClient) GetHexString() error {
	return nil
}

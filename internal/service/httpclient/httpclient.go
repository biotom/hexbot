//go:generate mockgen -package httpclient -source=httpclient.go -destination httpclient_mock.go

package httpclient

import "context"

type HexHTTPClient struct {

}

func (hc *HexHTTPClient) GetHexString(ctx context.Context) error {
	return nil
}

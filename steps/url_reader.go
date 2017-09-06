package steps

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"github.com/rai-project/pipeline"
)

type urlReader struct {
	base
}

func NewURLReader(ctx context.Context) (pipeline.Step, error) {
	var res urlReader
	return res.New(ctx)
}

func (p urlReader) New(ctx context.Context) (pipeline.Step, error) {
	return p, nil
}

func (p urlReader) Info() string {
	return "URLReader"
}

func (p urlReader) do(ctx context.Context, in0 interface{}) interface{} {
	in, ok := in0.(string)
	if !ok {
		return errors.Errorf("expecting a string for url reader Step, but got %v", in0)
	}
	resp, err := http.Get(in)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.Errorf("bad response code: %d", resp.StatusCode)
	}
	if resp.StatusCode != 200 {
		return errors.Errorf("bad response code: %d", resp.StatusCode)
	}

	res := new(bytes.Buffer)
	_, err = io.Copy(res, resp.Body)
	if err != nil {
		return errors.Errorf("unable to copy body")
	}
	return res
}

func (p urlReader) Close() error {
	return nil
}

func init() {
	pipeline.Register(urlReader{})
}

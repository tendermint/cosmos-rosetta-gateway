package clienttest

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

const (
	starportReleaseVersion    = "0.0.10"
	starportReleaseURLPattern = "https://github.com/tendermint/starport/releases/download/v%[1]v/starport_%[1]v_%[2]v_amd64.tar.gz"
	starportReleaseFileName   = "starport"
	starportBinaryName        = ".starport_test"
)

var installOnce sync.Once

func starportBinaryPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, starportBinaryName), nil
}

func starportURL(version string) string {
	return fmt.Sprintf(starportReleaseURLPattern, version, runtime.GOOS)
}

func installStarport(ctx context.Context) error {
	u := starportURL(starportReleaseVersion)
	path, err := starportBinaryPath()
	if err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return untarFile(starportReleaseFileName, f, resp.Body)
}

func untarFile(name string, out io.Writer, r io.Reader) error {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gzr.Close()
	tr := tar.NewReader(gzr)
	for {
		header, err := tr.Next()
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		}
		if header.Name != name || header.Typeflag != tar.TypeReg {
			continue
		}
		_, err = io.Copy(out, tr)
		return err
	}
}

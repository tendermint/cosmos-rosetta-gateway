package clienttest

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/tendermint/starport/starport/pkg/cmdrunner"
	"github.com/tendermint/starport/starport/pkg/cmdrunner/step"
	"github.com/tendermint/starport/starport/pkg/httpstatuschecker"
	"golang.org/x/sync/errgroup"
)

var isCI, _ = strconv.ParseBool(os.Getenv("CI"))

// Ctx returns a new context with deadline.
func Ctx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Minute*3)
}

type Environment struct {
	SDKAddr, TendermintAddr string

	appName     string
	serveCancel context.CancelFunc
	serveG      *errgroup.Group
}

// NewLaunchpad creates a new Launchpad testing environment.
func NewLaunchpad(ctx context.Context, appName string) (*Environment, error) {
	return New(ctx, appName)
}

// New creates a new testing environment.
func New(ctx context.Context, appName string) (*Environment, error) {
	e := &Environment{
		SDKAddr:        "http://localhost:1317",
		TendermintAddr: "http://localhost:26657",
		appName:        appName,
	}
	if e.checkServerAvailabity(ctx) {
		return nil, errors.New("a cosmos app is already running, stop it first")
	}
	e.Cleanup()
	var err error
	installOnce.Do(func() { err = errors.Wrap(installStarport(ctx), "cannot install starport") })
	if err != nil {
		return nil, err
	}
	bin, err := starportBinaryPath()
	if err != nil {
		return nil, err
	}
	appTemp, err := ioutil.TempDir("", "")
	if err != nil {
		return nil, err
	}
	ctx, e.serveCancel = context.WithCancel(ctx)

	errb := &bytes.Buffer{}
	e.serveG, ctx = errgroup.WithContext(ctx)

	opts := []cmdrunner.Option{
		cmdrunner.DefaultStderr(errb),
		cmdrunner.DefaultStdout(os.Stdout),
	}
	if isCI {
		opts = append(opts, cmdrunner.EndSignal(os.Kill))
	}
	e.serveG.Go(func() error {
		return cmdrunner.
			New(opts...).
			Run(ctx,
				step.New(
					step.Exec(bin, "app", "github.com/app/"+appName),
					step.Workdir(appTemp),
				),
				step.New(
					step.Exec(bin, "serve"),
					step.Workdir(filepath.Join(appTemp, appName)),
				),
			)
	})
	for {
		if err := ctx.Err(); err != nil {
			err = errors.Wrap(err, errb.String())
			return e, fmt.Errorf("cannot scaffold or serve scaffold, reason: %s", err.Error())
		}
		time.Sleep(time.Second * 5)
		if e.checkServerAvailabity(ctx) {
			// wait servers to fully get ready.
			time.Sleep(time.Second * 8)
			return e, nil
		}
		fmt.Println("rpc servers aren't avaiable yet")
	}
}

func (e *Environment) checkServerAvailabity(ctx context.Context) (ok bool) {
	for _, u := range []string{
		e.TendermintAddr,
		e.SDKAddr + "/node_info",
	} {
		if ok, err := httpstatuschecker.Check(ctx, u); err != nil || !ok {
			return false
		}
	}
	return true
}

// Appd returns app's daemon binary name.
func (e *Environment) Appd() string { return e.appName + "d" }

// Appcli returns app's cli binary name.
func (e *Environment) Appcli() string { return e.appName + "cli" }

// Cleanup shutdowns app servers and cleans up the testing environment.
func (e *Environment) Cleanup() {
	if e.serveCancel != nil {
		e.serveCancel()
		_ = e.serveG.Wait()
	}
	if appdPath, _ := exec.LookPath(e.Appd()); appdPath != "" {
		os.Remove(appdPath)
	}
	if appcliPath, _ := exec.LookPath(e.Appcli()); appcliPath != "" {
		os.Remove(appcliPath)
	}
	home, _ := os.UserHomeDir()
	os.RemoveAll(filepath.Join(home, "."+e.Appcli()))
	os.RemoveAll(filepath.Join(home, "."+e.Appd()))
}

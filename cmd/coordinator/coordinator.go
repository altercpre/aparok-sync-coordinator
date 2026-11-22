package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/altercpre/aparok-sync/acl"
	"github.com/altercpre/aparok-sync/app"
	"github.com/altercpre/aparok-sync/app/logger"
	"github.com/altercpre/aparok-sync/consensus/consensusclient"
	"github.com/altercpre/aparok-sync/metric"
	"github.com/altercpre/aparok-sync/net/peerservice"
	"github.com/altercpre/aparok-sync/net/pool"
	"github.com/altercpre/aparok-sync/net/rpc/server"
	"github.com/altercpre/aparok-sync/net/secureservice"
	"github.com/altercpre/aparok-sync/net/transport/quic"
	"github.com/altercpre/aparok-sync/net/transport/yamux"
	"github.com/altercpre/aparok-sync/nodeconf"
	"github.com/altercpre/aparok-sync/nodeconf/nodeconfstore"
	"go.uber.org/zap"

	"github.com/altercpre/aparok-sync-coordinator/account"
	"github.com/altercpre/aparok-sync-coordinator/accountlimit"
	"github.com/altercpre/aparok-sync-coordinator/acleventlog"
	"github.com/altercpre/aparok-sync-coordinator/config"
	"github.com/altercpre/aparok-sync-coordinator/coordinator"
	"github.com/altercpre/aparok-sync-coordinator/coordinatorlog"
	"github.com/altercpre/aparok-sync-coordinator/db"
	"github.com/altercpre/aparok-sync-coordinator/deletionlog"
	"github.com/altercpre/aparok-sync-coordinator/fileusage"
	"github.com/altercpre/aparok-sync-coordinator/identityrepo"
	"github.com/altercpre/aparok-sync-coordinator/inbox"
	"github.com/altercpre/aparok-sync-coordinator/invitestore"
	"github.com/altercpre/aparok-sync-coordinator/nodeconfsource"
	"github.com/altercpre/aparok-sync-coordinator/spacestatus"
	"github.com/altercpre/aparok-sync-coordinator/subscribe"

	// import this to keep govvv in go.mod on mod tidy
	_ "github.com/ahmetb/govvv/integration-test/app-different-package/mypkg"
)

var log = logger.NewNamed("main")

var (
	flagConfigFile = flag.String("c", "etc/aparok-sync-coordinator.yml", "path to config file")
	flagVersion    = flag.Bool("v", false, "show version and exit")
	flagHelp       = flag.Bool("h", false, "show help and exit")
)

func main() {
	flag.Parse()

	if *flagVersion {
		fmt.Println(app.AppName)
		fmt.Println(app.Version())
		fmt.Println(app.VersionDescription())
		return
	}
	if *flagHelp {
		flag.PrintDefaults()
		return
	}

	if debug, ok := os.LookupEnv("ANYPROF"); ok && debug != "" {
		go func() {
			http.ListenAndServe(debug, nil)
		}()
	}

	// create app
	ctx := context.Background()
	a := new(app.App)

	// open config file
	conf, err := config.NewFromFile(*flagConfigFile)
	if err != nil {
		log.Fatal("can't open config file", zap.Error(err))
	}

	// bootstrap components
	a.Register(conf)
	Bootstrap(a)

	// start app
	if err := a.Start(ctx); err != nil {
		log.Fatal("can't start app", zap.Error(err))
	}
	log.Info("app started", zap.String("version", a.Version()))

	// wait exit signal
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-exit
	log.Info("received exit signal, stop app...", zap.String("signal", fmt.Sprint(sig)))

	// close app
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	if err := a.Close(ctx); err != nil {
		log.Fatal("close error", zap.Error(err))
	} else {
		log.Info("goodbye!")
	}
	time.Sleep(time.Second / 3)
}

func Bootstrap(a *app.App) {
	a.Register(db.New()).
		Register(metric.New()).
		Register(account.New()).
		Register(nodeconfstore.New()).
		Register(nodeconf.New()).
		Register(nodeconfsource.New()).
		Register(deletionlog.New()).
		Register(peerservice.New()).
		Register(pool.New()).
		Register(secureservice.New()).
		Register(server.New()).
		Register(coordinatorlog.New()).
		Register(acleventlog.New()).
		Register(spacestatus.New()).
		Register(consensusclient.New()).
		Register(acl.New()).
		Register(subscribe.New()).
		Register(inbox.New()).
		Register(accountlimit.New()).
		Register(fileusage.New()).
		Register(identityrepo.New()).
		Register(invitestore.New()).
		Register(coordinator.New()).
		Register(yamux.New()).
		Register(quic.New())
}

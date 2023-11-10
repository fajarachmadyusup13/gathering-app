package console

import (
	"errors"
	"fmt"
	"os"
	"os/signal"

	"github.com/fajarachmadyusup13/gathering-app/internal/config"
	"github.com/fajarachmadyusup13/gathering-app/internal/db"
	"github.com/fajarachmadyusup13/gathering-app/internal/delivery/httpsvc"
	"github.com/fajarachmadyusup13/gathering-app/internal/repository"
	"github.com/fajarachmadyusup13/gathering-app/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Long:  "This subcommand start the server",
	Run:   run,
}

func init() {
	RootCmd.AddCommand(runCmd)
}

func run(cmd *cobra.Command, args []string) {

	db.InitializeMySQLConn()

	memberRepo := repository.NewMemberRepository(db.MySQL)
	gatheringRepo := repository.NewGatheringRepository(db.MySQL)
	invitationRepo := repository.NewInvitationRepository(db.MySQL)
	attendeeRepo := repository.NewAttendeeRepository(db.MySQL)

	memberUsecase := usecase.NewMemberUsecase(memberRepo)
	gatheringUsecase := usecase.NewGatheringUsecase(gatheringRepo)
	invitationUsecase := usecase.NewInvitationUsecase(invitationRepo, memberRepo, gatheringRepo, attendeeRepo)

	httpService := httpsvc.NewHTTPService()
	httpService.RegisterMemberUsecase(memberUsecase)
	httpService.RegisterGatheringUsecase(gatheringUsecase)
	httpService.RegisterInvitationUsecase(invitationUsecase)

	sigCh := make(chan os.Signal, 1)
	errCh := make(chan error, 1)
	signal.Notify(sigCh, os.Interrupt)

	go func() {
		<-sigCh
		errCh <- errors.New("received an interupt")
		db.StopTickerCh <- true
	}()

	go runHTTPServer(httpService, errCh)
	log.Error(<-errCh)

}

func runHTTPServer(httpService *httpsvc.HTTPService, errCh chan<- error) {
	g := gin.Default()

	httpService.InitRoutes(g)
	errCh <- g.Run(fmt.Sprintf("0.0.0.0:%s", config.HTTPPort()))
}

package logic

import (
	"context"
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

type YoutubeIntent struct {
	intent.BaseIntent
	*models.Youtube
}

func (s *Service) CreateYoutube(
	ctx context.Context, request *services.CreateYoutubeRequest,
) (*services.CreateYoutubeResponse, error) {
	i := &YoutubeIntent{
		Youtube: request.GetYoutube(),
	}

	err := s.handleCreate(ctx, i, nil, i.Youtube)
	if err != nil {
		return nil, err
	}

	return s.BaseService.CreateYoutube(ctx, request)
}

func (intent *YoutubeIntent) Evaluate(_ context.Context, _ *intent.EvaluateContext) error {
	url := intent.URL
	fmt.Println("Got youtube intent, URL: ", url)
	cmd := exec.Command("chromium-browser", url)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

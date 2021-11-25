package commands

import (
	"errors"
	"fmt"
	"subteez/config"
	"subteez/messages"
	"subteez/subteez"
)

func mainFiles(args []string, cfg config.Config) error {
	client := subteez.NewClient(cfg.GetServer())

	if len(args) < 2 {
		return errors.New(messages.NotEnoughArguments)
	}

	request := subteez.SubtitleDetailsRequest{
		ID:              args[1],
		LanguageFilters: cfg.GetLanguageFilters(),
	}
	response, err := client.GetDetails(request)
	if err != nil {
		return err
	}

	if len(response.Files) < 1 {
		return errors.New(messages.NoFileFound)
	}

	for _, item := range response.Files {
		fmt.Printf(messages.FileRow, item.ID, item.Language.GetTitle(), item.Title, item.Author)
	}

	return nil
}

package service

import (
	config "12-startups-one-month/config"
	"12-startups-one-month/graph/model"
	db "12-startups-one-month/internal"
	"12-startups-one-month/utils"
	"context"
	"fmt"

	"io"
)

type IFileService interface {
	UploadSingleFile(ctx context.Context, file *model.UploadInput) (*model.UploadResponse, error)
}

type FileService struct {
	server *config.Server
}

func NewFileService(server *config.Server) *FileService {
	return &FileService{
		server: server,
	}
}

func (s *FileService) UploadSingleFile(ctx context.Context, file *model.UploadInput) (*model.UploadResponse, error) {
	var content []byte
	var err error
	var res *model.UploadResponse

	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		// check if file extension is allowed
		if !utils.IsFileExtensionAllowed(file.File.Filename, []string{
			".jpg",
			".jpeg",
			".png",
		}) {
			return fmt.Errorf("file extension is not allowed")
		}

		content, err = io.ReadAll(file.File.File)
		if err != nil {
			return err
		}
		// check if content > 500kb
		if len(content) > 500000 {
			return fmt.Errorf("file is too large")
		}
		// check if it's a valid image
		if !utils.IsImage(content, []string{
			"image/jpeg",
			"image/png",
		}) {
			return fmt.Errorf("file has wrong format")
		}
		// create path
		dst, fn, err := utils.CreatePathFile(file.File.Filename)
		if err != nil {
			fmt.Printf("Error create file => %v\n", err)
			return err
		}
		// put content in file
		_, err = dst.Write(content)
		if err != nil {
			fmt.Printf("Error write file => %v\n", err)
			return err
		}
		// close file
		defer dst.Close()
		if file.Height != nil || file.Width != nil {
			if *file.Height > 0 && *file.Width > 0 {
				go utils.CropImage(fn, int(*file.Width), int(*file.Height))
			}
		}
		arg := db.CreateFileParams{
			Name: utils.NullS(file.File.Filename),
			Size: utils.NullI64(int64(len(content))),
			Url:  utils.NullS("/" + fn),
			Mime: utils.NullS(file.File.ContentType),
		}
		newFile, err := q.CreateFile(ctx, arg)
		if err != nil {
			return fmt.Errorf("ERROR_CREATE_FILE: %v", err)
		}
		res = &model.UploadResponse{
			Name:    newFile.Name.String,
			URL:     newFile.Url.String,
			Size:    int(newFile.Size.Int64),
			Success: true,
		}
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_CREATE_FILE", err)
	}
	return res, nil
}

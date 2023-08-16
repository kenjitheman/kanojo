package gdrive

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	drive "google.golang.org/api/drive/v3"
)

func getUserClient(credentialFile, tokenFile string, scopes ...string) (*http.Client, error) {
	b, err := os.ReadFile(credentialFile)
	if err != nil {
		return nil, err
	}
	c := struct {
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	}{}
	err = json.Unmarshal(b, &c)
	if err != nil {
		return nil, err
	}
	config := &jwt.Config{
		Email:      c.Email,
		PrivateKey: []byte(c.PrivateKey),
		Scopes:     append([]string{drive.DriveScope}, scopes...),
		TokenURL:   google.JWTTokenURL,
	}

	tok, err := getTokenFromFile(tokenFile)
	if err != nil {
		return nil, err
	}
	client := config.Client(oauth2.NoContext, tok)
	return client, nil
}

func getTokenFromFile(tokenFile string) (*oauth2.Token, error) {
	tokData, err := os.ReadFile(tokenFile)
	if err != nil {
		return nil, err
	}
	var tok oauth2.Token
	err = json.Unmarshal(tokData, &tok)
	if err != nil {
		return nil, err
	}
	return &tok, nil
}

func saveFile(srv *drive.Service, file *os.File, filename string) error {
	baseMimeType := "text/plain"
	fileInf, err := file.Stat()
	if err != nil {
		return err
	}
	defer file.Close()
	f := &drive.File{Name: filename}
	_, err = srv.Files.Create(f).
		Media(file, googleapi.ContentType(baseMimeType)).
		Do()
	if err != nil {
		return err
	}
	return nil
}

func Gdrive(credentialFile, tokenFile, fileName string) error {
	client, err := getUserClient(credentialFile, tokenFile)
	if err != nil {
		return err
	}

	srv, err := drive.NewService(context.Background(), withClient(client))
	if err != nil {
		return err
	}

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	err = saveFile(srv, file, filepath.Base(fileName))
	if err != nil {
		return err
	}

	fmt.Println("[SUCCESS] file saved successfully.")
	return nil
}

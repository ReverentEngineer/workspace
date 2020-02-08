package main

import (
  "context"
  "errors"
  "fmt"
  "github.com/digitalocean/godo"
  "golang.org/x/oauth2"
  "gopkg.in/yaml.v3"
  "io/ioutil"
  "log"
  "os"
)

const WorkspaceTag = "workspace"

type ManagerConfig struct {
  Access_Token string
  Slug         string
  Size         string
  Region       string
  Ssh_Fingerprint  string
}

func (c ManagerConfig) Token() (*oauth2.Token, error) {
  token := &oauth2.Token{
    AccessToken: c.Access_Token,
  }
  return token, nil
}


func HandleArgs(args []string, config *ManagerConfig) error {
  if len(args) == 0 {
    return errors.New("No commands provided.")
  }

  oauthClient := oauth2.NewClient(context.Background(), config)
  client := godo.NewClient(oauthClient)

  var err error
  switch command := args[0]; command {
  case "start":
    workspaces, _, _ := client.Droplets.ListByTag(context.TODO(), WorkspaceTag, nil)
    if len(workspaces) > 0 {
      fmt.Println("Workspace already started.")
      return nil
    }
    createRequest := &godo.DropletCreateRequest{
      Name:   WorkspaceTag,
      Region: (*config).Region,
      Size:   (*config).Size,
      SSHKeys: []godo.DropletCreateSSHKey{
        godo.DropletCreateSSHKey{
          Fingerprint: (*config).Ssh_Fingerprint,
        },
      },
      Tags: []string{WorkspaceTag},
      Image: godo.DropletCreateImage{
        Slug: (*config).Slug,
      },
    }
    client.Droplets.Create(context.TODO(), createRequest)
  case "stop":
    client.Droplets.DeleteByTag(context.TODO(), WorkspaceTag)
  default:
    err = errors.New("Invalid command.")
  }

  return err
}

func PrintUsage() {
  fmt.Println("usage: workspace <start|stop>")
}

func main() {
  var config ManagerConfig
  config_path := os.Getenv("HOME") + "/.config/workspace/config.yaml"
  config_data, err := ioutil.ReadFile(config_path)
  if err != nil {
    log.Fatal("No config file found.")
    os.Exit(1)
  }
  yaml.Unmarshal(config_data, &config)
  err = HandleArgs(os.Args[1:], &config)
  if err != nil {
    PrintUsage()
    os.Exit(1)
  }
}

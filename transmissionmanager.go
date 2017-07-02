package main

import (
  "fmt"
  "os/exec"
  "os"
  "time"
  "github.com/lnguyen/go-transmission/transmission"
)

func MoveDownloadedFile(source string, dest string) error {
  fmt.Println("MoveDownloadedFile: " + source + " to " + dest)
  _, err := exec.Command("mv", source, dest).Output()
  return err
}

func HandleFinishedTorrent(client transmission.TransmissionClient, torrent transmission.Torrent) error {
  fmt.Println("HandleFinishedTorrent")
  _, err := client.RemoveTorrent(torrent.ID, false)
  if err != nil {
    return err
  }
  return nil
}

func ProcessAllTorrents(client transmission.TransmissionClient, uri string, username string, password string, sourcePath string, destPath string) error {
  torrents, err := client.GetTorrents()
  if err != nil {
    fmt.Println("Error while getting list of torrents")
    return err
  }
  for _, torrent := range torrents {
    fmt.Println(torrent)
    if torrent.LeftUntilDone == 0 && torrent.PercentDone == 1 {
      fmt.Println(torrent.Name + " is finished")
      err := HandleFinishedTorrent(client, torrent)
      if err != nil {
        fmt.Println("Error while removing torrent from transmission " + torrent.Name)
        return err
      }
      err = MoveDownloadedFile(sourcePath + "/" + torrent.Name, destPath)
      if err != nil {
        fmt.Println("Error while moving downloaded files " + torrent.Name)
        return err
      }
    } else {
      fmt.Println(torrent.Name + " is not finished, skipping ...")
    }
  }
  return nil
}

func main() {
  if len(os.Args) != 6 {
    fmt.Println("usage: " + os.Args[0] + " uri username password source dest")
    fmt.Println("\tusername: transmission web uri (e.g., http://transmission.local:8080)")
    fmt.Println("\tusername: transmission web username")
    fmt.Println("\tpassword: transmission web password")
    fmt.Println("\tsource: path to transmission downloads")
    fmt.Println("\tdest: where to move downloaded files after torrent is removed")
    os.Exit(1)
  }
  var uri string = os.Args[1]
  var username string = os.Args[2]
  var password string = os.Args[3]
  var sourcePath string = os.Args[4]
  var destPath string = os.Args[5]
  var client transmission.TransmissionClient
  for {
    fmt.Println("Initializing transmission client")
    client = transmission.New(uri, username, password)
    for {
      time.Sleep(30 * time.Second)
      fmt.Println("Processing torrents")
      err := ProcessAllTorrents(client, uri, username, password, sourcePath, destPath)
      if err != nil {
        fmt.Println(err)
        break
      }
    }
  }
}

package cache

import (
  "io/ioutil"
  "os"
  "bytes"
  "context"
  "fmt"
  "github.com/redis/go-redis/v9"
)

// Memory > Storage > Server


// Memory
func savePage() error{
  client := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
    Password: "",
    DB: 0,
  })
  
  ctx := context.Background()

  err := client.Set(ctx, "key", "value", 0).Err()
  if err != nil {
    return err
  }
  
  val, err := client.Get(ctx, "key").Result()
  if err != nil {
    return err
  }

  fmt.Println(val)
  return nil
}


func GetPage(index []byte) ([]byte, error){
  var data []byte

  homeDir, err := os.UserHomeDir()
  if err != nil {
    return nil, err
  }

  // Creating the path
  buf := &bytes.Buffer{}
  buf.WriteString(homeDir)
  buf.WriteString("/.cache/poketch/test")
  pathToCache := buf.String()

  data, err = ioutil.ReadFile(pathToCache)
  if err != nil {
    return nil, err
  }

  return data, nil
}

func WritePage(data []byte, index []byte) error {
  homeDir, err := os.UserHomeDir()
  if err != nil {
    return err
  }
  
  // Creating the path
  buf := &bytes.Buffer{}
  buf.WriteString(homeDir)
  buf.WriteString("/.cache/poketch/test")
  pathToCache := buf.String()
  
  file, err := os.OpenFile(pathToCache, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    return err
  }

  _, err = file.Write(data)
  if err != nil {
    return err
  }

  err = file.Close()
  if err != nil {
    return err
  }
  return nil
}





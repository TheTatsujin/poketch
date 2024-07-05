package pokeapi

import (
  "io/ioutil"
  "net/http"
  "encoding/json"
  "github.com/TheTatsujin/poketch/mod/apifetch/cache"
)

type jsonObject map[string]interface{}

const pokeapiURL string = "https://pokeapi.co/api/v2/"

func GetJson(url string) (jsonObject, error) {
  var rawData []byte
  var jsonData jsonObject
  var err error
  var isCached bool = true

  // See if page is in disk
  if isCached {
    rawData, err = cache.GetPage(nil)
    if err != nil {
      return nil, err
    }

  } else{    
    response, err := http.Get(url)
    if err != nil {
      return nil, err
    }
    defer response.Body.Close()

    rawData, err = ioutil.ReadAll(response.Body)
    if err != nil {
      return nil, err
    }

    if !json.Valid(rawData) {
      return nil, err
    }

    if err := cache.WritePage(rawData, nil); err != nil {
      return nil, err
    }
  }

  err = json.Unmarshal(rawData, &jsonData)
  if err != nil {
    return nil, err
  }

  return jsonData, nil
}


package persist

import (
	"context"
	"errors"
	"gopkg.in/olivere/elastic.v5"
)

type Item struct {
	Id    string
	Value string
	Type  string
}

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item Item, result *string) error {
	err := save(s.Client, s.Index, item)
	if err == nil {
		*result = "ok"
	}
	return nil
}

func save(client *elastic.Client, index string, item Item) error {
	if item.Value == "" {
		return errors.New("can`t be null")
	}

	indexService := client.Index().Index(index).Type("string").BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}

	return nil
	/*
		resp,err := client.Get().Index(index).Type("string").Id(id).Do(context.Background())
		if err != nil{
			return err
		}
		var itemObj Item
		err = json.Unmarshal(*resp.Source,&itemObj)
		if err != nil{
			return err
		}
	*/
}

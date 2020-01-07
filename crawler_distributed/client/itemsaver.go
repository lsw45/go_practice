package client

func ItemSaver(host string) {
	client, err := rpcsupport.NewClient(host)

	out := make(chan persist.Item)

	go func() {
		item := <-out

		result := ""
		client.Call("ItemSaverService.Save", item, &result)

	}()
}

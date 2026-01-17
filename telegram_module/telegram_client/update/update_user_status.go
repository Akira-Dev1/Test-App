package update

import (
	"time"
	"net/http"
	"log"
)

func UpdateUserStatus() {
	ticker := time.NewTicker(10 * time.Second)
	
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	go func() {
		for range ticker.C {
			url := "http://tg_nginx/login/check"

			resp, err := client.Get(url)
			if err != nil {
				log.Printf("Status update failed: %v", err)
				continue
			}
			resp.Body.Close()
			
			log.Printf("Status check sent to %s, Response: %d", url, resp.StatusCode)
		}
	}()

}
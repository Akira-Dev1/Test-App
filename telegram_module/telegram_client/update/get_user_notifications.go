package update

import (
	"time"
	"net/http"
	"log"
)

func GetUserNotifications() {
	ticker := time.NewTicker(10 * time.Second)
	
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	go func() {
		for range ticker.C {
			url := "http://tg_nginx/notifications"

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
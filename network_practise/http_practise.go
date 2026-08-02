package networkpractise

// данные пакеты нужны для системы проверки
import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func httpRequest() {
	client := http.Client{Timeout: time.Millisecond * 500}
    resp, err := client.Get("http://127.0.0.1:5555/get")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {л
		log.Fatal("unexpected status code: %d", resp.StatusCode)
		os.Exit(1)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error while parsing Body: %v", err)
	}
	fmt.Printf("%s", data)
}
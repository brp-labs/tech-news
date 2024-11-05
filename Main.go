package main

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"net/http"
)

type RSS struct {
    Channel Channel `xml:"channel"`
}

type Channel struct {
    Title       string  `xml:"title"`
    Description string  `xml:"description"`
    Items       []Item  `xml:"item"`
}

type Item struct {
    Title       string `xml:"title"`
    Description string `xml:"description"`
    Link        string `xml:"link"`
    Category    string `xml:"category"`
    PubDate     string `xml:"pubDate"`
    Thumbnail   Thumbnail `xml:"thumbnail"`
}

type Thumbnail struct {
    URL    string `xml:"url,attr"`
    Width  string `xml:"width,attr"`
    Height string `xml:"height,attr"`
}

func getRSSFeed(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    feedURL := "https://techxplore.com/rss-feed/breaking/machine-learning-ai-news/"

    client := &http.Client{}
    req, err := http.NewRequest("GET", feedURL, nil)
    if err != nil {
        log.Printf("Error creating request: %v", err)
        http.Error(w, "Error creating request", http.StatusInternalServerError)
        return
    }

    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

    resp, err := client.Do(req)
    if err != nil {
        log.Printf("Error fetching RSS feed: %v", err)
        http.Error(w, "Error fetching RSS feed", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        log.Printf("Error fetching RSS feed: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
        http.Error(w, "Error fetching RSS feed", resp.StatusCode)
        return
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Printf("Error reading RSS data: %v", err)
        http.Error(w, "Error reading RSS data", http.StatusInternalServerError)
        return
    }

    log.Printf("RSS data received: %s", string(body))

    var rss RSS
    if err := xml.Unmarshal(body, &rss); err != nil {
        log.Printf("Error parsing XML: %v", err)
        http.Error(w, "Error parsing XML", http.StatusInternalServerError)
        return
    }

    var articles []map[string]string
    for _, item := range rss.Channel.Items {
        article := map[string]string{
            "title":       item.Title,
            "description": item.Description,
            "link":        item.Link,
            "category":    item.Category,
            "pubDate":     item.PubDate,
            "thumbnail":   item.Thumbnail.URL,
        }
        articles = append(articles, article)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(articles)
}

func main() {
    http.HandleFunc("/api/rss", getRSSFeed)
    log.Println("Server runs at http://localhost:8080/api/rss")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

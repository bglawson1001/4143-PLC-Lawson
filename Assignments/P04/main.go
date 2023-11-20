/************************************************************************
 *
 *  Author:           Brayden Lawson
 *  Title:            Concurrent Image Downloader
 *  Course:           4143-101
 *  Semester:         Fall 2023
 *
 *  Description:
 * Golang program that will download 9 images concurrently and sequentially and write how long
 * each took to download all 9 images.
 *
 *
 *  Usage:
 *        Used to download images, concurrent is faster.
 *
 *
 *  Files: main.go, go.mod, go.sum. concurrent_images, sequential_images
 ************************************************************************/



package main

import (
  "fmt"
  "io"
  "net/http"
  "os"
  "path/filepath"
  "sync"
  "time"
)

func main() {
  // Create folders for sequential and concurrent downloads
  sequentialFolder := "sequential_images"
  concurrentFolder := "concurrent_images"
  createFolder(sequentialFolder)
  createFolder(concurrentFolder)

  //Image urls for this assignment.
  urls := []string{
    "https://unsplash.com/photos/hvdnff_bieQ/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
    "https://unsplash.com/photos/HQaZKCDaax0/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
    "https://images.unsplash.com/photo-1698778573682-346d219402b5?ixlib=rb-4.0.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640",
    "https://unsplash.com/photos/Bs2jGUWu4f8/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
   "https://cdn.pixabay.com/photo/2023/09/20/04/04/sea-urchin-8263832_1280.jpg",
    "https://cdn.pixabay.com/photo/2023/10/20/13/48/tamarin-8329530_1280.png",
    "https://cdn.pixabay.com/photo/2023/11/04/07/57/owl-8364426_1280.jpg",
    "https://cdn.stocksnap.io/img-thumbs/960w/futuristic-circuits_OZPRB0FQZK.jpg",
    "https://images.pexels.com/photos/18587783/pexels-photo-18587783/free-photo-of-the-inside-of-a-library-with-a-skylight.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
  }

  // Sequential download
  startTime := time.Now()
  for _, url := range urls {
    downloadImageSequential(url, sequentialFolder)
  }
  sequentialDuration := time.Since(startTime)
  fmt.Println("")
  fmt.Printf("Sequential download time: %s\n", sequentialDuration)
  fmt.Println("")

  // Concurrent download
  startTime = time.Now()
  var wg sync.WaitGroup
  for _, url := range urls {
    wg.Add(1)
    go func(u string) {
      defer wg.Done()
      downloadImageConcurrent(u, concurrentFolder)
    }(url)
  }
  wg.Wait()
  concurrentDuration := time.Since(startTime)
  fmt.Println("")
  fmt.Printf("Concurrent download time: %s\n", concurrentDuration)
  fmt.Println("")
}

// Function to download the image sequentially
func downloadImageSequential(url string, folder string) {
  // Create a new `http.Request` object.
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    fmt.Println(err)
    return
  }

  // Create a new `http.Client` object.
  client := &http.Client{}

  // Do the request and get the response.
  resp, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }

  // Check the response status code.
  if resp.StatusCode != http.StatusOK {
    fmt.Println("Response status code:", resp.StatusCode)
    return
  }

  // Create a unique filename based on the URL with a .jpg extension.
  filename := filepath.Join(folder, "sequential_image_"+extractFilename(url)+".jpg")
  f, err := os.Create(filename)
  if err != nil {
    fmt.Println(err)
    return
  }

  // Copy the image from the response body to the file.
  _, err = io.Copy(f, resp.Body)
  if err != nil {
    fmt.Println(err)
    return
  }

  // Close the file.
  f.Close()

  // Print a success message.
  fmt.Println("Sequential image saved to", filename)
}

// Function to download the image concurrently
func downloadImageConcurrent(url string, folder string) {
  // Create a new `http.Request` object.
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    fmt.Println(err)
    return
  }

  // Create a new `http.Client` object.
  client := &http.Client{}

  // Do the request and get the response.
  resp, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }

  // Check the response status code.
  if resp.StatusCode != http.StatusOK {
    fmt.Println("Response status code:", resp.StatusCode)
    return
  }

  // Create a unique filename based on the URL with a .jpg extension.
  filename := filepath.Join(folder, "concurrent_image_"+extractFilename(url)+".jpg")
  f, err := os.Create(filename)
  if err != nil {
    fmt.Println(err)
    return
  }

  // Copy the image from the response body to the file.
  _, err = io.Copy(f, resp.Body)
  if err != nil {
    fmt.Println(err)
    return
  }

  // Close the file.
  f.Close()

  // Print a success message.
  fmt.Println("Concurrent image saved to", filename)
}

// Helper function to create a folder if it doesn't exist
func createFolder(folder string) {
  if _, err := os.Stat(folder); os.IsNotExist(err) {
    os.Mkdir(folder, os.ModePerm)
  }
}

// Helper function to extract filename from URL
func extractFilename(url string) string {
  // Use filepath.Base to extract the filename
  return filepath.Base(url)
}

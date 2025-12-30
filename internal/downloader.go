package internal

//
// import (
// 	"fmt"
// 	"sync"
// )
//
// type Job struct {
// 	Category string
// 	Item     string
// 	Index    int
// }
//
// func GenerateCards(cfg *Config, outDir string, imagesPerItem, workers int) error {
// 	jobs := make(chan Job)
// 	var wg sync.WaitGroup
//
// 	for i := 0; i < workers; i++ {
// 		wg.Add(1)
// 		go worker(jobs, &wg, outDir)
// 	}
//
// 	for category, items := range cfg.Categories {
// 		for _, item := range items {
// 			for i := 1; i <= imagesPerItem; i++ {
// 				jobs <- Job{category, item, i}
// 			}
// 		}
// 	}
//
// 	close(jobs)
// 	wg.Wait()
// 	fmt.Println("✅ Done")
// 	return nil
// }

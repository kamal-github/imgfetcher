package internal

import "sync"

func GenerateCards(cfg *Config, outDir string, imagesPerItem, workers int) error {
	jobs := make(chan Job)
	var wg sync.WaitGroup

	// start workers
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(jobs, &wg, outDir)
	}

	// enqueue jobs
	for category, items := range cfg.Categories {
		for _, item := range items {
			for i := 1; i <= imagesPerItem; i++ {
				jobs <- Job{
					Category: category,
					Item:     item,
					Index:    i,
				}
			}
		}
	}

	close(jobs)
	wg.Wait()
	return nil
}

func worker(jobs <-chan Job, wg *sync.WaitGroup, outDir string) {
	defer wg.Done()

	for job := range jobs {
		_ = processJob(job, outDir)
	}
}

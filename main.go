package main

import (
	"shadow_release/repo"
	"shadow_release/shadow_repo"
	"sync"
)

// func (s *Tool)

// func parseConfig() Config {
// 	configContent, err := os.ReadFile("./config.json")
// 	if err != nil {
// 		panic(err)
// 	}
//
//
// 	var config Config
//
// 	if err := json.Unmarshal(configContent, &config); err != nil {
// 		panic(err)
// 	}
//
// 	return config
// }
//
// func fetchRepo(url string) {
// 	cmd := exec.Command("git", "clone", url, "tmp/repo")
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	cmd.Run()
// }

func main() {
	wait_group := sync.WaitGroup{}
	wait_group.Add(2)

	go func() {
		defer wait_group.Done()
		repo.Server()
	}()
	go func() {
		defer wait_group.Done()
		shadow_repo.Server()
	}()

	wait_group.Wait()
}




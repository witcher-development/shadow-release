package main

import "shadow_release/repo"

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
	repo.Server()
}




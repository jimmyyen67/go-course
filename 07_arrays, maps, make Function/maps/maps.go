package maps

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "https://www.google.com",
		"Amazon Web Services": "https://www.aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	websites["linkedin"] = "https://www.linkedin.com"
	fmt.Println(websites)

	delete(websites, "Amazon Web Services")
	fmt.Println(websites)
}

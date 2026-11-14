package main

import (
	"fmt"
	"github.com/ztaylorpossible/gator/internal/config"
)

func main() {
	conf, err := config.Read();

	if err != nil {
		fmt.Printf("Error reading config: %v\n", err);
		return;
	}

	conf.SetUser("zach");
	if err != nil {
		fmt.Printf("Error setting user: %v\n", err);
		return;
	}
	
	conf, err = config.Read();

	if err != nil {
		fmt.Printf("Error reading config: %v\n", err);
		return;
	}

	fmt.Printf("Database url: %s\nUsername: %s\n", conf.Db_url, conf.Current_user_name);
}

package main

import (
    "github.com/bwmarrin/discordgo"
	"GoGrebball/commands"
	"fmt"
)

func main() {
	fmt.Println("Grebball starting ...")
    _ = LoadConfig();
    commands.Ping();

    bot, err := discordgo.New("token");
    if err != nil {
        fmt.Println("[ERR] Could not create new discordgo instance", err);
    }
    fmt.Println("[SUCCESS] new discordbot instance ready: ", bot);
}

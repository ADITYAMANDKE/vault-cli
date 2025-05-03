package main

import (
	"flag"
	"fmt"
	"learngo/models"
	"learngo/vault"
	"log"
	"os"
	"time"
)

func main() {
	key := []byte(os.Getenv("VAULT_KEY"))
	if len(key) != 32 {
		log.Fatal("Invalid or missing VAULT_KEY environment variable, please read READ.ME on Github")
	}
	vaultPath := "vault_data.enc"
	v, err := vault.LoadVault(vaultPath, key)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Vault not found.. Creating a new vault")
			v = vault.NewVault()
		} else {
			log.Fatal(err)
		}
	}
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	genCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	addSite := addCmd.String("site", "", "Site name")
	addUsername := addCmd.String("username", "", "Username")
	addPassword := addCmd.String("password", "", "Password")
	addNote := addCmd.String("note", "", "Optional note")
	getSite := getCmd.String("site", "", "Site to retrieve")
	delSite := deleteCmd.String("site", "", "Site to delete")
	genLength := genCmd.Int("length", 12, "Length of password")
	if len(os.Args) < 2 || os.Args[1] == "--help" || os.Args[1] == "-h" {
		fmt.Println(`Usage: vault-cli [command] [flags]
	
	Commands:
	  add       Add a new password entry
	  get       Retrieve an entry by site
	  delete    Delete an entry by site
	  list      List all stored site names
	  generate  Generate a strong password
	
	Use "vault-cli [command] --help" for more information about a command.`)
		os.Exit(0)
	}
	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		entry := models.Entry{
			Site:       *addSite,
			Username:   *addUsername,
			Password:   *addPassword,
			ModifiedAt: time.Now(),
			Note:       *addNote,
		}
		v.AddEntry(entry)
		fmt.Println("Entry added.")
	case "get":
		getCmd.Parse(os.Args[2:])
		entry, found := v.GetEntry(*getSite)
		if !found {
			fmt.Println("Entry not found.")
			return
		}
		fmt.Printf("Site: %s\nUsername: %s\nPassword: %s\nNote: %s\nModified: %s\n",
			entry.Site, entry.Username, entry.Password, entry.Note, entry.ModifiedAt.Format(time.RFC1123))
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		ok := v.DeleteEntry(*delSite)
		if ok {
			fmt.Println("Entry deleted.")
		} else {
			fmt.Println("Entry not found.")
		}
	case "generate":
		genCmd.Parse(os.Args[2:])
		pwd, err := vault.GenerateRandomPassword(*genLength)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Generated password:", pwd)
	case "list":
		listCmd.Parse(os.Args[2:])
		sites := v.ListSites()
		for _, site := range sites {
			fmt.Println(site)
		}
	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
	switch os.Args[1] {
	case "add", "delete":
		err := vault.SaveVault(v, vaultPath, key)
		if err != nil {
			log.Fatal("Failed to save vault:", err)
		}
	}

}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/go-ldap/ldap/v3"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Invalid number of arguments. Expected 2, got %d", len(os.Args)-1)
	}

	var conf Config
	_, err := toml.DecodeFile(os.Args[1], &conf)
	if err != nil {
		log.Fatal(err)
	}
	l, err := ldap.DialURL(conf.LdapUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	err = l.UnauthenticatedBind("")
	if err != nil {
		log.Fatal(err)
	}
	filter := fmt.Sprintf("(uid=%s)", os.Args[2])
	sr := ldap.NewSearchRequest(conf.BaseDn, 2, 0, 1, 30, false, filter, []string{"sshPublicKey"}, nil)

	res, err := l.Search(sr)
	if err != nil {
		log.Fatal(err)
	}

	for _, ele := range res.Entries {
		keys := ele.GetAttributeValues("sshPublicKey")
		for _, k := range keys {
			fmt.Println(k)
		}
	}
}

type Config struct {
	LdapUrl string `toml:"ldap_url"`
	BaseDn  string `toml:"ldap_base_dn"`
}

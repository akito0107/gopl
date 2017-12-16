package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var in string
	var alg = flag.String("alg", "sha256", "")
	fmt.Scan(&in)
	flag.Parse()

	switch *alg {
	case "sha256":
		fmt.Fprintf(os.Stdout, "%x \n", sha256.Sum256([]byte(in)))
		break
	case "sha384":
		fmt.Fprintf(os.Stdout, "%x \n", sha512.Sum384([]byte(in)))
		break
	case "sha512":
		fmt.Fprintf(os.Stdout, "%x \n", sha512.Sum512([]byte(in)))
		break
	default:
		log.Fatalf("unsupported alg type %s \n", alg)
	}

}

package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Get the URL from the user.
    fmt.Print("Enter an HTTPS URL: ")
    url := ""
    fmt.Scanf("%s", &url)

    // Create a new HTTP client.
    client := &http.Client{}

    // Grab the TLS configuration for the URL.
    tlsConfig, err := client.Transport.(*http.Transport).TLSClientConfig()
    if err != nil {
        fmt.Println(err)
        return
    }

    // Print the supported server-side SSL cipher suites.
    for _, cipherSuite := range tlsConfig.CipherSuites {
        fmt.Println(cipherSuite)
    }
}

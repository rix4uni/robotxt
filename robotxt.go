package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rix4uni/robotxt/banner"
)

func main() {
	// Define flags
	outputFile := flag.String("o", "", "File to save the output (default: print to stdout)")
	userAgent := flag.String("H", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36", "Custom User-Agent for requests")
	typesCount := flag.Bool("types-count", false, "Print the count of types at the end")
	complete := flag.Bool("complete", false, "Include the full URL in the output")
	Type := flag.String("type", "", "Specify which one to extract: 'Disallow' or 'Allow'")
	timeout := flag.Duration("timeout", 10*time.Second, "Timeout duration for HTTP requests")
	delay := flag.Duration("delay", 0*time.Second, "Delay between requests (default 0s)")
	silent := flag.Bool("silent", false, "silent mode.")
	versionFlag := flag.Bool("version", false, "Print the version of the tool and exit.")
	verbose := flag.Bool("verbose", false, "Enable verbose output")
	flag.Parse()

	if *versionFlag {
		banner.PrintBanner()
		banner.PrintVersion()
		return
	}

	if !*silent {
		banner.PrintBanner()
	}

	// Prepare output writer
	var outputWriter io.Writer = os.Stdout
	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			return
		}
		defer file.Close()
		outputWriter = io.MultiWriter(os.Stdout, file)
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: *timeout,
	}

	// Read URLs from stdin
	scanner := bufio.NewScanner(os.Stdin)
	disallowCount, allowCount := 0, 0
	for scanner.Scan() {
		url := scanner.Text()
		if url == "" {
			continue
		}

		// Fetch robots.txt
		robotsURL := strings.TrimRight(url, "/") + "/robots.txt"
		if *verbose {
			fmt.Fprintf(os.Stderr, "Fetching: %s\n", robotsURL)
		}
		req, err := http.NewRequest("GET", robotsURL, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating request for %s: %v\n", robotsURL, err)
			continue
		}
		req.Header.Set("User-Agent", *userAgent)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching %s: %v\n", robotsURL, err)
			continue
		}
		defer resp.Body.Close()

		// Check if robots.txt exists
		if resp.StatusCode != http.StatusOK {
			fmt.Fprintf(os.Stderr, "Failed to retrieve %s: %s\n", robotsURL, resp.Status)
			continue
		}

		// Parse robots.txt
		parsedDisallow, parsedAllow := parseRobotsTxt(resp.Body, url, outputWriter, *complete, *Type, *verbose)
		disallowCount += parsedDisallow
		allowCount += parsedAllow

		// Add delay between requests if specified
		if *delay > 0 {
			time.Sleep(*delay)
		}
	}

	// Print types count if the flag is enabled
	if *typesCount {
		fmt.Fprintf(outputWriter, "\nTotal Disallow: %d\n", disallowCount)
		fmt.Fprintf(outputWriter, "Total Allow: %d\n", allowCount)
	}
}

// parseRobotsTxt parses and extracts Disallow and Allow rules from robots.txt
func parseRobotsTxt(body io.Reader, baseURL string, writer io.Writer, complete bool, Type string, verbose bool) (int, int) {
	scanner := bufio.NewScanner(body)
	disallowCount, allowCount := 0, 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Filter lines by rule type if specified
		if Type != "" {
			if !strings.HasPrefix(line, Type+":") {
				continue
			}
		}

		// Process Disallow or Allow lines
		if strings.HasPrefix(line, "Disallow:") {
			disallowCount++
		} else if strings.HasPrefix(line, "Allow:") {
			allowCount++
		} else {
			continue
		}

		// Include the base URL if complete flag is set
		if complete {
			rule := strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
			fmt.Fprintf(writer, "%s: %s%s\n", strings.Split(line, ":")[0], baseURL, rule)
		} else {
			fmt.Fprintln(writer, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading robots.txt: %v\n", err)
	}
	return disallowCount, allowCount
}

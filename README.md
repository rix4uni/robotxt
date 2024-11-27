## robotxt

Extract endpoints marked as Allow and Disallow in robots.txt

## Installation
```
go install github.com/rix4uni/robotxt@latest
```

## Download prebuilt binaries
```
wget https://github.com/rix4uni/robotxt/releases/download/v0.0.1/robotxt-linux-amd64-0.0.1.tgz
tar -xvzf robotxt-linux-amd64-0.0.1.tgz
rm -rf robotxt-linux-amd64-0.0.1.tgz
mv robotxt ~/go/bin/robotxt
```
Or download [binary release](https://github.com/rix4uni/robotxt/releases) for your platform.

## Compile from source
```
git clone --depth 1 github.com/rix4uni/robotxt.git
cd robotxt; go install
```

## Usage
```
Usage of robotxt:
  -H string
        Custom User-Agent for requests (default "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")
  -complete
        Include the full URL in the output
  -delay duration
        Delay between requests (default 0s)
  -o string
        File to save the output (default: print to stdout)
  -silent
        silent mode.
  -timeout duration
        Timeout duration for HTTP requests (default 10s)
  -type string
        Specify which one to extract: 'Disallow' or 'Allow'
  -types-count
        Print the count of types at the end
  -verbose
        Enable verbose output
  -version
        Print the version of the tool and exit.
```

## Usage
Single Target:
```
▶ echo "https://www.google.com" | robotxt -silent
```

Multiple Targets:
```
▶ cat targets.txt
https://www.google.com
https://www.facebook.com

▶ cat targets.txt | robotxt -silent
```
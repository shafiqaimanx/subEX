# subex
_knownproject_ has been rename to **subex**

<p align="center">
    <br>
    <img alt="Screenshot" src="https://github.com/shafiqaimanx/subex/blob/main/img/subex-banner.png"/>
    <br>
</p>

<img src='https://img.shields.io/badge/MADE%20WITH-GO-teal?style=flat-square&logo=go'/> <img src='https://img.shields.io/badge/PLATFORM-LINUX-green?style=flat-square&logo=linux'/> <img 
src='https://img.shields.io/badge/PLATFORM-WINDOWS-blue?style=flat-square&logo=windows'/> <img src='https://img.shields.io/badge/PLATFORM-DARWIN-silver?style=flat-square&logo=apple'/> <img src='https://img.shields.io/badge/LICENSE-MIT-orange?style=flat-square&logo=creativecommons'/>

`subex` is a CLI tool for finding any subdomains in passive way possible.

# Disclaimer
Be careful while using this tool because it can be blocked easily. If you get blocked, just wait a couple of minutes and run it again. 
> The author assumes no responsibility.

# Description
`subex` is a tool for passively enumerating subdomains. For now, it's only supports `DNSdumpster` and `GOOGLE` as an engine. Futhermore, by supplying the `-i` flag, you can retrieve any amount of subdomains in the desired (google) pages. Anyways, enjoy! 😉

# Installation
`subex` requires **go1.21** to install successfully. Run the following command to install the latest version

```bash
go install -v github.com/shafiqaimanx/subex@latest
```

But, if you want to build on your own: _follow the step below_
```bash
git clone https://github.com/shafiqaimanx/subex.git
cd subex
go build -o subex main.go
```

Well, you can get the prebuilt release right [here](https://github.com/shafiqaimanx/subEX/releases/tag/v0.2.0)!

# Usage
```bash
[INFO] Usage: subex -d <example.com>
Options:
        -d DOMAIN           Domain to enumerate
        -i INTERACTION      Page interaction [default:10]
        -o OUTPUT           Output results to file [default:subex.txt]
```


# License
License [MIT](https://raw.githubusercontent.com/Doct3rJohn/subEX/main/LICENSE)

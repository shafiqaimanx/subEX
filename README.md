# subEX
_knownproject_ has been rename to subEX

<p align="center">
    <br>
    <img alt="Screenshot" src="https://github.com/Doct3rJohn/subEX/blob/main/img/subEX-banner.png"/>
    <br>
</p>

<img src='https://img.shields.io/badge/MADE%20WITH-GO-teal?style=flat-square&logo=go'/> <img src='https://img.shields.io/badge/PLATFORM-LINUX-green?style=flat-square&logo=linux'/> <img 
src='https://img.shields.io/badge/PLATFORM-WINDOWS-blue?style=flat-square&logo=windows'/> <img src='https://img.shields.io/badge/PLATFORM-DARWIN-silver?style=flat-square&logo=apple'/> <img src='https://img.shields.io/badge/LICENSE-MIT-orange?style=flat-square&logo=creativecommons'/>

`subEX` is the CLI tool for findings any subdomain by using the dorking technique or in other word, in passive way for grabbing subdomains.

# Disclaimer
Be careful while using this tool because it can be blocked easily. If you get blocked, just wait a couple of minutes and run it again. The author assumes no responsibility.

# Description
`subEX` is the tool for passively enumerating subdomains. For now, it only supports the `GOOGLE` engine, and it retrieves the subdomains for about `10` pages long. However, by supplying the `-i` flag, you can retrieve any amount of subdomains in the desired pages. Anyways, enjoy! 😉

# Installation
`subEX` requires **go1.19** to install successfully. Run the following command to install the latest version

```bash
go install -v github.com/Doct3rJohn/subEX@latest
```

But, if you want to build on your own: _follow the step below_
```bash
git clone https://github.com/Doct3rJohn/subEX.git
cd subEX
go build -o subEX main.go
```

Well, you can get the prebuilt release right [here](https://github.com/Doct3rJohn/subEX/releases/tag/v0.1.1)!

# Usage
```bash
[INFO] Usage: subEX -d <example.com>
Options:
        -d DOMAIN           Domain to enumerate
        -i INTERACTION      Page interaction [default:10]
```


# License
License [MIT](https://raw.githubusercontent.com/Doct3rJohn/subEX/main/LICENSE)

# `🌐` KnownProject

<p align="center">
    <br>
    <img alt="Screenshot" src="https://github.com/Doct3rJohn/knownproject/blob/main/img/knownproject-banner.png"/>
    <br>
</p>

<img src='https://img.shields.io/badge/MADE%20WITH-GO-teal?style=flat-square&logo=go'/> <img src='https://img.shields.io/badge/PLATFORM-LINUX-green?style=flat-square&logo=linux'/> <img 
src='https://img.shields.io/badge/PLATFORM-WINDOWS-blue?style=flat-square&logo=windows'/> <img src='https://img.shields.io/badge/LICENSE-MIT-orange?style=flat-square&logo=creativecommons'/>

`KnownProject` is the CLI tool for findings any subdomain by using the dorking technique or in other word, in passive way for grabbing subdomains.

# Disclaimer
Be careful while using this tool because it can be blocked easily. If you get blocked, just wait a couple of minutes and run it again. On top of that, it only supports the `GOOGLE` engine for now, and it only retrieves the data/subdomains for about `10` pages of Google. Anyways, enjoy! 😉

# Installation
Well, you can get the prebuilt release right here! <br>
But, if you want to build on your own: _follow the step below_
```bash
git clone https://github.com/Doct3rJohn/knownproject.git
cd knownproject
go build main.go
```

# Usage
```bash
[+] Usage: knownproject -d <example.com>
[+] Options:
        -d DOMAIN       Domain name to enumerate subdomains
```

# License
License [MIT](https://raw.githubusercontent.com/Doct3rJohn/knownproject/main/LICENSE)
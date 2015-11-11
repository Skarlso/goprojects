Packer Build Wrapper
====================

I had to do this too many times to not write a shorthand. And I wanted to use Go instaed of a bash alias.

App will also perform a validation on the provided json file for an early failure.

I will implement some semantic checks for URLS in the JSON file, and some logical checks around if the provided scripts exist or not.

First, build the app:
```bash
go build packer_build.go
```

Copy it to a PATH location.

And, use it:
```bash
packer_build win8x64.json
```

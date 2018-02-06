# All 'dem icons for the freecloud! ☁️

> Generating icon files is a breeze with this magic masterpiece!✨

## You bring

- A working installation of [Go](https://golang.org)
- This project in `$GOPATH/src/github.com/freecloudio/icons`
- Cloned with submodules


## You get

- An `icons.ts` file, which is automatically generated from the `icons.txt` list.
- Joy 🎉

## What you need to do

First, install Go. If you are a cool person, you already have this 😉
Then, clone this repo by executing

```sh
cd $GOPATH/src/github.com/freecloudio
git clone  --recursive -j2 https://github.com/freecloudio/icons
cd icons
```

Then, edit the `icons.txt` file by adding icons to it or removing unneeded ones. You can find all available icons in `MaterialDesign/icons/svg`. The icon names are their filenames without their extension.

After editing the `icons.txt` file, just run `go run main.go` and there you have it!

*Note:* The parser expects the `icons.txt` file to have `\n` as line separator, so please watch out!
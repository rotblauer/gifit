# `$ gifit`: stdout worth a ~~thousand pictures~~ gif

### Install.
- `go get github.com/rotblauer/gifit`
- `cd $GOPATH/src/github.com/rotblauer/gifit`
- `go install`
<br>
or ...
<br>
- `git clone <this repo>`
- `cd <wherever you cloned it to> && go build`
- `mv gifit $GOPATH/bin/` or wherever else you keep your bin shi

### Use.
`$ gifit` finds a random search result from [Giphy](https://github.com/Giphy/GiphyAPI) and sends it to markdown format in your **stdout**, from which point you can behave as irresponsibly as you wish with said markdowned(default)/embeddable gif. 
<br>
_Why_ would I ever want _that_? 
<br>
1. Put a gif in a commit message ([you are markdowning your commit messages, right?](https://github.com/rotblauer/gitea)):
```shell
$ git commit -m `gifit shipit`
```

2. Illustrate your manuscripts.
```shell
$ gifit squirrels spinning >> Senior_thesis_final_draft.md
```

3. Copy an embeddable/fancily/socially previewable url of a hilarious cats gif to your clipboard.
```shell
$ gifit -e cats in glue | pbcopy
```

4. Get rambunctuous with your [Slack CLI](https://github.com/candrholdings/slack-cli).
```shell
$ slackcli -g professionalcolleagues -m "`gifit -e rat race`"
```

### Options
```shell
    # -s : still image
    gifit -s awesome cats
    > ![awesome+cats](http://media1.giphy.com/media/3o6Zt6dKB6ik0llg0o/giphy_s.gif)
    
    # -e : embed url (display previewly/properly in social posts, etc)
    gifit -e hilarious hamsters
    > http://giphy.com/embed/4HZbQBHDiUwIo
```

![cuz we haz 2](./Giphy Attribution Marks/Animated Logos/Badge/Poweredby_640px_Badge.gif)
Should I Do That
================

Go implementation of the twitter bot [https://twitter.com/shouldidothat](@shouldidothat)

It's pretty shitty.


### How to configure?

Create a json file called ``conf.json`` and put the credentials like this:

```json
{
  "ApiKey": "RxIf8blabalbalb8pnzK7",
  "ApiSecret": "rr6WjeyGwqqblablabalaZwbbsDpKH",
  "AccessToken": "2512654958-blablabalD4Bln9o29wDLnXRnDFK",
  "AccessTokenSecret": "llIjjNKlXZWqSblablabalDgZctDG1QqzN"
}
```

then ``go build`` and now you're ready to run ``./shouldidothat``

You can also specify the path for the config file with the argument ``-c /path/to/config.json``

## How to build docker image?

You can build as usual, with ``docker build -t shouldidothat .``


## How to run docker image?

The image will launch the shouldidothat daemon automatically, but it's expecting
the configuration file to be in ``/etc/shouldidothat/conf.json`

You should run the image with a command like

```
docker run -v /host/shouldidothat/conf.json:/etc/shouldidothat/conf.json shouldidothat
```



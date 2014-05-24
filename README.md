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


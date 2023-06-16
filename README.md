# Mnemonic Encoding Wordlist as a Service

>There are only two hard things in Computer Science: cache invalidation and naming things.
>
>-- Phil Karlton


The nice folks at [mnx.io](http://mnx.io/) wrote a fun [blog post](http://mnx.io/blog/a-proper-server-naming-scheme/) about server naming schemes.

One of the pearls of wisdom is the recommendation to use Oren Tirosh’s [mnemonic encoding project](http://web.archive.org/web/20090918202746/http://tothink.com/mnemonic/wordlist.html) as a wordlist. 

GENIUS!

I’d never heard of that particular wordlist before and immediately downloaded it and built a [shell function](https://gist.github.com/alexlovelltroy/119c32a12f6aca28c3f3) to give me a few examples for codenames.


## Python implementation

There's a flask application in the[python](/python) directory that serves out the list

## Go implementaion

There's a go module you can import with `go get github.com/alexlovelltroy/mnemonic` with an implementation in the [cmd/mnemonic-webservice](/cmd/mnemonic-webservice/) directory.

## Bonus

The Docker project, which is now called moby, has it's own method of generating names based on scientists.  I've always been a fan of that too.

https://github.com/moby/moby/blob/master/pkg/namesgenerator/names-generator.go


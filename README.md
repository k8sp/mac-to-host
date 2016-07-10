# mac-to-hostname

This program maps any MAC address to unique a human readable hostname.


## About MAC Address

According to [Wikipedia](https://en.wikipedia.org/wiki/MAC_address):

> The standard (IEEE 802) format for printing MAC-48 addresses in
> human-friendly form is six groups of two hexadecimal digits,
> separated by hyphens (-) in transmission order
> (e.g. 01-23-45-67-89-ab).

So the total number of unique MAC addresses is 256^6.


## About Hostname

`hostname` is part of FQDN (fully qualifieid domain name).

This
[link](http://stackoverflow.com/questions/8724954/what-is-the-maximum-number-of-characters-for-a-host-name-in-unix)
explains that

> According to RFC 1035 the length of a FQDN is limited to 255
> characters, and each label (node delimited by a dot in the hostname)
> is limited to 63 characters, so in effect the limit you're after is
> 63.

So a hostname contains up to 63 characters.

Also, this
[link](https://technet.microsoft.com/en-us/library/cc959336.aspx)
explains that

> According to RFC 1123, the only characters that can be used in DNS
> labels are "A" to "Z", "a" to "z", "0" to "9", and the hyphen
> ("-"). (The period [.] is also used in DNS names, but only between
> DNS labels and at the end of an FQDN.)


This program adopts a even more restrictive convention:

1. Each character is either an lower case alphabetical letter, "a" to
   "z", or a hyphen ("-"), and
2. hyphen is only used to separate words, (a word is a sequence of
   alphabetical letters).

By this convention, the total number of hostnames depends on the
length of words and the vocabulary size of words.


## English Words

The idea used in this program is to convert each of the 6 bytes of an
MAC address into a word (plus a hyhen if it is not the last word).  We
would like that

1. The first word is an adverb
2. The 2nd, 3rd, 4th, 5th words are adjectives, and
3. the last one is a noun.

For example:

> famously-red-fresh-lovely-shining-apple

### Adverbs

We downloaded a long list of English adverbs from
[EnchantedLearning.com](http://www.enchantedlearning.com/wordlist/adverbs.shtml).

The following Bash command eliminates words containing non
alphabetical letters and choose the shortest 255 of them:

```
$ cat adverbs | gawk '/^[a-z]+$/ {printf("%d %s\n", length($1), $1);}' | sort -n -k 1 | head -n 255 | cut -f 2 -d ' ' | tail
dreamily
entirely
evermore
famously
fiercely
finitely
fluently
formally
formerly
greedily
```

The longest among the chosen has length 8.  It is good since we expect
that each word in a hostname, plus a hyphen, is shorter than 63/6.

### Adjectives

We downloaded a long list of English adjectives from
[EnchantedLearning.com](http://www.enchantedlearning.com/wordlist/adjectives.shtml).
Then we get the shortest 4*255=1020 of them:

```
$ cat adjactives | gawk '/^[a-z]+$/ {printf("%d %s\n", length($1), $1);}' | sort -n -k 1 | head -n 1020 | cut -f 2 -d ' ' | tail
attentive
authentic
automatic
beautiful
bewitched
bountiful
breakable
brilliant
cavernous
cluttered
```

The longest one has length 9 -- shorter than 63/6 - 1.  Nice.

### Nouns

We downloaded a long list of English nouns from
http://www.desiquintans.com/nounlist.  The shorest 255 of them are:

```
$ cat nouns | gawk '/^[a-z]+$/ {printf("%d %s\n", length($1), $1);}' | sort -n -k 1 | head -n 255 | cut -f 2 -d ' ' | tail
bolt
bomb
bone
book
boot
bore
bowl
brow
bulb
bull
```

Very short!

## The Program



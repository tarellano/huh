# Have yo(u) heard

Have you heard (huh) is a last.fm integration which allows you to quickly check if your friends have heard a track, album, or artist.

## How to run

```
cd huh
make build
./bin/huh --api-key=$LASTFM_API_KEY --username=<your-lastfm-username> --track <track-name> --album <album-name> --artist <artist-name>
```
The track, album, and artist are all optional as long as at least one of them is provided.

Example:
```
tarellano@luna huh % ./bin/huh --api-key=$LASTFM_API_KEY --username=tarellano8 --artist "Tame Impala"   
Have you heard Tame Impala (https://www.last.fm/music/Tame+Impala)?
Let's see...
akvsh listened to Tame Impala 887 times.
werdyl listened to Tame Impala 427 times.
bellavirgilio listened to Tame Impala 322 times.
...
kitty_kh listened to Tame Impala 2 times.
mattyykim listened to Tame Impala 0 times.
bleepbloopdhc listened to Tame Impala 0 times.
```

## Design Doc

https://docs.google.com/document/d/1h9zeoeh_ag0aQPEHT4eC9Nkhn-5PrL4-XfweewCHkMw/edit?usp=sharing

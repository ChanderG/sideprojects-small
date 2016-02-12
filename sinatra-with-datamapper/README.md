# sinatra-with-datamapper

### Running
To setup dev environment: 
```
docker build -t chanderg/sinatra-datamapper .
```

To run the server:
```
docker run -it --rm --net=host -v `pwd`:/ws -w /ws chanderg/sinatra-datamapper ruby server.rb
```
or instead use the equivalent `runner.sh`.

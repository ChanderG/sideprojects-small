#! /bin/bash

docker run -it --rm --net=host -v `pwd`:/ws -w /ws chanderg/sinatra-datamapper ruby server.rb

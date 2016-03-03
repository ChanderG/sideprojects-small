if [ "$1" == "" ]; then
  echo "Incorrect usage."
  echo "./runner.sh <option> where option one of:"
  echo "  cli: start docker container"
  echo "  express: run express command"
  exit
fi

if [ "$1" == "cli" ]; then
  docker run -it --rm --net=host -v `pwd`:/ws -w /ws -e http_proxy="http://10.3.100.207:8080/" -e https_proxy="http://10.3.100.207:8080" node /bin/bash
  exit
fi

if [ "$1" == "express" ]; then
  node_modules/express-generator/bin/express "${@:2}"
fi

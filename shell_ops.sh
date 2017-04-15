#!/usr/local/bin/dash

destroy() {
  gopid=$(pgrep "go_server")
  gostatus=$?
  if [ $gostatus = "0" ]
  then
    kill $gopid
    echo "--- KILLED: $gopid ---"
  else
    echo "--- NO GO INSTANCE ---"
  fi
}

build() {
  go build
  echo $?
}

start() {
  ./go_server
}

started() {
  gopid=$(pgrep "go_server")
  gostatus=$?
  if [ $gostatus = "0" ]
  then
    echo "--- STARTED: $(pgrep "go_server") ---"
  fi
}

case $1 in
  1 )
    destroy
    ;;
  2 )
    build
    ;;
  3 )
    start
    ;;
  4 )
    started
    ;;
  * )
    ;;
esac

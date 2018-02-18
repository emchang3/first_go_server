#!/usr/local/bin/dash

destroy() {
  gopid=$(pgrep "go_server")
  gostatus=$?
  if [ $gostatus = "0" ]
  then
    kill $gopid
    echo $gopid
  else
    echo "NOGO"
  fi
}

build() {
  go build
  echo $?
}

start() {
  ./first_go_server
}

started() {
  gopid=$(pgrep "go_server")
  gostatus=$?
  if [ $gostatus = "0" ]
  then
    echo $gopid
  else
    echo "NOGO"
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

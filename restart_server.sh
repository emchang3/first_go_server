ps | grep "go" | while read -r line ; do
  name=$(echo $line | awk '{ print $4 }')
  if [ $name = "./go_server" ]
  then
    pid=$(echo $line | awk '{ print $1 }')
    echo "Found and killed server process: $pid."
    kill -9 $pid
    break
  fi
done

./go_server
echo "Server (re)started."

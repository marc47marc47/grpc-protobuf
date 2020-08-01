

compi(){
  local inflie=$1
  local outfile=$2
  go build -ldflags "-s -w" -o bin/${outfile} cmd/${inflie}
  if [ $? -ne 0 ]; then
    echo "Failed to build $outfile\n" >&2
    exit 1
  fi
}

compi server.go server.exe
compi client.go client.exe
exit 0

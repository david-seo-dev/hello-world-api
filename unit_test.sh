#!/bin/bash

curl_root()
{
    echo "root"
    root_res=`curl http://localhost:8080`
}

curl_status()
{
    echo "status"
    status_res=`curl http://localhost:8080/status`
}

is_expected_result_root()
{
    echo "result_root = $root_res"
}

is_expected_result_status()
{
    echo "result_status = $status_res"
}

PARAMS=""
while (( "$#" )); do
  case "$1" in
    -s|--sha)
      LASTSHA=$2
      shift 2
      ;;
    --) # end argument parsing
      shift
      break
      ;;
    -*|--*=) # unsupported flags
      echo "Error: Unsupported flag $1" >&2
      exit 1
      ;;
    *) # preserve positional arguments
      PARAMS="$PARAMS $1"
      shift
      ;;
  esac
done

echo "arguments entered: $LASTSHA"
curl_root
curl_status
is_expected_result_root
is_expected_result_status
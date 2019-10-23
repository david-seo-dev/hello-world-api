#!/bin/bash

curl_root()
{
    root_res=`curl -s http://localhost:8080`
}

curl_status()
{
    status_res=`curl -s http://localhost:8080/status`
}

is_expected_result_root()
{
    root_match="false"
    if [ "$root_res" = "Hello, World" ]; then
        root_match="true"
    else
        root_match="false"
    fi
}

is_expected_result_status()
{
    status_match="false"
    if [ ! -z "$status_res" ]; then
        status_match="true"
    else 
        status_match="false"
    fi
}

does_expected_results_pass() {
    if [ "$root_match" = "true" ] && [ "$status_match" = "true" ]; then
        echo "basic unit tests PASSED"
    else
        echo "basic unit test FAILED"
        exit 1
    fi
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
does_expected_results_pass
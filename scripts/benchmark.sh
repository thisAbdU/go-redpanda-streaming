#!/bin/bash

# Set default values
STREAM_ID="test_stream"
PAYLOAD='{"payload": "Hello, World!"}'
NUM_REQUESTS=1000
CONCURRENT_REQUESTS=100
API_URL="http://localhost:8080/stream/send/$STREAM_ID"

# Check if hey is installed
if ! command -v hey &> /dev/null
then
    echo "hey could not be found. Please install it first."
    exit
fi

# Run the benchmark
echo "Starting benchmark..."
hey -n $NUM_REQUESTS -c $CONCURRENT_REQUESTS -m POST -d "$PAYLOAD" "$API_URL"

echo "Benchmark completed."
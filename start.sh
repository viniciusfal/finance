#!/bin/sh
cd backend
if [ -f "./bin/server" ]; then
    ./bin/server
else
    echo "Error: bin/server not found in backend directory"
    ls -la backend/
    exit 1
fi


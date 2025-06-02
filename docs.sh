#!/bin/bash

# Script to manage MkDocs documentation

case "$1" in
  serve)
    echo "Starting MkDocs development server..."
    mkdocs serve
    ;;
  build)
    echo "Building MkDocs documentation..."
    mkdocs build
    echo "Documentation built in the 'site/' directory."
    ;;
  *)
    echo "Usage: $0 {serve|build}"
    echo "  serve: Start the MkDocs development server"
    echo "  build: Build the MkDocs documentation site"
    exit 1
    ;;
esac

exit 0
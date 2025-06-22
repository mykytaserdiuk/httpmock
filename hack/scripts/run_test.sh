if go test $(go list ./...) | grep -v '[no test files]' | grep FAIL; then

  exit 1
else
  exit 0
fi
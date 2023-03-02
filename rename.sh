
go mod edit -module github.com/chestarss/elk
#-- rename all imported module
find . -type f -name '*.go' \
  -exec sed -i -e 's,github.com/masseelch/elk,github.com/chestarss/elk,g' {} \;

# main version
VERSION ?= $(shell git describe --tags --always --dirty)

# git commit Hash
COMMIT_HASH ?= $(shell git show -s --format=%H)

# build time 
BUILD_TIME ?= $(shell date +%Y%m%d%H%M%S)

# go file list 
GOFILES := $(shell find . ! -path "./vendor/*" -name "*.go")

# build environments
BUILD_ENV := 

# additional options use for test 
TEST_OPTS := -v

# additional options use for benchmark
BENCHMARK_OPTS := -cpu 1,2,3,4,5,6,7,8 -benchmem 

# sonar report output
REPORT_FOLDER := sonar
TEST_REPORT := ${REPORT_FOLDER}/test.report 
COVER_REPORT := ${REPORT_FOLDER}/cover.report
GOLANGCI_LINT_REPORT := ${REPORT_FOLDER}/golangci-lint.xml 
GOLINT_REPORT := ${REPORT_FOLDER}/golint.report 

.PHONY: format test benchmark sonar clean

# UT
test: 
	${BUILD_ENV} go test ${TEST_OPTS} ./...

# format
format:
	@for f in ${GOFILES} ; do 											\
		gofmt -w $${f};													\
	done																\

# benchmark
benchmark:
	go test -bench . -run ^$$ ${BENCHMARK_OPTS}  ./...

# sonar
sonar: 
	mkdir -p ${REPORT_FOLDER}
	go test -json ./... > ${TEST_REPORT}
	go test -coverprofile=${COVER_REPORT} ./... 
	golangci-lint run --out-format checkstyle  ./... > ${GOLANGCI_LINT_REPORT}
	golint ./... > ${GOLINT_REPORT}
	# sonar-scanner

# clean
clean:
	-rm -f ${TEST_REPORT}
	-rm -f ${COVER_REPORT}
	-rm -f ${GOLANGCI_LINT_REPORT}
	-rm -f ${GOLINT_REPORT}
	-go clean 
	-go clean -cache
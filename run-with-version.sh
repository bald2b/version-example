go run  -ldflags "-X 'main.version=1.0' -X 'main.gitBranch=`git branch --show-current`' -X 'main.buildDate=`date`'" main.go
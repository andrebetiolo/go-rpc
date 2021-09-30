# Load enviroment variables
# include .env

# Export enviroment variables to commands
export

help:: ## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | sort | fgrep -v fgrep | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run-server:: ## Run the server
	@ go run server/main.go

run-client:: ## Run the client
	@ go run client/main.go

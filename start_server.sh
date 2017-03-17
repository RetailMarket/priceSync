if [[ ! -e out/ ]];
		then
			mkdir out/
		fi 

		go build -o out/build app/priceSync/main.go; ./out/build

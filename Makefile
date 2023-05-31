file:='file=@./testData/matrix.csv'

full-flow:
	curl -F $(file) "localhost:8080/echo"
	@echo ""

	curl -F $(file) "localhost:8080/invert"
	@echo ""

	curl -F $(file) "localhost:8080/flatten"
	@echo ""

	curl -F $(file) "localhost:8080/sum"
	@echo ""

	curl -F $(file) "localhost:8080/multiply"

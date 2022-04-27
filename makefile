genmocks:
	source scripts/gen_mocks.sh && generate_mocks

test: genmocks
	source scripts/test.sh && test_all

deploy: test
	source scripts/deploy.sh && deploy

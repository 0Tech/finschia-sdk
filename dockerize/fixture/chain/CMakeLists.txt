set(FIXTURE_DAEMON_VERSION local CACHE STRING "binary version for the chains (local)")

configure_file(generate-compose.py.in generate-compose.py
  @ONLY)
configure_file(setup-chains.sh.in setup-chains.sh
  @ONLY)
configure_file(cleanup-chains.sh cleanup-chains.sh
  COPYONLY)

add_custom_target(chain)

function(add_chains
	name
	num_chains
	num_regions
	num_sentries
	timeout
  )

  get_target_property(CHAIN_BINARY_DIR chain BINARY_DIR)
  add_test(
	NAME setup_chains_${name}
	COMMAND timeout ${timeout} sh -c "PROJECT_NAME=${name} NUM_CHAINS=${num_chains} NUM_REGIONS=${num_regions} NUM_SENTRIES=${num_sentries} sh ${CHAIN_BINARY_DIR}/setup-chains.sh")
  set_tests_properties(setup_chains_${name}
	PROPERTIES FIXTURES_SETUP ${name})
  add_test(
	NAME cleanup_chains_${name}
	COMMAND sh -c "PROJECT_NAME=${name} sh ${CHAIN_BINARY_DIR}/cleanup-chains.sh")
  set_tests_properties(cleanup_chains_${name}
	PROPERTIES FIXTURES_CLEANUP ${name})
endfunction(add_chains)

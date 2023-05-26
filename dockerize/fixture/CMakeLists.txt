set(FIXTURE_NUM_CHAINS 2 CACHE STRING "number of the chains") # TO REMOVE
set(FIXTURE_NUM_REGIONS 4 CACHE STRING "number of the regions")
set(FIXTURE_NUM_SENTRIES 2 CACHE STRING "number of the sentries")

set(FIXTURE_CLIENT_KEYRING_BACKEND test CACHE STRING "keyring backend")
mark_as_advanced(FIXTURE_CLIENT_KEYRING_BACKEND)

set(FIXTURE_CONFIG_TIMEOUT_COMMIT 1s CACHE STRING "timeout_commit")
set(FIXTURE_CONFIG_TIMEOUT_PRECOMMIT 1s CACHE STRING "timeout_precommit")
set(FIXTURE_CONFIG_TIMEOUT_PREVOTE 1s CACHE STRING "timeout_prevote")
set(FIXTURE_CONFIG_TIMEOUT_PROPOSE 1s CACHE STRING "timeout_propose")
mark_as_advanced(
  FIXTURE_CONFIG_TIMEOUT_COMMIT
  FIXTURE_CONFIG_TIMEOUT_PRECOMMIT
  FIXTURE_CONFIG_TIMEOUT_PREVOTE
  FIXTURE_CONFIG_TIMEOUT_PROPOSE)

set(FIXTURE_STATE_BOND_DENOM cony CACHE STRING "bond denom")
mark_as_advanced(FIXTURE_STATE_BOND_DENOM)
set(FIXTURE_STATE_BALANCE 1000000000 CACHE STRING "initial balance of each operator (in ${FIXTURE_STATE_BOND_DENOM})")
set(FIXTURE_STATE_DELEGATION 1000000 CACHE STRING "initial delegation of each operator to its validator (in ${FIXTURE_STATE_BOND_DENOM}, MUST be >1000000)")
set(FIXTURE_STATE_MIN_DEPOSIT 1 CACHE STRING "x/gov min_deposit in ${FIXTURE_STATE_BOND_DENOM}")
set(FIXTURE_STATE_VOTING_PERIOD 10s CACHE STRING "x/gov voting_period")
set(FIXTURE_STATE_UNBONDING_TIME 10s CACHE STRING "x/staking unbonding_time")

set(DAEMON_ALLOW_DOWNLOAD_BINARIES true CACHE STRING "whether automatically download the binary or not")

configure_file(generate-compose.py.in generate-compose.py
  @ONLY)
configure_file(daemon.env.in daemon.env
  @ONLY)
configure_file(setup-chains.sh.in setup-chains.sh
  @ONLY)
configure_file(cleanup-chains.sh cleanup-chains.sh
  COPYONLY)

add_test(
  NAME setup_chains
  COMMAND timeout 3m sh setup-chains.sh)
set_tests_properties(setup_chains
  PROPERTIES FIXTURES_SETUP CHAIN)
add_test(
  NAME cleanup_chains
  COMMAND sh cleanup-chains.sh)
set_tests_properties(cleanup_chains
  PROPERTIES FIXTURES_CLEANUP CHAIN)

add_subdirectory(upgrade)
add_subdirectory(blackout)

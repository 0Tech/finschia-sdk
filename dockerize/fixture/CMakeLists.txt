# environment variables (certain fixtures may override these)
set(FIXTURE_NUM_REGIONS 4 CACHE STRING "number of the regions in fixtures")
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

set(FIXTURE_DAEMON_VERSION local CACHE STRING "version of daemon. `local` is locally built one.")
set(FIXTURE_DAEMON_VERSION_NEW default CACHE STRING "version of new daemon. `local` is locally built one, `dummy` is dummy binary, `default` is scenario default one (`dummy` or `local`).")

add_subdirectory(daemon)
add_subdirectory(runtime)

add_subdirectory(chain)

add_subdirectory(upgrade)
add_subdirectory(blackout)

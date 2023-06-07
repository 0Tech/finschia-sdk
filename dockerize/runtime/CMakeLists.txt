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

execute_process(
  COMMAND sh -c "printf $(id -u)"
  OUTPUT_VARIABLE DAEMON_UID)
execute_process(
  COMMAND sh -c "printf $(id -g)"
  OUTPUT_VARIABLE DAEMON_GID)
configure_file(Dockerfile.in Dockerfile
  @ONLY)

configure_file(init binaries/init
  COPYONLY)
configure_file(setup binaries/setup
  COPYONLY)
configure_file(healthcheck.in binaries/healthcheck
  @ONLY)
configure_file(install-image.sh.in install-image.sh
  @ONLY)

get_target_property(DAEMON_BINARY_DIR build_daemon BINARY_DIR)
get_target_property(COSMOVISOR_BINARY_DIR build_cosmovisor BINARY_DIR)
add_custom_command(
  DEPENDS install-image.sh Dockerfile binaries/init binaries/setup binaries/healthcheck build
  OUTPUT install-image.touch
  COMMAND sh install-image.sh ${DAEMON_BINARY_DIR}/daemon ${COSMOVISOR_BINARY_DIR}/cosmovisor install-image.touch
  COMMENT "Installing the Docker image"
  VERBATIM)
add_custom_target(install_image
  DEPENDS install-image.touch)
install(CODE "execute_process(COMMAND ${CMAKE_COMMAND} --build ${CMAKE_BINARY_DIR} --target install_image)")

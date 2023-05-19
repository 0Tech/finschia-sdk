set(DAEMON_NUM_INSTANCES 4)
set(DAEMON_KEY_MNEMONIC "")
set(DAEMON_NET_BEGIN 192.168.0.2/16 CACHE STRING "start address of the network")

set(CHAIN_ID testnet)
set(DAEMON_KEYRING_BACKEND test)

set(DAEMON_TIMEOUT_COMMIT 1s)
set(DAEMON_TIMEOUT_PRECOMMIT 1s)
set(DAEMON_TIMEOUT_PREVOTE 1s)
set(DAEMON_TIMEOUT_PROPOSE 1s)

set(CHAIN_MIN_DEPOSIT 1)
set(CHAIN_VOTING_PERIOD 5s)
set(CHAIN_UNBONDING_TIME 5s)
set(CHAIN_UPGRADE_NAME v1-Algae)

set(DAEMON_ALLOW_DOWNLOAD_BINARIES false)

get_target_property(DEPLOY_SOURCE_DIR deploy SOURCE_DIR)
configure_file(${DEPLOY_SOURCE_DIR}/daemon.env.in daemon.env
  @ONLY)
get_target_property(TOOL_SOURCE_DIR tool SOURCE_DIR)
configure_file(upgrade.sh.in upgrade.sh
  @ONLY)
configure_file(dummyd.in dummyd
  @ONLY)

get_target_property(DEPLOY_BINARY_DIR deploy BINARY_DIR)
add_custom_command(
  OUTPUT docker-compose.yml
  BYPRODUCTS dsimd
  DEPENDS deploy ${DEPLOY_BINARY_DIR}/dsimd ${DEPLOY_BINARY_DIR}/generate-compose.py
  COMMAND sh -c "set -e; cp ${DEPLOY_BINARY_DIR}/dsimd .; [ -f docker-compose.yml ] && ./dsimd down || true; ${DEPLOY_BINARY_DIR}/generate-compose.py ${DAEMON_NUM_INSTANCES} --network ${DAEMON_NET_BEGIN} -o docker-compose.yml"
  COMMENT "Generating Docker compose file: dummy"
  VERBATIM)
add_custom_target(dummy
  DEPENDS docker-compose.yml)

add_test(
  NAME test_upgrade_dummy
  COMMAND sh upgrade.sh)

set(DAEMON_NUM_INSTANCES 4 CACHE STRING "number of the instances")
set(DAEMON_KEY_MNEMONIC "mind flame tobacco sense move hammer drift crime ring globe art gaze cinnamon helmet cruise special produce notable negative wait path scrap recall have" CACHE STRING "mnemonic for keys (set empty values to get random keys)")
set(DAEMON_NET_BEGIN 192.168.0.2/16 CACHE STRING "start address of the network")
mark_as_advanced(DAEMON_NET_BEGIN)
set(DAEMON_KEYRING_BACKEND test CACHE STRING "keyring backend, options are: os, file, test")
mark_as_advanced(DAEMON_KEYRING_BACKEND)

set(DAEMON_TIMEOUT_COMMIT 1s CACHE STRING "timeout_commit")
set(DAEMON_TIMEOUT_PRECOMMIT 1s CACHE STRING "timeout_precommit")
set(DAEMON_TIMEOUT_PREVOTE 1s CACHE STRING "timeout_prevote")
set(DAEMON_TIMEOUT_PROPOSE 1s CACHE STRING "timeout_propose")

set(CHAIN_ID testnet CACHE STRING "chain id")
set(CHAIN_MIN_DEPOSIT 1000000 CACHE STRING "voting period")
set(CHAIN_VOTING_PERIOD 1m CACHE STRING "voting period")
set(CHAIN_UNBONDING_TIME 1m CACHE STRING "unbonding time")

set(DAEMON_ALLOW_DOWNLOAD_BINARIES true CACHE BOOL "allow cosmovisor downloading new binaries")

get_target_property(DEPLOY_SOURCE_DIR deploy SOURCE_DIR)
configure_file(${DEPLOY_SOURCE_DIR}/daemon.env.in daemon.env
  @ONLY)

get_target_property(DEPLOY_BINARY_DIR deploy BINARY_DIR)
add_custom_command(
  OUTPUT docker-compose.yml
  BYPRODUCTS dsimd
  DEPENDS deploy ${DEPLOY_BINARY_DIR}/dsimd ${DEPLOY_BINARY_DIR}/generate-compose.py
  COMMAND sh -c "set -e; cp ${DEPLOY_BINARY_DIR}/dsimd .; [ -f docker-compose.yml ] && ./dsimd down || true; ${DEPLOY_BINARY_DIR}/generate-compose.py ${DAEMON_NUM_INSTANCES} --network ${DAEMON_NET_BEGIN} -o docker-compose.yml"
  COMMENT "Generating Docker compose file: simple"
  VERBATIM)
add_custom_target(simple
  DEPENDS docker-compose.yml)

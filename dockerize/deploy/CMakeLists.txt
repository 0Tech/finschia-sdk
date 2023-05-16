set(DAEMON_NUM_INSTANCES 7 CACHE STRING "number of the instances")
set(DAEMON_KEY_MNEMONIC "mind flame tobacco sense move hammer drift crime ring globe art gaze cinnamon helmet cruise special produce notable negative wait path scrap recall have" CACHE STRING "mnemonic for keys")
set(DAEMON_NET_BEGIN 192.168.0.2/16 CACHE STRING "start address of the network")

set(CHAIN_ID testnet CACHE STRING "chain id")
set(DAEMON_KEYRING_BACKEND test CACHE STRING "keyring backend, options are: os, file, test")

set(DAEMON_TIMEOUT_COMMIT 1s CACHE STRING "timeout_commit")
set(DAEMON_TIMEOUT_PRECOMMIT 1s CACHE STRING "timeout_precommit")
set(DAEMON_TIMEOUT_PREVOTE 1s CACHE STRING "timeout_prevote")
set(DAEMON_TIMEOUT_PROPOSE 1s CACHE STRING "timeout_propose")

set(CHAIN_VOTING_PERIOD 172800s CACHE STRING "voting period")
set(CHAIN_UNBONDING_TIME 1814400s CACHE STRING "unbonding time")

set(DAEMON_ALLOW_DOWNLOAD_BINARIES true CACHE BOOL "allow cosmovisor downloading new binaries")

configure_file(generate-compose.py.in generate-compose.py
  @ONLY)
configure_file(daemon.env.in daemon.env
  @ONLY)
configure_file(dsimd.in dsimd
  @ONLY)

add_custom_command(
  OUTPUT docker-compose.yml
  DEPENDS generate-compose.py
  COMMAND sh -c "[ -f docker-compose.yml ] && ./dsimd down || true; ./generate-compose.py ${DAEMON_NUM_INSTANCES} --network ${DAEMON_NET_BEGIN} -o docker-compose.yml"
  COMMENT "Generating Docker compose file"
  VERBATIM)
add_custom_target(deploy
  DEPENDS docker-compose.yml install_image)

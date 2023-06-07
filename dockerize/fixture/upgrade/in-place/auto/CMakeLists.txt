set(DAEMON_ALLOW_DOWNLOAD_BINARIES true)

configure_file(../upgrade-chain.sh.in upgrade-chain.sh
  @ONLY)
configure_file(../dummyd dummyd
  COPYONLY)
configure_file(${IMAGE_SOURCE_DIR}/daemon.env.in daemon.env
  @ONLY)

add_test(
  NAME upgrade_chain_auto
  COMMAND timeout 5m sh upgrade-chain.sh ChainsUpgradeAuto 0)
add_chains(ChainsUpgradeAuto 1 ${FIXTURE_NUM_REGIONS} 1 90s)
set_tests_properties(upgrade_chain_auto
  PROPERTIES FIXTURES_REQUIRED ChainsUpgradeAuto)

set(FIXTURE_NUM_REGIONS 4)
set(DAEMON_ALLOW_DOWNLOAD_BINARIES false)

configure_file(../upgrade-chain.sh.in upgrade-chain.sh
  @ONLY)
configure_file(../dummyd bundle/${DAEMON_NAME}
  COPYONLY)
get_target_property(IMAGE_SOURCE_DIR install_image SOURCE_DIR)
configure_file(${IMAGE_SOURCE_DIR}/daemon.env.in daemon.env
  @ONLY)

add_test(
  NAME upgrade_chain_manual
  COMMAND timeout 2m sh upgrade-chain.sh ChainsUpgradeManual 0)
add_chains(ChainsUpgradeManual 1 ${FIXTURE_NUM_REGIONS} 1 2m)
set_tests_properties(upgrade_chain_manual
  PROPERTIES FIXTURES_REQUIRED ChainsUpgradeManual)

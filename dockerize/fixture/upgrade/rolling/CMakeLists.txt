set(FIXTURE_NUM_REGIONS 7)

configure_file(upgrade-chain.sh.in upgrade-chain.sh
  @ONLY)
# configure_file(dummyd bundle/${DAEMON_NAME}
#   COPYONLY)
get_target_property(IMAGE_SOURCE_DIR install_image SOURCE_DIR)
configure_file(${IMAGE_SOURCE_DIR}/daemon.env.in daemon.env
  @ONLY)

add_test(
  NAME upgrade_chain_rolling
  COMMAND timeout 5m sh upgrade-chain.sh ChainsUpgradeRolling 0)
add_chains(ChainsUpgradeRolling 1 ${FIXTURE_NUM_REGIONS} 1 5m)
set_tests_properties(upgrade_chain_rolling
  PROPERTIES FIXTURES_REQUIRED ChainsUpgradeRolling)

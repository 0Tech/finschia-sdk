configure_file(upgrade-chain.sh.in upgrade-chain.sh
  @ONLY)
configure_file(${IMAGE_SOURCE_DIR}/daemon.env.in daemon.env
  @ONLY)

add_test(
  NAME upgrade_chain_rolling
  COMMAND timeout 3m sh upgrade-chain.sh ChainsUpgradeRolling 0)
add_chains(ChainsUpgradeRolling 1 ${FIXTURE_NUM_REGIONS} 1 3m)
set_tests_properties(upgrade_chain_rolling
  PROPERTIES FIXTURES_REQUIRED ChainsUpgradeRolling)

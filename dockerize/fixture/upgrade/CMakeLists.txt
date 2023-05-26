set(FIXTURE_CHAIN_UPGRADE_NAME v1-Algae)

configure_file(upgrade-chain.sh.in upgrade-chain.sh
  @ONLY)
configure_file(dummyd dummyd
  COPYONLY)

add_test(
  NAME upgrade_chain
  COMMAND timeout 2m sh upgrade-chain.sh 1)
set_tests_properties(upgrade_chain
  PROPERTIES FIXTURES_REQUIRED CHAIN)

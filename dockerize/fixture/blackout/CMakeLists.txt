set(FIXTURE_NUM_REGIONS 4)

configure_file(shutdown-sentries.sh.in shutdown-sentries.sh
  @ONLY)
get_target_property(IMAGE_SOURCE_DIR install_image SOURCE_DIR)
configure_file(${IMAGE_SOURCE_DIR}/daemon.env.in daemon.env
  @ONLY)

add_test(
  NAME shutdown_sentries
  COMMAND timeout 2m sh shutdown-sentries.sh ChainsShutdownSentries 0)
add_chains(ChainsShutdownSentries 1 ${FIXTURE_NUM_REGIONS} 1 2m)
set_tests_properties(shutdown_sentries
  PROPERTIES FIXTURES_REQUIRED ChainsShutdownSentries)

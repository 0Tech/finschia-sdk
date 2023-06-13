configure_file(restart-sentries.sh.in restart-sentries.sh
  @ONLY)
get_target_property(IMAGE_SOURCE_DIR install_image SOURCE_DIR)
configure_file(${IMAGE_SOURCE_DIR}/daemon.env.in daemon.env
  @ONLY)

add_test(
  NAME restart_sentries
  COMMAND timeout 2m sh restart-sentries.sh ChainsRestartSentries 0)
add_chains(ChainsRestartSentries 1 ${FIXTURE_NUM_REGIONS} 1 3m)
set_tests_properties(restart_sentries
  PROPERTIES FIXTURES_REQUIRED ChainsRestartSentries)

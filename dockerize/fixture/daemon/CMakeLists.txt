configure_file(prepare-daemons.sh.in prepare-daemons.sh
  @ONLY)
configure_file(dummyd dummy/${DAEMON_NAME}
  COPYONLY)

get_target_property(DAEMON_BINARY_DIR build_daemon BINARY_DIR)
add_test(
  NAME prepare_daemons
  COMMAND sh prepare-daemons.sh ${DAEMON_BINARY_DIR}/bundle.tar.gz)
set_tests_properties(prepare_daemons
  PROPERTIES FIXTURES_SETUP Daemons)

add_custom_target(prepare_daemons)

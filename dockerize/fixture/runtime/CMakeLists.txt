execute_process(
  COMMAND sh -c "printf $(id -u)"
  OUTPUT_VARIABLE DAEMON_UID)
execute_process(
  COMMAND sh -c "printf $(id -g)"
  OUTPUT_VARIABLE DAEMON_GID)
configure_file(Dockerfile.in Dockerfile
  @ONLY)

configure_file(init binaries/init
  COPYONLY)
configure_file(setup binaries/setup
  COPYONLY)
configure_file(healthcheck.in binaries/healthcheck
  @ONLY)

get_target_property(DAEMON_BINARY_DIR prepare_daemons BINARY_DIR)
configure_file(install-image.sh.in install-image.sh
  @ONLY)
configure_file(uninstall-image.sh.in uninstall-image.sh
  @ONLY)

get_target_property(COSMOVISOR_BINARY_DIR build_cosmovisor BINARY_DIR)
add_test(
  NAME install_image
  COMMAND sh install-image.sh ${COSMOVISOR_BINARY_DIR}/cosmovisor ${DAEMON_BINARY_DIR}/${FIXTURE_DAEMON_VERSION}/bundle.tar.gz)
set_tests_properties(install_image
  PROPERTIES FIXTURES_SETUP ImageSetup)
set_tests_properties(install_image
  PROPERTIES FIXTURES_REQUIRED Daemons)

add_test(
  NAME uninstall_image
  COMMAND sh uninstall-image.sh)
set_tests_properties(uninstall_image
  PROPERTIES FIXTURES_CLEANUP ImageCleanup)

add_custom_target(install_image)

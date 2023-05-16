execute_process(
  COMMAND sh -c "printf $(id -u)"
  OUTPUT_VARIABLE DAEMON_UID)
execute_process(
  COMMAND sh -c "printf $(id -g)"
  OUTPUT_VARIABLE DAEMON_GID)
configure_file(Dockerfile.in Dockerfile
  @ONLY)
unset(DAEMON_GID)
unset(DAEMON_UID)

configure_file(init init
  COPYONLY)
configure_file(install-image.sh.in install-image.sh
  @ONLY)

get_target_property(DAEMON_BINARY_DIR build_daemon BINARY_DIR)
get_target_property(COSMOVISOR_BINARY_DIR build_cosmovisor BINARY_DIR)
add_custom_command(
  DEPENDS install-image.sh Dockerfile init build
  OUTPUT install-image.touch
  COMMAND ./install-image.sh ${DAEMON_BINARY_DIR}/daemon ${COSMOVISOR_BINARY_DIR}/cosmovisor install-image.touch
  COMMENT "Installing the Docker image"
  VERBATIM)
add_custom_target(install_image
  DEPENDS install-image.touch)
install(CODE "execute_process(COMMAND ${CMAKE_COMMAND} --build ${CMAKE_BINARY_DIR} --target install_image)")

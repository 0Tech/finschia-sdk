configure_file(Dockerfile.in Dockerfile
  @ONLY)
configure_file(build.sh.in build.sh
  @ONLY)

add_custom_target(generate_daemon_source_fingerprint
  BYPRODUCTS source.sum
  DEPENDS generate-source-fingerprint.sh
  COMMAND ${CMAKE_CURRENT_SOURCE_DIR}/generate-source-fingerprint.sh source.sum ${GO_SOURCE_DIR} ${PROJECT_DIR_NAME}
  VERBATIM)
add_custom_command(
  DEPENDS up-to-date.sh generate_daemon_source_fingerprint
  OUTPUT source.touch
  BYPRODUCTS source.sum~
  COMMAND ${CMAKE_CURRENT_SOURCE_DIR}/up-to-date.sh source.sum source.touch
  COMMENT "Checking daemon source update"
  VERBATIM)
add_custom_command(
  OUTPUT daemon
  DEPENDS build.sh Dockerfile source.touch
  COMMAND ./build.sh Dockerfile ${GO_SOURCE_DIR}
  COMMENT "Building daemon"
  VERBATIM)
add_custom_target(build_daemon
  DEPENDS daemon)

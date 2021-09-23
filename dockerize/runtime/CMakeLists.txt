execute_process(
  COMMAND sh -c "printf $(id -u)"
  OUTPUT_VARIABLE UID)
  # OUTPUT_QUIET)
execute_process(
  COMMAND sh -c "printf $(id -g)"
  OUTPUT_VARIABLE GID)
  # OUTPUT_QUIET)
configure_file(Dockerfile.in Dockerfile
  @ONLY)
unset(UID)
unset(GID)

set(BINARY simd)
configure_file(init.in init
  @ONLY)

add_custom_target(install_simd
  DEPENDS install-simd.sh Dockerfile
  COMMAND ${CMAKE_CURRENT_SOURCE_DIR}/install-simd.sh ${CMAKE_BINARY_DIR}/${BINARY}/${BINARY})
unset(BINARY)

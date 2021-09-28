execute_process(
  COMMAND sh -c "printf $(id -u)"
  OUTPUT_VARIABLE SIMD_UID)
execute_process(
  COMMAND sh -c "printf $(id -g)"
  OUTPUT_VARIABLE SIMD_GID)
configure_file(Dockerfile.in Dockerfile
  @ONLY)
unset(SIMD_GID)
unset(SIMD_UID)

set(SIMD_BINARY simd)
configure_file(init.in init
  @ONLY)

add_custom_target(install_simd
  DEPENDS install-simd.sh Dockerfile init
  COMMAND ${CMAKE_CURRENT_SOURCE_DIR}/install-simd.sh ${CMAKE_BINARY_DIR}/${SIMD_BINARY}/${SIMD_BINARY})

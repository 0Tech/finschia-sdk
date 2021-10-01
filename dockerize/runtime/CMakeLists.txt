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

configure_file(init init
  COPYONLY)

get_target_property(BUILD_SIMD_DIR build_simd BINARY_DIR)
add_custom_command(
  DEPENDS install-simd.sh Dockerfile init build_simd ${BUILD_SIMD_DIR}/simd
  OUTPUT install-simd.touch
  COMMAND ${CMAKE_CURRENT_SOURCE_DIR}/install-simd.sh ${BUILD_SIMD_DIR}/simd install-simd.touch
  COMMENT "Installing a Docker image of simd"
  VERBATIM)
add_custom_target(install_simd
  DEPENDS install-simd.touch)

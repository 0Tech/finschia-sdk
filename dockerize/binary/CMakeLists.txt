set(SIMD_BUILDER_NAME cosmossdk/rbuilder:latest CACHE STRING "Image name of the builder")
mark_as_advanced(SIMD_BUILDER_NAME)

file(RELATIVE_PATH BUILD_SCRIPT_RELPATH ${GO_SOURCE_DIR} ${CMAKE_CURRENT_SOURCE_DIR}/build-simd-native.sh)
configure_file(build-simd.sh.in build-simd.sh
  @ONLY)

add_custom_target(generate_source_fingerprint
  DEPENDS generate-source-fingerprint.sh
  COMMAND ${CMAKE_CURRENT_SOURCE_DIR}/generate-source-fingerprint.sh source.sum ${GO_SOURCE_DIR} ${PROJECT_DIR_NAME}
  VERBATIM)
add_custom_command(
  DEPENDS up-to-date.sh generate_source_fingerprint source.sum
  OUTPUT source.touch
  COMMAND ${CMAKE_CURRENT_SOURCE_DIR}/up-to-date.sh source.sum source.touch
  COMMENT "Checking source update"
  VERBATIM)
add_custom_command(
  OUTPUT simd build_report artifact.tar.gz
  DEPENDS build-simd.sh source.touch
  COMMAND ./build-simd.sh ${GO_SOURCE_DIR} ${SIMD_BUILDER_NAME}
  COMMENT "Building simd"
  VERBATIM)
add_custom_target(build_simd ALL
  DEPENDS simd)
